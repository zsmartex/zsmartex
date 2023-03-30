// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: common/code/v1/code.proto

package codev1

import (
	resource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	_ "github.com/zsmartex/protoc-gen-gorm/gormpb"
	v1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
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

type CodeType int32

const (
	CodeType_email CodeType = 0
	CodeType_phone CodeType = 1
)

// Enum value maps for CodeType.
var (
	CodeType_name = map[int32]string{
		0: "email",
		1: "phone",
	}
	CodeType_value = map[string]int32{
		"email": 0,
		"phone": 1,
	}
)

func (x CodeType) Enum() *CodeType {
	p := new(CodeType)
	*p = x
	return p
}

func (x CodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_code_v1_code_proto_enumTypes[0].Descriptor()
}

func (CodeType) Type() protoreflect.EnumType {
	return &file_common_code_v1_code_proto_enumTypes[0]
}

func (x CodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeType.Descriptor instead.
func (CodeType) EnumDescriptor() ([]byte, []int) {
	return file_common_code_v1_code_proto_rawDescGZIP(), []int{0}
}

type CodeCategory int32

const (
	CodeCategory_login    CodeCategory = 0
	CodeCategory_register CodeCategory = 1
)

// Enum value maps for CodeCategory.
var (
	CodeCategory_name = map[int32]string{
		0: "login",
		1: "register",
	}
	CodeCategory_value = map[string]int32{
		"login":    0,
		"register": 1,
	}
)

func (x CodeCategory) Enum() *CodeCategory {
	p := new(CodeCategory)
	*p = x
	return p
}

func (x CodeCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_common_code_v1_code_proto_enumTypes[1].Descriptor()
}

func (CodeCategory) Type() protoreflect.EnumType {
	return &file_common_code_v1_code_proto_enumTypes[1]
}

func (x CodeCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeCategory.Descriptor instead.
func (CodeCategory) EnumDescriptor() ([]byte, []int) {
	return file_common_code_v1_code_proto_rawDescGZIP(), []int{1}
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *resource.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId         *resource.Identifier   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Code           string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type           CodeType               `protobuf:"varint,4,opt,name=type,proto3,enum=common.code.v1.CodeType" json:"type,omitempty"`
	Category       CodeCategory           `protobuf:"varint,5,opt,name=category,proto3,enum=common.code.v1.CodeCategory" json:"category,omitempty"`
	EmailEncrypted string                 `protobuf:"bytes,6,opt,name=email_encrypted,json=emailEncrypted,proto3" json:"email_encrypted,omitempty"`
	EmailIndex     int64                  `protobuf:"varint,7,opt,name=email_index,json=emailIndex,proto3" json:"email_index,omitempty"`
	PhoneEncrypted string                 `protobuf:"bytes,8,opt,name=phone_encrypted,json=phoneEncrypted,proto3" json:"phone_encrypted,omitempty"`
	PhoneIndex     int64                  `protobuf:"varint,9,opt,name=phone_index,json=phoneIndex,proto3" json:"phone_index,omitempty"`
	AttemptCount   int64                  `protobuf:"varint,10,opt,name=attempt_count,json=attemptCount,proto3" json:"attempt_count,omitempty"`
	Data           []byte                 `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
	ValidatedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=validated_at,json=validatedAt,proto3" json:"validated_at,omitempty"`
	ExpiredAt      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	User           *v1.User               `protobuf:"bytes,14,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_code_v1_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_common_code_v1_code_proto_msgTypes[0]
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
	return file_common_code_v1_code_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Code) GetUserId() *resource.Identifier {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Code) GetType() CodeType {
	if x != nil {
		return x.Type
	}
	return CodeType_email
}

func (x *Code) GetCategory() CodeCategory {
	if x != nil {
		return x.Category
	}
	return CodeCategory_login
}

func (x *Code) GetEmailEncrypted() string {
	if x != nil {
		return x.EmailEncrypted
	}
	return ""
}

func (x *Code) GetEmailIndex() int64 {
	if x != nil {
		return x.EmailIndex
	}
	return 0
}

func (x *Code) GetPhoneEncrypted() string {
	if x != nil {
		return x.PhoneEncrypted
	}
	return ""
}

func (x *Code) GetPhoneIndex() int64 {
	if x != nil {
		return x.PhoneIndex
	}
	return 0
}

func (x *Code) GetAttemptCount() int64 {
	if x != nil {
		return x.AttemptCount
	}
	return 0
}

func (x *Code) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Code) GetValidatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidatedAt
	}
	return nil
}

func (x *Code) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *Code) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_common_code_v1_code_proto protoreflect.FileDescriptor

var file_common_code_v1_code_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xac, 0x05, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x10, 0xba, 0xb9, 0x19,
	0x0c, 0x0a, 0x0a, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x48, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x0a,
	0x06, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x0f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x40, 0x01, 0x52, 0x0e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0b,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x40, 0x01, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2f, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x40, 0x01, 0x52, 0x0e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x40, 0x01, 0x52, 0x0a, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x40, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x45, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x3a, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x08, 0x01, 0x28, 0x01, 0x2a, 0x20,
	0x0a, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x01,
	0x2a, 0x27, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10, 0x01, 0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x65, 0x78, 0x2f, 0x7a, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x64, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_code_v1_code_proto_rawDescOnce sync.Once
	file_common_code_v1_code_proto_rawDescData = file_common_code_v1_code_proto_rawDesc
)

func file_common_code_v1_code_proto_rawDescGZIP() []byte {
	file_common_code_v1_code_proto_rawDescOnce.Do(func() {
		file_common_code_v1_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_code_v1_code_proto_rawDescData)
	})
	return file_common_code_v1_code_proto_rawDescData
}

var file_common_code_v1_code_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_code_v1_code_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_code_v1_code_proto_goTypes = []interface{}{
	(CodeType)(0),                 // 0: common.code.v1.CodeType
	(CodeCategory)(0),             // 1: common.code.v1.CodeCategory
	(*Code)(nil),                  // 2: common.code.v1.Code
	(*resource.Identifier)(nil),   // 3: atlas.rpc.Identifier
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*v1.User)(nil),               // 5: common.user.v1.User
}
var file_common_code_v1_code_proto_depIdxs = []int32{
	3, // 0: common.code.v1.Code.id:type_name -> atlas.rpc.Identifier
	3, // 1: common.code.v1.Code.user_id:type_name -> atlas.rpc.Identifier
	0, // 2: common.code.v1.Code.type:type_name -> common.code.v1.CodeType
	1, // 3: common.code.v1.Code.category:type_name -> common.code.v1.CodeCategory
	4, // 4: common.code.v1.Code.validated_at:type_name -> google.protobuf.Timestamp
	4, // 5: common.code.v1.Code.expired_at:type_name -> google.protobuf.Timestamp
	5, // 6: common.code.v1.Code.user:type_name -> common.user.v1.User
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_common_code_v1_code_proto_init() }
func file_common_code_v1_code_proto_init() {
	if File_common_code_v1_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_code_v1_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_common_code_v1_code_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_code_v1_code_proto_goTypes,
		DependencyIndexes: file_common_code_v1_code_proto_depIdxs,
		EnumInfos:         file_common_code_v1_code_proto_enumTypes,
		MessageInfos:      file_common_code_v1_code_proto_msgTypes,
	}.Build()
	File_common_code_v1_code_proto = out.File
	file_common_code_v1_code_proto_rawDesc = nil
	file_common_code_v1_code_proto_goTypes = nil
	file_common_code_v1_code_proto_depIdxs = nil
}
