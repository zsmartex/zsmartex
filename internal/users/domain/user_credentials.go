package domain

import (
	"time"

	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
)

type UserCredentials struct {
	UserID    uint64
	Type      userv1.UserCredentials_Type
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
