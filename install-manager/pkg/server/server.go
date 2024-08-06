package server

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/install-manager/pkg/config"
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManager struct {
	Server     server.Server
	Config     *config.InstallManagerConfig
	HelmClient *helm.HelmClient
}

func (install *InstallManager) Close() error {
	if err := install.Server.Close(); err != nil {
		return err
	}

	return nil
}

func (s *InstallManager) Run() {
	s.Server.Run()
}

func New(component util.Components, cfg *config.InstallManagerConfig) server.Server {
	s := grpc.NewServer()
	h := helm.New()
	h.AddRepository("bitnami", "https://charts.bitnami.com/bitnami", true)
	h.AddRepository("argo", "https://argoproj.github.io/argo-helm", true)

	svr := &InstallManager{
		Server:     s,
		HelmClient: h,
		Config:     cfg,
	}

	svr.SetEnvFromConfig()

	pbServer := &grpc.ArgoCDServer{
		HelmClient: h,
	}
	pb.RegisterArgoCDServer(s.Server, pbServer)

	return svr
}

func (c *InstallManager) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("INSTALL_MANAGER_PORT", strconv.Itoa(c.Config.Service.Port))
}
