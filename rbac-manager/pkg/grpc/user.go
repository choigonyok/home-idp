package grpc

import (
	"context"
	"log"

	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoReply, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.GetUserInfoReply{
		Name:        "test",
		Email:       "test@test.com",
		Password:    "Passwordtest",
		ProjectName: []string{"P1", "P2"},
		RoleName:    []string{"R1", "R2"},
	}, nil
}

func (s *UserServiceServer) ListProjectUsers(ctx context.Context, in *pb.ListProjectUsersRequest) (*pb.ListProjectUsersReply, error) {
	return &pb.ListProjectUsersReply{
		UserId: &pb.ProjectUser{
			UserId:   1,
			UserName: "test",
		},
	}, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.Success, error) {
	return &pb.Success{
		Succeed: true,
	}, nil
}

func (s *UserServiceServer) PutUser(ctx context.Context, in *pb.PutUserRequest) (*pb.Success, error) {
	return &pb.Success{
		Succeed: true,
	}, nil
}

func (s *UserServiceServer) PutUserInfo(ctx context.Context, in *pb.PutUserRequest) (*pb.Success, error) {
	return &pb.Success{
		Succeed: true,
	}, nil
}
