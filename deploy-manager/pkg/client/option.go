package client

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/docker/docker/api/types/registry"
	dockercli "github.com/docker/docker/client"
)

func WithDockerClient() client.ClientOption {
	return useDockerClient()
}

func useDockerClient() client.ClientOption {
	return newDockerClientOption(func(cli client.ClientSet) {
		client, _ := dockercli.NewClientWithOpts(dockercli.FromEnv, dockercli.WithVersion("1.43"))

		cfg := registry.AuthConfig{
			Username: env.Get("DEPLOY_MANAGER_DOCKER_USERNAME"),
			Password: env.Get("DEPLOY_MANAGER_DOCKER_PASSWORD"),
		}

		config, _ := registry.EncodeAuthConfig(cfg)

		i := &docker.DockerClient{
			Client:         client,
			AuthCredential: &config,
		}

		cli.Set(util.DockerClient, i)
		// client.LoginWithEnv()
		return
	})
}

func newDockerClientOption(f func(cli client.ClientSet)) *client.GrpcClientOption {
	return &client.GrpcClientOption{
		F: f,
	}
}
