package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	rootCommand *cobra.Command
)

func init() {
	rootCommand = &cobra.Command{
		Use:   "dpull",
		Short: "a docker image pull tool to help pull image",
	}

	rootCommand.AddCommand(
		InitCommand,
		PullCommand,
	)

	dir, _ := homedir.Dir()
	configFilePath = filepath.Join(dir, "dpull", "config.yaml")
	rootCommand.PersistentFlags().StringVarP(&configFilePath, "config", "c", configFilePath, "config file path")
}

func Execute() error {
	return rootCommand.Execute()
}
