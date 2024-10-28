package service

import (
	pkgclient "github.com/choigonyok/home-idp/pkg/client"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/trace-manager/pkg/client"
	"github.com/choigonyok/home-idp/trace-manager/pkg/grpc"
	pb "github.com/choigonyok/home-idp/trace-manager/pkg/proto"
)

type TraceManager struct {
	ClientSet *client.TraceManagerClientSet
	Server    *grpc.TraceManagerServer
}

func New(port int, opts ...pkgclient.ClientOption) *TraceManager {
	cs := client.EmptyClientSet()
	for _, opt := range opts {
		opt.Apply(cs)
	}

	return &TraceManager{
		Server:    grpc.NewServer(port),
		ClientSet: cs,
	}
}

func (svc *TraceManager) Stop() {
	if err := svc.Server.Stop(); err != nil {
		return
	}
	svc.ClientSet.StorageClient.Close()
}

func (svc *TraceManager) Start() {
	tracePbServer := &grpc.TraceServiceServer{
		StorageClient: svc.ClientSet.StorageClient,
	}

	pb.RegisterTraceServiceServer(svc.Server.Grpc, tracePbServer)

	svc.ClientSet.StorageClient.CreateAdminUser(env.Get("HOME_IDP_ADMIN_GIT_USERNAME"))

	svc.Server.Run()
}
