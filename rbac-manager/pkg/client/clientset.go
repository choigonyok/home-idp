package client

import (
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/storage"
)

type RbacManagerClientSet struct {
	StorageClient *storage.RbacManagerStorageClient
	MailClient    mail.MailClient
}

func EmptyClientSet() *RbacManagerClientSet {
	return &RbacManagerClientSet{}
}

func (cs *RbacManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.StorageClient:
		tmp := &storage.RbacManagerStorageClient{}
		tmp.Set(i)
		cs.StorageClient = tmp
		return
	// case util.GitClient:
	// 	tmp := &git.InstallManagerGitClient{}
	// 	tmp.Set(i)
	// 	cs.GitClient = tmp
	// 	return
	// case util.HttpClient:
	// 	tmp := &http.InstallManagerHttpClient{}
	// 	tmp.Set(i)
	// 	cs.HttpClient = tmp
	// 	return
	// case util.KubeClient:
	// 	tmp := &kube.InstallManagerKubeClient{}
	// 	tmp.Set(i)
	// 	cs.KubeClient = tmp
	// 	return
	default:
		return
	}
}
