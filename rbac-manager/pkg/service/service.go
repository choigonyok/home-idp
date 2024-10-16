package service

import (
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/client"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type RbacManager struct {
	ClientSet *client.RbacManagerClientSet
	Server    *grpc.RbacManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *RbacManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &RbacManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *RbacManager) Stop() {
	if err := svc.Server.Stop(); err != nil {
		return
	}
	svc.ClientSet.StorageClient.Close()
}

func (svc *RbacManager) Start() {
	pbServer := &grpc.LoginServiceServer{
		StorageClient: svc.ClientSet.StorageClient,
	}

	rbacPbServer := &grpc.RbacServiceServer{
		StorageClient: svc.ClientSet.StorageClient,
	}

	pb.RegisterLoginServiceServer(svc.Server.Grpc, pbServer)
	pb.RegisterRbacServiceServer(svc.Server.Grpc, rbacPbServer)

	svc.ClientSet.StorageClient.CreateAdminUser(env.Get("HOME_IDP_ADMIN_GIT_USERNAME"))

	svc.Server.Run()
}
