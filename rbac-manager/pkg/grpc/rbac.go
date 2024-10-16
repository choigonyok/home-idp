package grpc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
	defer rs.Close()

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

func (svr *RbacServiceServer) GetRole(ctx context.Context, in *pb.GetRoleRequest) (*pb.GetRoleReply, error) {
	r := svr.StorageClient.DB().QueryRow(`SELECT users.role_id, roles.name AS role_name FROM users JOIN roles ON users.role_id = roles.id WHERE users.id = ` + in.GetUserId())

	role := pb.Role{}
	r.Scan(&role.Id, &role.Name)

	return &pb.GetRoleReply{Role: &role}, nil
}

func (svr *RbacServiceServer) GetRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	r, err := svr.StorageClient.DB().Query(`SELECT id, name FROM roles ORDER BY create_time DESC`)
	if err != nil {
		fmt.Println("TEST GETROLES QUERY ERR:", err)
		return nil, err
	}
	defer r.Close()

	roles := []*pb.Role{}
	for r.Next() {
		role := pb.Role{}
		r.Scan(&role.Id, &role.Name)
		roles = append(roles, &role)
	}

	return &pb.GetRolesReply{Roles: roles}, nil
}

func (svr *RbacServiceServer) GetPolicies(ctx context.Context, in *pb.GetPoliciesRequest) (*pb.GetPoliciesReply, error) {
	r, err := svr.StorageClient.DB().Query(`SELECT policy_id FROM rolepolicymapping WHERE role_id = ` + in.RoleId)
	if err != nil {
		fmt.Println("TEST GETPOLICIES FROM MAPPING QUERY ERR:", err)
		return nil, err
	}
	defer r.Close()

	ids := []string{}
	id := ""

	for r.Next() {
		r.Scan(&id)
		ids = append(ids, id)
	}

	p := pb.Policy{}
	ps := []*pb.Policy{}

	row, err := svr.StorageClient.DB().Query(`SELECT id, name, policy FROM policies WHERE id IN (` + strings.Join(ids, ", ") + `) ORDER BY create_time DESC`)
	if err != nil {
		fmt.Println("TEST GETPOLICIES QUERY ERR:", err)
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&p.Id, &p.Name, &p.Json)
		ps = append(ps, &p)
	}

	return &pb.GetPoliciesReply{
		Policies: ps,
	}, nil
}
