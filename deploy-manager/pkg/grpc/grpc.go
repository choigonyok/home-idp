package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	"google.golang.org/grpc"
)

type DeployManagerServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewServer() *DeployManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("DEPLOY_MANAGER_PORT"))

	s := &DeployManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener: l,
	}

	return s
}

func (s *DeployManagerServer) Close() error {
	s.Server.Stop()
	return s.Listener.Close()
}

func (s *DeployManagerServer) Run() {
	s.Server.Serve(s.Listener)
}
