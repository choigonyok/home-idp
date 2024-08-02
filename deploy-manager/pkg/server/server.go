package server

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/docker"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/grpc"
	"github.com/choigonyok/home-idp/deploy-manager/pkg/kube"
	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/util"
)

type DeployManager struct {
	Server       server.Server
	Config       config.Config
	KubeClient   *kube.KubeClient
	DockerClient *docker.DockerClient
}

func (deploy *DeployManager) Close() error {
	if err := deploy.Server.Close(); err != nil {
		return err
	}
	if err := deploy.DockerClient.Close(); err != nil {
		return err
	}

	return nil
}

func (s *DeployManager) Run() {
	s.Server.Run()
}

func New(component util.Components, cfg config.Config) server.Server {
	s := grpc.NewServer()
	dc := docker.New()
	dc.LoginWithEnv()

	svr := &DeployManager{
		Server:       s,
		Config:       cfg,
		KubeClient:   kube.NewKubeClient(),
		DockerClient: dc,
	}

	pbServer := &grpc.ManifestServiceServer{
		// StorageClient: svr.StorageClient,
	}
	pb.RegisterManifestServiceServer(s.Server, pbServer)

	return svr
}
