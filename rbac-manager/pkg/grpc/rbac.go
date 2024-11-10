package grpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/choigonyok/home-idp/pkg/storage"
	"github.com/choigonyok/home-idp/pkg/trace"
	"github.com/choigonyok/home-idp/rbac-manager/pkg/git"
	pb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RbacServiceServer struct {
	pb.UnimplementedRbacServiceServer
	StorageClient storage.StorageClient
	TraceClient   *trace.TraceClient
	GitClient     *git.RbacGitClient
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

func (svr *RbacServiceServer) GetRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.GetRolesReply{
			Error: &pb.Error{
				Message:    "Your signing up request still waiting administrator's approval",
				StatusCode: 403,
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "roles", "GET") {
		return &pb.GetRolesReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

	r, err := svr.StorageClient.DB().Query(`SELECT id, name, create_time FROM roles ORDER BY create_time DESC`)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	rps := []*pb.RolePolicy{}

	for r.Next() {
		role := pb.Role{}
		r.Scan(&role.Id, &role.Name, &role.CreateTime)

		r2, err := svr.StorageClient.DB().Query(`SELECT policies.id, policies.name FROM policies JOIN rolepolicymapping ON policies.id = rolepolicymapping.policy_id WHERE rolepolicymapping.role_id = '` + role.Id + `'`)
		if err != nil {
			return nil, err
		}
		defer r2.Close()

		ps := []*pb.Policy{}

		tmp := pb.RolePolicy{
			Role: &role,
		}

		for r2.Next() {
			id, name := "", ""
			r2.Scan(&id, &name)
			p := pb.Policy{
				Id:   id,
				Name: name,
			}
			ps = append(ps, &p)
		}

		tmp.Policies = ps
		rps = append(rps, &tmp)
	}

	return &pb.GetRolesReply{
		RolePolicies: rps,
	}, nil
}

func (svr *RbacServiceServer) GetPolicies(ctx context.Context, in *pb.GetPoliciesRequest) (*pb.GetPoliciesReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.GetPoliciesReply{
			Error: &pb.Error{
				Message:    "Your signing up request still waiting administrator's approval",
				StatusCode: 403,
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "policies", "GET") {
		return &pb.GetPoliciesReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

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

func (svr *RbacServiceServer) GetProjects(ctx context.Context, in *pb.GetProjectsRequest) (*pb.GetProjectsReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.GetProjectsReply{
			Error: &pb.Error{
				Message:    "Your signing up request still waiting administrator's approval",
				StatusCode: 403,
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "projects", "GET") {
		return &pb.GetProjectsReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

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

func (svr *RbacServiceServer) GetUsersInProject(ctx context.Context, in *pb.GetUsersInProjectRequest) (*pb.GetUsersInProjectReply, error) {
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

	return &pb.GetUsersInProjectReply{
		Users: usrs,
	}, nil
}

func (svr *RbacServiceServer) CheckPolicy(uid float64, target, action string) bool {
	userId := strconv.FormatFloat(uid, 'e', -1, 64)
	r, err := svr.StorageClient.DB().Query(`SELECT policy FROM policies LEFT JOIN rolepolicymapping ON rolepolicymapping.policy_id = policies.id LEFT JOIN roles ON roles.id = rolepolicymapping.role_id JOIN users ON users.role_id = roles.id WHERE users.id = ` + userId)
	if err != nil {
		return false
	}
	defer r.Close()

	policies := []map[string]interface{}{}

	for r.Next() {
		policy := make(map[string]interface{})
		j := ""
		r.Scan(&j)
		json.Unmarshal([]byte(j), &policy)
		fmt.Println(policy)
		policies = append(policies, policy)
	}

	for _, p := range policies {
		if p["policy"].(map[string]interface{})["effect"].(string) == "Allow" {
			if p["policy"].(map[string]interface{})["target"].(string) == "*" || p["policy"].(map[string]interface{})["target"].(string) == target {
				if p["policy"].(map[string]interface{})["action"].(string) == "*" || p["policy"].(map[string]interface{})["action"].(string) == action {
					return true
				}
			}
		}
	}

	return false
}

func (svr *RbacServiceServer) CheckApplicant(uid float64) bool {
	userId := strconv.FormatFloat(uid, 'e', -1, 64)
	roleName := ""

	r := svr.StorageClient.DB().QueryRow(`SELECT roles.name FROM roles JOIN users ON users.role_id = roles.id WHERE users.id = ` + userId)

	r.Scan(&roleName)

	if roleName == "applicant" {
		return true
	}
	return false
}

func (svr *RbacServiceServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.GetUsersReply{
			Users: nil,
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "Your signing up request still waiting administrator's approval",
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "users", "GET") {
		return &pb.GetUsersReply{
			Users: nil,
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

	r, err := svr.StorageClient.DB().Query(`SELECT users.id AS user_id, users.name AS username, roles.id AS role_id, roles.name AS role_name, users.create_time, projects.name AS project_name, projects.id AS project_id FROM users JOIN roles ON users.role_id = roles.id LEFT JOIN userprojectmapping ON userprojectmapping.user_id = users.id LEFT JOIN projects ON userprojectmapping.project_id = projects.id`)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	users := []*pb.User{}

	for r.Next() {
		user := pb.User{}
		r.Scan(&user.Id, &user.Name, &user.RoleId, &user.RoleName, &user.CreateTime, &user.ProjectName, &user.ProjectId)
		users = append(users, &user)
	}

	return &pb.GetUsersReply{
		Users: users,
	}, nil
}

func (svr *RbacServiceServer) PostPolicy(ctx context.Context, in *pb.PostPolicyRequest) (*emptypb.Empty, error) {
	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO policies (id, name, policy) VALUES ('` + uuid.NewString() + `', '` + in.Policy.GetName() + `', '` + in.Policy.GetJson() + `')`); err != nil {
		return nil, err
	}
	return nil, nil
}

func (svr *RbacServiceServer) PostProject(ctx context.Context, in *pb.PostProjectRequest) (*pb.PostProjectReply, error) {
	projectName := in.GetProjectName()
	uid := strconv.FormatFloat(in.GetCreatorId(), 'e', -1, 64)

	r := svr.StorageClient.DB().QueryRow(`INSERT INTO projects (id, name, creator_id) VALUES ('` + uuid.NewString() + `', '` + projectName + `', ` + uid + `) RETURNING id`)

	projectId := ""
	r.Scan(&projectId)

	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO userprojectmapping (user_id, project_id) VALUES (` + uid + `, '` + projectId + `')`); err != nil {
		fmt.Println("ERR POSTING NEW PROJECT QUERY :", err)
		return nil, err
	}

	return nil, nil
}

func (svr *RbacServiceServer) UpdateUserRole(ctx context.Context, in *pb.UpdateUserRoleRequest) (*pb.UpdateUserRoleReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.UpdateUserRoleReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "Your signing up request still waiting administrator's approval",
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "users", "GET") {
		return &pb.UpdateUserRoleReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

	userId := in.GetUser().GetId()
	username := in.GetUser().GetName()
	roleId := in.GetRole().GetId()

	if _, err := svr.StorageClient.DB().Exec(`UPDATE users SET name = '` + username + `', role_id = '` + roleId + `' WHERE id = ` + strconv.FormatFloat(userId, 'e', -1, 64)); err != nil {
		fmt.Println("ERR CREATING NEW ROLE :", err)
		return &pb.UpdateUserRoleReply{
			Error: &pb.Error{
				StatusCode: 500,
				Message:    err.Error(),
			},
		}, err
	}

	return &pb.UpdateUserRoleReply{
		Error: &pb.Error{
			StatusCode: 200,
			Message:    "",
		},
	}, nil
}

func (svr *RbacServiceServer) PostRole(ctx context.Context, in *pb.PostRoleRequest) (*pb.PostRoleReply, error) {
	if svr.CheckApplicant(in.GetUid()) {
		return &pb.PostRoleReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "Your signing up request still waiting administrator's approval",
			},
		}, nil
	}

	if !svr.CheckPolicy(in.GetUid(), "users", "GET") {
		return &pb.PostRoleReply{
			Error: &pb.Error{
				StatusCode: 403,
				Message:    "You do not have permission to access this page. Please contact to your administrator.",
			},
		}, nil
	}

	roleName := in.GetRole().GetName()
	id := uuid.NewString()

	if _, err := svr.StorageClient.DB().Exec(`INSERT INTO roles (id, name) VALUES ('` + id + `', '` + roleName + `')`); err != nil {
		fmt.Println("ERR CREATING NEW ROLE :", err)
		return nil, err
	}

	for _, p := range in.GetPolicies() {
		if _, err := svr.StorageClient.DB().Exec(`INSERT INTO rolepolicymapping (role_id, policy_id) VALUES ('` + id + `', '` + p.GetId() + `')`); err != nil {
			fmt.Println("ERR CREATING NEW ROLE POLICY MAPPING :", err)
			return nil, err
		}
	}

	return nil, nil
}

func (svr *RbacServiceServer) PostUser(ctx context.Context, in *pb.PostUserRequest) (*emptypb.Empty, error) {
	uid := strconv.FormatFloat(in.GetUser().GetId(), 'e', -1, 64)
	username := in.GetUser().GetName()
	r := svr.StorageClient.DB().QueryRow(`SELECT id FROM roles WHERE name = 'applicant'`)
	roleId := ""
	r.Scan(&roleId)

	_, err := svr.StorageClient.DB().Exec(`INSERT INTO users  (id, name, role_id) VALUES (` + uid + `, '` + username + `', '` + roleId + `')`)

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
	repo := in.GetDockerfile().GetRepository()
	content := in.GetDockerfile().GetContent()
	traceId := in.GetDockerfile().GetTraceId()
	uid := strconv.FormatFloat(in.Dockerfile.GetCreatorId(), 'e', -1, 64)

	if _, err = svr.StorageClient.DB().Exec(`INSERT INTO dockerfiles (id, image_name, image_version, creator_id, repository, content, trace_id) VALUES ('` + uuid.NewString() + `', '` + imgName + `', '` + imgVersion + `', ` + uid + `, '` + repo + `', '` + content + `', '` + traceId + `')`); err != nil {
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

func (svr *RbacServiceServer) GetPolicyJson(ctx context.Context, in *pb.GetPolicyJsonRequest) (*pb.GetPolicyJsonReply, error) {
	pid := in.GetPolicyId()

	r := svr.StorageClient.DB().QueryRow(`SELECT id, name, policy FROM policies WHERE id = '` + pid + `'`)
	p := pb.Policy{}
	r.Scan(&p.Id, &p.Name, &p.Json)

	return &pb.GetPolicyJsonReply{
		Policy: &p,
	}, nil
}

func (svr *RbacServiceServer) IsUserExist(ctx context.Context, in *pb.IsUserExistRequest) (*pb.IsUserExistReply, error) {
	uid := strconv.FormatFloat(in.GetUserId(), 'e', -1, 64)

	r := svr.StorageClient.DB().QueryRow(`SELECT name, role_id, create_time FROM users WHERE id = ` + uid + ``)
	p := pb.User{}

	err := r.Scan(&p.Name, &p.RoleId, &p.CreateTime)
	if err == sql.ErrNoRows {
		return &pb.IsUserExistReply{
			Found: false,
		}, nil
	}
	if err != nil {
		return &pb.IsUserExistReply{
			Found: false,
		}, err
	}

	return &pb.IsUserExistReply{
		Found: true,
	}, nil
}
