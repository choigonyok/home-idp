package client

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/git"
	"github.com/choigonyok/home-idp/install-manager/pkg/http"
	"github.com/choigonyok/home-idp/install-manager/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManagerClientSet struct {
	// GrpcClient map[util.Components]client.GrpcClient
	HelmClient *helm.HelmClient
	KubeClient *kube.InstallManagerKubeClient
	HttpClient *http.InstallManagerHttpClient
	GitClient  *git.InstallManagerGitClient
}

func EmptyClientSet() *InstallManagerClientSet {
	return &InstallManagerClientSet{
		// GrpcClient: make(map[util.Components]client.GrpcClient),
	}
}

func (cs *InstallManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	// case util.GrpcRbacManagerClient:
	// 	tmp := &grpc.InstallManagerGrpcClient{}
	// 	tmp.Set(i)
	// 	cs.GrpcClient[util.RbacManager] = tmp
	// 	return
	case util.HelmClient:
		tmp := &helm.HelmClient{}
		tmp.Set(i)
		cs.HelmClient = tmp
		return
	case util.GitClient:
		tmp := &git.InstallManagerGitClient{}
		tmp.Set(i)
		cs.GitClient = tmp
		return
	case util.HttpClient:
		tmp := &http.InstallManagerHttpClient{}
		tmp.Set(i)
		cs.HttpClient = tmp
		return
	case util.KubeClient:
		tmp := &kube.InstallManagerKubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	default:
		return
	}
}
