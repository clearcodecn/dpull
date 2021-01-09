package cmd

import (
	"bytes"
	"context"
	"errors"
	"github.com/clearcodecn/dpull"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
	return initConfig()
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
		option = dpull.DefaultOption
		var buf = bytes.NewBuffer(nil)
		err := yaml.NewEncoder(buf).Encode(option)
		if err != nil {
			return err
		}
		os.MkdirAll(filepath.Dir(configFilePath), 0755)
		err = ioutil.WriteFile(configFilePath, buf.Bytes(), 0755)
		if err != nil {
			return err
		}
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

	color.Green("init config success: %s", configFilePath)

	return initRepo(option.RepoOption)
}

// create repo
func initRepo(option dpull.RepoOption) error {
	client, err := dpull.NewGitClient(option, dpull.GitClientOption{})
	if err != nil {
		return err
	}

	color.Green("init repo ...")
	err = client.Clone(context.Background())
	if err != nil {
		if !errors.Is(err, dpull.ErrRepoAlreadyExist) {
			return err
		}
	}
	color.Green("init repo success !")
	return nil
}
