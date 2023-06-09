// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: samantra/common/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{0}
}

type NullFloat64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Float64 float64 `protobuf:"fixed64,1,opt,name=Float64,proto3" json:"Float64,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *NullFloat64) Reset() {
	*x = NullFloat64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NullFloat64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NullFloat64) ProtoMessage() {}

func (x *NullFloat64) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NullFloat64.ProtoReflect.Descriptor instead.
func (*NullFloat64) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *NullFloat64) GetFloat64() float64 {
	if x != nil {
		return x.Float64
	}
	return 0
}

func (x *NullFloat64) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type NullInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int64 int64 `protobuf:"varint,1,opt,name=Int64,proto3" json:"Int64,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *NullInt64) Reset() {
	*x = NullInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NullInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NullInt64) ProtoMessage() {}

func (x *NullInt64) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NullInt64.ProtoReflect.Descriptor instead.
func (*NullInt64) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *NullInt64) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *NullInt64) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type NullString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string `protobuf:"bytes,1,opt,name=String,proto3" json:"String,omitempty"`
	Valid   bool   `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *NullString) Reset() {
	*x = NullString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NullString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NullString) ProtoMessage() {}

func (x *NullString) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NullString.ProtoReflect.Descriptor instead.
func (*NullString) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *NullString) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *NullString) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Type      int64  `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[4]
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
	return file_samantra_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *Unit) GetID() int64 {
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

func (x *Unit) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *NullFloat64 `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Unit  *Unit        `protobuf:"bytes,2,opt,name=Unit,proto3" json:"Unit,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[5]
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
	return file_samantra_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *Amount) GetValue() *NullFloat64 {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Amount) GetUnit() *Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type PlantQuantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Quantity    *Amount `protobuf:"bytes,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	QuotaWeight float32 `protobuf:"fixed32,3,opt,name=QuotaWeight,proto3" json:"QuotaWeight,omitempty"`
	QuotaQty    float32 `protobuf:"fixed32,4,opt,name=QuotaQty,proto3" json:"QuotaQty,omitempty"`
}

func (x *PlantQuantity) Reset() {
	*x = PlantQuantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantQuantity) ProtoMessage() {}

func (x *PlantQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantQuantity.ProtoReflect.Descriptor instead.
func (*PlantQuantity) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *PlantQuantity) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PlantQuantity) GetQuantity() *Amount {
	if x != nil {
		return x.Quantity
	}
	return nil
}

func (x *PlantQuantity) GetQuotaWeight() float32 {
	if x != nil {
		return x.QuotaWeight
	}
	return 0
}

func (x *PlantQuantity) GetQuotaQty() float32 {
	if x != nil {
		return x.QuotaQty
	}
	return 0
}

type PlantQuantities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantQuantities []*PlantQuantity `protobuf:"bytes,1,rep,name=PlantQuantities,proto3" json:"PlantQuantities,omitempty"`
}

func (x *PlantQuantities) Reset() {
	*x = PlantQuantities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_common_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantQuantities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantQuantities) ProtoMessage() {}

func (x *PlantQuantities) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_common_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantQuantities.ProtoReflect.Descriptor instead.
func (*PlantQuantities) Descriptor() ([]byte, []int) {
	return file_samantra_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *PlantQuantities) GetPlantQuantities() []*PlantQuantity {
	if x != nil {
		return x.PlantQuantities
	}
	return nil
}

var File_samantra_common_common_proto protoreflect.FileDescriptor

var file_samantra_common_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x0a, 0x0b, 0x4e, 0x75, 0x6c, 0x6c,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36,
	0x34, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0x3a, 0x0a, 0x0a, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x67, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36,
	0x34, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x55,
	0x6e, 0x69, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x51, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x51, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0f, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0x51, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74,
	0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_common_common_proto_rawDescOnce sync.Once
	file_samantra_common_common_proto_rawDescData = file_samantra_common_common_proto_rawDesc
)

func file_samantra_common_common_proto_rawDescGZIP() []byte {
	file_samantra_common_common_proto_rawDescOnce.Do(func() {
		file_samantra_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_common_common_proto_rawDescData)
	})
	return file_samantra_common_common_proto_rawDescData
}

var file_samantra_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_samantra_common_common_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: samantra.common.Empty
	(*NullFloat64)(nil),     // 1: samantra.common.NullFloat64
	(*NullInt64)(nil),       // 2: samantra.common.NullInt64
	(*NullString)(nil),      // 3: samantra.common.NullString
	(*Unit)(nil),            // 4: samantra.common.Unit
	(*Amount)(nil),          // 5: samantra.common.Amount
	(*PlantQuantity)(nil),   // 6: samantra.common.PlantQuantity
	(*PlantQuantities)(nil), // 7: samantra.common.PlantQuantities
}
var file_samantra_common_common_proto_depIdxs = []int32{
	1, // 0: samantra.common.Amount.Value:type_name -> samantra.common.NullFloat64
	4, // 1: samantra.common.Amount.Unit:type_name -> samantra.common.Unit
	5, // 2: samantra.common.PlantQuantity.Quantity:type_name -> samantra.common.Amount
	6, // 3: samantra.common.PlantQuantities.PlantQuantities:type_name -> samantra.common.PlantQuantity
	0, // 4: samantra.common.EmptyStandIn.EmptyFunction:input_type -> samantra.common.Empty
	0, // 5: samantra.common.EmptyStandIn.EmptyFunction:output_type -> samantra.common.Empty
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_samantra_common_common_proto_init() }
func file_samantra_common_common_proto_init() {
	if File_samantra_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_samantra_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NullFloat64); i {
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
		file_samantra_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NullInt64); i {
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
		file_samantra_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NullString); i {
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
		file_samantra_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_samantra_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_samantra_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantQuantity); i {
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
		file_samantra_common_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantQuantities); i {
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
			RawDescriptor: file_samantra_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_common_common_proto_goTypes,
		DependencyIndexes: file_samantra_common_common_proto_depIdxs,
		MessageInfos:      file_samantra_common_common_proto_msgTypes,
	}.Build()
	File_samantra_common_common_proto = out.File
	file_samantra_common_common_proto_rawDesc = nil
	file_samantra_common_common_proto_goTypes = nil
	file_samantra_common_common_proto_depIdxs = nil
}
