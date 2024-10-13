package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/server"
	"google.golang.org/grpc"
)

type RbacManagerServer struct {
	Grpc     *grpc.Server
	Listener net.Listener
}

func NewServer(port int) *RbacManagerServer {
	return &RbacManagerServer{
		Grpc:     server.NewServer(),
		Listener: server.NewListener(port),
	}
}

func (svr *RbacManagerServer) Stop() error {
	svr.Grpc.GracefulStop()
	return svr.Listener.Close()
}

func (svr *RbacManagerServer) Run() {
	svr.Grpc.Serve(svr.Listener)
}
