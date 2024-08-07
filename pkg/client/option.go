package client

import (
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClientOption struct {
	f func(ClientSet)
}

func (opt *grpcClientOption) Apply(cli ClientSet) error {
	opt.f(cli)
	return nil
}

func WithGrpcClient(host string, port int) ClientOption {
	return useGrpcClient(host, port)
}

func useGrpcClient(host string, port int) ClientOption {
	return newGrpcClientOption(func(cli ClientSet) {
		grpcOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// grpc.WithTransportCredentials(tlsOpt),
		}
		conn, _ := grpc.NewClient(host+":"+strconv.Itoa(port), grpcOptions...)
		cli.ApplyGrpcClient(cli.New(conn))

	})
}

func newGrpcClientOption(f func(cli ClientSet)) *grpcClientOption {
	return &grpcClientOption{
		f: f,
	}
}
