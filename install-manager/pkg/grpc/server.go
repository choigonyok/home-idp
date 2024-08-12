package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/server"
	"google.golang.org/grpc"
)

type InstallManagerServer struct {
	Grpc     *grpc.Server
	Listener net.Listener
}

func NewServer(port int) *InstallManagerServer {
	return &InstallManagerServer{
		Grpc:     server.NewServer(),
		Listener: server.NewListener(port),
	}
}

func (svr *InstallManagerServer) Run() {
	svr.Grpc.Serve(svr.Listener)
}

func (svr *InstallManagerServer) Stop() {
	svr.Grpc.GracefulStop()
	svr.Listener.Close()
}
