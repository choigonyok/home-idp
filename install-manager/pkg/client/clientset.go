package client

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManagerClientSet struct {
	GrpcClient    map[util.Components]client.GrpcClient
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
	HelmClient    *helm.HelmClient
}

func EmptyClientSet() *InstallManagerClientSet {
	return &InstallManagerClientSet{
		GrpcClient: make(map[util.Components]client.GrpcClient, client.ClientTotalCount),
	}
}

func (cs *InstallManagerClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.GrpcRbacManagerClient:
		tmp := &grpc.InstallManagerGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.RbacManager] = tmp
		return
	case util.HelmClient:
		tmp := &helm.HelmClient{}
		tmp.Set(i)
		cs.HelmClient = tmp
		return
	default:
		return
	}
}
