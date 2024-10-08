// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: deploy-manager/pkg/proto/manifest.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManifestServiceClient is the client API for ManifestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManifestServiceClient interface {
	ApplyManifest(ctx context.Context, in *ApplyManifestRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*SuccessReply, error)
}

type manifestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManifestServiceClient(cc grpc.ClientConnInterface) ManifestServiceClient {
	return &manifestServiceClient{cc}
}

func (c *manifestServiceClient) ApplyManifest(ctx context.Context, in *ApplyManifestRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/proto.ManifestService/ApplyManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manifestServiceClient) DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/proto.ManifestService/DeleteManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManifestServiceServer is the server API for ManifestService service.
// All implementations must embed UnimplementedManifestServiceServer
// for forward compatibility
type ManifestServiceServer interface {
	ApplyManifest(context.Context, *ApplyManifestRequest) (*SuccessReply, error)
	DeleteManifest(context.Context, *DeleteManifestRequest) (*SuccessReply, error)
	mustEmbedUnimplementedManifestServiceServer()
}

// UnimplementedManifestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManifestServiceServer struct {
}

func (UnimplementedManifestServiceServer) ApplyManifest(context.Context, *ApplyManifestRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyManifest not implemented")
}
func (UnimplementedManifestServiceServer) DeleteManifest(context.Context, *DeleteManifestRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManifest not implemented")
}
func (UnimplementedManifestServiceServer) mustEmbedUnimplementedManifestServiceServer() {}

// UnsafeManifestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManifestServiceServer will
// result in compilation errors.
type UnsafeManifestServiceServer interface {
	mustEmbedUnimplementedManifestServiceServer()
}

func RegisterManifestServiceServer(s grpc.ServiceRegistrar, srv ManifestServiceServer) {
	s.RegisterService(&ManifestService_ServiceDesc, srv)
}

func _ManifestService_ApplyManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifestServiceServer).ApplyManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManifestService/ApplyManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifestServiceServer).ApplyManifest(ctx, req.(*ApplyManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManifestService_DeleteManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifestServiceServer).DeleteManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManifestService/DeleteManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifestServiceServer).DeleteManifest(ctx, req.(*DeleteManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManifestService_ServiceDesc is the grpc.ServiceDesc for ManifestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManifestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ManifestService",
	HandlerType: (*ManifestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyManifest",
			Handler:    _ManifestService_ApplyManifest_Handler,
		},
		{
			MethodName: "DeleteManifest",
			Handler:    _ManifestService_DeleteManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploy-manager/pkg/proto/manifest.proto",
}
