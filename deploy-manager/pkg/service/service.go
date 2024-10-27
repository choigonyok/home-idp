package service

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/client"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
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
}

func (svc *DeployManager) Start() {
	pbBuildServer := &grpc.BuildServer{
		KubeClient: svc.ClientSet.KubeClient,
	}
	pbDeployServer := &grpc.DeployServer{
		GitClient:  svc.ClientSet.GitClient,
		KubeClient: svc.ClientSet.KubeClient,
	}

	pb.RegisterBuildServer(svc.Server.Server, pbBuildServer)
	pb.RegisterDeployServer(svc.Server.Server, pbDeployServer)

	svc.Server.Run()
}
