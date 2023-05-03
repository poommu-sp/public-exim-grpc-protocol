// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: samantra/services/rm2_services/rm2.proto

package rm2

import (
	proto "github.com/golang/protobuf/proto"
	otw "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/rm2_services/otw"
	phantha "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/rm2_services/phantha"
	processing "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/rm2_services/processing"
	warehouse "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/rm2_services/warehouse"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListWarehouseSummaryByLotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ProductCodes []string               `protobuf:"bytes,2,rep,name=product_codes,json=productCodes,proto3" json:"product_codes,omitempty"`
}

func (x *ListWarehouseSummaryByLotRequest) Reset() {
	*x = ListWarehouseSummaryByLotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWarehouseSummaryByLotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWarehouseSummaryByLotRequest) ProtoMessage() {}

func (x *ListWarehouseSummaryByLotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWarehouseSummaryByLotRequest.ProtoReflect.Descriptor instead.
func (*ListWarehouseSummaryByLotRequest) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{0}
}

func (x *ListWarehouseSummaryByLotRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListWarehouseSummaryByLotRequest) GetProductCodes() []string {
	if x != nil {
		return x.ProductCodes
	}
	return nil
}

type ListWarehouseSummaryByLotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouses []*warehouse.Warehouse `protobuf:"bytes,1,rep,name=warehouses,proto3" json:"warehouses,omitempty"`
}

func (x *ListWarehouseSummaryByLotResponse) Reset() {
	*x = ListWarehouseSummaryByLotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWarehouseSummaryByLotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWarehouseSummaryByLotResponse) ProtoMessage() {}

func (x *ListWarehouseSummaryByLotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWarehouseSummaryByLotResponse.ProtoReflect.Descriptor instead.
func (*ListWarehouseSummaryByLotResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{1}
}

func (x *ListWarehouseSummaryByLotResponse) GetWarehouses() []*warehouse.Warehouse {
	if x != nil {
		return x.Warehouses
	}
	return nil
}

type ListPhanthaOTWRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ProductCodes []string               `protobuf:"bytes,2,rep,name=product_codes,json=productCodes,proto3" json:"product_codes,omitempty"`
}

func (x *ListPhanthaOTWRequest) Reset() {
	*x = ListPhanthaOTWRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhanthaOTWRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhanthaOTWRequest) ProtoMessage() {}

func (x *ListPhanthaOTWRequest) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhanthaOTWRequest.ProtoReflect.Descriptor instead.
func (*ListPhanthaOTWRequest) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{2}
}

func (x *ListPhanthaOTWRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListPhanthaOTWRequest) GetProductCodes() []string {
	if x != nil {
		return x.ProductCodes
	}
	return nil
}

type ListPhanthaOTWResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phantas []*phantha.Phantha `protobuf:"bytes,1,rep,name=phantas,proto3" json:"phantas,omitempty"`
	Otws    []*otw.OTW         `protobuf:"bytes,2,rep,name=otws,proto3" json:"otws,omitempty"`
}

func (x *ListPhanthaOTWResponse) Reset() {
	*x = ListPhanthaOTWResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhanthaOTWResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhanthaOTWResponse) ProtoMessage() {}

func (x *ListPhanthaOTWResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhanthaOTWResponse.ProtoReflect.Descriptor instead.
func (*ListPhanthaOTWResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{3}
}

func (x *ListPhanthaOTWResponse) GetPhantas() []*phantha.Phantha {
	if x != nil {
		return x.Phantas
	}
	return nil
}

func (x *ListPhanthaOTWResponse) GetOtws() []*otw.OTW {
	if x != nil {
		return x.Otws
	}
	return nil
}

type ListProcessingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ProductCodes []string               `protobuf:"bytes,2,rep,name=product_codes,json=productCodes,proto3" json:"product_codes,omitempty"`
}

func (x *ListProcessingRequest) Reset() {
	*x = ListProcessingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProcessingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProcessingRequest) ProtoMessage() {}

