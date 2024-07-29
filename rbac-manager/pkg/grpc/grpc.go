package grpc

import (
	"context"
	"log"
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"google.golang.org/grpc"
)

type RbacGrpcServer struct {
	Server   *grpc.Server
	PbUser   pb.UnimplementedUserServer
	PbRole   pb.UnimplementedRoleServer
	Listener net.Listener
}

func (s *RbacGrpcServer) PutRole(ctx context.Context, in *pb.RoleRequest) (*pb.RoleReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.RoleReply{Message: "Hello " + in.GetName()}, nil
}

func NewRbacGrpcServer() *RbacGrpcServer {
	l, _ := net.Listen("tcp", ":"+env.Get("RBAC_MANAGER_PORT"))

	s := &RbacGrpcServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener: l,
	}
	pb.RegisterRoleServer(s.Server, s.PbRole)
	pb.RegisterUserServer(s.Server, s.PbUser)

	return s
}
