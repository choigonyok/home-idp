package client

import (
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/util"
)

func WithKubeClient() client.ClientOption {
	return useKubeClient()
}

func useKubeClient() client.ClientOption {
	return newKubeClientOption(func(cli client.ClientSet) {
		i := kube.NewKubeClient()
		cli.Set(util.KubeClient, i)
		return
	})
}

func newKubeClientOption(f func(cli client.ClientSet)) *client.GrpcClientOption {
	return &client.GrpcClientOption{
		F: f,
	}
}
