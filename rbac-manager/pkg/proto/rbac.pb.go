// protoc --go_out=rbac-manager/pkg/proto --go-grpc_out=rbac-manager/pkg/proto rbac-manager/pkg/proto/rbac.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: rbac-manager/pkg/proto/rbac.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_CREATE   Action = 0
	Action_RETRIEVE Action = 1
	Action_UPDATE   Action = 2
	Action_DELETE   Action = 3
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "CREATE",
		1: "RETRIEVE",
		2: "UPDATE",
		3: "DELETE",
	}
	Action_value = map[string]int32{
		"CREATE":   0,
		"RETRIEVE": 1,
		"UPDATE":   2,
		"DELETE":   3,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_rbac_manager_pkg_proto_rbac_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_rbac_manager_pkg_proto_rbac_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{0}
}

type Result int32

const (
	Result_ASK   Result = 0
	Result_ALLOW Result = 1
	Result_DENY  Result = 2
	Result_ERROR Result = 3
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "ASK",
		1: "ALLOW",
		2: "DENY",
		3: "ERROR",
	}
	Result_value = map[string]int32{
		"ASK":   0,
		"ALLOW": 1,
		"DENY":  2,
		"ERROR": 3,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_rbac_manager_pkg_proto_rbac_proto_enumTypes[1].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_rbac_manager_pkg_proto_rbac_proto_enumTypes[1]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{1}
}

type RbacRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Target   string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Action   Action `protobuf:"varint,3,opt,name=action,proto3,enum=proto.Action" json:"action,omitempty"`
}

func (x *RbacRequest) Reset() {
	*x = RbacRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RbacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RbacRequest) ProtoMessage() {}

func (x *RbacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RbacRequest.ProtoReflect.Descriptor instead.
func (*RbacRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *RbacRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RbacRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *RbacRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_CREATE
}

type RbacReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=proto.Result" json:"result,omitempty"`
}

func (x *RbacReply) Reset() {
	*x = RbacReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RbacReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RbacReply) ProtoMessage() {}

func (x *RbacReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RbacReply.ProtoReflect.Descriptor instead.
func (*RbacReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *RbacReply) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_ASK
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesRequest.ProtoReflect.Descriptor instead.
func (*GetRolesRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{3}
}

type GetRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *GetRolesReply) Reset() {
	*x = GetRolesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesReply) ProtoMessage() {}

func (x *GetRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesReply.ProtoReflect.Descriptor instead.
func (*GetRolesReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{4}
}

func (x *GetRolesReply) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{5}
}

func (x *GetRoleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *GetRoleReply) Reset() {
	*x = GetRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReply) ProtoMessage() {}

