package grpc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/util"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type LoginServiceServer struct {
	pb.UnimplementedLoginServiceServer
	StorageClient storage.StorageClient
}

func (svr *LoginServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	inputUsername := in.User.GetUsername()

	r := svr.StorageClient.DB().QueryRow(`SELECT password_hash FROM users WHERE name = '` + inputUsername + `'`)

	// if there is error in query
	if r.Err() != nil {
		return &pb.LoginReply{Success: false}, r.Err()
	}

	pw := ""
	// if there is no match username and password
	if r.Scan(&pw) == sql.ErrNoRows {
		fmt.Println("TEST USER " + inputUsername + " NOT EXISTS!")
		return &pb.LoginReply{Success: false}, nil
	}

	if pw == util.Hash(in.User.GetPassword()) {
		fmt.Println("TEST USER " + inputUsername + " WITH PASSWORD " + in.User.Password + " LOGIN SUCCESS!")
		return &pb.LoginReply{Success: true}, nil
	} else {
		fmt.Println("TEST USER " + inputUsername + " WITH PASSWORD " + in.User.Password + " LOGIN FAILED!")
		return &pb.LoginReply{Success: true}, nil
	}
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
