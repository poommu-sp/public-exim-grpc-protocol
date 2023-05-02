// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: daily_price/daily_price.proto

package daily_price

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[0]
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
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{0}
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{1}
}

func (x *GetByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DailyPriceDocument *DailyPriceDocument `protobuf:"bytes,1,opt,name=daily_price_document,json=dailyPriceDocument,proto3" json:"daily_price_document,omitempty"`
}

func (x *GetByIdResponse) Reset() {
	*x = GetByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResponse) ProtoMessage() {}

func (x *GetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResponse.ProtoReflect.Descriptor instead.
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIdResponse) GetDailyPriceDocument() *DailyPriceDocument {
	if x != nil {
		return x.DailyPriceDocument
	}
	return nil
}

type ApproveAndFreezeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApproveAndFreezeRequest) Reset() {
	*x = ApproveAndFreezeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveAndFreezeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveAndFreezeRequest) ProtoMessage() {}

func (x *ApproveAndFreezeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveAndFreezeRequest.ProtoReflect.Descriptor instead.
func (*ApproveAndFreezeRequest) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveAndFreezeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DailyPriceDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy       string                 `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedBy       string                 `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	No              string                 `protobuf:"bytes,7,opt,name=no,proto3" json:"no,omitempty"`
	EffectiveDate   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=effective_date,json=effectiveDate,proto3" json:"effective_date,omitempty"`
	WinfeedDivision *WinfeedDivision       `protobuf:"bytes,9,opt,name=WinfeedDivision,proto3" json:"WinfeedDivision,omitempty"`
	Status          int32                  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DailyPriceDocument) Reset() {
	*x = DailyPriceDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyPriceDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyPriceDocument) ProtoMessage() {}

func (x *DailyPriceDocument) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyPriceDocument.ProtoReflect.Descriptor instead.
func (*DailyPriceDocument) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{4}
}

func (x *DailyPriceDocument) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DailyPriceDocument) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DailyPriceDocument) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DailyPriceDocument) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DailyPriceDocument) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *DailyPriceDocument) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DailyPriceDocument) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *DailyPriceDocument) GetEffectiveDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EffectiveDate
	}
	return nil
}

func (x *DailyPriceDocument) GetWinfeedDivision() *WinfeedDivision {
	if x != nil {
		return x.WinfeedDivision
	}
	return nil
}

