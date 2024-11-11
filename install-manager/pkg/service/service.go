package service

import (
	"fmt"
	"time"

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

	svc.Server.Run()
	return
}

func (svc *InstallManager) installDefaultServices() {
	svc.installPostgresql()

	if env.Get("DEFAULT_REGISTRY_ENABLED") == "true" {
		svc.installHarbor()
	}
	// if env.Get("DEFAULT_CI_ENABLED") == "true" {
	// 	cli := helm.NewJenkinsClient(env.Get("HOME_IDP_NAMESPACE"), "home-idp-ci")
	// 	cli.Install(*svc.ClientSet.HelmClient)
	// }
	if env.Get("DEFAULT_CD_ENABLED") == "true" {
		svc.installArgoCD()
	}
}

func (svc *InstallManager) installPostgresql() {
	// TODO: transfer sql file to comfigmap

	c := helm.NewPostgresClient(env.Get("HOME_IDP_NAMESPACE"), "home-idp-postgres")
	c.Install(*svc.ClientSet.HelmClient)
}

func (svc *InstallManager) installArgoCD() {
	c := helm.NewArgoCDClient(env.Get("HOME_IDP_NAMESPACE"), "home-idp-cd")

	c.Install(*svc.ClientSet.HelmClient)

	for {
		ok := svc.ClientSet.KubeClient.IsArgoCDRunning(env.Get("HOME_IDP_NAMESPACE"))
		if ok {
			fmt.Println("@@@TEST ARGOCD HEALTH CHECK SUCCESS@@@")
			break
		}
		fmt.Println("TEST ARGOCD NOT RUNNING YET")
		time.Sleep(time.Second * 1)
	}

	pw := svc.ClientSet.KubeClient.GetArgoCDPassword()
	fmt.Println("TEST ARGOCD PASSWORD:", pw)

	token, err := svc.ClientSet.HttpClient.CreateArgoCDSessionToken(pw)
	if err != nil {
		fmt.Println("TEST ARGOCD SESSION CREATE ERR:", err)
	}

	if err := svc.ClientSet.HttpClient.CreateArgoCDRepository(pw, token); err != nil {
		fmt.Println("TEST ARGOCD REPOSITORY CREATE ERR:", err)
	}

	svc.ClientSet.GitClient.CreateArgoCDApplicationManifest(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_NAMESPACE"))
}

func (svc *InstallManager) installHarbor() {
	c := helm.NewHarborClient(env.Get("HOME_IDP_NAMESPACE"), "home-idp-harbor")

	if err := svc.ClientSet.KubeClient.ApplyHarborCredentialSecret(); err != nil {
		fmt.Println("TEST APPLY HARBOR CRED MANIFEST ERR:", err)
	}

	c.Install(*svc.ClientSet.HelmClient)

	for {
		ok, err := svc.ClientSet.HttpClient.IsHarborHealthy()
		fmt.Println("TEST HARBOR HEALTH CHECK REQUEST ERR: ", err)
		if ok {
			fmt.Println("@@@TEST HARBOR HEALTH CHECK SUCCESS@@@")
			break
		}
		time.Sleep(time.Second * 1)
	}
}
