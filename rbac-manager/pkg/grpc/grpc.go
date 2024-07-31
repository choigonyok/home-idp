package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	"google.golang.org/grpc"
)

type RbacManagerServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewServer() *RbacManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("RBAC_MANAGER_PORT"))

	s := &RbacManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener: l,
	}

	return s
}

func (s *RbacManagerServer) Close() error {
	s.Server.Stop()
	return s.Listener.Close()
}

func (s *RbacManagerServer) Run() {
	s.Server.Serve(s.Listener)
}
