package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"log"
	"os"

	gssh "golang.org/x/crypto/ssh"
)

func main() {

	data := []byte(``)

	publicKeys, err := ssh.NewPublicKeys("git", data, "")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := publicKeys.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf.User = "clearcodecn"
	conf.Auth = []gssh.AuthMethod{
		gssh.KeyboardInteractive(SshInteractive),
	}
	publicKeys.HostKeyCallback = gssh.InsecureIgnoreHostKey()

	repo, err := git.PlainClone("./repo", false, &git.CloneOptions{
		URL:      "git@codechina.csdn.net:clearcodecn/dpull.git",
		Auth:     publicKeys,
		Depth:    1,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	// The second parameter is unused
	for n, _ := range questions {
		fmt.Println(n)
	}

	return answers, nil
}
