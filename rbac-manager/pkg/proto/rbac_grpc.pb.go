// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: rbac-manager/pkg/proto/rbac.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RbacServiceClient is the client API for RbacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RbacServiceClient interface {
	Check(ctx context.Context, in *RbacRequest, opts ...grpc.CallOption) (*RbacReply, error)
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesReply, error)
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error)
	PostRole(ctx context.Context, in *PostRoleRequest, opts ...grpc.CallOption) (*PostRoleReply, error)
	PostPolicy(ctx context.Context, in *PostPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyReply, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyReply, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyReply, error)
	GetPolicies(ctx context.Context, in *GetPoliciesRequest, opts ...grpc.CallOption) (*GetPoliciesReply, error)
	GetPolicyJson(ctx context.Context, in *GetPolicyJsonRequest, opts ...grpc.CallOption) (*GetPolicyJsonReply, error)
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsReply, error)
	GetUsersInProject(ctx context.Context, in *GetUsersInProjectRequest, opts ...grpc.CallOption) (*GetUsersInProjectReply, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error)
	IsUserExist(ctx context.Context, in *IsUserExistRequest, opts ...grpc.CallOption) (*IsUserExistReply, error)
	PostUser(ctx context.Context, in *PostUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*UpdateUserRoleReply, error)
	PostProject(ctx context.Context, in *PostProjectRequest, opts ...grpc.CallOption) (*PostProjectReply, error)
	PutUser(ctx context.Context, in *PutUserRequest, opts ...grpc.CallOption) (*PutUserReply, error)
	GetDockerfiles(ctx context.Context, in *GetDockerfilesRequest, opts ...grpc.CallOption) (*GetDockerfilesReply, error)
	PostDockerfile(ctx context.Context, in *PostDockerfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTraceId(ctx context.Context, in *GetTraceIdRequest, opts ...grpc.CallOption) (*GetTraceIdReply, error)
	GetTraceIdByDockerfileId(ctx context.Context, in *GetTraceIdByDockerfileIdRequest, opts ...grpc.CallOption) (*GetTraceIdByDockerfileIdReply, error)
}

type rbacServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRbacServiceClient(cc grpc.ClientConnInterface) RbacServiceClient {
	return &rbacServiceClient{cc}
}

