package grpc

import (
	gatewaygrpc "github.com/choigonyok/home-idp/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	rbacManagerPort = "5105"
	rbacManagerHost = "localhost"
)

type GrpcClient struct {
	Conn *grpc.ClientConn
}

func NewClient() *GrpcClient {
	grpc := &GrpcClient{
		Conn: gatewaygrpc.NewClient(rbacManagerHost, rbacManagerPort),
	}

	return grpc
}

func (g *GrpcClient) Close() error {
	return g.Conn.Close()
}
