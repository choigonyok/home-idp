package grpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/choigonyok/home-idp/pkg/storage"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
)

type RbacServiceServer struct {
	pb.UnimplementedRbacServiceServer
	StorageClient storage.StorageClient
}

func (svr *RbacServiceServer) Check(ctx context.Context, in *pb.RbacRequest) (*pb.RbacReply, error) {
	user := in.GetUsername()
	action := in.GetAction().String()
	target := in.GetTarget()

	fmt.Println(user, action, target)

	r := svr.StorageClient.DB().QueryRow(`SELECT role_id FROM users WHERE name='` + user + `'`)
	roleId := ""
	if err := r.Scan(&roleId); err != nil {
		fmt.Println("TEST SCAN ERR:", err)
		return &pb.RbacReply{Result: pb.Result_ERROR}, err
	}

	rs, _ := svr.StorageClient.DB().Query(`SELECT policy_id FROM rolepolicymapping WHERE role_id=` + roleId)

	pids := []int{}
	pid := 0
	for rs.Next() {
		rs.Scan(&pid)
		pids = append(pids, pid)
	}

	ps := []string{}
	p := ""
	for _, v := range pids {
		row := svr.StorageClient.DB().QueryRow(`SELECT policy FROM policies WHERE id=` + strconv.Itoa(v))
		row.Scan(&p)
		ps = append(ps, p)
	}

	fmt.Println("TEST EVERY POLICIES FOR USER "+user+" : ", ps)

	return &pb.RbacReply{Result: pb.Result_ALLOW}, nil
}
