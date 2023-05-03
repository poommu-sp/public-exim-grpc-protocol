// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package warehouse_stock_report

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

// WarehouseStockReportClient is the client API for WarehouseStockReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseStockReportClient interface {
	ListWarehouseStockReport(ctx context.Context, in *ListWarehouseStockReportPayload, opts ...grpc.CallOption) (*ListWarehouseStockReportResponse, error)
}

type warehouseStockReportClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseStockReportClient(cc grpc.ClientConnInterface) WarehouseStockReportClient {
	return &warehouseStockReportClient{cc}
}

func (c *warehouseStockReportClient) ListWarehouseStockReport(ctx context.Context, in *ListWarehouseStockReportPayload, opts ...grpc.CallOption) (*ListWarehouseStockReportResponse, error) {
	out := new(ListWarehouseStockReportResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.winfeed.warehouse_stock_report.WarehouseStockReport/ListWarehouseStockReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseStockReportServer is the server API for WarehouseStockReport service.
// All implementations must embed UnimplementedWarehouseStockReportServer
// for forward compatibility
type WarehouseStockReportServer interface {
	ListWarehouseStockReport(context.Context, *ListWarehouseStockReportPayload) (*ListWarehouseStockReportResponse, error)
	mustEmbedUnimplementedWarehouseStockReportServer()
}

// UnimplementedWarehouseStockReportServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseStockReportServer struct {
}

func (UnimplementedWarehouseStockReportServer) ListWarehouseStockReport(context.Context, *ListWarehouseStockReportPayload) (*ListWarehouseStockReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWarehouseStockReport not implemented")
}
func (UnimplementedWarehouseStockReportServer) mustEmbedUnimplementedWarehouseStockReportServer() {}

// UnsafeWarehouseStockReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseStockReportServer will
// result in compilation errors.
type UnsafeWarehouseStockReportServer interface {
	mustEmbedUnimplementedWarehouseStockReportServer()
}

func RegisterWarehouseStockReportServer(s grpc.ServiceRegistrar, srv WarehouseStockReportServer) {
	s.RegisterService(&WarehouseStockReport_ServiceDesc, srv)
}

func _WarehouseStockReport_ListWarehouseStockReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWarehouseStockReportPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseStockReportServer).ListWarehouseStockReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.winfeed.warehouse_stock_report.WarehouseStockReport/ListWarehouseStockReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseStockReportServer).ListWarehouseStockReport(ctx, req.(*ListWarehouseStockReportPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// WarehouseStockReport_ServiceDesc is the grpc.ServiceDesc for WarehouseStockReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarehouseStockReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "samantra.services.winfeed.warehouse_stock_report.WarehouseStockReport",
	HandlerType: (*WarehouseStockReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWarehouseStockReport",
			Handler:    _WarehouseStockReport_ListWarehouseStockReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samantra/services/winfeed_services/warehouse_stock_report_service/warehouse_stock_report.proto",
}
