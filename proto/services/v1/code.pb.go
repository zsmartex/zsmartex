// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
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

type GetPendingCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CodeType     v1.Code_Type     `protobuf:"varint,2,opt,name=code_type,json=codeType,proto3,enum=common.v1.Code_Type" json:"code_type,omitempty"`
	CodeCategory v1.Code_Category `protobuf:"varint,3,opt,name=code_category,json=codeCategory,proto3,enum=common.v1.Code_Category" json:"code_category,omitempty"`
}

func (x *GetPendingCodeRequest) Reset() {
	*x = GetPendingCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingCodeRequest) ProtoMessage() {}

func (x *GetPendingCodeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPendingCodeRequest.ProtoReflect.Descriptor instead.
func (*GetPendingCodeRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{0}
}

func (x *GetPendingCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetPendingCodeRequest) GetCodeType() v1.Code_Type {
	if x != nil {
		return x.CodeType
	}
	return v1.Code_Type(0)
}

func (x *GetPendingCodeRequest) GetCodeCategory() v1.Code_Category {
	if x != nil {
		return x.CodeCategory
	}
	return v1.Code_Category(0)
}

type GetPendingCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetPendingCodeResponse) Reset() {
	*x = GetPendingCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingCodeResponse) ProtoMessage() {}

func (x *GetPendingCodeResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPendingCodeResponse.ProtoReflect.Descriptor instead.
func (*GetPendingCodeResponse) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{1}
}

func (x *GetPendingCodeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GenerateCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CodeType     v1.Code_Type     `protobuf:"varint,2,opt,name=code_type,json=codeType,proto3,enum=common.v1.Code_Type" json:"code_type,omitempty"`
	CodeCategory v1.Code_Category `protobuf:"varint,3,opt,name=code_category,json=codeCategory,proto3,enum=common.v1.Code_Category" json:"code_category,omitempty"`
	Data         []byte           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GenerateCodeRequest) Reset() {
	*x = GenerateCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeRequest) ProtoMessage() {}

func (x *GenerateCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[2]
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
	return file_services_v1_code_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GenerateCodeRequest) GetCodeType() v1.Code_Type {
	if x != nil {
		return x.CodeType
	}
	return v1.Code_Type(0)
}

func (x *GenerateCodeRequest) GetCodeCategory() v1.Code_Category {
	if x != nil {
		return x.CodeCategory
	}
	return v1.Code_Category(0)
}

func (x *GenerateCodeRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GenerateCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GenerateCodeResponse) Reset() {
	*x = GenerateCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeResponse) ProtoMessage() {}

func (x *GenerateCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[3]
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
	return file_services_v1_code_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateCodeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CheckCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CodeType     v1.Code_Type     `protobuf:"varint,2,opt,name=code_type,json=codeType,proto3,enum=common.v1.Code_Type" json:"code_type,omitempty"`
	CodeCategory v1.Code_Category `protobuf:"varint,3,opt,name=code_category,json=codeCategory,proto3,enum=common.v1.Code_Category" json:"code_category,omitempty"`
}

func (x *CheckCodeRequest) Reset() {
	*x = CheckCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCodeRequest) ProtoMessage() {}

func (x *CheckCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCodeRequest.ProtoReflect.Descriptor instead.
func (*CheckCodeRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{4}
}

func (x *CheckCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckCodeRequest) GetCodeType() v1.Code_Type {
	if x != nil {
		return x.CodeType
	}
	return v1.Code_Type(0)
}

func (x *CheckCodeRequest) GetCodeCategory() v1.Code_Category {
	if x != nil {
		return x.CodeCategory
	}
	return v1.Code_Category(0)
}

type CheckCodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*CheckCodeRequest `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *CheckCodesRequest) Reset() {
	*x = CheckCodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCodesRequest) ProtoMessage() {}

func (x *CheckCodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCodesRequest.ProtoReflect.Descriptor instead.
func (*CheckCodesRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{5}
}

func (x *CheckCodesRequest) GetCodes() []*CheckCodeRequest {
	if x != nil {
		return x.Codes
	}
	return nil
}

type CheckCodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*v1.Code `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *CheckCodesResponse) Reset() {
	*x = CheckCodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCodesResponse) ProtoMessage() {}

func (x *CheckCodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCodesResponse.ProtoReflect.Descriptor instead.
func (*CheckCodesResponse) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{6}
}

func (x *CheckCodesResponse) GetCodes() []*v1.Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

type ValidateCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CodeType     string `protobuf:"bytes,2,opt,name=code_type,json=codeType,proto3" json:"code_type,omitempty"`
	CodeCategory string `protobuf:"bytes,3,opt,name=code_category,json=codeCategory,proto3" json:"code_category,omitempty"`
}

func (x *ValidateCodeRequest) Reset() {
	*x = ValidateCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCodeRequest) ProtoMessage() {}

