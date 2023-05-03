// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package morning_report

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

// MorningReportClient is the client API for MorningReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MorningReportClient interface {
	InsertDailySnapshot(ctx context.Context, in *MorningReportInsertDailySnapshotPayLoad, opts ...grpc.CallOption) (*MorningReportInsertDailySnapshotResponse, error)
}

type morningReportClient struct {
	cc grpc.ClientConnInterface
}

func NewMorningReportClient(cc grpc.ClientConnInterface) MorningReportClient {
	return &morningReportClient{cc}
}

func (c *morningReportClient) InsertDailySnapshot(ctx context.Context, in *MorningReportInsertDailySnapshotPayLoad, opts ...grpc.CallOption) (*MorningReportInsertDailySnapshotResponse, error) {
	out := new(MorningReportInsertDailySnapshotResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.winfeed.morning_report.MorningReport/InsertDailySnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MorningReportServer is the server API for MorningReport service.
// All implementations must embed UnimplementedMorningReportServer
// for forward compatibility
type MorningReportServer interface {
	InsertDailySnapshot(context.Context, *MorningReportInsertDailySnapshotPayLoad) (*MorningReportInsertDailySnapshotResponse, error)
	mustEmbedUnimplementedMorningReportServer()
}

// UnimplementedMorningReportServer must be embedded to have forward compatible implementations.
type UnimplementedMorningReportServer struct {
}

func (UnimplementedMorningReportServer) InsertDailySnapshot(context.Context, *MorningReportInsertDailySnapshotPayLoad) (*MorningReportInsertDailySnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertDailySnapshot not implemented")
}
func (UnimplementedMorningReportServer) mustEmbedUnimplementedMorningReportServer() {}

// UnsafeMorningReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MorningReportServer will
// result in compilation errors.
type UnsafeMorningReportServer interface {
	mustEmbedUnimplementedMorningReportServer()
}

func RegisterMorningReportServer(s grpc.ServiceRegistrar, srv MorningReportServer) {
	s.RegisterService(&MorningReport_ServiceDesc, srv)
}

func _MorningReport_InsertDailySnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MorningReportInsertDailySnapshotPayLoad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorningReportServer).InsertDailySnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.winfeed.morning_report.MorningReport/InsertDailySnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorningReportServer).InsertDailySnapshot(ctx, req.(*MorningReportInsertDailySnapshotPayLoad))
	}
	return interceptor(ctx, in, info, handler)
}

// MorningReport_ServiceDesc is the grpc.ServiceDesc for MorningReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MorningReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "samantra.services.winfeed.morning_report.MorningReport",
	HandlerType: (*MorningReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertDailySnapshot",
			Handler:    _MorningReport_InsertDailySnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samantra/services/winfeed_services/morning_report_service/morning_report.proto",
}