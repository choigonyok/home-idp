package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/trace"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RbacServiceServer struct {
	pb.UnimplementedRbacServiceServer
	StorageClient storage.StorageClient
	TraceClient   *trace.TraceClient
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
	userName := in.GetUserName()
	r := svr.StorageClient.DB().QueryRow(`SELECT roles.id, roles.name, roles.create_time FROM roles JOIN users ON users.role_id = roles.id WHERE users.name = '` + userName + `'`)

	role := pb.Role{}
	r.Scan(&role.Id, &role.Name, &role.CreateTime)

	return &pb.GetRoleReply{Role: &role}, nil
}

func (svr *RbacServiceServer) GetRoles(ctx context.Context, in *emptypb.Empty) (*pb.GetRolesReply, error) {
	// reqId := ""

	// md, ok := metadata.FromIncomingContext(ctx)

	// if ok {
	// 	fmt.Println("Received x-trace-id:", md.Get("x-trace-id"))
	// 	fmt.Println("Received x-request-time:", md.Get("x-request-time"))

	// 	if md.Get("x-trace-id") != nil {
	// 		reqId = md.Get("x-trace-id")[0]
	// 	}
	// } else {
	// 	fmt.Println("No x-request-time found")
	// }

	// times := md.Get("x-request-time")
	// times = append(times, time.Now().Format("2006-01-02T15:04:05.999Z-SERVER"))
	// trailer := metadata.Pairs(
	// 	"x-envoy-upstream-cluster", "rbac_manager_service_cluster",
	// 	"x-trace-id", reqId,
	// 	"x-request-time", strings.Join(times, ", "),
	// )

	// grpc.SetTrailer(ctx, trailer)

	r, err := svr.StorageClient.DB().Query(`SELECT id, name, create_time FROM roles ORDER BY create_time DESC`)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	roles := []*pb.Role{}
	for r.Next() {
		role := pb.Role{}
		r.Scan(&role.Id, &role.Name, &role.CreateTime)
		roles = append(roles, &role)
	}

	return &pb.GetRolesReply{
		Roles: roles,
	}, nil
}

func (svr *RbacServiceServer) GetPolicies(ctx context.Context, in *emptypb.Empty) (*pb.GetPoliciesReply, error) {
	r, err := svr.StorageClient.DB().Query(`SELECT id, name, policy FROM policies ORDER BY create_time ASC`)
	if err != nil {
		fmt.Println("TEST GETPOLICIES FROM MAPPING QUERY ERR:", err)
		return nil, err
	}
	defer r.Close()

	ps := []*pb.Policy{}
	for r.Next() {
		p := pb.Policy{}
		r.Scan(&p.Id, &p.Name, &p.Json)
		ps = append(ps, &p)
	}

	return &pb.GetPoliciesReply{
		Policies: ps,
	}, nil
}

