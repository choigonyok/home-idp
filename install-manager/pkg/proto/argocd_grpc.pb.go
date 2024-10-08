// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: install-manager/pkg/proto/argocd.proto

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

// ArgoCDClient is the client API for ArgoCD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArgoCDClient interface {
	InstallArgoCDChart(ctx context.Context, in *InstallArgoCDChartRequest, opts ...grpc.CallOption) (*InstallArgoCDChartReply, error)
	UninstallArgoCDChart(ctx context.Context, in *UninstallArgoCDChartRequest, opts ...grpc.CallOption) (*UninstallArgoCDChartReply, error)
}

type argoCDClient struct {
	cc grpc.ClientConnInterface
}

func NewArgoCDClient(cc grpc.ClientConnInterface) ArgoCDClient {
	return &argoCDClient{cc}
}

func (c *argoCDClient) InstallArgoCDChart(ctx context.Context, in *InstallArgoCDChartRequest, opts ...grpc.CallOption) (*InstallArgoCDChartReply, error) {
	out := new(InstallArgoCDChartReply)
	err := c.cc.Invoke(ctx, "/proto.ArgoCD/InstallArgoCDChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *argoCDClient) UninstallArgoCDChart(ctx context.Context, in *UninstallArgoCDChartRequest, opts ...grpc.CallOption) (*UninstallArgoCDChartReply, error) {
	out := new(UninstallArgoCDChartReply)
	err := c.cc.Invoke(ctx, "/proto.ArgoCD/UninstallArgoCDChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArgoCDServer is the server API for ArgoCD service.
// All implementations must embed UnimplementedArgoCDServer
// for forward compatibility
type ArgoCDServer interface {
	InstallArgoCDChart(context.Context, *InstallArgoCDChartRequest) (*InstallArgoCDChartReply, error)
	UninstallArgoCDChart(context.Context, *UninstallArgoCDChartRequest) (*UninstallArgoCDChartReply, error)
	mustEmbedUnimplementedArgoCDServer()
}

// UnimplementedArgoCDServer must be embedded to have forward compatible implementations.
type UnimplementedArgoCDServer struct {
}

func (UnimplementedArgoCDServer) InstallArgoCDChart(context.Context, *InstallArgoCDChartRequest) (*InstallArgoCDChartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallArgoCDChart not implemented")
}
func (UnimplementedArgoCDServer) UninstallArgoCDChart(context.Context, *UninstallArgoCDChartRequest) (*UninstallArgoCDChartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallArgoCDChart not implemented")
}
func (UnimplementedArgoCDServer) mustEmbedUnimplementedArgoCDServer() {}

// UnsafeArgoCDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArgoCDServer will
// result in compilation errors.
type UnsafeArgoCDServer interface {
	mustEmbedUnimplementedArgoCDServer()
}

func RegisterArgoCDServer(s grpc.ServiceRegistrar, srv ArgoCDServer) {
	s.RegisterService(&ArgoCD_ServiceDesc, srv)
}

func _ArgoCD_InstallArgoCDChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallArgoCDChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgoCDServer).InstallArgoCDChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ArgoCD/InstallArgoCDChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgoCDServer).InstallArgoCDChart(ctx, req.(*InstallArgoCDChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArgoCD_UninstallArgoCDChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallArgoCDChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgoCDServer).UninstallArgoCDChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ArgoCD/UninstallArgoCDChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgoCDServer).UninstallArgoCDChart(ctx, req.(*UninstallArgoCDChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArgoCD_ServiceDesc is the grpc.ServiceDesc for ArgoCD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArgoCD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ArgoCD",
	HandlerType: (*ArgoCDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstallArgoCDChart",
			Handler:    _ArgoCD_InstallArgoCDChart_Handler,
		},
		{
			MethodName: "UninstallArgoCDChart",
			Handler:    _ArgoCD_UninstallArgoCDChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "install-manager/pkg/proto/argocd.proto",
}
