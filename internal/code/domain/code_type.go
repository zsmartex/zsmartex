package domain

import commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"

type CodeType string

const (
	CodeTypeEmail CodeType = "email"
	CodeTypePhone CodeType = "phone"
)

func (c CodeType) ProtobufValue() commonv1.Code_Type {
	switch c {
	case CodeTypeEmail:
		return commonv1.Code_TYPE_EMAIL
	case CodeTypePhone:
		return commonv1.Code_TYPE_PHONE
	default:
		return commonv1.Code_TYPE_UNKNOWN
	}
}
