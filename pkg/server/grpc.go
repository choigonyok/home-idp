package server

import (
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server interface{}

func NewServer() *grpc.Server {
	svr := grpc.NewServer(
		grpc.MaxConcurrentStreams(100),
		// grpc.ConnectionTimeout(time.Duration(30)),
	)
	return svr
}

func NewListener(port int) net.Listener {
	l, _ := net.Listen("tcp", ":"+strconv.Itoa(port))
	return l
}
