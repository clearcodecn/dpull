package dpull

import (
	"context"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGit(t *testing.T) {
	option := DefaultRepoOption
	os.RemoveAll(option.StorePath)
	git, err := NewGitClient(option)
	require.Nil(t, err)

	err = git.Clone(context.Background())
	require.Nil(t, err)

	err = git.Pull(context.Background())
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "already up-to-date")

	_, err = os.Stat(option.StorePath)
	require.Nil(t, err)

	_, err = os.Stat(filepath.Join(option.StorePath, ".git"))
	require.Nil(t, err)

	data := "FROM ubuntu"
	dockerfile := filepath.Join(option.StorePath, "Dockerfile")
	err = ioutil.WriteFile(dockerfile, []byte(data), 0755)
	require.Nil(t, err)

	err = git.AddAndCommit([]string{dockerfile}, "change docker file")
	require.Nil(t, err)

	// tags:release-v$version
	err = git.TagAndPush("release-vgit-test3", "change docker file")
	require.Nil(t, err)
}
