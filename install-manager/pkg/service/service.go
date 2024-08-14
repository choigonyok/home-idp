package service

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/client"
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	"github.com/choigonyok/home-idp/install-manager/pkg/helm"

	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/env"
)

type InstallManager struct {
	ClientSet *client.InstallManagerClientSet
	Server    *grpc.InstallManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *InstallManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &InstallManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *InstallManager) Stop() {
	svc.Server.Stop()
	return
}

func (svc *InstallManager) Start() {
	svc.ClientSet.HelmClient.AddRepository("bitnami", "https://charts.bitnami.com/bitnami", true)
	svc.ClientSet.HelmClient.AddRepository("argo", "https://argoproj.github.io/argo-helm", true)
	svc.ClientSet.HelmClient.AddRepository("harbor", "https://helm.goharbor.io", true)
	svc.ClientSet.HelmClient.AddRepository("jenkins", "https://charts.jenkins.io", true)

	pbServer := &grpc.ArgoCDServer{
		HelmClient: svc.ClientSet.HelmClient,
	}

	pb.RegisterArgoCDServer(svc.Server.Grpc, pbServer)

	svc.installDefaultServices()

	svc.ClientSet.HttpClient.RequestJenkinsCrumb()

	svc.Server.Run()
	return
}

func (svc *InstallManager) installDefaultServices() {
	if env.Get("DEFAULT_REGISTRY_ENABLED") == "true" {
		cli := helm.NewHarborClient(env.Get("GLOBAL_NAMESPACE"), "home-idp")
		cli.Install(*svc.ClientSet.HelmClient)
	}
	if env.Get("DEFAULT_CD_ENABLED") == "true" {
		cli := helm.NewArgoCDClient(env.Get("GLOBAL_NAMESPACE"), "home-idp-cd")
		cli.Install(*svc.ClientSet.HelmClient)
	}
	if env.Get("DEFAULT_CI_ENABLED") == "true" {
		cli := helm.NewJenkinsClient(env.Get("GLOBAL_NAMESPACE"), "home-idp-ci")
		cli.Install(*svc.ClientSet.HelmClient)
	}
}
