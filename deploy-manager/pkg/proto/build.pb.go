// protoc --go_out=rbac-manager/pkg/proto --go-grpc_out=rbac-manager/pkg/proto rbac-manager/pkg/proto/rbac.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: deploy-manager/pkg/proto/build.proto

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_build_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type BuildDockerfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *Image `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *BuildDockerfileRequest) Reset() {
	*x = BuildDockerfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildDockerfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDockerfileRequest) ProtoMessage() {}

func (x *BuildDockerfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDockerfileRequest.ProtoReflect.Descriptor instead.
func (*BuildDockerfileRequest) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_build_proto_rawDescGZIP(), []int{1}
}

func (x *BuildDockerfileRequest) GetImg() *Image {
	if x != nil {
		return x.Img
	}
	return nil
}

type BuildDockerfileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
}

func (x *BuildDockerfileReply) Reset() {
	*x = BuildDockerfileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildDockerfileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDockerfileReply) ProtoMessage() {}

func (x *BuildDockerfileReply) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_manager_pkg_proto_build_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDockerfileReply.ProtoReflect.Descriptor instead.
func (*BuildDockerfileReply) Descriptor() ([]byte, []int) {
	return file_deploy_manager_pkg_proto_build_proto_rawDescGZIP(), []int{2}
}

func (x *BuildDockerfileReply) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

var File_deploy_manager_pkg_proto_build_proto protoreflect.FileDescriptor

var file_deploy_manager_pkg_proto_build_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x16, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x22, 0x30,
	0x0a, 0x14, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64,
	0x32, 0x58, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x4f, 0x0a, 0x0f, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deploy_manager_pkg_proto_build_proto_rawDescOnce sync.Once
	file_deploy_manager_pkg_proto_build_proto_rawDescData = file_deploy_manager_pkg_proto_build_proto_rawDesc
)

func file_deploy_manager_pkg_proto_build_proto_rawDescGZIP() []byte {
	file_deploy_manager_pkg_proto_build_proto_rawDescOnce.Do(func() {
		file_deploy_manager_pkg_proto_build_proto_rawDescData = protoimpl.X.CompressGZIP(file_deploy_manager_pkg_proto_build_proto_rawDescData)
	})
	return file_deploy_manager_pkg_proto_build_proto_rawDescData
}

var file_deploy_manager_pkg_proto_build_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deploy_manager_pkg_proto_build_proto_goTypes = []any{
	(*Image)(nil),                  // 0: proto.Image
	(*BuildDockerfileRequest)(nil), // 1: proto.BuildDockerfileRequest
	(*BuildDockerfileReply)(nil),   // 2: proto.BuildDockerfileReply
}
var file_deploy_manager_pkg_proto_build_proto_depIdxs = []int32{
	0, // 0: proto.BuildDockerfileRequest.img:type_name -> proto.Image
	1, // 1: proto.Build.BuildDockerfile:input_type -> proto.BuildDockerfileRequest
	2, // 2: proto.Build.BuildDockerfile:output_type -> proto.BuildDockerfileReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_deploy_manager_pkg_proto_build_proto_init() }
func file_deploy_manager_pkg_proto_build_proto_init() {
	if File_deploy_manager_pkg_proto_build_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deploy_manager_pkg_proto_build_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Image); i {
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
		file_deploy_manager_pkg_proto_build_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BuildDockerfileRequest); i {
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
		file_deploy_manager_pkg_proto_build_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BuildDockerfileReply); i {
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
			RawDescriptor: file_deploy_manager_pkg_proto_build_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deploy_manager_pkg_proto_build_proto_goTypes,
		DependencyIndexes: file_deploy_manager_pkg_proto_build_proto_depIdxs,
		MessageInfos:      file_deploy_manager_pkg_proto_build_proto_msgTypes,
	}.Build()
	File_deploy_manager_pkg_proto_build_proto = out.File
	file_deploy_manager_pkg_proto_build_proto_rawDesc = nil
	file_deploy_manager_pkg_proto_build_proto_goTypes = nil
	file_deploy_manager_pkg_proto_build_proto_depIdxs = nil
}