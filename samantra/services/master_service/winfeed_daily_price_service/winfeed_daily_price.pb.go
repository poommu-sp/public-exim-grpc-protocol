// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: samantra/services/master_service/winfeed_daily_price_service/winfeed_daily_price.proto

package winfeed_daily_price

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

type WinfeedDailyPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceivedDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ReceivedDate,proto3" json:"ReceivedDate,omitempty"`
	ReceivedTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ReceivedTime,proto3" json:"ReceivedTime,omitempty"`
	PlantCode    string                 `protobuf:"bytes,3,opt,name=PlantCode,proto3" json:"PlantCode,omitempty"`
	PlantName    string                 `protobuf:"bytes,4,opt,name=PlantName,proto3" json:"PlantName,omitempty"`
	OrgCode      string                 `protobuf:"bytes,5,opt,name=OrgCode,proto3" json:"OrgCode,omitempty"`
	OrgName      string                 `protobuf:"bytes,6,opt,name=OrgName,proto3" json:"OrgName,omitempty"`
	ProductGroup string                 `protobuf:"bytes,7,opt,name=ProductGroup,proto3" json:"ProductGroup,omitempty"`
	ProductCode  string                 `protobuf:"bytes,8,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	LotNO        string                 `protobuf:"bytes,9,opt,name=LotNO,proto3" json:"LotNO,omitempty"`
	DivisionCode string                 `protobuf:"bytes,10,opt,name=DivisionCode,proto3" json:"DivisionCode,omitempty"`
	UnitPrice    float64                `protobuf:"fixed64,11,opt,name=UnitPrice,proto3" json:"UnitPrice,omitempty"`
	RequestDate  *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=RequestDate,proto3" json:"RequestDate,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
}

func (x *WinfeedDailyPrice) Reset() {
	*x = WinfeedDailyPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinfeedDailyPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinfeedDailyPrice) ProtoMessage() {}

func (x *WinfeedDailyPrice) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinfeedDailyPrice.ProtoReflect.Descriptor instead.
func (*WinfeedDailyPrice) Descriptor() ([]byte, []int) {
	return file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescGZIP(), []int{0}
}

func (x *WinfeedDailyPrice) GetReceivedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedDate
	}
	return nil
}

func (x *WinfeedDailyPrice) GetReceivedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedTime
	}
	return nil
}

func (x *WinfeedDailyPrice) GetPlantCode() string {
	if x != nil {
		return x.PlantCode
	}
	return ""
}

func (x *WinfeedDailyPrice) GetPlantName() string {
	if x != nil {
		return x.PlantName
	}
	return ""
}

func (x *WinfeedDailyPrice) GetOrgCode() string {
	if x != nil {
		return x.OrgCode
	}
	return ""
}

func (x *WinfeedDailyPrice) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *WinfeedDailyPrice) GetProductGroup() string {
	if x != nil {
		return x.ProductGroup
	}
	return ""
}

func (x *WinfeedDailyPrice) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *WinfeedDailyPrice) GetLotNO() string {
	if x != nil {
		return x.LotNO
	}
	return ""
}

func (x *WinfeedDailyPrice) GetDivisionCode() string {
	if x != nil {
		return x.DivisionCode
	}
	return ""
}

func (x *WinfeedDailyPrice) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *WinfeedDailyPrice) GetRequestDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestDate
	}
	return nil
}

func (x *WinfeedDailyPrice) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

var File_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto protoreflect.FileDescriptor

var file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDesc = []byte{
	0x0a, 0x56, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x77,
	0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x77, 0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x11, 0x57, 0x69, 0x6e, 0x66,
	0x65, 0x65, 0x64, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x67, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74, 0x4e, 0x4f, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x42, 0x88, 0x01, 0x5a, 0x85, 0x01, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x73, 0x61,
	0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x77,
	0x69, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x77, 0x69, 0x6e, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescOnce sync.Once
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescData = file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDesc
)

func file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescGZIP() []byte {
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescOnce.Do(func() {
		file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescData)
	})
	return file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDescData
}

var file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_goTypes = []interface{}{
	(*WinfeedDailyPrice)(nil),     // 0: samantra.services.master.winfeed_daily_price.WinfeedDailyPrice
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_depIdxs = []int32{
	1, // 0: samantra.services.master.winfeed_daily_price.WinfeedDailyPrice.ReceivedDate:type_name -> google.protobuf.Timestamp
	1, // 1: samantra.services.master.winfeed_daily_price.WinfeedDailyPrice.ReceivedTime:type_name -> google.protobuf.Timestamp
	1, // 2: samantra.services.master.winfeed_daily_price.WinfeedDailyPrice.RequestDate:type_name -> google.protobuf.Timestamp
	1, // 3: samantra.services.master.winfeed_daily_price.WinfeedDailyPrice.StartDate:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_init()
}
func file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_init() {
	if File_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinfeedDailyPrice); i {
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
			RawDescriptor: file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_goTypes,
		DependencyIndexes: file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_depIdxs,
		MessageInfos:      file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_msgTypes,
	}.Build()
	File_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto = out.File
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_rawDesc = nil
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_goTypes = nil
	file_samantra_services_master_service_winfeed_daily_price_service_winfeed_daily_price_proto_depIdxs = nil
}