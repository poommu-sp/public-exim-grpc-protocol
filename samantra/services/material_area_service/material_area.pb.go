// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: samantra/services/material_area_service/material_area.proto

package material_area

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	master_service "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/master_service"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MaterialAreaPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date       *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	PlantCodes []string             `protobuf:"bytes,2,rep,name=PlantCodes,proto3" json:"PlantCodes,omitempty"`
}

func (x *MaterialAreaPayload) Reset() {
	*x = MaterialAreaPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialAreaPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialAreaPayload) ProtoMessage() {}

func (x *MaterialAreaPayload) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialAreaPayload.ProtoReflect.Descriptor instead.
func (*MaterialAreaPayload) Descriptor() ([]byte, []int) {
	return file_samantra_services_material_area_service_material_area_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialAreaPayload) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *MaterialAreaPayload) GetPlantCodes() []string {
	if x != nil {
		return x.PlantCodes
	}
	return nil
}

type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedBulk  float32 `protobuf:"fixed32,1,opt,name=NeedBulk,proto3" json:"NeedBulk,omitempty"`
	NeedBag   float32 `protobuf:"fixed32,2,opt,name=NeedBag,proto3" json:"NeedBag,omitempty"`
	ExistBulk float32 `protobuf:"fixed32,3,opt,name=ExistBulk,proto3" json:"ExistBulk,omitempty"`
	ExistBag  float32 `protobuf:"fixed32,4,opt,name=ExistBag,proto3" json:"ExistBag,omitempty"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_samantra_services_material_area_service_material_area_proto_rawDescGZIP(), []int{1}
}

func (x *Area) GetNeedBulk() float32 {
	if x != nil {
		return x.NeedBulk
	}
	return 0
}

func (x *Area) GetNeedBag() float32 {
	if x != nil {
		return x.NeedBag
	}
	return 0
}

func (x *Area) GetExistBulk() float32 {
	if x != nil {
		return x.ExistBulk
	}
	return 0
}

func (x *Area) GetExistBag() float32 {
	if x != nil {
		return x.ExistBag
	}
	return 0
}

type MaterialAreaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      *timestamp.Timestamp  `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Plant     *master_service.Plant `protobuf:"bytes,2,opt,name=Plant,proto3" json:"Plant,omitempty"`
	LimitBulk float32               `protobuf:"fixed32,3,opt,name=LimitBulk,proto3" json:"LimitBulk,omitempty"`
	LimitBag  float32               `protobuf:"fixed32,4,opt,name=LimitBag,proto3" json:"LimitBag,omitempty"`
	Areas     []*Area               `protobuf:"bytes,5,rep,name=Areas,proto3" json:"Areas,omitempty"`
}

func (x *MaterialAreaInfo) Reset() {
	*x = MaterialAreaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialAreaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialAreaInfo) ProtoMessage() {}

func (x *MaterialAreaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialAreaInfo.ProtoReflect.Descriptor instead.
func (*MaterialAreaInfo) Descriptor() ([]byte, []int) {
	return file_samantra_services_material_area_service_material_area_proto_rawDescGZIP(), []int{2}
}

func (x *MaterialAreaInfo) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *MaterialAreaInfo) GetPlant() *master_service.Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

func (x *MaterialAreaInfo) GetLimitBulk() float32 {
	if x != nil {
		return x.LimitBulk
	}
	return 0
}

func (x *MaterialAreaInfo) GetLimitBag() float32 {
	if x != nil {
		return x.LimitBag
	}
	return 0
}

func (x *MaterialAreaInfo) GetAreas() []*Area {
	if x != nil {
		return x.Areas
	}
	return nil
}

type MaterialAreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialAreas []*MaterialAreaInfo `protobuf:"bytes,1,rep,name=MaterialAreas,proto3" json:"MaterialAreas,omitempty"`
}

func (x *MaterialAreaResponse) Reset() {
	*x = MaterialAreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialAreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialAreaResponse) ProtoMessage() {}

