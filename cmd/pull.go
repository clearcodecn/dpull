package cmd

import (
	"context"
	"github.com/clearcodecn/dpull"
	"github.com/spf13/cobra"
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
