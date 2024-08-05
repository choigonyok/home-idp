package grpc

import (
	gatewaygrpc "github.com/choigonyok/home-idp/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	rbacManagerPort    = "5105"
	rbacManagerHost    = "localhost"
	installManagerPort = "5107"
	installManagerHost = "localhost"
)

type GrpcClient struct {
	RbacConn    *grpc.ClientConn
	InstallConn *grpc.ClientConn
}

func NewClient() *GrpcClient {
	grpc := &GrpcClient{
		RbacConn:    gatewaygrpc.NewClient(rbacManagerHost, rbacManagerPort),
		InstallConn: gatewaygrpc.NewClient(installManagerHost, installManagerPort),
	}

	return grpc
}

func (g *GrpcClient) Close() error {
	if err := g.RbacConn.Close(); err != nil {
		return err
	}
	if err := g.InstallConn.Close(); err != nil {
		return err
	}
	return nil
}
