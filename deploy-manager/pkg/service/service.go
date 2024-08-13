package service

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/client"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/grpc"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
)

type DeployManager struct {
	ClientSet *client.DeployManagerClientSet
	Server    *grpc.DeployManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *DeployManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &DeployManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *DeployManager) Stop() {
	svc.Server.Close()
	return
}

func (svc *DeployManager) Start() {
	// pbServer := &grpc.ArgoCDServer{
	// 	HelmClient: svc.ClientSet.HelmClient,
	// }

	// pb.RegisterArgoCDServer(svc.Server.Grpc, pbServer)
	svc.Server.Run()
	return
}
