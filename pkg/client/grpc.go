package client

import (
	"google.golang.org/grpc"
)

type GrpcClient interface {
	Close()
	GetConnection() *grpc.ClientConn
}