func (x *MaterialAreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_material_area_service_material_area_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialAreaResponse.ProtoReflect.Descriptor instead.
func (*MaterialAreaResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_material_area_service_material_area_proto_rawDescGZIP(), []int{3}
}

func (x *MaterialAreaResponse) GetMaterialAreas() []*MaterialAreaInfo {
	if x != nil {
		return x.MaterialAreas
	}
	return nil
}

var File_samantra_services_material_area_service_material_area_proto protoreflect.FileDescriptor

var file_samantra_services_material_area_service_material_area_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x73,
	0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x13, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x65, 0x65, 0x64, 0x42, 0x75, 0x6c, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x4e, 0x65, 0x65, 0x64, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x65,
	0x64, 0x42, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x4e, 0x65, 0x65, 0x64,
	0x42, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x75, 0x6c, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x75, 0x6c,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x61, 0x67, 0x22, 0xf0, 0x01,
	0x0a, 0x10, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x61, 0x67, 0x12, 0x3b, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x05, 0x41, 0x72, 0x65, 0x61, 0x73,
	0x22, 0x6f, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61,
	0x73, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72,
	0x65, 0x61, 0x12, 0x93, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a,
	0x35, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6c, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74,
	0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_material_area_service_material_area_proto_rawDescOnce sync.Once
	file_samantra_services_material_area_service_material_area_proto_rawDescData = file_samantra_services_material_area_service_material_area_proto_rawDesc
)

func file_samantra_services_material_area_service_material_area_proto_rawDescGZIP() []byte {
	file_samantra_services_material_area_service_material_area_proto_rawDescOnce.Do(func() {
		file_samantra_services_material_area_service_material_area_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_material_area_service_material_area_proto_rawDescData)
	})
	return file_samantra_services_material_area_service_material_area_proto_rawDescData
}

var file_samantra_services_material_area_service_material_area_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_samantra_services_material_area_service_material_area_proto_goTypes = []interface{}{
	(*MaterialAreaPayload)(nil),  // 0: samantra.services.material_area.MaterialAreaPayload
	(*Area)(nil),                 // 1: samantra.services.material_area.Area
	(*MaterialAreaInfo)(nil),     // 2: samantra.services.material_area.MaterialAreaInfo
	(*MaterialAreaResponse)(nil), // 3: samantra.services.material_area.MaterialAreaResponse
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*master_service.Plant)(nil), // 5: samantra.services.master.Plant
}
var file_samantra_services_material_area_service_material_area_proto_depIdxs = []int32{
	4, // 0: samantra.services.material_area.MaterialAreaPayload.Date:type_name -> google.protobuf.Timestamp
	4, // 1: samantra.services.material_area.MaterialAreaInfo.Date:type_name -> google.protobuf.Timestamp
	5, // 2: samantra.services.material_area.MaterialAreaInfo.Plant:type_name -> samantra.services.master.Plant
	1, // 3: samantra.services.material_area.MaterialAreaInfo.Areas:type_name -> samantra.services.material_area.Area
	2, // 4: samantra.services.material_area.MaterialAreaResponse.MaterialAreas:type_name -> samantra.services.material_area.MaterialAreaInfo
	0, // 5: samantra.services.material_area.MaterialArea.GetMaterialAreaByDateAndPlantCodes:input_type -> samantra.services.material_area.MaterialAreaPayload
	3, // 6: samantra.services.material_area.MaterialArea.GetMaterialAreaByDateAndPlantCodes:output_type -> samantra.services.material_area.MaterialAreaResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_samantra_services_material_area_service_material_area_proto_init() }
func file_samantra_services_material_area_service_material_area_proto_init() {
	if File_samantra_services_material_area_service_material_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_material_area_service_material_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialAreaPayload); i {
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
		file_samantra_services_material_area_service_material_area_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		file_samantra_services_material_area_service_material_area_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialAreaInfo); i {
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
		file_samantra_services_material_area_service_material_area_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialAreaResponse); i {
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
			RawDescriptor: file_samantra_services_material_area_service_material_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_services_material_area_service_material_area_proto_goTypes,
		DependencyIndexes: file_samantra_services_material_area_service_material_area_proto_depIdxs,
		MessageInfos:      file_samantra_services_material_area_service_material_area_proto_msgTypes,
	}.Build()
	File_samantra_services_material_area_service_material_area_proto = out.File
	file_samantra_services_material_area_service_material_area_proto_rawDesc = nil
	file_samantra_services_material_area_service_material_area_proto_goTypes = nil
	file_samantra_services_material_area_service_material_area_proto_depIdxs = nil
}
