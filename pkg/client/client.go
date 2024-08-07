package client

import "google.golang.org/grpc"

type ClientOption interface {
	Apply(ClientSet) error
}

type ClientSet interface {
	ApplyGrpcClient(conn GrpcClient)
	New(conn *grpc.ClientConn) GrpcClient
}