func (x *GetRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReply.ProtoReflect.Descriptor instead.
func (*GetRoleReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{6}
}

func (x *GetRoleReply) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Json string `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{7}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type GetPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *GetPoliciesRequest) Reset() {
	*x = GetPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliciesRequest) ProtoMessage() {}

func (x *GetPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliciesRequest.ProtoReflect.Descriptor instead.
func (*GetPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{8}
}

func (x *GetPoliciesRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type GetPoliciesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *GetPoliciesReply) Reset() {
	*x = GetPoliciesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliciesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliciesReply) ProtoMessage() {}

func (x *GetPoliciesReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliciesReply.ProtoReflect.Descriptor instead.
func (*GetPoliciesReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{9}
}

func (x *GetPoliciesReply) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{10}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type GetProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetProjectsRequest) Reset() {
	*x = GetProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectsRequest) ProtoMessage() {}

func (x *GetProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectsRequest.ProtoReflect.Descriptor instead.
func (*GetProjectsRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{11}
}

func (x *GetProjectsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetProjectsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *GetProjectsReply) Reset() {
	*x = GetProjectsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectsReply) ProtoMessage() {}

func (x *GetProjectsReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_rbac_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectsReply.ProtoReflect.Descriptor instead.
func (*GetProjectsReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP(), []int{12}
}

func (x *GetProjectsReply) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

var File_rbac_manager_pkg_proto_rbac_proto protoreflect.FileDescriptor

var file_rbac_manager_pkg_proto_rbac_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x62, 0x61, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0b, 0x52, 0x62,
	0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x09, 0x52, 0x62, 0x61, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2c,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2a, 0x3a, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x54, 0x52, 0x49, 0x45, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x2a, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x32, 0xbd, 0x02, 0x0a, 0x0b,
	0x52, 0x62, 0x61, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x62, 0x61,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x62, 0x61, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rbac_manager_pkg_proto_rbac_proto_rawDescOnce sync.Once
	file_rbac_manager_pkg_proto_rbac_proto_rawDescData = file_rbac_manager_pkg_proto_rbac_proto_rawDesc
)

func file_rbac_manager_pkg_proto_rbac_proto_rawDescGZIP() []byte {
	file_rbac_manager_pkg_proto_rbac_proto_rawDescOnce.Do(func() {
		file_rbac_manager_pkg_proto_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_rbac_manager_pkg_proto_rbac_proto_rawDescData)
	})
	return file_rbac_manager_pkg_proto_rbac_proto_rawDescData
}

var file_rbac_manager_pkg_proto_rbac_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rbac_manager_pkg_proto_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rbac_manager_pkg_proto_rbac_proto_goTypes = []any{
	(Action)(0),                // 0: proto.Action
	(Result)(0),                // 1: proto.Result
	(*RbacRequest)(nil),        // 2: proto.RbacRequest
	(*RbacReply)(nil),          // 3: proto.RbacReply
	(*Role)(nil),               // 4: proto.Role
	(*GetRolesRequest)(nil),    // 5: proto.GetRolesRequest
	(*GetRolesReply)(nil),      // 6: proto.GetRolesReply
	(*GetRoleRequest)(nil),     // 7: proto.GetRoleRequest
	(*GetRoleReply)(nil),       // 8: proto.GetRoleReply
	(*Policy)(nil),             // 9: proto.Policy
	(*GetPoliciesRequest)(nil), // 10: proto.GetPoliciesRequest
	(*GetPoliciesReply)(nil),   // 11: proto.GetPoliciesReply
	(*Project)(nil),            // 12: proto.Project
	(*GetProjectsRequest)(nil), // 13: proto.GetProjectsRequest
	(*GetProjectsReply)(nil),   // 14: proto.GetProjectsReply
}
var file_rbac_manager_pkg_proto_rbac_proto_depIdxs = []int32{
	0,  // 0: proto.RbacRequest.action:type_name -> proto.Action
	1,  // 1: proto.RbacReply.result:type_name -> proto.Result
	4,  // 2: proto.GetRolesReply.roles:type_name -> proto.Role
	4,  // 3: proto.GetRoleReply.role:type_name -> proto.Role
	9,  // 4: proto.GetPoliciesReply.policies:type_name -> proto.Policy
	12, // 5: proto.GetProjectsReply.projects:type_name -> proto.Project
	2,  // 6: proto.RbacService.Check:input_type -> proto.RbacRequest
	5,  // 7: proto.RbacService.GetRoles:input_type -> proto.GetRolesRequest
	7,  // 8: proto.RbacService.GetRole:input_type -> proto.GetRoleRequest
	10, // 9: proto.RbacService.GetPolicies:input_type -> proto.GetPoliciesRequest
	13, // 10: proto.RbacService.GetProjects:input_type -> proto.GetProjectsRequest
	3,  // 11: proto.RbacService.Check:output_type -> proto.RbacReply
	6,  // 12: proto.RbacService.GetRoles:output_type -> proto.GetRolesReply
	8,  // 13: proto.RbacService.GetRole:output_type -> proto.GetRoleReply
	11, // 14: proto.RbacService.GetPolicies:output_type -> proto.GetPoliciesReply
	14, // 15: proto.RbacService.GetProjects:output_type -> proto.GetProjectsReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rbac_manager_pkg_proto_rbac_proto_init() }
func file_rbac_manager_pkg_proto_rbac_proto_init() {
	if File_rbac_manager_pkg_proto_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RbacRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RbacReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRolesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetRolesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoleReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoliciesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoliciesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Project); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetProjectsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rbac_manager_pkg_proto_rbac_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetProjectsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rbac_manager_pkg_proto_rbac_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rbac_manager_pkg_proto_rbac_proto_goTypes,
		DependencyIndexes: file_rbac_manager_pkg_proto_rbac_proto_depIdxs,
		EnumInfos:         file_rbac_manager_pkg_proto_rbac_proto_enumTypes,
		MessageInfos:      file_rbac_manager_pkg_proto_rbac_proto_msgTypes,
	}.Build()
	File_rbac_manager_pkg_proto_rbac_proto = out.File
	file_rbac_manager_pkg_proto_rbac_proto_rawDesc = nil
	file_rbac_manager_pkg_proto_rbac_proto_goTypes = nil
	file_rbac_manager_pkg_proto_rbac_proto_depIdxs = nil
}
