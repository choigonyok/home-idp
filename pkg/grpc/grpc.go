package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewListener(port string) net.Listener {
	l, _ := net.Listen("tcp", ":"+port)
	return l
}

func NewClientConn(port string) *grpc.ClientConn {
	// tlsOpt, _ := credentials.NewClientTLSFromFile()
	grpcOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithTransportCredentials(tlsOpt),
	}
	conn, _ := grpc.NewClient("localhost:5103", grpcOptions...)
	return conn
}