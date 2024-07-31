package server

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/secret-manager/pkg/grpc"
	secretstorage "github.com/choigonyok/home-idp/secret-manager/pkg/storage"
)

type SecretServer struct {
	Grpc          *grpc.SecretManagerServer
	StorageClient storage.StorageClient
	MailClient    *mail.SmtpClient
}

func (s *SecretServer) Close() error {
	if err := s.Grpc.Listener.Close(); err != nil {
		return err
	}
	if err := s.StorageClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *SecretServer) Run() {
	s.Grpc.Server.Serve(s.Grpc.Listener)
}

func New(component util.Components) server.Server {
	sc, _ := secretstorage.NewClient(component)
	svr := &SecretServer{
		Grpc:          grpc.NewServer(),
		StorageClient: sc,
	}

	log.Printf("---Start installing mail server...")
	if config.Enabled(component, "mail") {
		mc := mail.NewClient(component)
		svr.MailClient = mc
	}

	return svr
}
