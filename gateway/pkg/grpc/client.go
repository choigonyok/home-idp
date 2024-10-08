package grpc

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

func (gc *GatewayGrpcClient) Set(i interface{}) {
	gc.Client = parseGrpcClientFromInterface(i)
}

func parseGrpcClientFromInterface(i interface{}) *grpc.ClientConn {
	client := i.(*grpc.ClientConn)
	return client
}