func (svr *RbacServiceServer) GetPolicy(ctx context.Context, in *pb.GetPolicyRequest) (*pb.GetPolicyReply, error) {
	r, err := svr.StorageClient.DB().Query(`SELECT policy_id FROM rolepolicymapping WHERE role_id = '` + in.RoleId + `'`)
	if err != nil {
		fmt.Println("TEST GETPOLICY FROM MAPPING QUERY ERR:", err)
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

	row, err := svr.StorageClient.DB().Query(`SELECT id, name, policy FROM policies WHERE id IN ('` + strings.Join(ids, "', '") + `') ORDER BY create_time DESC`)
	if err != nil {
		fmt.Println("TEST GETPOLICIES QUERY ERR:", err)
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&p.Id, &p.Name, &p.Json)
		ps = append(ps, &p)
	}

	return &pb.GetPolicyReply{
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

func (svr *RbacServiceServer) GetDockerfiles(ctx context.Context, in *pb.GetDockerfilesRequest) (*pb.GetDockerfilesReply, error) {
	userName := in.GetUserName()
	r, err := svr.StorageClient.DB().Query(`SELECT dockerfiles.id, dockerfiles.image_name, dockerfiles.image_version, dockerfiles.repository, dockerfiles.creator_id, dockerfiles.content, dockerfiles.trace_id FROM dockerfiles JOIN users ON dockerfiles.creator_id = users.id WHERE users.name = '` + userName + `'`)
	if err != nil {
		fmt.Println("ERR GETTING DOCKERFILES QUERY :", err)
		return nil, err
	}
	defer r.Close()

	ds := []*pb.Dockerfile{}

	for r.Next() {
		d := pb.Dockerfile{}
		r.Scan(&d.Id, &d.ImageName, &d.ImageVersion, &d.Repository, &d.CreatorId, &d.Content, &d.TraceId)
		ds = append(ds, &d)
	}

	return &pb.GetDockerfilesReply{
		Dockerfiles: ds,
	}, nil
}

func (svr *RbacServiceServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	pid := ""
	row := svr.StorageClient.DB().QueryRow(`SELECT id FROM projects WHERE name = '` + in.ProjectName + `'`)
	row.Scan(&pid)

	r, err := svr.StorageClient.DB().Query(`SELECT users.id, users.name, users.create_time, users.role_id FROM users JOIN userprojectmapping ON users.id = userprojectmapping.user_id WHERE userprojectmapping.project_id = '` + pid + `'`)

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

func (svr *RbacServiceServer) PostPolicy(ctx context.Context, in *pb.PostPolicyRequest) (*emptypb.Empty, error) {
	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO policies (id, name, policy) VALUES ('` + uuid.NewString() + `', '` + in.Policy.GetName() + `', '` + in.Policy.GetJson() + `')`); err != nil {
		return nil, err
	}
	return nil, nil
}

func (svr *RbacServiceServer) PostProject(ctx context.Context, in *pb.PostProjectRequest) (*pb.PostProjectReply, error) {
	creatorId := in.GetCreatorId()
	projectName := in.GetProjectName()

	r := svr.StorageClient.DB().QueryRow(`INSERT INTO projects (id, name, creator_id) VALUES ('` + uuid.NewString() + `', '` + projectName + `', '` + creatorId + `') RETURNING id`)

	projectId := ""
	r.Scan(&projectId)

	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO userprojectmapping (user_id, project_id) VALUES ('` + creatorId + `', '` + projectId + `')`); err != nil {
		fmt.Println("ERR POSTING NEW PROJECT QUERY :", err)
		return nil, err
	}

	return nil, nil
}

func (svr *RbacServiceServer) PostRole(ctx context.Context, in *pb.PostRoleRequest) (*emptypb.Empty, error) {
	roleName := in.GetRoleName()

	id := uuid.NewString()

	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO roles (id, name) VALUES ('` + id + `', '` + roleName + `')`); err != nil {
		fmt.Println("ERR CREATING NEW ROLE :", err)
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

	fmt.Println("username:", userName)
	fmt.Println("roleid:", roleId)
	fmt.Println("project name:", projectName)

	id := ""
	rr := svr.StorageClient.DB().QueryRow(`SELECT id FROM users WHERE name = '` + userName + `'`)

	if rr.Scan(&id) == sql.ErrNoRows {
		id = uuid.NewString()
		if _, err := svr.StorageClient.DB().Exec(`INSERT INTO users  (id, name, role_id) VALUES ('` + id + `', '` + userName + `', '` + roleId + `')`); err != nil {
			return nil, err
		}
	}

	_, err := svr.StorageClient.DB().Exec(`INSERT INTO userprojectmapping (user_id, project_id) VALUES ('` + id + `', '` + projectId + `' )`)

	return nil, err
}

func (svr *RbacServiceServer) PostDockerfile(ctx context.Context, in *pb.PostDockerfileRequest) (*emptypb.Empty, error) {
	storeDockerfileSpan := svr.TraceClient.NewSpanFromIncomingContext(ctx)
	err := storeDockerfileSpan.Start(ctx)
	if err != nil {
		fmt.Println("SPAN START ERR:", err)
	}

	imgName := in.GetDockerfile().GetImageName()
	imgVersion := in.GetDockerfile().GetImageVersion()
	creatorId := in.GetDockerfile().GetCreatorId()
	repo := in.GetDockerfile().GetRepository()
	content := in.GetDockerfile().GetContent()
	traceId := in.GetDockerfile().GetTraceId()

	if _, err = svr.StorageClient.DB().Exec(`INSERT INTO dockerfiles (id, image_name, image_version, creator_id, repository, content, trace_id) VALUES ('` + uuid.NewString() + `', '` + imgName + `', '` + imgVersion + `', '` + creatorId + `', '` + repo + `', '` + content + `', '` + traceId + `')`); err != nil {
		fmt.Println(err)
	}

	err = storeDockerfileSpan.Stop()
	if err != nil {
		fmt.Println("TRACE STOP ERR:", err)
	}

	return &emptypb.Empty{}, err
}

func (svr *RbacServiceServer) PutUser(ctx context.Context, in *pb.PutUserRequest) (*pb.PutUserReply, error) {
	usr := in.GetUser()

	r := svr.StorageClient.DB().QueryRow(`SELECT id FROM roles WHERE name = '` + usr.GetRoleId() + `'`)
	roleId := ""
	r.Scan(&roleId)

	_, err := svr.StorageClient.DB().Exec(`UPDATE users SET role_id = ` + roleId + ` WHERE name = '` + usr.GetName() + `'`)

	return &pb.PutUserReply{}, err
}

func (svr *RbacServiceServer) GetTraceId(ctx context.Context, in *pb.GetTraceIdRequest) (*pb.GetTraceIdReply, error) {
	name := in.GetImageName()
	version := in.GetImageVersion()

	r := svr.StorageClient.DB().QueryRow(`SELECT trace_id, repository FROM dockerfiles WHERE image_name = '` + name + `' and image_version = '` + version + `'`)
	traceId, repository := "", ""
	r.Scan(&traceId, &repository)

	return &pb.GetTraceIdReply{
		TraceId:    traceId,
		Repository: repository,
	}, nil
}

func (svr *RbacServiceServer) GetTraceIdByDockerfileId(ctx context.Context, in *pb.GetTraceIdByDockerfileIdRequest) (*pb.GetTraceIdByDockerfileIdReply, error) {
	dId := in.GetDockerfileId()

	r := svr.StorageClient.DB().QueryRow(`SELECT trace_id FROM dockerfiles WHERE id = '` + dId + `'`)
	traceId := ""
	r.Scan(&traceId)

	return &pb.GetTraceIdByDockerfileIdReply{
		TraceId: traceId,
	}, nil
}
