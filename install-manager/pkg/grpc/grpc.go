package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	"google.golang.org/grpc"
)

type InstallManagerServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewServer() *InstallManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("INSTALL_MANAGER_PORT"))

	s := &InstallManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener: l,
	}

	return s
}

func (s *InstallManagerServer) Close() error {
	s.Server.Stop()
	return s.Listener.Close()
}

func (s *InstallManagerServer) Run() {
	s.Server.Serve(s.Listener)
}
