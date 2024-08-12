package client

import (
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

const clientCount = 5

type GatewayClientSet struct {
	GrpcClient    map[util.Components]client.GrpcClient
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
}

func EmptyClientSet() *GatewayClientSet {
	m := make(map[util.Components]client.GrpcClient, clientCount)

	for i := range m {
		m[i] = nil
	}

	return &GatewayClientSet{
		GrpcClient: make(map[util.Components]client.GrpcClient, clientCount),
	}
}

func (cs *GatewayClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.GrpcInstallManagerClient:
		tmp := &GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.InstallManager] = tmp
		return
	case util.GrpcRbacManagerClient:
		tmp := &GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.RbacManager] = tmp
		return
	case util.GrpcDeployManagerClient:
		tmp := &GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.DeployManager] = tmp
		return
	case util.GrpcSecretManagerClient:
		tmp := &GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.SecretManager] = tmp
		return
	default:
		return
	}
}
