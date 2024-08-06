package server

import (
	"log"
	"strconv"

	globalconfig "github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/mail"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/config"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	rbacstorage "github.com/choigonyok/home-idp/rbac-manager/pkg/storage"
)

type RbacManager struct {
	Server        server.Server
	StorageClient storage.StorageClient
	MailClient    mail.MailClient
	Config        *config.RbacManagerConfig
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

func New(component util.Components, cfg *config.RbacManagerConfig) server.Server {
	s := grpc.NewServer()
	sc, _ := rbacstorage.NewClient(component)

	svr := &RbacManager{
		Server:        s,
		StorageClient: sc,
		Config:        cfg,
	}

	svr.SetEnvFromConfig()

	log.Printf("---Start installing mail server...")
	if globalconfig.Enabled(component, "mail") {
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

func (c *RbacManager) SetEnvFromConfig() {
	log.Printf("Start injecting appropriate environments variables...")
	env.Set("RBAC_MANAGER_PORT", strconv.Itoa(c.Config.Service.Port))
	env.Set("RBAC_MANAGER_STORAGE_TYPE", c.Config.Storage.Type)
	env.Set("RBAC_MANAGER_STORAGE_HOST", c.Config.Storage.Host)
	env.Set("RBAC_MANAGER_STORAGE_USERNAME", c.Config.Storage.Username)
	env.Set("RBAC_MANAGER_STORAGE_PASSWORD", c.Config.Storage.Password)
	env.Set("RBAC_MANAGER_STORAGE_DATABASE", c.Config.Storage.Database)
	env.Set("RBAC_MANAGER_STORAGE_PORT", strconv.Itoa(c.Config.Storage.Port))
	if c.Config.Smtp.Enabled == true {
		env.Set("RBAC_MANAGER_SMTP_HOST", c.Config.Smtp.Config.Host)
		env.Set("RBAC_MANAGER_SMTP_PORT", c.Config.Smtp.Config.Port)
		env.Set("RBAC_MANAGER_SMTP_USER", c.Config.Smtp.Config.User)
		env.Set("RBAC_MANAGER_SMTP_PASSWORD", c.Config.Smtp.Config.Password)
		env.Set("RBAC_MANAGER_SMTP_DOMAIN", c.Config.Smtp.Config.Domain)
		env.Set("RBAC_MANAGER_SMTP_ENABLED", strconv.FormatBool(c.Config.Smtp.Enabled))
	}
}
