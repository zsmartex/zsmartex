// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: services/v1/code.proto

package servicesv1

import (
	v1 "github.com/zsmartex/zsmartex/proto/common/v1"
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

type GenerateCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeType v1.CodeType `protobuf:"varint,1,opt,name=code_type,proto3,enum=common.v1.CodeType" json:"code_type,omitempty"`
}

func (x *GenerateCodeRequest) Reset() {
	*x = GenerateCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeRequest) ProtoMessage() {}

func (x *GenerateCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCodeRequest.ProtoReflect.Descriptor instead.
func (*GenerateCodeRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateCodeRequest) GetCodeType() v1.CodeType {
	if x != nil {
		return x.CodeType
	}
	return v1.CodeType(0)
}

type GenerateCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateCodeResponse) Reset() {
	*x = GenerateCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeResponse) ProtoMessage() {}

func (x *GenerateCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCodeResponse.ProtoReflect.Descriptor instead.
func (*GenerateCodeResponse) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{1}
}

var File_services_v1_code_proto protoreflect.FileDescriptor

var file_services_v1_code_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x13, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x62, 0x0a,
	0x0b, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xa4, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x7a, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x78,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_v1_code_proto_rawDescOnce sync.Once
	file_services_v1_code_proto_rawDescData = file_services_v1_code_proto_rawDesc
)

func file_services_v1_code_proto_rawDescGZIP() []byte {
	file_services_v1_code_proto_rawDescOnce.Do(func() {
		file_services_v1_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_v1_code_proto_rawDescData)
	})
	return file_services_v1_code_proto_rawDescData
}

var file_services_v1_code_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_v1_code_proto_goTypes = []interface{}{
	(*GenerateCodeRequest)(nil),  // 0: services.v1.GenerateCodeRequest
	(*GenerateCodeResponse)(nil), // 1: services.v1.GenerateCodeResponse
	(v1.CodeType)(0),             // 2: common.v1.CodeType
}
var file_services_v1_code_proto_depIdxs = []int32{
	2, // 0: services.v1.GenerateCodeRequest.code_type:type_name -> common.v1.CodeType
	0, // 1: services.v1.CodeService.GenerateCode:input_type -> services.v1.GenerateCodeRequest
	1, // 2: services.v1.CodeService.GenerateCode:output_type -> services.v1.GenerateCodeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_v1_code_proto_init() }
func file_services_v1_code_proto_init() {
	if File_services_v1_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_v1_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCodeRequest); i {
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
		file_services_v1_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCodeResponse); i {
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
			RawDescriptor: file_services_v1_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_v1_code_proto_goTypes,
		DependencyIndexes: file_services_v1_code_proto_depIdxs,
		MessageInfos:      file_services_v1_code_proto_msgTypes,
	}.Build()
	File_services_v1_code_proto = out.File
	file_services_v1_code_proto_rawDesc = nil
	file_services_v1_code_proto_goTypes = nil
	file_services_v1_code_proto_depIdxs = nil
}
