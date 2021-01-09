package main

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	gssh "golang.org/x/crypto/ssh"
	"log"
)

var defaultPrivateKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEA1fZpC6/CjrdDl7DbQmKohLboTYWc9PjclN4FYhPsizRbxBR098RU
7gTlbJ3vKMIlqHlMCUY2OwEb1rpoAuTu5JEoQl1HKX8eKq41CsfLrkc+VZuTqdteMPqnhK
JQqrRXnx3CMsWO6KED6v+3UIkRTENxrVyeVXKsKzUsnlVU47KzVEl+RTofHMo4+CaCBFFs
8ZyebGOHBTKi0JUeLD6PcXrzXCwqPQbCj+8JEda1zx68MU5GArqFuxsgEXI0jsJ3hJxM6x
DsfdZl+PTdpkkII2kl5gN4HfInrjHc+rp+yvQ0F4CNm2Lx5SvMjiG46eDZDxOCdlqgYKki
sPco4zQTOwAAA7irpi1bq6YtWwAAAAdzc2gtcnNhAAABAQDV9mkLr8KOt0OXsNtCYqiEtu
hNhZz0+NyU3gViE+yLNFvEFHT3xFTuBOVsne8owiWoeUwJRjY7ARvWumgC5O7kkShCXUcp
fx4qrjUKx8uuRz5Vm5Op214w+qeEolCqtFefHcIyxY7ooQPq/7dQiRFMQ3GtXJ5VcqwrNS
yeVVTjsrNUSX5FOh8cyjj4JoIEUWzxnJ5sY4cFMqLQlR4sPo9xevNcLCo9BsKP7wkR1rXP
HrwxTkYCuoW7GyARcjSOwneEnEzrEOx91mX49N2mSQgjaSXmA3gd8ieuMdz6un7K9DQXgI
2bYvHlK8yOIbjp4NkPE4J2WqBgqSKw9yjjNBM7AAAAAwEAAQAAAQBIP8bE7XqzGms2o7/G
MO5asjDLTJztk8NYeYgz0CqF7w41rfq5V5CeNwUJomMJzlVNCHiGgTD6x6sQ3S0WHRwWDn
Ybwseu2X/kRaMfmsvKc8A2xCwepTavL1S10uGOYwtbbX8QCenx370k82iBR2eR6wxN0AKf
M/OzO2dvp7zcjfLjMRKFSbroDOm33mMEqOSfpGJthCWMpHSZzTNlygFrYFfj6ouCzcNJLY
ibUsoDCelDkhBCDEkIsuGfx3Kcja/AkESrwBvdqDs1VXiwgCnsVue8Un4liNl/UPG5ITkl
O+0E61LE3XVw2oCm9cawfRNTDHLuJJUi7+/aYsSuxQlBAAAAgQCv+Ms/eFMD8JV1NtH0G9
QHJXpiyS84G+hFrWSN7QZejrFOEFH0MvLVJypF47k0Ca0P3qh5nLLfHrtoDcM8uOAVtbE6
LRWjknLJKwprQn5X4wY2iZKLU5mEj3riIeP0Z5RDsePaSKlAmy7yrfbMq4mlN9xzY672am
UIV0f57FEROAAAAIEA+WdxDkSbnYmcHcu5bvrU242WlqIisj0Z38wLHRVsO9Ci8Sx6LoiY
4wE8XOf6BLToAvltmnmZsVsoOAymR3M4Mx+EpwgZ3082RfiOA8QJ0iII7EsmYBVwzs2XxD
H7/MofseeL5hSFEcRFYs7hvKlQ4QHQ53MDD7pO7AwKMqAmgpEAAACBANufBDDoBc5jl2ac
bSIt1ZSZtejbvMgYDgMl1YeLUX3KDUDEJ22iUnTALSg1e7yImLli30LAWP3G4cgtCnVVGV
whKjBVEBbPNiEAXupJm9e4ZTLJ627K/4rNjk79sQobbJQhj6VrVQuquNFVTiAiEmOKUB4A
mDE4JKT55aN8kocLAAAAAAEC
-----END OPENSSH PRIVATE KEY-----`

func main() {
	auth, err := ssh.NewPublicKeys(`git`, []byte(defaultPrivateKey), ``)
	if err != nil {
		log.Println(err)
	}
	auth.HostKeyCallback = gssh.InsecureIgnoreHostKey()

	repo, err := git.PlainOpen("/Users/zhangmengjin/.dpull/repo")
	if err != nil {
		log.Fatal(err)
	}

	tag := `release-vazhzLmdjci5pby9wYXVzZTozLjI_`

	t, _ := repo.Worktree()
	err = t.PullContext(context.Background(), &git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.NewTagReferenceName(tag),
		SingleBranch:  true,
		Depth:         1,
		Auth:          auth,
	})
	log.Println(err)

}
