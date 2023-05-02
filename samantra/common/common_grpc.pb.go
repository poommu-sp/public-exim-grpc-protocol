// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: samantra/common/common.proto

package common

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

// EmptyStandInClient is the client API for EmptyStandIn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmptyStandInClient interface {
	EmptyFunction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type emptyStandInClient struct {
	cc grpc.ClientConnInterface
}

func NewEmptyStandInClient(cc grpc.ClientConnInterface) EmptyStandInClient {
	return &emptyStandInClient{cc}
}

func (c *emptyStandInClient) EmptyFunction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/samantra.common.EmptyStandIn/EmptyFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmptyStandInServer is the server API for EmptyStandIn service.
// All implementations must embed UnimplementedEmptyStandInServer
// for forward compatibility
type EmptyStandInServer interface {
	EmptyFunction(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedEmptyStandInServer()
}

// UnimplementedEmptyStandInServer must be embedded to have forward compatible implementations.
type UnimplementedEmptyStandInServer struct {
}

func (UnimplementedEmptyStandInServer) EmptyFunction(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyFunction not implemented")
}
func (UnimplementedEmptyStandInServer) mustEmbedUnimplementedEmptyStandInServer() {}

// UnsafeEmptyStandInServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmptyStandInServer will
// result in compilation errors.
type UnsafeEmptyStandInServer interface {
	mustEmbedUnimplementedEmptyStandInServer()
}

func RegisterEmptyStandInServer(s grpc.ServiceRegistrar, srv EmptyStandInServer) {
	s.RegisterService(&EmptyStandIn_ServiceDesc, srv)
}

func _EmptyStandIn_EmptyFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmptyStandInServer).EmptyFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.common.EmptyStandIn/EmptyFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmptyStandInServer).EmptyFunction(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// EmptyStandIn_ServiceDesc is the grpc.ServiceDesc for EmptyStandIn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmptyStandIn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "samantra.common.EmptyStandIn",
	HandlerType: (*EmptyStandInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyFunction",
			Handler:    _EmptyStandIn_EmptyFunction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samantra/common/common.proto",
}