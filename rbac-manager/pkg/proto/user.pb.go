// protoc --go_out=rbac-manager/pkg/proto --go-grpc_out=rbac-manager/pkg/proto rbac-manager/pkg/proto/rbac.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: rbac-manager/pkg/proto/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash string `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	ProjectId    int32  `protobuf:"varint,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *User) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ProjectUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *ProjectUser) Reset() {
	*x = ProjectUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUser) ProtoMessage() {}

func (x *ProjectUser) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUser.ProtoReflect.Descriptor instead.
func (*ProjectUser) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectUser) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ProjectUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *Success) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

type ListProjectUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ListProjectUsersRequest) Reset() {
	*x = ListProjectUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectUsersRequest) ProtoMessage() {}

func (x *ListProjectUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectUsersRequest.ProtoReflect.Descriptor instead.
func (*ListProjectUsersRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *ListProjectUsersRequest) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PutUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ProjectId int32  `protobuf:"varint,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *PutUserRequest) Reset() {
	*x = PutUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutUserRequest) ProtoMessage() {}

func (x *PutUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutUserRequest.ProtoReflect.Descriptor instead.
func (*PutUserRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *PutUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PutUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PutUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PutUserRequest) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserInfoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password    string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ProjectName []string `protobuf:"bytes,5,rep,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	RoleName    []string `protobuf:"bytes,6,rep,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *GetUserInfoReply) Reset() {
	*x = GetUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReply) ProtoMessage() {}

func (x *GetUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReply.ProtoReflect.Descriptor instead.
func (*GetUserInfoReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserInfoReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserInfoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserInfoReply) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetUserInfoReply) GetProjectName() []string {
	if x != nil {
		return x.ProjectName
	}
	return nil
}

func (x *GetUserInfoReply) GetRoleName() []string {
	if x != nil {
		return x.RoleName
	}
	return nil
}

type ListProjectUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *ProjectUser `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName *ProjectUser `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *ListProjectUsersReply) Reset() {
	*x = ListProjectUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectUsersReply) ProtoMessage() {}

func (x *ListProjectUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_rbac_manager_pkg_proto_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectUsersReply.ProtoReflect.Descriptor instead.
func (*ListProjectUsersReply) Descriptor() ([]byte, []int) {
	return file_rbac_manager_pkg_proto_user_proto_rawDescGZIP(), []int{8}
}

func (x *ListProjectUsersReply) GetUserId() *ProjectUser {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ListProjectUsersReply) GetUserName() *ProjectUser {
	if x != nil {
		return x.UserName
	}
	return nil
}

var File_rbac_manager_pkg_proto_user_proto protoreflect.FileDescriptor

var file_rbac_manager_pkg_proto_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x62, 0x61, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x43, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x50,
	0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xcc, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rbac_manager_pkg_proto_user_proto_rawDescOnce sync.Once
	file_rbac_manager_pkg_proto_user_proto_rawDescData = file_rbac_manager_pkg_proto_user_proto_rawDesc
)

func file_rbac_manager_pkg_proto_user_proto_rawDescGZIP() []byte {
	file_rbac_manager_pkg_proto_user_proto_rawDescOnce.Do(func() {
		file_rbac_manager_pkg_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rbac_manager_pkg_proto_user_proto_rawDescData)
	})
	return file_rbac_manager_pkg_proto_user_proto_rawDescData
}

var file_rbac_manager_pkg_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rbac_manager_pkg_proto_user_proto_goTypes = []any{
	(*User)(nil),                    // 0: proto.User
	(*ProjectUser)(nil),             // 1: proto.ProjectUser
	(*Success)(nil),                 // 2: proto.Success
	(*ListProjectUsersRequest)(nil), // 3: proto.ListProjectUsersRequest
	(*DeleteUserRequest)(nil),       // 4: proto.DeleteUserRequest
	(*PutUserRequest)(nil),          // 5: proto.PutUserRequest
	(*GetUserInfoRequest)(nil),      // 6: proto.GetUserInfoRequest
	(*GetUserInfoReply)(nil),        // 7: proto.GetUserInfoReply
	(*ListProjectUsersReply)(nil),   // 8: proto.ListProjectUsersReply
}
var file_rbac_manager_pkg_proto_user_proto_depIdxs = []int32{
	1, // 0: proto.ListProjectUsersReply.user_id:type_name -> proto.ProjectUser
	1, // 1: proto.ListProjectUsersReply.user_name:type_name -> proto.ProjectUser
	3, // 2: proto.UserService.ListProjectUsers:input_type -> proto.ListProjectUsersRequest
	4, // 3: proto.UserService.DeleteUser:input_type -> proto.DeleteUserRequest
	5, // 4: proto.UserService.PutUser:input_type -> proto.PutUserRequest
	6, // 5: proto.UserService.GetUserInfo:input_type -> proto.GetUserInfoRequest
	5, // 6: proto.UserService.PutUserInfo:input_type -> proto.PutUserRequest
	8, // 7: proto.UserService.ListProjectUsers:output_type -> proto.ListProjectUsersReply
	2, // 8: proto.UserService.DeleteUser:output_type -> proto.Success
	2, // 9: proto.UserService.PutUser:output_type -> proto.Success
	7, // 10: proto.UserService.GetUserInfo:output_type -> proto.GetUserInfoReply
	2, // 11: proto.UserService.PutUserInfo:output_type -> proto.Success
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rbac_manager_pkg_proto_user_proto_init() }
func file_rbac_manager_pkg_proto_user_proto_init() {
	if File_rbac_manager_pkg_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rbac_manager_pkg_proto_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProjectUser); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Success); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListProjectUsersRequest); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserRequest); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PutUserRequest); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoRequest); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoReply); i {
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
		file_rbac_manager_pkg_proto_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListProjectUsersReply); i {
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
			RawDescriptor: file_rbac_manager_pkg_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rbac_manager_pkg_proto_user_proto_goTypes,
		DependencyIndexes: file_rbac_manager_pkg_proto_user_proto_depIdxs,
		MessageInfos:      file_rbac_manager_pkg_proto_user_proto_msgTypes,
	}.Build()
	File_rbac_manager_pkg_proto_user_proto = out.File
	file_rbac_manager_pkg_proto_user_proto_rawDesc = nil
	file_rbac_manager_pkg_proto_user_proto_goTypes = nil
	file_rbac_manager_pkg_proto_user_proto_depIdxs = nil
}
