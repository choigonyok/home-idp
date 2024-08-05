package grpc

import (
	"context"
	"time"

	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

func (g *GrpcClient) PutUser(email, name, password string, projectId int32) (*pb.Success, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()
	r, err := pb.NewUserServiceClient(g.RbacConn).PutUser(ctx, &pb.PutUserRequest{
		Email:     email,
		Name:      name,
		Password:  password,
		ProjectId: projectId,
	})

	return r, err
}
