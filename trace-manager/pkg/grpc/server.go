package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/server"
	"google.golang.org/grpc"
)

type TraceManagerServer struct {
	Grpc     *grpc.Server
	Listener net.Listener
}

func NewServer(port int) *TraceManagerServer {
	return &TraceManagerServer{
		Grpc:     server.NewServer(),
		Listener: server.NewListener(port),
	}
}

func (svr *TraceManagerServer) Stop() error {
	svr.Grpc.GracefulStop()
	return svr.Listener.Close()
}

func (svr *TraceManagerServer) Run() {
	svr.Grpc.Serve(svr.Listener)
}
