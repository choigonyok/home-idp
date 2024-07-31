package grpc

import (
	"context"
	"log"
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	pb "github.com/choigonyok/home-idp/secret-manager/pkg/proto"
	"google.golang.org/grpc"
)

type SecretManagerServer struct {
	Server    *grpc.Server
	PbGreeter pb.UnimplementedGreeterServer
	Listener  net.Listener
}

func (s *SecretManagerServer) GreeterServer(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewServer() *SecretManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("SECRET_MANAGER_PORT"))

	s := &SecretManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener: l,
	}
	pb.RegisterGreeterServer(s.Server, s.PbGreeter)

	return s
}
