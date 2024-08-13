package client

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/docker"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/util"
)

type DeployManagerClientSet struct {
	GrpcClient   map[util.Components]client.GrpcClient
	MailClient   mail.MailClient
	KubeClient   *kube.KubeClient
	DockerClient *docker.DockerClient
}

func EmptyClientSet() *DeployManagerClientSet {
	return &DeployManagerClientSet{
		GrpcClient: make(map[util.Components]client.GrpcClient, client.ClientTotalCount),
	}
}

func (cs *DeployManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.GrpcRbacManagerClient:
		tmp := &grpc.InstallManagerGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.RbacManager] = tmp
		return
	case util.DockerClient:
		tmp := &docker.DockerClient{}
		tmp.Set(i)
		cs.DockerClient = tmp
		return
	default:
		return
	}
}
