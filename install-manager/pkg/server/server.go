package server

import (
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
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

	// pbServer := &grpc.UserServiceServer{
	// 	StorageClient: svr.StorageClient,
	// }
	// pb.RegisterUserServiceServer(s.Server, pbServer)

	return svr
}
