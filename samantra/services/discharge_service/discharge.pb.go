// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: samantra/services/discharge_service/discharge.proto

package discharge

import (
	master_service "github.com/poommu-sp/public-exim-grpc-protocol/samantra/services/master_service"
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

type ListDischargePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=RequestDate,proto3" json:"RequestDate,omitempty"`
}

func (x *ListDischargePayload) Reset() {
	*x = ListDischargePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDischargePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDischargePayload) ProtoMessage() {}

func (x *ListDischargePayload) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDischargePayload.ProtoReflect.Descriptor instead.
func (*ListDischargePayload) Descriptor() ([]byte, []int) {
	return file_samantra_services_discharge_service_discharge_proto_rawDescGZIP(), []int{0}
}

func (x *ListDischargePayload) GetRequestDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestDate
	}
	return nil
}

type ListDischargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListDischarge []*DischargeItem `protobuf:"bytes,1,rep,name=ListDischarge,proto3" json:"ListDischarge,omitempty"`
}

func (x *ListDischargeResponse) Reset() {
	*x = ListDischargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDischargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDischargeResponse) ProtoMessage() {}

func (x *ListDischargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDischargeResponse.ProtoReflect.Descriptor instead.
func (*ListDischargeResponse) Descriptor() ([]byte, []int) {
	return file_samantra_services_discharge_service_discharge_proto_rawDescGZIP(), []int{1}
}

func (x *ListDischargeResponse) GetListDischarge() []*DischargeItem {
	if x != nil {
		return x.ListDischarge
	}
	return nil
}

type DischargeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotNO       string                `protobuf:"bytes,1,opt,name=LotNO,proto3" json:"LotNO,omitempty"`
	ProductCode string                `protobuf:"bytes,2,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	Plant       *master_service.Plant `protobuf:"bytes,3,opt,name=Plant,proto3" json:"Plant,omitempty"`
	QuotaWGH    float64               `protobuf:"fixed64,4,opt,name=QuotaWGH,proto3" json:"QuotaWGH,omitempty"`
}

func (x *DischargeItem) Reset() {
	*x = DischargeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DischargeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DischargeItem) ProtoMessage() {}

func (x *DischargeItem) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_discharge_service_discharge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DischargeItem.ProtoReflect.Descriptor instead.
func (*DischargeItem) Descriptor() ([]byte, []int) {
	return file_samantra_services_discharge_service_discharge_proto_rawDescGZIP(), []int{2}
}

func (x *DischargeItem) GetLotNO() string {
	if x != nil {
		return x.LotNO
	}
	return ""
}

func (x *DischargeItem) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *DischargeItem) GetPlant() *master_service.Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

func (x *DischargeItem) GetQuotaWGH() float64 {
	if x != nil {
		return x.QuotaWGH
	}
	return 0
}

var File_samantra_services_discharge_service_discharge_proto protoreflect.FileDescriptor

var file_samantra_services_discharge_service_discharge_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x69, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e,
	0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a,
	0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73,
	0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x57, 0x47, 0x48,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x57, 0x47, 0x48,
	0x32, 0x85, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x78,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12,
	0x31, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74,
	0x2f, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_discharge_service_discharge_proto_rawDescOnce sync.Once
	file_samantra_services_discharge_service_discharge_proto_rawDescData = file_samantra_services_discharge_service_discharge_proto_rawDesc
)

func file_samantra_services_discharge_service_discharge_proto_rawDescGZIP() []byte {
	file_samantra_services_discharge_service_discharge_proto_rawDescOnce.Do(func() {
		file_samantra_services_discharge_service_discharge_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_discharge_service_discharge_proto_rawDescData)
	})
	return file_samantra_services_discharge_service_discharge_proto_rawDescData
}

var file_samantra_services_discharge_service_discharge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_samantra_services_discharge_service_discharge_proto_goTypes = []interface{}{
	(*ListDischargePayload)(nil),  // 0: samantra.services.discharge.ListDischargePayload
	(*ListDischargeResponse)(nil), // 1: samantra.services.discharge.ListDischargeResponse
	(*DischargeItem)(nil),         // 2: samantra.services.discharge.DischargeItem
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*master_service.Plant)(nil),  // 4: samantra.services.master.Plant
}
var file_samantra_services_discharge_service_discharge_proto_depIdxs = []int32{
	3, // 0: samantra.services.discharge.ListDischargePayload.RequestDate:type_name -> google.protobuf.Timestamp
	2, // 1: samantra.services.discharge.ListDischargeResponse.ListDischarge:type_name -> samantra.services.discharge.DischargeItem
	4, // 2: samantra.services.discharge.DischargeItem.Plant:type_name -> samantra.services.master.Plant
	0, // 3: samantra.services.discharge.Discharge.ListDischarge:input_type -> samantra.services.discharge.ListDischargePayload
	1, // 4: samantra.services.discharge.Discharge.ListDischarge:output_type -> samantra.services.discharge.ListDischargeResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_samantra_services_discharge_service_discharge_proto_init() }
func file_samantra_services_discharge_service_discharge_proto_init() {
	if File_samantra_services_discharge_service_discharge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_discharge_service_discharge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDischargePayload); i {
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
		file_samantra_services_discharge_service_discharge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDischargeResponse); i {
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
		file_samantra_services_discharge_service_discharge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DischargeItem); i {
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
			RawDescriptor: file_samantra_services_discharge_service_discharge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samantra_services_discharge_service_discharge_proto_goTypes,
		DependencyIndexes: file_samantra_services_discharge_service_discharge_proto_depIdxs,
		MessageInfos:      file_samantra_services_discharge_service_discharge_proto_msgTypes,
	}.Build()
	File_samantra_services_discharge_service_discharge_proto = out.File
	file_samantra_services_discharge_service_discharge_proto_rawDesc = nil
	file_samantra_services_discharge_service_discharge_proto_goTypes = nil
	file_samantra_services_discharge_service_discharge_proto_depIdxs = nil
}
