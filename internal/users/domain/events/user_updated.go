package events

import userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"

type UserUpdatedEvent struct {
	Email          string           `json:"email,omitempty"`
	Phone          string           `json:"phone,omitempty"`
	PasswordDigest string           `json:"password_digest,omitempty"`
	Role           userv1.UserRole  `json:"role,omitempty"`
	State          userv1.UserState `json:"state,omitempty"`
}
