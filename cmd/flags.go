package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configFilePath string

	serverAddress string = "https://dpull.clearcode.cn"
)


func applyFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&serverAddress, "server", "s", "", "server address")
}
