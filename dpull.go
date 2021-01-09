package dpull

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/containerd/containerd/errdefs"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"time"
)

type Proxy struct {
	Option       Option
	DockerClient *Client
	GitClient    *GitClient
}

type PullOption struct {
	ForceProxy bool
}

func (p *Proxy) Pull(ctx context.Context, ref string, opt PullOption) error {
	// 1. get domain
	domain, _ := splitDockerDomain(ref)
	if domain == defaultDomain && !opt.ForceProxy {
		color.Green("pulling image %s ...", ref)
		err := p.DockerClient.Pull(ref, os.Stdout)
		if err != nil {
			return err
		}
		color.Green("pulling image %s done", ref)

		return nil
	}
	imgTag := strings.Replace(base64.StdEncoding.EncodeToString([]byte(ref)), "=", "_", -1)

	color.Green("pulling from proxy")
	color.Green("(1/5) modify dockerfile...")
	filename, err := p.GitClient.ModifyDockerfile(ref, imgTag)
	if err != nil {
		return err
	}
	color.Green("(2/5) commit changes...")
	msg := fmt.Sprintf("modify image %s", ref)
	err = p.GitClient.AddAndCommit([]string{filename}, msg)
	if err != nil {
		return err
	}
	color.Green("(3/5) push tag to remote...")
	tag, err := p.GitClient.TagAndPush(imgTag, msg)

	imgWithTag := fmt.Sprintf("%s:%s", p.Option.MirrorOption.ImageBasePath, tag)

	color.Green("(4/5) wait remote building...")
	now := time.Now()
	buf := bytes.NewBuffer(nil)
	for {
		err := p.DockerClient.Pull(imgWithTag, buf)
		if err != nil {
			time.Sleep(time.Second)
			if errdefs.IsNotFound(err) {
				color.Green("(5/5) waiting %s, press ^-C to cancel \r", time.Now().Sub(now).String())
			}
			continue
		}
		color.Green("(4/5) pulling...")
		io.Copy(os.Stdout, buf)
		break
	}
	err = p.DockerClient.Tag(imgWithTag, ref)
	if err != nil {
		return errors.Wrap(err, "failed to tag image: "+imgWithTag+" to "+ref)
	}

	err = p.DockerClient.Remove(imgWithTag)
	if err != nil {
		return errors.Wrap(err, "failed to remove image: "+imgWithTag)
	}

	color.Green("download %s success", ref)

	return nil
}
