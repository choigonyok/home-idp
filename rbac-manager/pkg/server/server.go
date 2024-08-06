package server

import (
	"log"

	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	rbacstorage "github.com/choigonyok/home-idp/rbac-manager/pkg/storage"
)

type RbacManager struct {
	Server        server.Server
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
	Config        config.Config
}

func (rbac *RbacManager) Close() error {
	if err := rbac.Server.Close(); err != nil {
		return err
	}

	if err := rbac.StorageClient.Close(); err != nil {
		return err
	}
	return nil
}

func (s *RbacManager) Run() {
	s.Server.Run()
}

func New(component util.Components, cfg config.Config) server.Server {
	s := grpc.NewServer()
	sc, _ := rbacstorage.NewClient(component)

	svr := &RbacManager{
		Server:        s,
		StorageClient: sc,
		Config:        cfg,
	}

	log.Printf("---Start installing mail server...")
	if config.Enabled(component, "mail") {
		mc := mail.NewClient(component)
		svr.MailClient = mc
	}

	pbServer := &grpc.UserServiceServer{
		StorageClient: svr.StorageClient,
	}
	pb.RegisterUserServiceServer(s.Server, pbServer)

	return svr
}

// func (s *RbacServer) TestSendEmail() error {
// 	return s.MailClient.SendMail([]string{"achoistic98@naver.com"})
// }