func (c *rbacServiceClient) Check(ctx context.Context, in *RbacRequest, opts ...grpc.CallOption) (*RbacReply, error) {
	out := new(RbacReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesReply, error) {
	out := new(GetRolesReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error) {
	out := new(GetRoleReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PostRole(ctx context.Context, in *PostRoleRequest, opts ...grpc.CallOption) (*PostRoleReply, error) {
	out := new(PostRoleReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PostRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PostPolicy(ctx context.Context, in *PostPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PostPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyReply, error) {
	out := new(DeletePolicyReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error) {
	out := new(DeleteRoleReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error) {
	out := new(DeleteProjectReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyReply, error) {
	out := new(GetPolicyReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyReply, error) {
	out := new(UpdatePolicyReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/UpdatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetPolicies(ctx context.Context, in *GetPoliciesRequest, opts ...grpc.CallOption) (*GetPoliciesReply, error) {
	out := new(GetPoliciesReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetPolicyJson(ctx context.Context, in *GetPolicyJsonRequest, opts ...grpc.CallOption) (*GetPolicyJsonReply, error) {
	out := new(GetPolicyJsonReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetPolicyJson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsReply, error) {
	out := new(GetProjectsReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetUsersInProject(ctx context.Context, in *GetUsersInProjectRequest, opts ...grpc.CallOption) (*GetUsersInProjectReply, error) {
	out := new(GetUsersInProjectReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetUsersInProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error) {
	out := new(GetUsersReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) IsUserExist(ctx context.Context, in *IsUserExistRequest, opts ...grpc.CallOption) (*IsUserExistReply, error) {
	out := new(IsUserExistReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/IsUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PostUser(ctx context.Context, in *PostUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PostUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*UpdateUserRoleReply, error) {
	out := new(UpdateUserRoleReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PostProject(ctx context.Context, in *PostProjectRequest, opts ...grpc.CallOption) (*PostProjectReply, error) {
	out := new(PostProjectReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PostProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PutUser(ctx context.Context, in *PutUserRequest, opts ...grpc.CallOption) (*PutUserReply, error) {
	out := new(PutUserReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetDockerfiles(ctx context.Context, in *GetDockerfilesRequest, opts ...grpc.CallOption) (*GetDockerfilesReply, error) {
	out := new(GetDockerfilesReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetDockerfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) PostDockerfile(ctx context.Context, in *PostDockerfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RbacService/PostDockerfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetTraceId(ctx context.Context, in *GetTraceIdRequest, opts ...grpc.CallOption) (*GetTraceIdReply, error) {
	out := new(GetTraceIdReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetTraceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetTraceIdByDockerfileId(ctx context.Context, in *GetTraceIdByDockerfileIdRequest, opts ...grpc.CallOption) (*GetTraceIdByDockerfileIdReply, error) {
	out := new(GetTraceIdByDockerfileIdReply)
	err := c.cc.Invoke(ctx, "/proto.RbacService/GetTraceIdByDockerfileId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RbacServiceServer is the server API for RbacService service.
// All implementations must embed UnimplementedRbacServiceServer
// for forward compatibility
type RbacServiceServer interface {
	Check(context.Context, *RbacRequest) (*RbacReply, error)
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesReply, error)
	GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error)
	PostRole(context.Context, *PostRoleRequest) (*PostRoleReply, error)
	PostPolicy(context.Context, *PostPolicyRequest) (*emptypb.Empty, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyReply, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyReply, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyReply, error)
	GetPolicies(context.Context, *GetPoliciesRequest) (*GetPoliciesReply, error)
	GetPolicyJson(context.Context, *GetPolicyJsonRequest) (*GetPolicyJsonReply, error)
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsReply, error)
	GetUsersInProject(context.Context, *GetUsersInProjectRequest) (*GetUsersInProjectReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	IsUserExist(context.Context, *IsUserExistRequest) (*IsUserExistReply, error)
	PostUser(context.Context, *PostUserRequest) (*emptypb.Empty, error)
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*UpdateUserRoleReply, error)
	PostProject(context.Context, *PostProjectRequest) (*PostProjectReply, error)
	PutUser(context.Context, *PutUserRequest) (*PutUserReply, error)
	GetDockerfiles(context.Context, *GetDockerfilesRequest) (*GetDockerfilesReply, error)
	PostDockerfile(context.Context, *PostDockerfileRequest) (*emptypb.Empty, error)
	GetTraceId(context.Context, *GetTraceIdRequest) (*GetTraceIdReply, error)
	GetTraceIdByDockerfileId(context.Context, *GetTraceIdByDockerfileIdRequest) (*GetTraceIdByDockerfileIdReply, error)
	mustEmbedUnimplementedRbacServiceServer()
}

// UnimplementedRbacServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRbacServiceServer struct {
}

func (UnimplementedRbacServiceServer) Check(context.Context, *RbacRequest) (*RbacReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedRbacServiceServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedRbacServiceServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRbacServiceServer) PostRole(context.Context, *PostRoleRequest) (*PostRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRole not implemented")
}
func (UnimplementedRbacServiceServer) PostPolicy(context.Context, *PostPolicyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPolicy not implemented")
}
func (UnimplementedRbacServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedRbacServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRbacServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedRbacServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedRbacServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedRbacServiceServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedRbacServiceServer) GetPolicies(context.Context, *GetPoliciesRequest) (*GetPoliciesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicies not implemented")
}
func (UnimplementedRbacServiceServer) GetPolicyJson(context.Context, *GetPolicyJsonRequest) (*GetPolicyJsonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyJson not implemented")
}
func (UnimplementedRbacServiceServer) GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedRbacServiceServer) GetUsersInProject(context.Context, *GetUsersInProjectRequest) (*GetUsersInProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInProject not implemented")
}
func (UnimplementedRbacServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedRbacServiceServer) IsUserExist(context.Context, *IsUserExistRequest) (*IsUserExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserExist not implemented")
}
func (UnimplementedRbacServiceServer) PostUser(context.Context, *PostUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUser not implemented")
}
func (UnimplementedRbacServiceServer) UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*UpdateUserRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedRbacServiceServer) PostProject(context.Context, *PostProjectRequest) (*PostProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProject not implemented")
}
func (UnimplementedRbacServiceServer) PutUser(context.Context, *PutUserRequest) (*PutUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutUser not implemented")
}
func (UnimplementedRbacServiceServer) GetDockerfiles(context.Context, *GetDockerfilesRequest) (*GetDockerfilesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockerfiles not implemented")
}
func (UnimplementedRbacServiceServer) PostDockerfile(context.Context, *PostDockerfileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDockerfile not implemented")
}
func (UnimplementedRbacServiceServer) GetTraceId(context.Context, *GetTraceIdRequest) (*GetTraceIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceId not implemented")
}
func (UnimplementedRbacServiceServer) GetTraceIdByDockerfileId(context.Context, *GetTraceIdByDockerfileIdRequest) (*GetTraceIdByDockerfileIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceIdByDockerfileId not implemented")
}
func (UnimplementedRbacServiceServer) mustEmbedUnimplementedRbacServiceServer() {}

// UnsafeRbacServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RbacServiceServer will
// result in compilation errors.
type UnsafeRbacServiceServer interface {
	mustEmbedUnimplementedRbacServiceServer()
}

func RegisterRbacServiceServer(s grpc.ServiceRegistrar, srv RbacServiceServer) {
	s.RegisterService(&RbacService_ServiceDesc, srv)
}

func _RbacService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RbacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).Check(ctx, req.(*RbacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PostRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PostRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PostRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PostRole(ctx, req.(*PostRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PostPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PostPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PostPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PostPolicy(ctx, req.(*PostPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/UpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetPolicies(ctx, req.(*GetPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetPolicyJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetPolicyJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetPolicyJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetPolicyJson(ctx, req.(*GetPolicyJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetUsersInProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetUsersInProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetUsersInProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetUsersInProject(ctx, req.(*GetUsersInProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_IsUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).IsUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/IsUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).IsUserExist(ctx, req.(*IsUserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PostUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PostUser(ctx, req.(*PostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).UpdateUserRole(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PostProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PostProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PostProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PostProject(ctx, req.(*PostProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PutUser(ctx, req.(*PutUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetDockerfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDockerfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetDockerfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetDockerfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetDockerfiles(ctx, req.(*GetDockerfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_PostDockerfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDockerfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).PostDockerfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/PostDockerfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).PostDockerfile(ctx, req.(*PostDockerfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetTraceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTraceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetTraceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetTraceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetTraceId(ctx, req.(*GetTraceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetTraceIdByDockerfileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTraceIdByDockerfileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetTraceIdByDockerfileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RbacService/GetTraceIdByDockerfileId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetTraceIdByDockerfileId(ctx, req.(*GetTraceIdByDockerfileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RbacService_ServiceDesc is the grpc.ServiceDesc for RbacService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RbacService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RbacService",
	HandlerType: (*RbacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _RbacService_Check_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _RbacService_GetRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RbacService_GetRole_Handler,
		},
		{
			MethodName: "PostRole",
			Handler:    _RbacService_PostRole_Handler,
		},
		{
			MethodName: "PostPolicy",
			Handler:    _RbacService_PostPolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _RbacService_DeletePolicy_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RbacService_DeleteRole_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _RbacService_DeleteProject_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _RbacService_DeleteUser_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _RbacService_GetPolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _RbacService_UpdatePolicy_Handler,
		},
		{
			MethodName: "GetPolicies",
			Handler:    _RbacService_GetPolicies_Handler,
		},
		{
			MethodName: "GetPolicyJson",
			Handler:    _RbacService_GetPolicyJson_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _RbacService_GetProjects_Handler,
		},
		{
			MethodName: "GetUsersInProject",
			Handler:    _RbacService_GetUsersInProject_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _RbacService_GetUsers_Handler,
		},
		{
			MethodName: "IsUserExist",
			Handler:    _RbacService_IsUserExist_Handler,
		},
		{
			MethodName: "PostUser",
			Handler:    _RbacService_PostUser_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _RbacService_UpdateUserRole_Handler,
		},
		{
			MethodName: "PostProject",
			Handler:    _RbacService_PostProject_Handler,
		},
		{
			MethodName: "PutUser",
			Handler:    _RbacService_PutUser_Handler,
		},
		{
			MethodName: "GetDockerfiles",
			Handler:    _RbacService_GetDockerfiles_Handler,
		},
		{
			MethodName: "PostDockerfile",
			Handler:    _RbacService_PostDockerfile_Handler,
		},
		{
			MethodName: "GetTraceId",
			Handler:    _RbacService_GetTraceId_Handler,
		},
		{
			MethodName: "GetTraceIdByDockerfileId",
			Handler:    _RbacService_GetTraceIdByDockerfileId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rbac-manager/pkg/proto/rbac.proto",
}
