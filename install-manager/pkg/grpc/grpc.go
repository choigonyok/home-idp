package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
	"google.golang.org/grpc"
)

type InstallManagerServer struct {
	Server     *grpc.Server
	Listener   net.Listener
	HelmClient *helm.HelmClient
}

func NewServer() *InstallManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("INSTALL_MANAGER_PORT"))
	h := helm.New()

	s := &InstallManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener:   l,
		HelmClient: h,
	}

	h.AddRepository("bitnami", "https://charts.bitnami.com/bitnami", true)
	h.AddRepository("argo", "https://argoproj.github.io/argo-helm", true)

	return s
}

func (s *InstallManagerServer) Close() error {
	s.Server.Stop()
	return s.Listener.Close()
}

func (s *InstallManagerServer) Run() {
	s.Server.Serve(s.Listener)
}
