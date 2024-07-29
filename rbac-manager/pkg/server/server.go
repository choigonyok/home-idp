package server

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
	rbacstorage "github.com/choigonyok/home-idp/rbac-manager/pkg/storage"
)

type RbacServer struct {
	Grpc          *grpc.GrpcServer
	StorageClient storage.StorageClient
	MailClient    *mail.SmtpClient
}

func (s *RbacServer) Close() error {
	if err := s.Grpc.Listener.Close(); err != nil {
		return err
	}
	if err := s.StorageClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *RbacServer) Run() {
	s.Grpc.Server.Serve(s.Grpc.Listener)
}

func New(component util.Components) server.Server {
	sc, _ := rbacstorage.NewClient(component)
	svr := &RbacServer{
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

func (s *RbacServer) TestSendEmail() error {
	return s.MailClient.SendMail([]string{"achoistic98@naver.com"})
}
