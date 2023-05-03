// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: samantra/services/winfeed_services/warehouse_stock_report_service/warehouse_stock_report.proto

package warehouse_stock_report

import (
	common "github.com/poommu-sp/public-exim-grpc-protocol/samantra/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListWarehouseStockReportPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=RequestDate,proto3" json:"RequestDate,omitempty"`
}

func (x *ListWarehouseStockReportPayload) Reset() {
	*x = ListWarehouseStockReportPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWarehouseStockReportPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWarehouseStockReportPayload) ProtoMessage() {}

func (x *ListWarehouseStockReportPayload) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWarehouseStockReportPayload.ProtoReflect.Descriptor instead.
func (*ListWarehouseStockReportPayload) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{0}
}

func (x *ListWarehouseStockReportPayload) GetRequestDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestDate
	}
	return nil
}

type ListWarehouseStockReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListWarehouseStockReport []*WarehouseStock `protobuf:"bytes,1,rep,name=ListWarehouseStockReport,proto3" json:"ListWarehouseStockReport,omitempty"`
}

func (x *ListWarehouseStockReportResponse) Reset() {
	*x = ListWarehouseStockReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWarehouseStockReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWarehouseStockReportResponse) ProtoMessage() {}

func (x *ListWarehouseStockReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWarehouseStockReportResponse.ProtoReflect.Descriptor instead.
func (*ListWarehouseStockReportResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{1}
}

func (x *ListWarehouseStockReportResponse) GetListWarehouseStockReport() []*WarehouseStock {
	if x != nil {
		return x.ListWarehouseStockReport
	}
	return nil
}

type WarehouseStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameInReport            string                  `protobuf:"bytes,1,opt,name=NameInReport,proto3" json:"NameInReport,omitempty"`
	SubNameInReport         string                  `protobuf:"bytes,2,opt,name=SubNameInReport,proto3" json:"SubNameInReport,omitempty"`
	LotNO                   string                  `protobuf:"bytes,3,opt,name=LotNO,proto3" json:"LotNO,omitempty"`
	Warehouse               string                  `protobuf:"bytes,4,opt,name=Warehouse,proto3" json:"Warehouse,omitempty"`
	Age                     int32                   `protobuf:"varint,5,opt,name=Age,proto3" json:"Age,omitempty"`
	Hold                    float64                 `protobuf:"fixed64,6,opt,name=Hold,proto3" json:"Hold,omitempty"`
	PlantQuantities         []*common.PlantQuantity `protobuf:"bytes,7,rep,name=PlantQuantities,proto3" json:"PlantQuantities,omitempty"`
	ProductCode             string                  `protobuf:"bytes,8,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	WarehouseSmokedDate     *timestamppb.Timestamp  `protobuf:"bytes,9,opt,name=WarehouseSmokedDate,proto3" json:"WarehouseSmokedDate,omitempty"`
	WarehouseDGasDate       *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=WarehouseDGasDate,proto3" json:"WarehouseDGasDate,omitempty"`
	WarehouseLastedDGasDate *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=WarehouseLastedDGasDate,proto3" json:"WarehouseLastedDGasDate,omitempty"`
	Remark                  string                  `protobuf:"bytes,12,opt,name=Remark,proto3" json:"Remark,omitempty"`
}

func (x *WarehouseStock) Reset() {
	*x = WarehouseStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseStock) ProtoMessage() {}

func (x *WarehouseStock) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseStock.ProtoReflect.Descriptor instead.
func (*WarehouseStock) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{2}
}

func (x *WarehouseStock) GetNameInReport() string {
	if x != nil {
		return x.NameInReport
	}
	return ""
}

func (x *WarehouseStock) GetSubNameInReport() string {
	if x != nil {
		return x.SubNameInReport
	}
	return ""
}

func (x *WarehouseStock) GetLotNO() string {
	if x != nil {
		return x.LotNO
	}
	return ""
}

func (x *WarehouseStock) GetWarehouse() string {
	if x != nil {
		return x.Warehouse
	}
	return ""
}

func (x *WarehouseStock) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *WarehouseStock) GetHold() float64 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *WarehouseStock) GetPlantQuantities() []*common.PlantQuantity {
	if x != nil {
		return x.PlantQuantities
	}
	return nil
}

func (x *WarehouseStock) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *WarehouseStock) GetWarehouseSmokedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.WarehouseSmokedDate
	}
	return nil
}

func (x *WarehouseStock) GetWarehouseDGasDate() *timestamppb.Timestamp {
	if x != nil {
		return x.WarehouseDGasDate
	}
	return nil
}

func (x *WarehouseStock) GetWarehouseLastedDGasDate() *timestamppb.Timestamp {
	if x != nil {
		return x.WarehouseLastedDGasDate
	}
	return nil
}

