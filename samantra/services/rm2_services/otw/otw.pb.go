// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: samantra/services/rm2_services/otw/otw.proto

package otw

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/poommu-sp/public-exim-grpc-protocol/samantra/common"
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

type OTW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotNo      string                  `protobuf:"bytes,1,opt,name=lot_no,json=lotNo,proto3" json:"lot_no,omitempty"`
	Quantities []*common.PlantQuantity `protobuf:"bytes,2,rep,name=Quantities,proto3" json:"Quantities,omitempty"`
}

func (x *OTW) Reset() {
	*x = OTW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samantra_services_rm2_services_otw_otw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTW) ProtoMessage() {}

func (x *OTW) ProtoReflect() protoreflect.Message {
	mi := &file_samantra_services_rm2_services_otw_otw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTW.ProtoReflect.Descriptor instead.
func (*OTW) Descriptor() ([]byte, []int) {
	return file_samantra_services_rm2_services_otw_otw_proto_rawDescGZIP(), []int{0}
}

func (x *OTW) GetLotNo() string {
	if x != nil {
		return x.LotNo
	}
	return ""
}

func (x *OTW) GetQuantities() []*common.PlantQuantity {
	if x != nil {
		return x.Quantities
	}
	return nil
}

var File_samantra_services_rm2_services_otw_otw_proto protoreflect.FileDescriptor

var file_samantra_services_rm2_services_otw_otw_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6f, 0x74, 0x77, 0x2f, 0x6f, 0x74, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x73, 0x61, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x72, 0x6d, 0x32, 0x2e, 0x6f, 0x74, 0x77, 0x1a, 0x1c, 0x73, 0x61, 0x6d, 0x61, 0x6e,
	0x74, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x03, 0x4f, 0x54, 0x57, 0x12, 0x15,
	0x0a, 0x06, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x74, 0x4e, 0x6f, 0x12, 0x3e, 0x0a, 0x0a, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x6d, 0x61,
	0x6e, 0x74, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x74, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x73, 0x61,
	0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x72, 0x6d, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x74, 0x77,
	0x3b, 0x6f, 0x74, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samantra_services_rm2_services_otw_otw_proto_rawDescOnce sync.Once
	file_samantra_services_rm2_services_otw_otw_proto_rawDescData = file_samantra_services_rm2_services_otw_otw_proto_rawDesc
)

func file_samantra_services_rm2_services_otw_otw_proto_rawDescGZIP() []byte {
	file_samantra_services_rm2_services_otw_otw_proto_rawDescOnce.Do(func() {
		file_samantra_services_rm2_services_otw_otw_proto_rawDescData = protoimpl.X.CompressGZIP(file_samantra_services_rm2_services_otw_otw_proto_rawDescData)
	})
	return file_samantra_services_rm2_services_otw_otw_proto_rawDescData
}

var file_samantra_services_rm2_services_otw_otw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_samantra_services_rm2_services_otw_otw_proto_goTypes = []interface{}{
	(*OTW)(nil),                  // 0: samantra.services.rm2.otw.OTW
	(*common.PlantQuantity)(nil), // 1: samantra.common.PlantQuantity
}
var file_samantra_services_rm2_services_otw_otw_proto_depIdxs = []int32{
	1, // 0: samantra.services.rm2.otw.OTW.Quantities:type_name -> samantra.common.PlantQuantity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_samantra_services_rm2_services_otw_otw_proto_init() }
func file_samantra_services_rm2_services_otw_otw_proto_init() {
	if File_samantra_services_rm2_services_otw_otw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samantra_services_rm2_services_otw_otw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTW); i {
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
			RawDescriptor: file_samantra_services_rm2_services_otw_otw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_samantra_services_rm2_services_otw_otw_proto_goTypes,
		DependencyIndexes: file_samantra_services_rm2_services_otw_otw_proto_depIdxs,
		MessageInfos:      file_samantra_services_rm2_services_otw_otw_proto_msgTypes,
	}.Build()
	File_samantra_services_rm2_services_otw_otw_proto = out.File
	file_samantra_services_rm2_services_otw_otw_proto_rawDesc = nil
	file_samantra_services_rm2_services_otw_otw_proto_goTypes = nil
	file_samantra_services_rm2_services_otw_otw_proto_depIdxs = nil
}
