package client

import (
	"google.golang.org/grpc"
)

type GatewayGrpcClient struct {
	Client *grpc.ClientConn
}

func (gc *GatewayGrpcClient) Close() {
	return
}

func (gc *GatewayGrpcClient) GetConnection() *grpc.ClientConn {
	return gc.Client
}
