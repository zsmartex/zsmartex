package domain

import commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"

type UserState string

const (
	UserStatePending UserState = "pending"
	UserStateActive  UserState = "active"
	UserStateBanned  UserState = "banned"
)

func (u UserState) ProtobufValue() commonv1.User_State {
	switch u {
	case UserStatePending:
		return commonv1.User_STATE_PENDING
	case UserStateActive:
		return commonv1.User_STATE_ACTIVE
	case UserStateBanned:
		return commonv1.User_STATE_BANNED
	default:
		return commonv1.User_STATE_UNKNOWN
	}
}
