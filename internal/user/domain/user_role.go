package domain

import commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleMember UserRole = "member"
)

func (u UserRole) ProtobufValue() commonv1.User_Role {
	switch u {
	case UserRoleMember:
		return commonv1.User_ROLE_MEMBER
	case UserRoleAdmin:
		return commonv1.User_ROLE_ADMIN
	default:
		return commonv1.User_ROLE_UNKNOWN
	}
}
