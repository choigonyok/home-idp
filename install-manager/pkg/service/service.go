package service

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/client"
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
)

type InstallManager struct {
	ClientSet *client.InstallManagerClientSet
	Server    *grpc.InstallManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *InstallManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &InstallManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *InstallManager) Stop() {
	svc.Server.Stop()
	return
}

func (svc *InstallManager) Start() {
	svc.ClientSet.HelmClient.AddRepository("bitnami", "https://charts.bitnami.com/bitnami", true)
	svc.ClientSet.HelmClient.AddRepository("argo", "https://argoproj.github.io/argo-helm", true)

	pbServer := &grpc.ArgoCDServer{
		HelmClient: svc.ClientSet.HelmClient,
	}

	pb.RegisterArgoCDServer(svc.Server.Grpc, pbServer)
	svc.Server.Run()
	return
}