func (x *DailyPriceDocument) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type WinfeedDivision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CountryAlpha_3 string `protobuf:"bytes,3,opt,name=country_alpha_3,json=countryAlpha3,proto3" json:"country_alpha_3,omitempty"`
	CountryName    string `protobuf:"bytes,4,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	Code           string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *WinfeedDivision) Reset() {
	*x = WinfeedDivision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinfeedDivision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinfeedDivision) ProtoMessage() {}

func (x *WinfeedDivision) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinfeedDivision.ProtoReflect.Descriptor instead.
func (*WinfeedDivision) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{5}
}

func (x *WinfeedDivision) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WinfeedDivision) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WinfeedDivision) GetCountryAlpha_3() string {
	if x != nil {
		return x.CountryAlpha_3
	}
	return ""
}

func (x *WinfeedDivision) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *WinfeedDivision) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RejectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RejectRequest) Reset() {
	*x = RejectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daily_price_daily_price_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectRequest) ProtoMessage() {}

func (x *RejectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daily_price_daily_price_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectRequest.ProtoReflect.Descriptor instead.
func (*RejectRequest) Descriptor() ([]byte, []int) {
	return file_daily_price_daily_price_proto_rawDescGZIP(), []int{6}
}

func (x *RejectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_daily_price_daily_price_proto protoreflect.FileDescriptor

var file_daily_price_daily_price_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x29,
	0x0a, 0x17, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc6, 0x03, 0x0a, 0x12, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12,
	0x41, 0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x57, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x44, 0x69, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x69, 0x6e, 0x66, 0x65, 0x65,
	0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x57, 0x69, 0x6e, 0x66, 0x65,
	0x65, 0x64, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x44, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x33, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x33, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc5, 0x03, 0x0a, 0x0a, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x64, 0x46,
	0x72, 0x65, 0x65, 0x7a, 0x65, 0x12, 0x24, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72,
	0x65, 0x65, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x18, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x12, 0x24, 0x2e,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x12, 0x53, 0x65, 0x6e,
	0x64, 0x56, 0x69, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1b, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daily_price_daily_price_proto_rawDescOnce sync.Once
	file_daily_price_daily_price_proto_rawDescData = file_daily_price_daily_price_proto_rawDesc
)

func file_daily_price_daily_price_proto_rawDescGZIP() []byte {
	file_daily_price_daily_price_proto_rawDescOnce.Do(func() {
		file_daily_price_daily_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_daily_price_daily_price_proto_rawDescData)
	})
	return file_daily_price_daily_price_proto_rawDescData
}

var file_daily_price_daily_price_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_daily_price_daily_price_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: daily_price.Empty
	(*GetByIdRequest)(nil),          // 1: daily_price.GetByIdRequest
	(*GetByIdResponse)(nil),         // 2: daily_price.GetByIdResponse
	(*ApproveAndFreezeRequest)(nil), // 3: daily_price.ApproveAndFreezeRequest
	(*DailyPriceDocument)(nil),      // 4: daily_price.DailyPriceDocument
	(*WinfeedDivision)(nil),         // 5: daily_price.WinfeedDivision
	(*RejectRequest)(nil),           // 6: daily_price.RejectRequest
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_daily_price_daily_price_proto_depIdxs = []int32{
	4,  // 0: daily_price.GetByIdResponse.daily_price_document:type_name -> daily_price.DailyPriceDocument
	7,  // 1: daily_price.DailyPriceDocument.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: daily_price.DailyPriceDocument.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 3: daily_price.DailyPriceDocument.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // 4: daily_price.DailyPriceDocument.effective_date:type_name -> google.protobuf.Timestamp
	5,  // 5: daily_price.DailyPriceDocument.WinfeedDivision:type_name -> daily_price.WinfeedDivision
	1,  // 6: daily_price.DailyPrice.GetById:input_type -> daily_price.GetByIdRequest
	3,  // 7: daily_price.DailyPrice.ApproveAndFreeze:input_type -> daily_price.ApproveAndFreezeRequest
	3,  // 8: daily_price.DailyPrice.RollbackApproveAndFreeze:input_type -> daily_price.ApproveAndFreezeRequest
	6,  // 9: daily_price.DailyPrice.Reject:input_type -> daily_price.RejectRequest
	6,  // 10: daily_price.DailyPrice.RollbackReject:input_type -> daily_price.RejectRequest
	1,  // 11: daily_price.DailyPrice.SendVietnamMsgById:input_type -> daily_price.GetByIdRequest
	2,  // 12: daily_price.DailyPrice.GetById:output_type -> daily_price.GetByIdResponse
	0,  // 13: daily_price.DailyPrice.ApproveAndFreeze:output_type -> daily_price.Empty
	0,  // 14: daily_price.DailyPrice.RollbackApproveAndFreeze:output_type -> daily_price.Empty
	0,  // 15: daily_price.DailyPrice.Reject:output_type -> daily_price.Empty
	0,  // 16: daily_price.DailyPrice.RollbackReject:output_type -> daily_price.Empty
	0,  // 17: daily_price.DailyPrice.SendVietnamMsgById:output_type -> daily_price.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_daily_price_daily_price_proto_init() }
func file_daily_price_daily_price_proto_init() {
	if File_daily_price_daily_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daily_price_daily_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_daily_price_daily_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_daily_price_daily_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdResponse); i {
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
		file_daily_price_daily_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveAndFreezeRequest); i {
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
		file_daily_price_daily_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyPriceDocument); i {
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
		file_daily_price_daily_price_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinfeedDivision); i {
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
		file_daily_price_daily_price_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectRequest); i {
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
			RawDescriptor: file_daily_price_daily_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daily_price_daily_price_proto_goTypes,
		DependencyIndexes: file_daily_price_daily_price_proto_depIdxs,
		MessageInfos:      file_daily_price_daily_price_proto_msgTypes,
	}.Build()
	File_daily_price_daily_price_proto = out.File
	file_daily_price_daily_price_proto_rawDesc = nil
	file_daily_price_daily_price_proto_goTypes = nil
	file_daily_price_daily_price_proto_depIdxs = nil
}