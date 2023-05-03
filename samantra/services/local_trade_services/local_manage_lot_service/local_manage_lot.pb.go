// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: samantra/services/local_trade_services/local_manage_lot_service/local_manage_lot.proto

package local_manage_lot

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

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  float64 `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	UnitID int64   `protobuf:"varint,2,opt,name=UnitID,proto3" json:"UnitID,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[0]
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
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Amount) GetUnitID() int64 {
	if x != nil {
		return x.UnitID
	}
	return 0
}

type IncludingTransportView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantIDList []int64 `protobuf:"varint,1,rep,packed,name=PlantIDList,proto3" json:"PlantIDList,omitempty"`
	Cost        *Amount `protobuf:"bytes,2,opt,name=Cost,proto3" json:"Cost,omitempty"`
}

func (x *IncludingTransportView) Reset() {
	*x = IncludingTransportView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncludingTransportView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncludingTransportView) ProtoMessage() {}

func (x *IncludingTransportView) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncludingTransportView.ProtoReflect.Descriptor instead.
func (*IncludingTransportView) Descriptor() ([]byte, []int) {
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{1}
}

func (x *IncludingTransportView) GetPlantIDList() []int64 {
	if x != nil {
		return x.PlantIDList
	}
	return nil
}

func (x *IncludingTransportView) GetCost() *Amount {
	if x != nil {
		return x.Cost
	}
	return nil
}

type LocalLotForOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotNO             string  `protobuf:"bytes,1,opt,name=LotNO,proto3" json:"LotNO,omitempty"`
	UUID              string  `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ETA               string  `protobuf:"bytes,3,opt,name=ETA,proto3" json:"ETA,omitempty"`
	ETD               string  `protobuf:"bytes,4,opt,name=ETD,proto3" json:"ETD,omitempty"`
	Tonnes            float64 `protobuf:"fixed64,5,opt,name=Tonnes,proto3" json:"Tonnes,omitempty"`
	PlantTypeID       int64   `protobuf:"varint,6,opt,name=PlantTypeID,proto3" json:"PlantTypeID,omitempty"`
	PODNo             string  `protobuf:"bytes,7,opt,name=PODNo,proto3" json:"PODNo,omitempty"`
	BrokerContractNo  string  `protobuf:"bytes,8,opt,name=BrokerContractNo,proto3" json:"BrokerContractNo,omitempty"`
	PortName          string  `protobuf:"bytes,9,opt,name=PortName,proto3" json:"PortName,omitempty"`
	NumberOfContainer int64   `protobuf:"varint,10,opt,name=NumberOfContainer,proto3" json:"NumberOfContainer,omitempty"`
	SeaFreightID      int64   `protobuf:"varint,11,opt,name=SeaFreightID,proto3" json:"SeaFreightID,omitempty"`
}

func (x *LocalLotForOS) Reset() {
	*x = LocalLotForOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalLotForOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalLotForOS) ProtoMessage() {}

func (x *LocalLotForOS) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalLotForOS.ProtoReflect.Descriptor instead.
func (*LocalLotForOS) Descriptor() ([]byte, []int) {
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{2}
}

func (x *LocalLotForOS) GetLotNO() string {
	if x != nil {
		return x.LotNO
	}
	return ""
}

func (x *LocalLotForOS) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *LocalLotForOS) GetETA() string {
	if x != nil {
		return x.ETA
	}
	return ""
}

func (x *LocalLotForOS) GetETD() string {
	if x != nil {
		return x.ETD
	}
	return ""
}

func (x *LocalLotForOS) GetTonnes() float64 {
	if x != nil {
		return x.Tonnes
	}
	return 0
}

func (x *LocalLotForOS) GetPlantTypeID() int64 {
	if x != nil {
		return x.PlantTypeID
	}
	return 0
}

func (x *LocalLotForOS) GetPODNo() string {
	if x != nil {
		return x.PODNo
	}
	return ""
}

func (x *LocalLotForOS) GetBrokerContractNo() string {
	if x != nil {
		return x.BrokerContractNo
	}
	return ""
}

func (x *LocalLotForOS) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *LocalLotForOS) GetNumberOfContainer() int64 {
	if x != nil {
		return x.NumberOfContainer
	}
	return 0
}

func (x *LocalLotForOS) GetSeaFreightID() int64 {
	if x != nil {
		return x.SeaFreightID
	}
	return 0
}

type LocalManageLotForOSPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCodeList  []string `protobuf:"bytes,1,rep,name=ProductCodeList,proto3" json:"ProductCodeList,omitempty"`
	PlantTypeList    []int64  `protobuf:"varint,2,rep,packed,name=PlantTypeList,proto3" json:"PlantTypeList,omitempty"`
	ETAStartDateUnix int64    `protobuf:"varint,3,opt,name=ETAStartDateUnix,proto3" json:"ETAStartDateUnix,omitempty"`
	ETAEndDateUnix   int64    `protobuf:"varint,4,opt,name=ETAEndDateUnix,proto3" json:"ETAEndDateUnix,omitempty"`
}

func (x *LocalManageLotForOSPayload) Reset() {
	*x = LocalManageLotForOSPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalManageLotForOSPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalManageLotForOSPayload) ProtoMessage() {}

func (x *LocalManageLotForOSPayload) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalManageLotForOSPayload.ProtoReflect.Descriptor instead.
func (*LocalManageLotForOSPayload) Descriptor() ([]byte, []int) {
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{3}
}

func (x *LocalManageLotForOSPayload) GetProductCodeList() []string {
	if x != nil {
		return x.ProductCodeList
	}
	return nil
}

func (x *LocalManageLotForOSPayload) GetPlantTypeList() []int64 {
	if x != nil {
		return x.PlantTypeList
	}
	return nil
}

func (x *LocalManageLotForOSPayload) GetETAStartDateUnix() int64 {
	if x != nil {
		return x.ETAStartDateUnix
	}
	return 0
}

func (x *LocalManageLotForOSPayload) GetETAEndDateUnix() int64 {
	if x != nil {
		return x.ETAEndDateUnix
	}
	return 0
}

type LocalManageLotForOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCode        string                    `protobuf:"bytes,1,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	PlantTypeID        int64                     `protobuf:"varint,2,opt,name=PlantTypeID,proto3" json:"PlantTypeID,omitempty"`
	Price              *Amount                   `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
	ExchangeRate       *Amount                   `protobuf:"bytes,4,opt,name=ExchangeRate,proto3" json:"ExchangeRate,omitempty"`
	ExcludingTransport *Amount                   `protobuf:"bytes,5,opt,name=ExcludingTransport,proto3" json:"ExcludingTransport,omitempty"`
	IncludingTransport []*IncludingTransportView `protobuf:"bytes,6,rep,name=IncludingTransport,proto3" json:"IncludingTransport,omitempty"`
	Lots               []*LocalLotForOS          `protobuf:"bytes,7,rep,name=Lots,proto3" json:"Lots,omitempty"`
}

func (x *LocalManageLotForOS) Reset() {
	*x = LocalManageLotForOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalManageLotForOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalManageLotForOS) ProtoMessage() {}

func (x *LocalManageLotForOS) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalManageLotForOS.ProtoReflect.Descriptor instead.
func (*LocalManageLotForOS) Descriptor() ([]byte, []int) {
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{4}
}

func (x *LocalManageLotForOS) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *LocalManageLotForOS) GetPlantTypeID() int64 {
	if x != nil {
		return x.PlantTypeID
	}
	return 0
}

func (x *LocalManageLotForOS) GetPrice() *Amount {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *LocalManageLotForOS) GetExchangeRate() *Amount {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

func (x *LocalManageLotForOS) GetExcludingTransport() *Amount {
	if x != nil {
		return x.ExcludingTransport
	}
	return nil
}

func (x *LocalManageLotForOS) GetIncludingTransport() []*IncludingTransportView {
	if x != nil {
		return x.IncludingTransport
	}
	return nil
}

func (x *LocalManageLotForOS) GetLots() []*LocalLotForOS {
	if x != nil {
		return x.Lots
	}
	return nil
}

type LocalManageLotForOSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalManageLotForOSList []*LocalManageLotForOS `protobuf:"bytes,1,rep,name=LocalManageLotForOSList,proto3" json:"LocalManageLotForOSList,omitempty"`
}

func (x *LocalManageLotForOSResponse) Reset() {
	*x = LocalManageLotForOSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalManageLotForOSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalManageLotForOSResponse) ProtoMessage() {}

func (x *LocalManageLotForOSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalManageLotForOSResponse.ProtoReflect.Descriptor instead.
func (*LocalManageLotForOSResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP(), []int{5}
}

func (x *LocalManageLotForOSResponse) GetLocalManageLotForOSList() []*LocalManageLotForOS {
	if x != nil {
		return x.LocalManageLotForOSList
	}
	return nil
}

var File_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto protoreflect.FileDescriptor

var file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDesc = []byte{
	0x0a, 0x56, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x22, 0x36, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x6e, 0x69, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x44,
	0x22, 0x86, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x04, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61,
	0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x22, 0xc7, 0x02, 0x0a, 0x0d, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x6f, 0x74, 0x4e, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74, 0x4e,
	0x4f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x54, 0x41, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x45, 0x54, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x54, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x54, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x6e,
	0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x54, 0x6f, 0x6e, 0x6e, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x4f, 0x44, 0x4e, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x50, 0x4f, 0x44, 0x4e, 0x6f, 0x12, 0x2a, 0x0a, 0x10, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x53, 0x65, 0x61, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x49, 0x44, 0x22, 0xc0, 0x01, 0x0a, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x54, 0x41, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x45, 0x54,
	0x41, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x12, 0x26,
	0x0a, 0x0e, 0x45, 0x54, 0x41, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x45, 0x54, 0x41, 0x45, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x22, 0xb6, 0x04, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x12, 0x4c, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c,
	0x6f, 0x74, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0c,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x12,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e,
	0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x12, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x76, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f,
	0x74, 0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x12, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x51, 0x0a, 0x04,
	0x4c, 0x6f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x73, 0x61, 0x6d,
	0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x52, 0x04, 0x4c, 0x6f, 0x74, 0x73, 0x22,
	0x9c, 0x01, 0x0a, 0x1b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c,
	0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x7d, 0x0a, 0x17, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x43, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f,
	0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74,
	0x46, 0x6f, 0x72, 0x4f, 0x53, 0x52, 0x17, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xfe,
	0x02, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x74, 0x12, 0xb4, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x57, 0x61, 0x79, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x4a, 0x2e,
	0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72,
	0x4f, 0x53, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x4b, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb4, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x57, 0x61, 0x79, 0x42, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x4a, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x4c, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x53, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x4b, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f,
	0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x74,
	0x46, 0x6f, 0x72, 0x4f, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x88, 0x01, 0x5a, 0x85, 0x01, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescOnce sync.Once
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescData = file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDesc
)

func file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescGZIP() []byte {
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescOnce.Do(func() {
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescData)
	})
	return file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDescData
}

var file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_goTypes = []interface{}{
	(*Amount)(nil),                      // 0: samantra.services.local_trade.local_manage_lot.Amount
	(*IncludingTransportView)(nil),      // 1: samantra.services.local_trade.local_manage_lot.IncludingTransportView
	(*LocalLotForOS)(nil),               // 2: samantra.services.local_trade.local_manage_lot.LocalLotForOS
	(*LocalManageLotForOSPayload)(nil),  // 3: samantra.services.local_trade.local_manage_lot.LocalManageLotForOSPayload
	(*LocalManageLotForOS)(nil),         // 4: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS
	(*LocalManageLotForOSResponse)(nil), // 5: samantra.services.local_trade.local_manage_lot.LocalManageLotForOSResponse
}
var file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_depIdxs = []int32{
	0, // 0: samantra.services.local_trade.local_manage_lot.IncludingTransportView.Cost:type_name -> samantra.services.local_trade.local_manage_lot.Amount
	0, // 1: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS.Price:type_name -> samantra.services.local_trade.local_manage_lot.Amount
	0, // 2: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS.ExchangeRate:type_name -> samantra.services.local_trade.local_manage_lot.Amount
	0, // 3: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS.ExcludingTransport:type_name -> samantra.services.local_trade.local_manage_lot.Amount
	1, // 4: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS.IncludingTransport:type_name -> samantra.services.local_trade.local_manage_lot.IncludingTransportView
	2, // 5: samantra.services.local_trade.local_manage_lot.LocalManageLotForOS.Lots:type_name -> samantra.services.local_trade.local_manage_lot.LocalLotForOS
	4, // 6: samantra.services.local_trade.local_manage_lot.LocalManageLotForOSResponse.LocalManageLotForOSList:type_name -> samantra.services.local_trade.local_manage_lot.LocalManageLotForOS
	3, // 7: samantra.services.local_trade.local_manage_lot.LocalManageLot.GetInterTradeWayByParam:input_type -> samantra.services.local_trade.local_manage_lot.LocalManageLotForOSPayload
	3, // 8: samantra.services.local_trade.local_manage_lot.LocalManageLot.GetLocalTradeWayByParam:input_type -> samantra.services.local_trade.local_manage_lot.LocalManageLotForOSPayload
	5, // 9: samantra.services.local_trade.local_manage_lot.LocalManageLot.GetInterTradeWayByParam:output_type -> samantra.services.local_trade.local_manage_lot.LocalManageLotForOSResponse
	5, // 10: samantra.services.local_trade.local_manage_lot.LocalManageLot.GetLocalTradeWayByParam:output_type -> samantra.services.local_trade.local_manage_lot.LocalManageLotForOSResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_init()
}
func file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_init() {
	if File_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncludingTransportView); i {
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
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalLotForOS); i {
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
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalManageLotForOSPayload); i {
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
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalManageLotForOS); i {
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
		file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalManageLotForOSResponse); i {
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
			RawDescriptor: file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_goTypes,
		DependencyIndexes: file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_depIdxs,
		MessageInfos:      file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_msgTypes,
	}.Build()
	File_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto = out.File
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_rawDesc = nil
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_goTypes = nil
	file_samantra_services_local_trade_services_local_manage_lot_service_local_manage_lot_proto_depIdxs = nil
}