func (x *ListProcessingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProcessingRequest.ProtoReflect.Descriptor instead.
func (*ListProcessingRequest) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{4}
}

func (x *ListProcessingRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListProcessingRequest) GetProductCodes() []string {
	if x != nil {
		return x.ProductCodes
	}
	return nil
}

type ListProcessingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processings []*processing.Processing `protobuf:"bytes,1,rep,name=processings,proto3" json:"processings,omitempty"`
}

func (x *ListProcessingResponse) Reset() {
	*x = ListProcessingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProcessingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProcessingResponse) ProtoMessage() {}

func (x *ListProcessingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_rm2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProcessingResponse.ProtoReflect.Descriptor instead.
func (*ListProcessingResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_rm2_proto_rawDescGZIP(), []int{5}
}

func (x *ListProcessingResponse) GetProcessings() []*processing.Processing {
	if x != nil {
		return x.Processings
	}
	return nil
}

var File_samantra_services_rm2_services_rm2_proto protoreflect.FileDescriptor

var file_samantra_services_rm2_services_rm2_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x72, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d,
	0x32, 0x1a, 0x34, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x68,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x74, 0x77, 0x2f, 0x6f, 0x74, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x20,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x42, 0x79, 0x4c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x79, 0x4c,
	0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x0a, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68,
	0x61, 0x6e, 0x74, 0x68, 0x61, 0x4f, 0x54, 0x57, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x61,
	0x6e, 0x74, 0x68, 0x61, 0x4f, 0x54, 0x57, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x07, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61,
	0x2e, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61, 0x52, 0x07, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x61,
	0x73, 0x12, 0x32, 0x0a, 0x04, 0x6f, 0x74, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x6f, 0x74, 0x77, 0x2e, 0x4f, 0x54, 0x57, 0x52,
	0x04, 0x6f, 0x74, 0x77, 0x73, 0x22, 0x6c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0x68, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xfb, 0x02,
	0x0a, 0x03, 0x52, 0x4d, 0x32, 0x12, 0x90, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x79,
	0x4c, 0x6f, 0x74, 0x12, 0x37, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x42, 0x79, 0x4c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73,
	0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x79, 0x4c, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61, 0x4f, 0x54, 0x57, 0x12, 0x2c, 0x2e, 0x73, 0x61, 0x6d,
	0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72,
	0x6d, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61, 0x4f, 0x54,
	0x57, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e,
	0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x6d, 0x32,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x68, 0x61, 0x4f, 0x54, 0x57, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c, 0x2e, 0x73,
	0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x61, 0x6d,
	0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72,
	0x6d, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x59, 0x5a, 0x57, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x67, 0x69, 0x74, 0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x6d, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_rm2_services_rm2_proto_rawDescOnce sync.Once
	file_samantra_services_rm2_services_rm2_proto_rawDescData = file_samantra_services_rm2_services_rm2_proto_rawDesc
)

func file_samantra_services_rm2_services_rm2_proto_rawDescGZIP() []byte {
	file_samantra_services_rm2_services_rm2_proto_rawDescOnce.Do(func() {
		file_samantra_services_rm2_services_rm2_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_rm2_services_rm2_proto_rawDescData)
	})
	return file_samantra_services_rm2_services_rm2_proto_rawDescData
}

