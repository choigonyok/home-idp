package grpc

import (
	"net"

	"github.com/choigonyok/home-idp/pkg/env"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"google.golang.org/grpc"
)

type RbacManagerServer struct {
	// PbRole   pb.UnimplementedUserServiceServer
	Server       *grpc.Server
	PbUserServer *UserServiceServer
	Listener     net.Listener
}

func NewServer() *RbacManagerServer {
	l, _ := net.Listen("tcp", ":"+env.Get("RBAC_MANAGER_PORT"))

	s := &RbacManagerServer{
		Server: grpc.NewServer(
			grpc.MaxConcurrentStreams(100),
			// grpc.ConnectionTimeout(time.Duration(30)),
		),
		Listener:     l,
		PbUserServer: &UserServiceServer{},
	}

	pb.RegisterUserServiceServer(s.Server, s.PbUserServer)

	return s
}
