// export GOPATH=/Users/choigonyok/go
// export PATH=$PATH:$(go env GOPATH)/bin
// protoc --go_out=deploy-manager/pkg/proto --go-grpc_out=deploy-manager/pkg/proto deploy-manager/pkg/proto/deploy.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: deploy-manager/pkg/proto/deploy.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

type DeployReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
}

func (x *DeployReply) Reset() {
	*x = DeployReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployReply) ProtoMessage() {}

func (x *DeployReply) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployReply.ProtoReflect.Descriptor instead.
func (*DeployReply) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{1}
}

func (x *DeployReply) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

type DeployPodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod *Pod `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
}

func (x *DeployPodRequest) Reset() {
	*x = DeployPodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployPodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployPodRequest) ProtoMessage() {}

func (x *DeployPodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployPodRequest.ProtoReflect.Descriptor instead.
func (*DeployPodRequest) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{2}
}

func (x *DeployPodRequest) GetPod() *Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Image         string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ContainerPort string `protobuf:"bytes,4,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{3}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Pod) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Pod) GetContainerPort() string {
	if x != nil {
		return x.ContainerPort
	}
	return ""
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{4}
}

func (x *Secret) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Secret) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeploySecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string    `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Pusher    string    `protobuf:"bytes,2,opt,name=pusher,proto3" json:"pusher,omitempty"`
	Secrets   []*Secret `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *DeploySecretRequest) Reset() {
	*x = DeploySecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploySecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploySecretRequest) ProtoMessage() {}

func (x *DeploySecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_deploy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploySecretRequest.ProtoReflect.Descriptor instead.
func (*DeploySecretRequest) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP(), []int{5}
}

func (x *DeploySecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeploySecretRequest) GetPusher() string {
	if x != nil {
		return x.Pusher
	}
	return ""
}

func (x *DeploySecretRequest) GetSecrets() []*Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

var File_deploy_manager_pkg_proto_deploy_proto protoreflect.FileDescriptor

var file_deploy_manager_pkg_proto_deploy_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65,
	0x64, 0x22, 0x30, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03,
	0x70, 0x6f, 0x64, 0x22, 0x74, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x30, 0x0a, 0x06, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x74, 0x0a, 0x13, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x75, 0x73, 0x68, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x32, 0xc4, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x34, 0x0a, 0x06,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6f, 0x64, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deploy_manager_pkg_proto_deploy_proto_rawDescOnce sync.Once
	file_deploy_manager_pkg_proto_deploy_proto_rawDescData = file_deploy_manager_pkg_proto_deploy_proto_rawDesc
)

func file_deploy_manager_pkg_proto_deploy_proto_rawDescGZIP() []byte {
	file_deploy_manager_pkg_proto_deploy_proto_rawDescOnce.Do(func() {
		file_deploy_manager_pkg_proto_deploy_proto_rawDescData = protoimpl.X.CompressGZIP(file_deploy_manager_pkg_proto_deploy_proto_rawDescData)
	})
	return file_deploy_manager_pkg_proto_deploy_proto_rawDescData
}

var file_deploy_manager_pkg_proto_deploy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_deploy_manager_pkg_proto_deploy_proto_goTypes = []any{
	(*DeployRequest)(nil),       // 0: proto.DeployRequest
	(*DeployReply)(nil),         // 1: proto.DeployReply
	(*DeployPodRequest)(nil),    // 2: proto.DeployPodRequest
	(*Pod)(nil),                 // 3: proto.Pod
	(*Secret)(nil),              // 4: proto.Secret
	(*DeploySecretRequest)(nil), // 5: proto.DeploySecretRequest
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_deploy_manager_pkg_proto_deploy_proto_depIdxs = []int32{
	3, // 0: proto.DeployPodRequest.pod:type_name -> proto.Pod
	4, // 1: proto.DeploySecretRequest.secrets:type_name -> proto.Secret
	0, // 2: proto.Deploy.Deploy:input_type -> proto.DeployRequest
	2, // 3: proto.Deploy.DeployPod:input_type -> proto.DeployPodRequest
	5, // 4: proto.Deploy.DeploySecret:input_type -> proto.DeploySecretRequest
	1, // 5: proto.Deploy.Deploy:output_type -> proto.DeployReply
	6, // 6: proto.Deploy.DeployPod:output_type -> google.protobuf.Empty
	6, // 7: proto.Deploy.DeploySecret:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_deploy_manager_pkg_proto_deploy_proto_init() }
func file_deploy_manager_pkg_proto_deploy_proto_init() {
	if File_deploy_manager_pkg_proto_deploy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeployRequest); i {
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
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeployReply); i {
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
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeployPodRequest); i {
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
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Pod); i {
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
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Secret); i {
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
		file_deploy_manager_pkg_proto_deploy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeploySecretRequest); i {
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
			RawDescriptor: file_deploy_manager_pkg_proto_deploy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deploy_manager_pkg_proto_deploy_proto_goTypes,
		DependencyIndexes: file_deploy_manager_pkg_proto_deploy_proto_depIdxs,
		MessageInfos:      file_deploy_manager_pkg_proto_deploy_proto_msgTypes,
	}.Build()
	File_deploy_manager_pkg_proto_deploy_proto = out.File
	file_deploy_manager_pkg_proto_deploy_proto_rawDesc = nil
	file_deploy_manager_pkg_proto_deploy_proto_goTypes = nil
	file_deploy_manager_pkg_proto_deploy_proto_depIdxs = nil
}