func (x *WarehouseStock) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type Quantities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Quantity *Amount `protobuf:"bytes,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *Quantities) Reset() {
	*x = Quantities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quantities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quantities) ProtoMessage() {}

func (x *Quantities) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quantities.ProtoReflect.Descriptor instead.
func (*Quantities) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{3}
}

func (x *Quantities) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Quantities) GetQuantity() *Amount {
	if x != nil {
		return x.Quantity
	}
	return nil
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Unit  *Unit   `protobuf:"bytes,2,opt,name=Unit,proto3" json:"Unit,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{4}
}

func (x *Amount) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Amount) GetUnit() *Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Type      int32  `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP(), []int{5}
}

func (x *Unit) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Unit) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto protoreflect.FileDescriptor

var file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x30, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xaa, 0x04, 0x0a, 0x0e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x53, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x12, 0x1c, 0x0a, 0x09,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x6f, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x48, 0x6f, 0x6c, 0x64,
	0x12, 0x48, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4c, 0x0a, 0x13,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x6d, 0x6f, 0x6b, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x53, 0x6d, 0x6f, 0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x47, 0x61, 0x73, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x11, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x47, 0x61, 0x73,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x64, 0x44, 0x47, 0x61, 0x73, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x17, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4c, 0x61, 0x73, 0x74,
	0x65, 0x64, 0x44, 0x47, 0x61, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x22, 0x72, 0x0a, 0x0a, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x54, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x6a, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x55, 0x6e,
	0x69, 0x74, 0x22, 0x5c, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x32, 0xdc, 0x01, 0x0a, 0x14, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0xc3, 0x01, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x51, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x52, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x90, 0x01, 0x5a, 0x8d, 0x01, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x6e, 0x66,
	0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescOnce sync.Once
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescData = file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDesc
)

func file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescGZIP() []byte {
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescOnce.Do(func() {
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescData)
	})
	return file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDescData
}

var file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_goTypes = []interface{}{
	(*ListWarehouseStockReportPayload)(nil),  // 0: samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportPayload
	(*ListWarehouseStockReportResponse)(nil), // 1: samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportResponse
	(*WarehouseStock)(nil),                   // 2: samantra.services.winfeed.warehouse_stock_report.WarehouseStock
	(*Quantities)(nil),                       // 3: samantra.services.winfeed.warehouse_stock_report.Quantities
	(*Amount)(nil),                           // 4: samantra.services.winfeed.warehouse_stock_report.Amount
	(*Unit)(nil),                             // 5: samantra.services.winfeed.warehouse_stock_report.Unit
	(*timestamppb.Timestamp)(nil),            // 6: google.protobuf.Timestamp
	(*common.PlantQuantity)(nil),             // 7: samantra.common.PlantQuantity
}
var file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_depIdxs = []int32{
	6, // 0: samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportPayload.RequestDate:type_name -> google.protobuf.Timestamp
	2, // 1: samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportResponse.ListWarehouseStockReport:type_name -> samantra.services.winfeed.warehouse_stock_report.WarehouseStock
	7, // 2: samantra.services.winfeed.warehouse_stock_report.WarehouseStock.PlantQuantities:type_name -> samantra.common.PlantQuantity
	6, // 3: samantra.services.winfeed.warehouse_stock_report.WarehouseStock.WarehouseSmokedDate:type_name -> google.protobuf.Timestamp
	6, // 4: samantra.services.winfeed.warehouse_stock_report.WarehouseStock.WarehouseDGasDate:type_name -> google.protobuf.Timestamp
	6, // 5: samantra.services.winfeed.warehouse_stock_report.WarehouseStock.WarehouseLastedDGasDate:type_name -> google.protobuf.Timestamp
	4, // 6: samantra.services.winfeed.warehouse_stock_report.Quantities.Quantity:type_name -> samantra.services.winfeed.warehouse_stock_report.Amount
	5, // 7: samantra.services.winfeed.warehouse_stock_report.Amount.Unit:type_name -> samantra.services.winfeed.warehouse_stock_report.Unit
	0, // 8: samantra.services.winfeed.warehouse_stock_report.WarehouseStockReport.ListWarehouseStockReport:input_type -> samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportPayload
	1, // 9: samantra.services.winfeed.warehouse_stock_report.WarehouseStockReport.ListWarehouseStockReport:output_type -> samantra.services.winfeed.warehouse_stock_report.ListWarehouseStockReportResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_init()
}
func file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_init() {
	if File_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWarehouseStockReportPayload); i {
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
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWarehouseStockReportResponse); i {
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
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseStock); i {
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
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quantities); i {
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
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
			RawDescriptor: file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_goTypes,
		DependencyIndexes: file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_depIdxs,
		MessageInfos:      file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_msgTypes,
	}.Build()
	File_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto = out.File
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_rawDesc = nil
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_goTypes = nil
	file_samantra_services_winfeed_services_warehouse_stock_report_service_warehouse_stock_report_proto_depIdxs = nil
}
