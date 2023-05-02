// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: samantra/services/rm1_services/rm1.proto

package rm1

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

// RM1Client is the client API for RM1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RM1Client interface {
	ListOSLotInfo(ctx context.Context, in *ListOSLotInfoRequest, opts ...grpc.CallOption) (*ListOSLotInfoResponse, error)
	GetListOWRM1(ctx context.Context, in *ListOSLotInfoRequest, opts ...grpc.CallOption) (*ListRM1Response, error)
}

type rM1Client struct {
	cc grpc.ClientConnInterface
}

func NewRM1Client(cc grpc.ClientConnInterface) RM1Client {
	return &rM1Client{cc}
}

func (c *rM1Client) ListOSLotInfo(ctx context.Context, in *ListOSLotInfoRequest, opts ...grpc.CallOption) (*ListOSLotInfoResponse, error) {
	out := new(ListOSLotInfoResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.rm1.RM1/ListOSLotInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rM1Client) GetListOWRM1(ctx context.Context, in *ListOSLotInfoRequest, opts ...grpc.CallOption) (*ListRM1Response, error) {
	out := new(ListRM1Response)
	err := c.cc.Invoke(ctx, "/samantra.services.rm1.RM1/GetListOWRM1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RM1Server is the server API for RM1 service.
// All implementations must embed UnimplementedRM1Server
// for forward compatibility
type RM1Server interface {
	ListOSLotInfo(context.Context, *ListOSLotInfoRequest) (*ListOSLotInfoResponse, error)
	GetListOWRM1(context.Context, *ListOSLotInfoRequest) (*ListRM1Response, error)
	mustEmbedUnimplementedRM1Server()
}

// UnimplementedRM1Server must be embedded to have forward compatible implementations.
type UnimplementedRM1Server struct {
}

func (UnimplementedRM1Server) ListOSLotInfo(context.Context, *ListOSLotInfoRequest) (*ListOSLotInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOSLotInfo not implemented")
}
func (UnimplementedRM1Server) GetListOWRM1(context.Context, *ListOSLotInfoRequest) (*ListRM1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOWRM1 not implemented")
}
func (UnimplementedRM1Server) mustEmbedUnimplementedRM1Server() {}

// UnsafeRM1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RM1Server will
// result in compilation errors.
type UnsafeRM1Server interface {
	mustEmbedUnimplementedRM1Server()
}

func RegisterRM1Server(s grpc.ServiceRegistrar, srv RM1Server) {
	s.RegisterService(&RM1_ServiceDesc, srv)
}

func _RM1_ListOSLotInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOSLotInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RM1Server).ListOSLotInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.rm1.RM1/ListOSLotInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RM1Server).ListOSLotInfo(ctx, req.(*ListOSLotInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RM1_GetListOWRM1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOSLotInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RM1Server).GetListOWRM1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.rm1.RM1/GetListOWRM1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RM1Server).GetListOWRM1(ctx, req.(*ListOSLotInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RM1_ServiceDesc is the grpc.ServiceDesc for RM1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RM1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "samantra.services.rm1.RM1",
	HandlerType: (*RM1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOSLotInfo",
			Handler:    _RM1_ListOSLotInfo_Handler,
		},
		{
			MethodName: "GetListOWRM1",
			Handler:    _RM1_GetListOWRM1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samantra/services/rm1_services/rm1.proto",
}