package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/choigonyok/home-idp/pkg/storage"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (svr *RbacServiceServer) GetProjects(ctx context.Context, in *emptypb.Empty) (*pb.GetProjectsReply, error) {
	r, err := svr.StorageClient.DB().Query(`SELECT id, name, creator_id FROM projects`)
	if err != nil {
		fmt.Println("ERR GETTING PROJECT QUERY :", err)
		return nil, err
	}
	defer r.Close()

	projs := []*pb.Project{}

	for r.Next() {
		proj := pb.Project{}
		r.Scan(&proj.Id, &proj.Name, &proj.CreatorId)
		projs = append(projs, &proj)
	}

	return &pb.GetProjectsReply{
		Projects: projs,
	}, nil
}

func (svr *RbacServiceServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	pid := ""
	row := svr.StorageClient.DB().QueryRow(`SELECT id FROM projects WHERE name = '` + in.ProjectName + `'`)
	row.Scan(&pid)

	r, err := svr.StorageClient.DB().Query(`SELECT users.id, users.name, users.create_time, users.role_id FROM users JOIN userprojectmapping ON users.id = userprojectmapping.user_id WHERE userprojectmapping.project_id = ` + pid)
	if err != nil {
		fmt.Println("TEST GETUSERS QUERY ERR:", err)
		return nil, err
	}
	defer r.Close()

	usrs := []*pb.User{}

	for r.Next() {
		usr := pb.User{}
		r.Scan(&usr.Id, &usr.Name, &usr.CreateTime, &usr.RoleId)
		usrs = append(usrs, &usr)
	}

	return &pb.GetUsersReply{
		Users: usrs,
	}, nil
}

func (svr *RbacServiceServer) PostProject(ctx context.Context, in *pb.PostProjectRequest) (*pb.PostProjectReply, error) {
	creatorId := in.GetCreatorId()
	projectName := in.GetProjectName()

	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO projects (id, name, creator_id) VALUES ('` + uuid.New().String() + `', '` + projectName + `', '` + creatorId + `')`); err != nil {
		fmt.Println("ERR CREATING NEW PROJECT :", err)
		return nil, err
	}

	return nil, nil
}

func (svr *RbacServiceServer) PostUser(ctx context.Context, in *pb.PostUserRequest) (*emptypb.Empty, error) {
	usr := in.GetUser()
	userName := usr.GetName()
	roleId := usr.GetRoleId()
	projectName := in.GetProjectName()

	r := svr.StorageClient.DB().QueryRow(`SELECT id FROM projects WHERE name = '` + projectName + `'`)
	projectId := ""
	r.Scan(&projectId)

	id := ""
	rr := svr.StorageClient.DB().QueryRow(`SELECT id FROM users WHERE name = '` + userName + `'`)

	if rr.Err() == sql.ErrNoRows {
		id = uuid.NewString()
		if _, err := svr.StorageClient.DB().Exec(`INSERT INTO users  (id, name, role_id) VALUES ('` + id + `', '` + userName + `', '` + roleId + `')`); err != nil {
			return nil, err
		}
	} else {
		rr.Scan(&id)
	}

	_, err := svr.StorageClient.DB().Exec(`INSERT INTO userprojectmapping (user_id, project_id) VALUES (` + id + `, ` + projectId + ` )`)

	return nil, err
}

func (svr *RbacServiceServer) PutUser(ctx context.Context, in *pb.PutUserRequest) (*pb.PutUserReply, error) {
	usr := in.GetUser()

	r := svr.StorageClient.DB().QueryRow(`SELECT id FROM roles WHERE name = '` + usr.GetRoleId() + `'`)
	roleId := ""
	r.Scan(&roleId)

	_, err := svr.StorageClient.DB().Exec(`UPDATE users SET role_id = ` + roleId + ` WHERE name = '` + usr.GetName() + `'`)

	return &pb.PutUserReply{}, err
}
