package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/server"
	"google.golang.org/grpc"
)

type DeployManagerServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewServer(port int) *DeployManagerServer {
	return &DeployManagerServer{
		Server:   server.NewServer(),
		Listener: server.NewListener(port),
	}
}

func (s *DeployManagerServer) Close() error {
	s.Server.Stop()
	return s.Listener.Close()
}

func (s *DeployManagerServer) Run() {
	s.Server.Serve(s.Listener)
}
