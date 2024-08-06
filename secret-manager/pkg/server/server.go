package server

import (
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/secret-manager/pkg/grpc"
	secretstorage "github.com/choigonyok/home-idp/secret-manager/pkg/storage"
)

type SecretServer struct {
	Server        *grpc.SecretManagerServer
	StorageClient storage.StorageClient
	Config        config.Config
}

func (s *SecretServer) Close() error {
	if err := s.Server.Listener.Close(); err != nil {
		return err
	}
	if err := s.StorageClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *SecretServer) Run() {
	s.Server.Server.Serve(s.Server.Listener)
}

func New(component util.Components, cfg config.Config) server.Server {
	s := grpc.NewServer()
	sc, _ := secretstorage.NewClient(component)
	svr := &SecretServer{
		Server:        s,
		StorageClient: sc,
		Config:        cfg,
	}

	// pbServer := &grpc.UserServiceServer{
	// 	StorageClient: svr.StorageClient,
	// }
	// pb.RegisterUserServiceServer(s.Server, pbServer)

	return svr
}
