package grpc

import (
	"google.golang.org/grpc"
)

type InstallManagerGrpcClient struct {
	Client *grpc.ClientConn
}

func (gc *InstallManagerGrpcClient) Close() {
	return
}

func (gc *InstallManagerGrpcClient) GetConnection() *grpc.ClientConn {
	return gc.Client
}

func (gc *InstallManagerGrpcClient) Set(i interface{}) {
	gc.Client = parseGrpcClientFromInterface(i)
}

func parseGrpcClientFromInterface(i interface{}) *grpc.ClientConn {
	client := i.(*grpc.ClientConn)
	return client
}
