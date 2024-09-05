package client

import (
	"github.com/choigonyok/home-idp/gateway/pkg/git"
	"github.com/choigonyok/home-idp/gateway/pkg/grpc"
	"github.com/choigonyok/home-idp/gateway/pkg/http"
	"github.com/choigonyok/home-idp/gateway/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
)

type GatewayClientSet struct {
	GrpcClient    map[util.Components]client.GrpcClient
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
	KubeClient    *kube.GatewayKubeClient
	GitClient     *git.GatewayGitClient
	HttpClient    *http.GatewayHttpClient
}

func EmptyClientSet() *GatewayClientSet {
	return &GatewayClientSet{
		GrpcClient: make(map[util.Components]client.GrpcClient),
	}
}

func (cs *GatewayClientSet) Set(cli util.Clients, i interface{}) {
	switch cli {
	case util.GrpcInstallManagerClient:
		tmp := &grpc.GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.InstallManager] = tmp
		return
	case util.GrpcRbacManagerClient:
		tmp := &grpc.GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.RbacManager] = tmp
		return
	case util.GrpcDeployManagerClient:
		tmp := &grpc.GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.DeployManager] = tmp
		return
	case util.GrpcSecretManagerClient:
		tmp := &grpc.GatewayGrpcClient{}
		tmp.Set(i)
		cs.GrpcClient[util.SecretManager] = tmp
		return
	case util.KubeClient:
		tmp := &kube.GatewayKubeClient{}
		tmp.Set(i)
		cs.KubeClient = tmp
		return
	case util.GitClient:
		tmp := &git.GatewayGitClient{}
		tmp.Set(i)
		cs.GitClient = tmp
		return
	case util.HttpClient:
		tmp := &http.GatewayHttpClient{}
		tmp.Set(i)
		cs.HttpClient = tmp
		return
	default:
		return
	}
}
