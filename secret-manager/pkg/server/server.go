package server

import (
	"log"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/secret-manager/pkg/config"
	"github.com/choigonyok/home-idp/secret-manager/pkg/grpc"
	secretstorage "github.com/choigonyok/home-idp/secret-manager/pkg/storage"
)

type SecretManager struct {
	Server        *grpc.SecretManagerServer
	StorageClient storage.StorageClient
	Config        *config.SecretManagerConfig
}

func (s *SecretManager) Close() error {
	if err := s.Server.Listener.Close(); err != nil {
		return err
	}
	if err := s.StorageClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *SecretManager) Run() {
	s.Server.Server.Serve(s.Server.Listener)
}

func New(component util.Components, cfg *config.SecretManagerConfig) server.Server {
	s := grpc.NewServer()
	sc, _ := secretstorage.NewClient(component)
	svr := &SecretManager{
		Server:        s,
		StorageClient: sc,
		Config:        cfg,
	}

	svr.SetEnvFromConfig()

	// pbServer := &grpc.UserServiceServer{
	// 	StorageClient: svr.StorageClient,
	// }
	// pb.RegisterUserServiceServer(s.Server, pbServer)

	return svr
}

func (c *SecretManager) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("SECRET_MANAGER_PORT", strconv.Itoa(c.Config.Service.Port))
	env.Set("SECRET_MANAGER_STORAGE_TYPE", c.Config.Storage.Type)
	env.Set("SECRET_MANAGER_STORAGE_HOST", c.Config.Storage.Host)
	env.Set("SECRET_MANAGER_STORAGE_USERNAME", c.Config.Storage.Username)
	env.Set("SECRET_MANAGER_STORAGE_PASSWORD", c.Config.Storage.Password)
	env.Set("SECRET_MANAGER_STORAGE_DATABASE", c.Config.Storage.Database)
	env.Set("SECRET_MANAGER_STORAGE_PORT", strconv.Itoa(c.Config.Storage.Port))
}
