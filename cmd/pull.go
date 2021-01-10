package cmd

import (
	"bytes"
	"context"
	"github.com/clearcodecn/dpull"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	PullCommand *cobra.Command
)

func init() {
	PullCommand = &cobra.Command{
		Use:     "pull",
		Short:   "pull docker images",
		Example: "dpull pull image:a image:b image:c",
		RunE:    runPull,
	}
}

func runPull(cmd *cobra.Command, args []string) error {
	opt := dpull.DefaultOption
	if configFilePath != "" {
		data, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return errors.Wrap(err, "failed to read config")
		}
		err = yaml.NewDecoder(bytes.NewReader(data)).Decode(&opt)
		if err != nil {
			return errors.Wrap(err, "failed to parse config")
		}
	}
	gitClient, err := dpull.NewGitClient(opt.RepoOption, dpull.GitClientOption{})
	if err != nil {
		return err
	}

	proxy := dpull.Proxy{
		Option:       opt,
		DockerClient: dpull.DefaultClient,
		GitClient:    gitClient,
	}

	for _, a := range args {
		err := proxy.Pull(context.Background(), a, dpull.PullOption{

		})
		if err != nil {
			return err
		}
	}
	return nil
}
