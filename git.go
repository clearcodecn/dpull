package dpull

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/pkg/errors"
	gssh "golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GitClient struct {
	option RepoOption

	repo *git.Repository
}

func NewGitClient(option RepoOption) (*GitClient, error) {
	cli := new(GitClient)
	cli.option = option

	if cli.option.SSHUrl == "" {
		return nil, errors.New("repo ssh url can't be empty")
	}
	if cli.option.GitEmail == "" {
		return nil, errors.New("repo email can't be empty")
	}
	if cli.option.GitUsername == "" {
		return nil, errors.New("repo username can't be empty")
	}
	if cli.option.PrivateKey == "" {
		return nil, errors.New("repo private key can't be empty")
	}
	if cli.option.SSHUser == "" {
		return nil, errors.New("repo ssh user can't be empty")
	}
	if cli.option.StorePath == "" {
		return nil, errors.New("repo store path can't be empty")
	}

	err := os.MkdirAll(filepath.Dir(cli.option.StorePath), 0755)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (c *GitClient) Clone(ctx context.Context) error {
	ok, _ := c.checkRepoExist()
	if ok {
		return fmt.Errorf("path: %s already exist", c.option.StorePath)
	}
	auth, err := c.getAuth()
	if err != nil {
		return err
	}
	color.Green("git clone --depth=1 %s", c.option.SSHUrl)

	_, err = git.PlainCloneContext(ctx, c.option.StorePath, false, &git.CloneOptions{
		URL:      c.option.SSHUrl,
		Auth:     auth,
		Progress: c,
		Depth:    1,
	})

	if err != nil {
		return errors.Wrap(err, "failed to clone repo")
	}
	return nil
}

func (c *GitClient) checkRepoExist() (bool, error) {
	_, err := os.Stat(c.option.StorePath)
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		}
		if !os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (c *GitClient) lazyInit() error {
	if c.repo != nil {
		return nil
	}
	repo, err := git.PlainOpen(c.option.StorePath)
	if err != nil {
		return err
	}
	c.repo = repo
	return nil
}

func (c *GitClient) getAuth() (*ssh.PublicKeys, error) {
	publicKeys, err := ssh.NewPublicKeys(c.option.SSHUser, []byte(c.option.PrivateKey), c.option.PrivateKeyPassword)
	if err != nil {
		return nil, err
	}
	publicKeys.HostKeyCallback = gssh.InsecureIgnoreHostKey()
	return publicKeys, nil
}

func (c *GitClient) Pull(ctx context.Context) error {
	if err := c.lazyInit(); err != nil {
		return err
	}

	wt, err := c.repo.Worktree()
	if err != nil {
		return err
	}

	auth, err := c.getAuth()
	if err != nil {
		return err
	}

	err = wt.PullContext(ctx, &git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: "refs/heads/master",
		SingleBranch:  true,
		Auth:          auth,
		Progress:      c,
		Force:         true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GitClient) AddAndCommit(paths []string, msg string) error {
	if err := c.lazyInit(); err != nil {
		return err
	}
	t, err := c.repo.Worktree()
	if err != nil {
		return err
	}
	for _, p := range paths {
		p = strings.TrimPrefix(p, c.option.StorePath+string(os.PathSeparator))
		_, err = t.Add(p)
		if err != nil {
			return err
		}
	}
	user := &object.Signature{
		Name:  c.option.GitUsername,
		Email: c.option.GitEmail,
		When:  time.Now(),
	}
	_, err = t.Commit(msg, &git.CommitOptions{
		All:       true,
		Author:    user,
		Committer: user,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GitClient) TagAndPush(tag string, msg string) error {
	if err := c.lazyInit(); err != nil {
		return err
	}
	ref, err := c.repo.Head()
	if err != nil {
		return err
	}
	if tagExists(tag, c.repo) {
		return errors.New("tag already exist")
	}
	_, err = c.repo.CreateTag(tag, ref.Hash(), &git.CreateTagOptions{
		Message: msg,
	})
	if err != nil {
		return err
	}

	auth, err := c.getAuth()
	if err != nil {
		return err
	}
	po := &git.PushOptions{
		RemoteName: "origin",
		Progress:   c,
		RefSpecs:   []config.RefSpec{config.RefSpec("refs/tags/*:refs/tags/*")},
		Auth:       auth,
	}
	err = c.repo.Push(po)
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			return nil
		}
		return err
	}
	return nil
}

func (c *GitClient) Write(b []byte) (int, error) {
	return os.Stdout.Write([]byte(color.BlackString(string(b))))
}

func (c *GitClient) RemoveTag(tag string) error {
	if err := c.lazyInit(); err != nil {
		return err
	}
	err := c.repo.DeleteTag(tag)
	if err != nil {
		return err
	}
	c.repo.Push(&git.PushOptions{
		RefSpecs: []config.RefSpec{":refs/tags/*:refs/tags/*"},
	})

}

func tagExists(tag string, r *git.Repository) bool {
	tagFoundErr := "tag was found"
	tags, err := r.TagObjects()
	if err != nil {
		return false
	}
	res := false
	err = tags.ForEach(func(t *object.Tag) error {
		if t.Name == tag {
			res = true
			return fmt.Errorf(tagFoundErr)
		}
		return nil
	})
	if err != nil && err.Error() != tagFoundErr {
		return false
	}
	return res
}
