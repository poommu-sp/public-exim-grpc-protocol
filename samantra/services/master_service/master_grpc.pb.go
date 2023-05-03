// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: samantra/services/master_service/master.proto

package master

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

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	GetProductCodesWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductCodesResponse, error)
	GetWinfeedProductReportNamesWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WinfeedProductReportNamesResponse, error)
	GetWinfeedProductReportIDsWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WinfeedProductReportIDsResponse, error)
	GetProductCodesFromWinfeedProductReportName(ctx context.Context, in *WinfeedProductReportName, opts ...grpc.CallOption) (*ProductCodesResponse, error)
	GetProductCodesFromWinfeedProductReportID(ctx context.Context, in *WinfeedProductReportID, opts ...grpc.CallOption) (*ProductCodesResponse, error)
	GetAllPlants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error)
	GetPlants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error)
	GetPlantsAgro(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error)
	GetOrgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrgsResponse, error)
	GetOrgsAgro(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrgsResponse, error)
	GetSubOrgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubOrgsResponse, error)
	GetWinfeedProductComponentsWithSubIDs(ctx context.Context, in *GetWinfeedProductComponentsWithSubIDsRequest, opts ...grpc.CallOption) (*WinfeedProductComponentsResponse, error)
	GetWinfeedDivisionList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedDivisionListResponse, error)
	GetWinfeedDivisionListWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedDivisionListResponse, error)
	GetWinfeedDivision(ctx context.Context, in *GetWinfeedDivisionRequest, opts ...grpc.CallOption) (*GetWinfeedDivisionResponse, error)
	ListWinfeedProductByDivisionID(ctx context.Context, in *ListWinfeedProductByDivisionIDRequest, opts ...grpc.CallOption) (*ListWinfeedProductByDivisionIDResponse, error)
	ListWinfeedStock(ctx context.Context, in *ListWinfeedStockRequest, opts ...grpc.CallOption) (*ListWinfeedStockResponse, error)
	ListWinfeedPurchase(ctx context.Context, in *ListWinfeedPurchaseRequest, opts ...grpc.CallOption) (*ListWinfeedPurchaseResponse, error)
	ListWinfeedDailyPrice(ctx context.Context, in *ListWinfeedDailyPriceRequest, opts ...grpc.CallOption) (*ListWinfeedDailyPriceResponse, error)
	GetWinfeedGradeList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedGradeListResponse, error)
	ListWinfeedProduct(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListWinfeedProductByDivisionIDResponse, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetProductCodesWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductCodesResponse, error) {
	out := new(ProductCodesResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetProductCodesWithPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedProductReportNamesWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WinfeedProductReportNamesResponse, error) {
	out := new(WinfeedProductReportNamesResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedProductReportNamesWithPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedProductReportIDsWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WinfeedProductReportIDsResponse, error) {
	out := new(WinfeedProductReportIDsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedProductReportIDsWithPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetProductCodesFromWinfeedProductReportName(ctx context.Context, in *WinfeedProductReportName, opts ...grpc.CallOption) (*ProductCodesResponse, error) {
	out := new(ProductCodesResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetProductCodesFromWinfeedProductReportName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetProductCodesFromWinfeedProductReportID(ctx context.Context, in *WinfeedProductReportID, opts ...grpc.CallOption) (*ProductCodesResponse, error) {
	out := new(ProductCodesResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetProductCodesFromWinfeedProductReportID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetAllPlants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error) {
	out := new(PlantsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetAllPlants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetPlants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error) {
	out := new(PlantsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetPlants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetPlantsAgro(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlantsResponse, error) {
	out := new(PlantsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetPlantsAgro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetOrgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrgsResponse, error) {
	out := new(OrgsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetOrgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetOrgsAgro(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrgsResponse, error) {
	out := new(OrgsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetOrgsAgro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetSubOrgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubOrgsResponse, error) {
	out := new(SubOrgsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetSubOrgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedProductComponentsWithSubIDs(ctx context.Context, in *GetWinfeedProductComponentsWithSubIDsRequest, opts ...grpc.CallOption) (*WinfeedProductComponentsResponse, error) {
	out := new(WinfeedProductComponentsResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedProductComponentsWithSubIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedDivisionList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedDivisionListResponse, error) {
	out := new(GetWinfeedDivisionListResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedDivisionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedDivisionListWithPermission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedDivisionListResponse, error) {
	out := new(GetWinfeedDivisionListResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedDivisionListWithPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedDivision(ctx context.Context, in *GetWinfeedDivisionRequest, opts ...grpc.CallOption) (*GetWinfeedDivisionResponse, error) {
	out := new(GetWinfeedDivisionResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ListWinfeedProductByDivisionID(ctx context.Context, in *ListWinfeedProductByDivisionIDRequest, opts ...grpc.CallOption) (*ListWinfeedProductByDivisionIDResponse, error) {
	out := new(ListWinfeedProductByDivisionIDResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/ListWinfeedProductByDivisionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ListWinfeedStock(ctx context.Context, in *ListWinfeedStockRequest, opts ...grpc.CallOption) (*ListWinfeedStockResponse, error) {
	out := new(ListWinfeedStockResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/ListWinfeedStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ListWinfeedPurchase(ctx context.Context, in *ListWinfeedPurchaseRequest, opts ...grpc.CallOption) (*ListWinfeedPurchaseResponse, error) {
	out := new(ListWinfeedPurchaseResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/ListWinfeedPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ListWinfeedDailyPrice(ctx context.Context, in *ListWinfeedDailyPriceRequest, opts ...grpc.CallOption) (*ListWinfeedDailyPriceResponse, error) {
	out := new(ListWinfeedDailyPriceResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/ListWinfeedDailyPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetWinfeedGradeList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetWinfeedGradeListResponse, error) {
	out := new(GetWinfeedGradeListResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/GetWinfeedGradeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ListWinfeedProduct(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListWinfeedProductByDivisionIDResponse, error) {
	out := new(ListWinfeedProductByDivisionIDResponse)
	err := c.cc.Invoke(ctx, "/samantra.services.master.Master/ListWinfeedProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	GetUsers(context.Context, *Empty) (*UsersResponse, error)
	GetProductCodesWithPermission(context.Context, *Empty) (*ProductCodesResponse, error)
	GetWinfeedProductReportNamesWithPermission(context.Context, *Empty) (*WinfeedProductReportNamesResponse, error)
	GetWinfeedProductReportIDsWithPermission(context.Context, *Empty) (*WinfeedProductReportIDsResponse, error)
	GetProductCodesFromWinfeedProductReportName(context.Context, *WinfeedProductReportName) (*ProductCodesResponse, error)
	GetProductCodesFromWinfeedProductReportID(context.Context, *WinfeedProductReportID) (*ProductCodesResponse, error)
	GetAllPlants(context.Context, *Empty) (*PlantsResponse, error)
	GetPlants(context.Context, *Empty) (*PlantsResponse, error)
	GetPlantsAgro(context.Context, *Empty) (*PlantsResponse, error)
	GetOrgs(context.Context, *Empty) (*OrgsResponse, error)
	GetOrgsAgro(context.Context, *Empty) (*OrgsResponse, error)
	GetSubOrgs(context.Context, *Empty) (*SubOrgsResponse, error)
	GetWinfeedProductComponentsWithSubIDs(context.Context, *GetWinfeedProductComponentsWithSubIDsRequest) (*WinfeedProductComponentsResponse, error)
	GetWinfeedDivisionList(context.Context, *Empty) (*GetWinfeedDivisionListResponse, error)
	GetWinfeedDivisionListWithPermission(context.Context, *Empty) (*GetWinfeedDivisionListResponse, error)
	GetWinfeedDivision(context.Context, *GetWinfeedDivisionRequest) (*GetWinfeedDivisionResponse, error)
	ListWinfeedProductByDivisionID(context.Context, *ListWinfeedProductByDivisionIDRequest) (*ListWinfeedProductByDivisionIDResponse, error)
	ListWinfeedStock(context.Context, *ListWinfeedStockRequest) (*ListWinfeedStockResponse, error)
	ListWinfeedPurchase(context.Context, *ListWinfeedPurchaseRequest) (*ListWinfeedPurchaseResponse, error)
	ListWinfeedDailyPrice(context.Context, *ListWinfeedDailyPriceRequest) (*ListWinfeedDailyPriceResponse, error)
	GetWinfeedGradeList(context.Context, *Empty) (*GetWinfeedGradeListResponse, error)
	ListWinfeedProduct(context.Context, *Empty) (*ListWinfeedProductByDivisionIDResponse, error)
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) GetUsers(context.Context, *Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedMasterServer) GetProductCodesWithPermission(context.Context, *Empty) (*ProductCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCodesWithPermission not implemented")
}
func (UnimplementedMasterServer) GetWinfeedProductReportNamesWithPermission(context.Context, *Empty) (*WinfeedProductReportNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedProductReportNamesWithPermission not implemented")
}
func (UnimplementedMasterServer) GetWinfeedProductReportIDsWithPermission(context.Context, *Empty) (*WinfeedProductReportIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedProductReportIDsWithPermission not implemented")
}
func (UnimplementedMasterServer) GetProductCodesFromWinfeedProductReportName(context.Context, *WinfeedProductReportName) (*ProductCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCodesFromWinfeedProductReportName not implemented")
}
func (UnimplementedMasterServer) GetProductCodesFromWinfeedProductReportID(context.Context, *WinfeedProductReportID) (*ProductCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCodesFromWinfeedProductReportID not implemented")
}
func (UnimplementedMasterServer) GetAllPlants(context.Context, *Empty) (*PlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlants not implemented")
}
func (UnimplementedMasterServer) GetPlants(context.Context, *Empty) (*PlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlants not implemented")
}
func (UnimplementedMasterServer) GetPlantsAgro(context.Context, *Empty) (*PlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlantsAgro not implemented")
}
func (UnimplementedMasterServer) GetOrgs(context.Context, *Empty) (*OrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgs not implemented")
}
func (UnimplementedMasterServer) GetOrgsAgro(context.Context, *Empty) (*OrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgsAgro not implemented")
}
func (UnimplementedMasterServer) GetSubOrgs(context.Context, *Empty) (*SubOrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubOrgs not implemented")
}
func (UnimplementedMasterServer) GetWinfeedProductComponentsWithSubIDs(context.Context, *GetWinfeedProductComponentsWithSubIDsRequest) (*WinfeedProductComponentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedProductComponentsWithSubIDs not implemented")
}
func (UnimplementedMasterServer) GetWinfeedDivisionList(context.Context, *Empty) (*GetWinfeedDivisionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedDivisionList not implemented")
}
func (UnimplementedMasterServer) GetWinfeedDivisionListWithPermission(context.Context, *Empty) (*GetWinfeedDivisionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedDivisionListWithPermission not implemented")
}
func (UnimplementedMasterServer) GetWinfeedDivision(context.Context, *GetWinfeedDivisionRequest) (*GetWinfeedDivisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedDivision not implemented")
}
func (UnimplementedMasterServer) ListWinfeedProductByDivisionID(context.Context, *ListWinfeedProductByDivisionIDRequest) (*ListWinfeedProductByDivisionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWinfeedProductByDivisionID not implemented")
}
func (UnimplementedMasterServer) ListWinfeedStock(context.Context, *ListWinfeedStockRequest) (*ListWinfeedStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWinfeedStock not implemented")
}
func (UnimplementedMasterServer) ListWinfeedPurchase(context.Context, *ListWinfeedPurchaseRequest) (*ListWinfeedPurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWinfeedPurchase not implemented")
}
func (UnimplementedMasterServer) ListWinfeedDailyPrice(context.Context, *ListWinfeedDailyPriceRequest) (*ListWinfeedDailyPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWinfeedDailyPrice not implemented")
}
func (UnimplementedMasterServer) GetWinfeedGradeList(context.Context, *Empty) (*GetWinfeedGradeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinfeedGradeList not implemented")
}
func (UnimplementedMasterServer) ListWinfeedProduct(context.Context, *Empty) (*ListWinfeedProductByDivisionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWinfeedProduct not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetProductCodesWithPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetProductCodesWithPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetProductCodesWithPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetProductCodesWithPermission(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedProductReportNamesWithPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedProductReportNamesWithPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedProductReportNamesWithPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedProductReportNamesWithPermission(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedProductReportIDsWithPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedProductReportIDsWithPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedProductReportIDsWithPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedProductReportIDsWithPermission(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetProductCodesFromWinfeedProductReportName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinfeedProductReportName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetProductCodesFromWinfeedProductReportName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetProductCodesFromWinfeedProductReportName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetProductCodesFromWinfeedProductReportName(ctx, req.(*WinfeedProductReportName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetProductCodesFromWinfeedProductReportID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinfeedProductReportID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetProductCodesFromWinfeedProductReportID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetProductCodesFromWinfeedProductReportID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetProductCodesFromWinfeedProductReportID(ctx, req.(*WinfeedProductReportID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetAllPlants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetAllPlants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetAllPlants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetAllPlants(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetPlants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetPlants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetPlants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetPlants(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetPlantsAgro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetPlantsAgro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetPlantsAgro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetPlantsAgro(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetOrgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetOrgs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetOrgsAgro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetOrgsAgro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetOrgsAgro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetOrgsAgro(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetSubOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetSubOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetSubOrgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetSubOrgs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedProductComponentsWithSubIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWinfeedProductComponentsWithSubIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedProductComponentsWithSubIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedProductComponentsWithSubIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedProductComponentsWithSubIDs(ctx, req.(*GetWinfeedProductComponentsWithSubIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedDivisionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedDivisionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedDivisionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedDivisionList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedDivisionListWithPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedDivisionListWithPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedDivisionListWithPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedDivisionListWithPermission(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWinfeedDivisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedDivision(ctx, req.(*GetWinfeedDivisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ListWinfeedProductByDivisionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWinfeedProductByDivisionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ListWinfeedProductByDivisionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/ListWinfeedProductByDivisionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ListWinfeedProductByDivisionID(ctx, req.(*ListWinfeedProductByDivisionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ListWinfeedStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWinfeedStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ListWinfeedStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/ListWinfeedStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ListWinfeedStock(ctx, req.(*ListWinfeedStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ListWinfeedPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWinfeedPurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ListWinfeedPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/ListWinfeedPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ListWinfeedPurchase(ctx, req.(*ListWinfeedPurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ListWinfeedDailyPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWinfeedDailyPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ListWinfeedDailyPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/ListWinfeedDailyPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ListWinfeedDailyPrice(ctx, req.(*ListWinfeedDailyPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetWinfeedGradeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetWinfeedGradeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/GetWinfeedGradeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetWinfeedGradeList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ListWinfeedProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ListWinfeedProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samantra.services.master.Master/ListWinfeedProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ListWinfeedProduct(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "samantra.services.master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _Master_GetUsers_Handler,
		},
		{
			MethodName: "GetProductCodesWithPermission",
			Handler:    _Master_GetProductCodesWithPermission_Handler,
		},
		{
			MethodName: "GetWinfeedProductReportNamesWithPermission",
			Handler:    _Master_GetWinfeedProductReportNamesWithPermission_Handler,
		},
		{
			MethodName: "GetWinfeedProductReportIDsWithPermission",
			Handler:    _Master_GetWinfeedProductReportIDsWithPermission_Handler,
		},
		{
			MethodName: "GetProductCodesFromWinfeedProductReportName",
			Handler:    _Master_GetProductCodesFromWinfeedProductReportName_Handler,
		},
		{
			MethodName: "GetProductCodesFromWinfeedProductReportID",
			Handler:    _Master_GetProductCodesFromWinfeedProductReportID_Handler,
		},
		{
			MethodName: "GetAllPlants",
			Handler:    _Master_GetAllPlants_Handler,
		},
		{
			MethodName: "GetPlants",
			Handler:    _Master_GetPlants_Handler,
		},
		{
			MethodName: "GetPlantsAgro",
			Handler:    _Master_GetPlantsAgro_Handler,
		},
		{
			MethodName: "GetOrgs",
			Handler:    _Master_GetOrgs_Handler,
		},
		{
			MethodName: "GetOrgsAgro",
			Handler:    _Master_GetOrgsAgro_Handler,
		},
		{
			MethodName: "GetSubOrgs",
			Handler:    _Master_GetSubOrgs_Handler,
		},
		{
			MethodName: "GetWinfeedProductComponentsWithSubIDs",
			Handler:    _Master_GetWinfeedProductComponentsWithSubIDs_Handler,
		},
		{
			MethodName: "GetWinfeedDivisionList",
			Handler:    _Master_GetWinfeedDivisionList_Handler,
		},
		{
			MethodName: "GetWinfeedDivisionListWithPermission",
			Handler:    _Master_GetWinfeedDivisionListWithPermission_Handler,
		},
		{
			MethodName: "GetWinfeedDivision",
			Handler:    _Master_GetWinfeedDivision_Handler,
		},
		{
			MethodName: "ListWinfeedProductByDivisionID",
			Handler:    _Master_ListWinfeedProductByDivisionID_Handler,
		},
		{
			MethodName: "ListWinfeedStock",
			Handler:    _Master_ListWinfeedStock_Handler,
		},
		{
			MethodName: "ListWinfeedPurchase",
			Handler:    _Master_ListWinfeedPurchase_Handler,
		},
		{
			MethodName: "ListWinfeedDailyPrice",
			Handler:    _Master_ListWinfeedDailyPrice_Handler,
		},
		{
			MethodName: "GetWinfeedGradeList",
			Handler:    _Master_GetWinfeedGradeList_Handler,
		},
		{
			MethodName: "ListWinfeedProduct",
			Handler:    _Master_ListWinfeedProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "samantra/services/master_service/master.proto",
}