var file_samantra_services_rm2_services_rm2_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_samantra_services_rm2_services_rm2_proto_goTypes = []interface{}{
	(*ListWarehouseSummaryByLotRequest)(nil),  // 0: samantra.services.rm2.ListWarehouseSummaryByLotRequest
	(*ListWarehouseSummaryByLotResponse)(nil), // 1: samantra.services.rm2.ListWarehouseSummaryByLotResponse
	(*ListPhanthaOTWRequest)(nil),             // 2: samantra.services.rm2.ListPhanthaOTWRequest
	(*ListPhanthaOTWResponse)(nil),            // 3: samantra.services.rm2.ListPhanthaOTWResponse
	(*ListProcessingRequest)(nil),             // 4: samantra.services.rm2.ListProcessingRequest
	(*ListProcessingResponse)(nil),            // 5: samantra.services.rm2.ListProcessingResponse
	(*timestamppb.Timestamp)(nil),             // 6: google.protobuf.Timestamp
	(*warehouse.Warehouse)(nil),               // 7: samantra.services.rm2.warehouse.Warehouse
	(*phantha.Phantha)(nil),                   // 8: samantra.services.rm2.phantha.Phantha
	(*otw.OTW)(nil),                           // 9: samantra.services.rm2.otw.OTW
	(*processing.Processing)(nil),             // 10: samantra.services.rm2.processing.Processing
}
var file_samantra_services_rm2_services_rm2_proto_depIdxs = []int32{
	6,  // 0: samantra.services.rm2.ListWarehouseSummaryByLotRequest.date:type_name -> google.protobuf.Timestamp
	7,  // 1: samantra.services.rm2.ListWarehouseSummaryByLotResponse.warehouses:type_name -> samantra.services.rm2.warehouse.Warehouse
	6,  // 2: samantra.services.rm2.ListPhanthaOTWRequest.date:type_name -> google.protobuf.Timestamp
	8,  // 3: samantra.services.rm2.ListPhanthaOTWResponse.phantas:type_name -> samantra.services.rm2.phantha.Phantha
	9,  // 4: samantra.services.rm2.ListPhanthaOTWResponse.otws:type_name -> samantra.services.rm2.otw.OTW
	6,  // 5: samantra.services.rm2.ListProcessingRequest.date:type_name -> google.protobuf.Timestamp
	10, // 6: samantra.services.rm2.ListProcessingResponse.processings:type_name -> samantra.services.rm2.processing.Processing
	0,  // 7: samantra.services.rm2.RM2.ListWarehouseSummaryByLot:input_type -> samantra.services.rm2.ListWarehouseSummaryByLotRequest
	2,  // 8: samantra.services.rm2.RM2.ListPhanthaOTW:input_type -> samantra.services.rm2.ListPhanthaOTWRequest
	4,  // 9: samantra.services.rm2.RM2.ListProcessings:input_type -> samantra.services.rm2.ListProcessingRequest
	1,  // 10: samantra.services.rm2.RM2.ListWarehouseSummaryByLot:output_type -> samantra.services.rm2.ListWarehouseSummaryByLotResponse
	3,  // 11: samantra.services.rm2.RM2.ListPhanthaOTW:output_type -> samantra.services.rm2.ListPhanthaOTWResponse
	5,  // 12: samantra.services.rm2.RM2.ListProcessings:output_type -> samantra.services.rm2.ListProcessingResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_samantra_services_rm2_services_rm2_proto_init() }
func file_samantra_services_rm2_services_rm2_proto_init() {
	if File_samantra_services_rm2_services_rm2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_rm2_services_rm2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWarehouseSummaryByLotRequest); i {
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
		file_samantra_services_rm2_services_rm2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWarehouseSummaryByLotResponse); i {
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
		file_samantra_services_rm2_services_rm2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhanthaOTWRequest); i {
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
		file_samantra_services_rm2_services_rm2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhanthaOTWResponse); i {
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
		file_samantra_services_rm2_services_rm2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProcessingRequest); i {
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
		file_samantra_services_rm2_services_rm2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProcessingResponse); i {
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
			RawDescriptor: file_samantra_services_rm2_services_rm2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_services_rm2_services_rm2_proto_goTypes,
		DependencyIndexes: file_samantra_services_rm2_services_rm2_proto_depIdxs,
		MessageInfos:      file_samantra_services_rm2_services_rm2_proto_msgTypes,
	}.Build()
	File_samantra_services_rm2_services_rm2_proto = out.File
	file_samantra_services_rm2_services_rm2_proto_rawDesc = nil
	file_samantra_services_rm2_services_rm2_proto_goTypes = nil
	file_samantra_services_rm2_services_rm2_proto_depIdxs = nil
}
