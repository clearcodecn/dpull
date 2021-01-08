package dpull

import "github.com/docker/docker/client"

type Client struct {
	dockerDaemonClient *client.APIClient
}
