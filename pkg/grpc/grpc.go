package grpc

import (
	"net"

	pb "github.com/choigonyok/home-idp/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer struct {
	PbServer   pb.UnimplementedGreeterServer
	GrpcServer *grpc.Server
}

func NewListener(port string) net.Listener {
	l, _ := net.Listen("tcp", ":"+port)
	return l
}

func NewClientConn(dst, port string) *grpc.ClientConn {
	// tlsOpt, _ := credentials.NewClientTLSFromFile()
	grpcOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithTransportCredentials(tlsOpt),
	}
	conn, _ := grpc.NewClient(dst+":"+port, grpcOptions...)
	return conn
}

func NewServer() *GrpcServer {
	svr := &GrpcServer{
		GrpcServer: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
	}

	return svr
}
