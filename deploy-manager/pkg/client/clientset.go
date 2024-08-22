package client

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/util"
)

type DeployManagerClientSet struct {
	// GrpcClient map[util.Components]client.GrpcClient
	KubeClient *kube.DeployManagerKubeClient
}

func EmptyClientSet() *DeployManagerClientSet {
	return &DeployManagerClientSet{
		// GrpcClient: make(map[util.Components]client.GrpcClient),
	}
}

func (cs *DeployManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.KubeClient:
		tmp := &kube.DeployManagerKubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	default:
		return
	}
}
