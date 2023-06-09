// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: samantra/services/master_service/winfeed_stock_service/winfeed_stock.proto

package winfeed_stock

import (
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

type WinfeedStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ReceivedDate     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ReceivedDate,proto3" json:"ReceivedDate,omitempty"`
	ReceivedTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ReceivedTime,proto3" json:"ReceivedTime,omitempty"`
	OrgCode          string                 `protobuf:"bytes,4,opt,name=OrgCode,proto3" json:"OrgCode,omitempty"`
	StockDate        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=StockDate,proto3" json:"StockDate,omitempty"`
	Time             *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=Time,proto3" json:"Time,omitempty"`
	ProductGroup     string                 `protobuf:"bytes,7,opt,name=ProductGroup,proto3" json:"ProductGroup,omitempty"`
	ProductCode      string                 `protobuf:"bytes,8,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	GradeCode        string                 `protobuf:"bytes,9,opt,name=GradeCode,proto3" json:"GradeCode,omitempty"`
	LotNO            string                 `protobuf:"bytes,10,opt,name=LotNO,proto3" json:"LotNO,omitempty"`
	PurchaseWeight   float64                `protobuf:"fixed64,11,opt,name=PurchaseWeight,proto3" json:"PurchaseWeight,omitempty"`
	TransferRWeight  float64                `protobuf:"fixed64,12,opt,name=TransferRWeight,proto3" json:"TransferRWeight,omitempty"`
	ReturnRWeight    float64                `protobuf:"fixed64,13,opt,name=ReturnRWeight,proto3" json:"ReturnRWeight,omitempty"`
	TransformRWeight float64                `protobuf:"fixed64,14,opt,name=TransformRWeight,proto3" json:"TransformRWeight,omitempty"`
	Adjust           float64                `protobuf:"fixed64,15,opt,name=Adjust,proto3" json:"Adjust,omitempty"`
	IssuedWeight     float64                `protobuf:"fixed64,16,opt,name=IssuedWeight,proto3" json:"IssuedWeight,omitempty"`
	SaleWeight       float64                `protobuf:"fixed64,17,opt,name=SaleWeight,proto3" json:"SaleWeight,omitempty"`
	TransferIWeight  float64                `protobuf:"fixed64,18,opt,name=TransferIWeight,proto3" json:"TransferIWeight,omitempty"`
	ReturnIWeight    float64                `protobuf:"fixed64,19,opt,name=ReturnIWeight,proto3" json:"ReturnIWeight,omitempty"`
	TransformIWeight float64                `protobuf:"fixed64,20,opt,name=TransformIWeight,proto3" json:"TransformIWeight,omitempty"`
	DamageWeight     float64                `protobuf:"fixed64,21,opt,name=DamageWeight,proto3" json:"DamageWeight,omitempty"`
	WeightTotal      float64                `protobuf:"fixed64,22,opt,name=WeightTotal,proto3" json:"WeightTotal,omitempty"`
}

func (x *WinfeedStock) Reset() {
	*x = WinfeedStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinfeedStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinfeedStock) ProtoMessage() {}

func (x *WinfeedStock) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinfeedStock.ProtoReflect.Descriptor instead.
func (*WinfeedStock) Descriptor() ([]byte, []int) {
	return file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescGZIP(), []int{0}
}

func (x *WinfeedStock) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *WinfeedStock) GetReceivedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedDate
	}
	return nil
}

func (x *WinfeedStock) GetReceivedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedTime
	}
	return nil
}

func (x *WinfeedStock) GetOrgCode() string {
	if x != nil {
		return x.OrgCode
	}
	return ""
}

func (x *WinfeedStock) GetStockDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StockDate
	}
	return nil
}

func (x *WinfeedStock) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *WinfeedStock) GetProductGroup() string {
	if x != nil {
		return x.ProductGroup
	}
	return ""
}

func (x *WinfeedStock) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *WinfeedStock) GetGradeCode() string {
	if x != nil {
		return x.GradeCode
	}
	return ""
}

func (x *WinfeedStock) GetLotNO() string {
	if x != nil {
		return x.LotNO
	}
	return ""
}

func (x *WinfeedStock) GetPurchaseWeight() float64 {
	if x != nil {
		return x.PurchaseWeight
	}
	return 0
}

func (x *WinfeedStock) GetTransferRWeight() float64 {
	if x != nil {
		return x.TransferRWeight
	}
	return 0
}

func (x *WinfeedStock) GetReturnRWeight() float64 {
	if x != nil {
		return x.ReturnRWeight
	}
	return 0
}

func (x *WinfeedStock) GetTransformRWeight() float64 {
	if x != nil {
		return x.TransformRWeight
	}
	return 0
}

func (x *WinfeedStock) GetAdjust() float64 {
	if x != nil {
		return x.Adjust
	}
	return 0
}

func (x *WinfeedStock) GetIssuedWeight() float64 {
	if x != nil {
		return x.IssuedWeight
	}
	return 0
}

func (x *WinfeedStock) GetSaleWeight() float64 {
	if x != nil {
		return x.SaleWeight
	}
	return 0
}

func (x *WinfeedStock) GetTransferIWeight() float64 {
	if x != nil {
		return x.TransferIWeight
	}
	return 0
}

func (x *WinfeedStock) GetReturnIWeight() float64 {
	if x != nil {
		return x.ReturnIWeight
	}
	return 0
}

func (x *WinfeedStock) GetTransformIWeight() float64 {
	if x != nil {
		return x.TransformIWeight
	}
	return 0
}

func (x *WinfeedStock) GetDamageWeight() float64 {
	if x != nil {
		return x.DamageWeight
	}
	return 0
}

func (x *WinfeedStock) GetWeightTotal() float64 {
	if x != nil {
		return x.WeightTotal
	}
	return 0
}

var File_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto protoreflect.FileDescriptor

var file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x73, 0x61,
	0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x06, 0x0a, 0x0c, 0x57, 0x69, 0x6e, 0x66, 0x65, 0x65,
	0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x67, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c,
	0x6f, 0x74, 0x4e, 0x4f, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x64, 0x6a, 0x75,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x7b, 0x5a, 0x79, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x73,
	0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescOnce sync.Once
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescData = file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDesc
)

func file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescGZIP() []byte {
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescOnce.Do(func() {
		file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescData)
	})
	return file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDescData
}

var file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_goTypes = []interface{}{
	(*WinfeedStock)(nil),          // 0: samantra.services.master.winfeed_stock.WinfeedStock
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_depIdxs = []int32{
	1, // 0: samantra.services.master.winfeed_stock.WinfeedStock.ReceivedDate:type_name -> google.protobuf.Timestamp
	1, // 1: samantra.services.master.winfeed_stock.WinfeedStock.ReceivedTime:type_name -> google.protobuf.Timestamp
	1, // 2: samantra.services.master.winfeed_stock.WinfeedStock.StockDate:type_name -> google.protobuf.Timestamp
	1, // 3: samantra.services.master.winfeed_stock.WinfeedStock.Time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_init() }
func file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_init() {
	if File_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinfeedStock); i {
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
			RawDescriptor: file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_goTypes,
		DependencyIndexes: file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_depIdxs,
		MessageInfos:      file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_msgTypes,
	}.Build()
	File_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto = out.File
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_rawDesc = nil
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_goTypes = nil
	file_samantra_services_master_service_winfeed_stock_service_winfeed_stock_proto_depIdxs = nil
}
