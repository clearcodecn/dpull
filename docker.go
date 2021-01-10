package dpull

import (
	"bufio"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"io"
	"os"
	"os/exec"
	"strings"
)

var (
	legacyDefaultDomain = "index.docker.io"
	defaultDomain       = "docker.io"
	officialRepoName    = "library"
	defaultTag          = "latest"
)

func InitDocker() {
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("docker", "version")
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	var index = 0
	r := bufio.NewReader(buf)
	for {
		data, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if strings.Contains(string(data), "API version") {
			index++
			if index == 2 {
				version := strings.Split(strings.TrimPrefix(strings.Trim(strings.Trim(string(data), " "), "API version:"), " "), " ")[0]
				os.Setenv("DOCKER_API_VERSION", version)
			}
		}
	}
	var err error
	cli, err := client.NewClientWithOpts()
	if err == nil {
		err = client.FromEnv(cli)
		if err == nil {
			DefaultClient = &Client{dockerDaemonClient: cli}
		}
	}
}

type Client struct {
	dockerDaemonClient *client.Client
}

var (
	DefaultClient *Client
)

func (c *Client) Pull(ref string, w io.Writer) error {
	resp, err := c.dockerDaemonClient.ImagePull(context.Background(), ref, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()
	fd, isTerminal := term.GetFdInfo(w)
	if err := jsonmessage.DisplayJSONMessagesStream(resp, os.Stdout, fd, isTerminal, nil); err != nil {
		return err
	}
	return nil
}

// splitDockerDomain splits a repository name to domain and remotename string.
// If no valid domain is found, the default domain is used. Repository name
// needs to be already validated before.
func splitDockerDomain(name string) (domain, remainder string) {
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ".:") && name[:i] != "localhost") {
		domain, remainder = defaultDomain, name
	} else {
		domain, remainder = name[:i], name[i+1:]
	}
	if domain == legacyDefaultDomain {
		domain = defaultDomain
	}
	if domain == defaultDomain && !strings.ContainsRune(remainder, '/') {
		remainder = officialRepoName + "/" + remainder
	}
	return
}

func (c *Client) Tag(source string, dst string) error {
	return c.dockerDaemonClient.ImageTag(context.Background(), source, dst)
}

func (c *Client) Remove(source string) error {
	_, err := c.dockerDaemonClient.ImageRemove(context.Background(), source, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: false,
	})
	return err
}