func (x *ValidateCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCodeRequest.ProtoReflect.Descriptor instead.
func (*ValidateCodeRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidateCodeRequest) GetCodeType() string {
	if x != nil {
		return x.CodeType
	}
	return ""
}

func (x *ValidateCodeRequest) GetCodeCategory() string {
	if x != nil {
		return x.CodeCategory
	}
	return ""
}

type ValidateCodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*ValidateCodeRequest `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *ValidateCodesRequest) Reset() {
	*x = ValidateCodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCodesRequest) ProtoMessage() {}

func (x *ValidateCodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCodesRequest.ProtoReflect.Descriptor instead.
func (*ValidateCodesRequest) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateCodesRequest) GetCodes() []*ValidateCodeRequest {
	if x != nil {
		return x.Codes
	}
	return nil
}

type ValidateCodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*v1.Code `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *ValidateCodesResponse) Reset() {
	*x = ValidateCodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v1_code_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCodesResponse) ProtoMessage() {}

func (x *ValidateCodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v1_code_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCodesResponse.ProtoReflect.Descriptor instead.
func (*ValidateCodesResponse) Descriptor() ([]byte, []int) {
	return file_services_v1_code_proto_rawDescGZIP(), []int{9}
}

func (x *ValidateCodesResponse) GetCodes() []*v1.Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_services_v1_code_proto protoreflect.FileDescriptor

var file_services_v1_code_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31,
	0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xb4,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x9d, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x48, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x4e, 0x0a, 0x14, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xe4, 0x02, 0x0a, 0x0b, 0x43,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
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

var file_services_v1_code_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_v1_code_proto_goTypes = []interface{}{
	(*GetPendingCodeRequest)(nil),  // 0: services.v1.GetPendingCodeRequest
	(*GetPendingCodeResponse)(nil), // 1: services.v1.GetPendingCodeResponse
	(*GenerateCodeRequest)(nil),    // 2: services.v1.GenerateCodeRequest
	(*GenerateCodeResponse)(nil),   // 3: services.v1.GenerateCodeResponse
	(*CheckCodeRequest)(nil),       // 4: services.v1.CheckCodeRequest
	(*CheckCodesRequest)(nil),      // 5: services.v1.CheckCodesRequest
	(*CheckCodesResponse)(nil),     // 6: services.v1.CheckCodesResponse
	(*ValidateCodeRequest)(nil),    // 7: services.v1.ValidateCodeRequest
	(*ValidateCodesRequest)(nil),   // 8: services.v1.ValidateCodesRequest
	(*ValidateCodesResponse)(nil),  // 9: services.v1.ValidateCodesResponse
	(v1.Code_Type)(0),              // 10: common.v1.Code.Type
	(v1.Code_Category)(0),          // 11: common.v1.Code.Category
	(*v1.Code)(nil),                // 12: common.v1.Code
}
var file_services_v1_code_proto_depIdxs = []int32{
	10, // 0: services.v1.GetPendingCodeRequest.code_type:type_name -> common.v1.Code.Type
	11, // 1: services.v1.GetPendingCodeRequest.code_category:type_name -> common.v1.Code.Category
	10, // 2: services.v1.GenerateCodeRequest.code_type:type_name -> common.v1.Code.Type
	11, // 3: services.v1.GenerateCodeRequest.code_category:type_name -> common.v1.Code.Category
	10, // 4: services.v1.CheckCodeRequest.code_type:type_name -> common.v1.Code.Type
	11, // 5: services.v1.CheckCodeRequest.code_category:type_name -> common.v1.Code.Category
	4,  // 6: services.v1.CheckCodesRequest.codes:type_name -> services.v1.CheckCodeRequest
	12, // 7: services.v1.CheckCodesResponse.codes:type_name -> common.v1.Code
	7,  // 8: services.v1.ValidateCodesRequest.codes:type_name -> services.v1.ValidateCodeRequest
	12, // 9: services.v1.ValidateCodesResponse.codes:type_name -> common.v1.Code
	0,  // 10: services.v1.CodeService.GetPendingCode:input_type -> services.v1.GetPendingCodeRequest
	2,  // 11: services.v1.CodeService.GenerateCode:input_type -> services.v1.GenerateCodeRequest
	5,  // 12: services.v1.CodeService.CheckCodes:input_type -> services.v1.CheckCodesRequest
	8,  // 13: services.v1.CodeService.ValidateCodes:input_type -> services.v1.ValidateCodesRequest
	1,  // 14: services.v1.CodeService.GetPendingCode:output_type -> services.v1.GetPendingCodeResponse
	3,  // 15: services.v1.CodeService.GenerateCode:output_type -> services.v1.GenerateCodeResponse
	6,  // 16: services.v1.CodeService.CheckCodes:output_type -> services.v1.CheckCodesResponse
	9,  // 17: services.v1.CodeService.ValidateCodes:output_type -> services.v1.ValidateCodesResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_v1_code_proto_init() }
func file_services_v1_code_proto_init() {
	if File_services_v1_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_v1_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPendingCodeRequest); i {
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
			switch v := v.(*GetPendingCodeResponse); i {
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
		file_services_v1_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_v1_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_v1_code_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCodeRequest); i {
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
		file_services_v1_code_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCodesRequest); i {
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
		file_services_v1_code_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCodesResponse); i {
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
		file_services_v1_code_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCodeRequest); i {
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
		file_services_v1_code_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCodesRequest); i {
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
		file_services_v1_code_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCodesResponse); i {
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
			NumMessages:   10,
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
