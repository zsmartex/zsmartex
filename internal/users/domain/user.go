package domain

import (
	"time"

	"github.com/mbahjadol/null"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
)

type User struct {
	ID              uint64
	Username        null.String
	PasswordDigest  string
	Level           int32
	OTP             bool
	Role            userv1.UserRole
	State           userv1.UserState
	ReferralUID     null.String
	Data            []byte
	UserCredentials []UserCredentials
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewUser() {

}
