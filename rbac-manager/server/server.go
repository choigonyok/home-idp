package server

import (
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/grpc"
)

type RbacServer struct {
	Grpc *grpc.RbacGrpcServer
}

func (s *RbacServer) Stop() {
	s.Grpc.Server.Stop()
}

func IntegrateGrpcServerToServer(svr *server.Server) {
	svr.Server = &RbacServer{
		Grpc: grpc.NewRbacGrpcServer(),
	}
}

func (s *RbacServer) CloseListner() {
	s.Grpc.Listener.Close()
}

func (s *RbacServer) Serve() {
	s.Grpc.Server.Serve(s.Grpc.Listener)
}
