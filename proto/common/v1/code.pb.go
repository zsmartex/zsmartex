// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/v1/code.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code_Type int32

const (
	Code_TYPE_UNKNOWN  Code_Type = 0
	Code_TYPE_EMAIL    Code_Type = 1
	Code_TYPE_PASSWORD Code_Type = 2
)

// Enum value maps for Code_Type.
var (
	Code_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_EMAIL",
		2: "TYPE_PASSWORD",
	}
	Code_Type_value = map[string]int32{
		"TYPE_UNKNOWN":  0,
		"TYPE_EMAIL":    1,
		"TYPE_PASSWORD": 2,
	}
)

func (x Code_Type) Enum() *Code_Type {
	p := new(Code_Type)
	*p = x
	return p
}

func (x Code_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_code_proto_enumTypes[0].Descriptor()
}

func (Code_Type) Type() protoreflect.EnumType {
	return &file_common_v1_code_proto_enumTypes[0]
}

func (x Code_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code_Type.Descriptor instead.
func (Code_Type) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_code_proto_rawDescGZIP(), []int{0, 0}
}

type Code_Category int32

const (
	Code_CATEGORY_UNKNOWN Code_Category = 0
)

// Enum value maps for Code_Category.
var (
	Code_Category_name = map[int32]string{
		0: "CATEGORY_UNKNOWN",
	}
	Code_Category_value = map[string]int32{
		"CATEGORY_UNKNOWN": 0,
	}
)

func (x Code_Category) Enum() *Code_Category {
	p := new(Code_Category)
	*p = x
	return p
}

func (x Code_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_code_proto_enumTypes[1].Descriptor()
}

func (Code_Category) Type() protoreflect.EnumType {
	return &file_common_v1_code_proto_enumTypes[1]
}

func (x Code_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code_Category.Descriptor instead.
func (Code_Category) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_code_proto_rawDescGZIP(), []int{0, 1}
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email    *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone    *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Code     string                  `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Type     Code_Type               `protobuf:"varint,5,opt,name=type,proto3,enum=common.v1.Code_Type" json:"type,omitempty"`
	Category Code_Category           `protobuf:"varint,6,opt,name=category,proto3,enum=common.v1.Code_Category" json:"category,omitempty"`
	Data     []byte                  `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_common_v1_code_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Code) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *Code) GetPhone() *wrapperspb.StringValue {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Code) GetType() Code_Type {
	if x != nil {
		return x.Type
	}
	return Code_TYPE_UNKNOWN
}

func (x *Code) GetCategory() Code_Category {
	if x != nil {
		return x.Category
	}
	return Code_CATEGORY_UNKNOWN
}

func (x *Code) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_v1_code_proto protoreflect.FileDescriptor

var file_common_v1_code_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xee, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02,
	0x22, 0x20, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x42, 0x96, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x7a, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_v1_code_proto_rawDescOnce sync.Once
	file_common_v1_code_proto_rawDescData = file_common_v1_code_proto_rawDesc
)

func file_common_v1_code_proto_rawDescGZIP() []byte {
	file_common_v1_code_proto_rawDescOnce.Do(func() {
		file_common_v1_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_code_proto_rawDescData)
	})
	return file_common_v1_code_proto_rawDescData
}

var file_common_v1_code_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_v1_code_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_v1_code_proto_goTypes = []interface{}{
	(Code_Type)(0),                 // 0: common.v1.Code.Type
	(Code_Category)(0),             // 1: common.v1.Code.Category
	(*Code)(nil),                   // 2: common.v1.Code
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_common_v1_code_proto_depIdxs = []int32{
	3, // 0: common.v1.Code.email:type_name -> google.protobuf.StringValue
	3, // 1: common.v1.Code.phone:type_name -> google.protobuf.StringValue
	0, // 2: common.v1.Code.type:type_name -> common.v1.Code.Type
	1, // 3: common.v1.Code.category:type_name -> common.v1.Code.Category
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_v1_code_proto_init() }
func file_common_v1_code_proto_init() {
	if File_common_v1_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
			RawDescriptor: file_common_v1_code_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_code_proto_goTypes,
		DependencyIndexes: file_common_v1_code_proto_depIdxs,
		EnumInfos:         file_common_v1_code_proto_enumTypes,
		MessageInfos:      file_common_v1_code_proto_msgTypes,
	}.Build()
	File_common_v1_code_proto = out.File
	file_common_v1_code_proto_rawDesc = nil
	file_common_v1_code_proto_goTypes = nil
	file_common_v1_code_proto_depIdxs = nil
}