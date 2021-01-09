package dpull

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(DefaultMirrorOption.ImageBasePath + ":git-test3")
	err := DefaultClient.Pull(DefaultMirrorOption.ImageBasePath+":git-test3", os.Stdout)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
	require.Nil(t, err)
}
