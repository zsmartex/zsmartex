package domain

import commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"

type CodeCategory string

const (
	CodeCategoryConfirm CodeCategory = "confirm"
	CodeCategoryLogin   CodeCategory = "login"
)

func (c CodeCategory) ProtobufValue() commonv1.Code_Category {
	switch c {
	case CodeCategoryConfirm:
		return commonv1.Code_CATEGORY_CONFIRM
	case CodeCategoryLogin:
		return commonv1.Code_CATEGORY_LOGIN
	default:
		return commonv1.Code_CATEGORY_UNKNOWN
	}
}
