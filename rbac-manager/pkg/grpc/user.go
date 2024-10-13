package grpc

import (
	"context"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type UserServiceServer struct {
	pb.UnimplementedLoginServiceServer
	StorageClient storage.StorageClient
}

func (svr *UserServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	inputUsername := in.User.GetUsername()
	inputPassword := util.Hash(in.User.GetPassword())

	r := svr.StorageClient.DB().QueryRow(`SELECT * FROM users WHERE id = '` + inputUsername + `' and password_hash = '` + inputPassword + `'`)

	if r.Err() != nil {
		return &pb.LoginReply{Success: false}, r.Err()
	}

	fmt.Println("TEST USER " + inputUsername + " WITH PASSWORD " + in.User.Password + " LOGIN SUCCESS!")
	return &pb.LoginReply{Success: true}, nil
}

// func (s *UserServiceServer) ListProjectUsers(ctx context.Context, in *pb.ListProjectUsersRequest) (*pb.ListProjectUsersReply, error) {
// 	return &pb.ListProjectUsersReply{
// 		UserId: &pb.ProjectUser{
// 			UserId:   1,
// 			UserName: "test",
// 		},
// 	}, nil
// }

// func (s *UserServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.Success, error) {
// 	return &pb.Success{
// 		Succeed: true,
// 	}, nil
// }

// func (s *UserServiceServer) PutUser(ctx context.Context, in *pb.PutUserRequest) (*pb.Success, error) {
// 	_, err := s.StorageClient.DB().Exec(
// 		`INSERT INTO users (email, name, password_hash, project_id) VALUES ($1, $2, $3, $4)`,
// 		in.Email,
// 		in.Name,
// 		util.Hash(in.Password),
// 		in.ProjectId,
// 	)
// 	if err != nil {
// 		return &pb.Success{
// 			Succeed: false,
// 		}, err
// 	}

// 	return &pb.Success{
// 		Succeed: true,
// 	}, nil
// }

// func (s *UserServiceServer) PutUserInfo(ctx context.Context, in *pb.PutUserRequest) (*pb.Success, error) {
// 	return &pb.Success{
// 		Succeed: true,
// 	}, nil
// }
