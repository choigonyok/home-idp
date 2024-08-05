package server

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/install-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManager struct {
	Server server.Server
	Config config.Config
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

func New(component util.Components, cfg config.Config) server.Server {
	s := grpc.NewServer()

	svr := &InstallManager{
		Server: s,
	}

	pbServer := &grpc.ArgoCDServer{}
	pb.RegisterArgoCDServer(s.Server, pbServer)

	return svr
}
