package client

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/git"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/trace"
	"github.com/choigonyok/home-idp/pkg/util"
)

type DeployManagerClientSet struct {
	// GrpcClient map[util.Components]client.GrpcClient
	TraceClient *trace.TraceClient
	KubeClient  *kube.DeployManagerKubeClient
	GitClient   *git.DeployManagerGitClient
}

func EmptyClientSet() *DeployManagerClientSet {
	return &DeployManagerClientSet{}
}

func (cs *DeployManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.TraceClient:
		tmp := &trace.TraceClient{}
		tmp.Set(i)
		cs.TraceClient = tmp
		return
	case util.KubeClient:
		tmp := &kube.DeployManagerKubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	case util.GitClient:
		tmp := &git.DeployManagerGitClient{}
		tmp.Set(i)
		cs.GitClient = tmp
		return
	default:
		return
	}
}
