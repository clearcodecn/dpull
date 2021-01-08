package cmd

import (
	"bytes"
	"context"
	"github.com/clearcodecn/dpull"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	InitCommand *cobra.Command
)

func init() {
	InitCommand = &cobra.Command{
		Use:     "init",
		Short:   "init the dpull tool kit",
		Example: "dpull init",
		RunE:    initRun,
	}
}

func initRun(cmd *cobra.Command, args []string) error {
	// 1. 初始化 home config 目录
	initConfig()
	// 1. 初始化 git 仓库
	// 2.
}

func initConfig() error {
	var (
		exist  = true
		option dpull.Option
	)

	_, err := os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			exist = false
		} else {
			return err
		}
	}

	if !exist {
		option = dpull.DefaultConfigOption
	} else {
		data, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return err
		}
		err = yaml.NewDecoder(bytes.NewReader(data)).Decode(&option)
		if err != nil {
			return err
		}
	}

}

// create repo
func initRepo(option dpull.RepoOption) {

}
