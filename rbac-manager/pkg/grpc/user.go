package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	StorageClient storage.StorageClient
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
	_, err := s.StorageClient.DB().Exec(
		`INSERT INTO users (email, name, password_hash, project_id) VALUES ($1, $2, $3, $4)`,
		in.Email,
		in.Name,
		util.Hash(in.Password),
		in.ProjectId,
	)
	fmt.Println("ERROR:", err)
	if err != nil {
		return &pb.Success{
			Succeed: false,
		}, err
	}

	return &pb.Success{
		Succeed: true,
	}, nil
}

func (s *UserServiceServer) PutUserInfo(ctx context.Context, in *pb.PutUserRequest) (*pb.Success, error) {
	return &pb.Success{
		Succeed: true,
	}, nil
}
