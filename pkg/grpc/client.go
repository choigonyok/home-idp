package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer interface {
	Stop()
	CloseListner()
	Serve()
}

func NewClient(dst, port string) *grpc.ClientConn {
	// tlsOpt, _ := credentials.NewClientTLSFromFile()
	grpcOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithTransportCredentials(tlsOpt),
	}
	conn, _ := grpc.NewClient(dst+":"+port, grpcOptions...)
	return conn
}
