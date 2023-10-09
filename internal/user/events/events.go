package events

import (
	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

const (
	UserCreated = "user.created"
)

var ListEvents = []string{
	string(UserCreated),
}

func RegisterEvents(r setup.EventRegistry) {
	codec.Register[UserCreatedData](r, string(UserCreated))
}

type UserCreatedData struct {
	UID            string
	Email          string
	PasswordDigest string
	Role           domain.UserRole
	State          domain.UserState
	Labels         []domain.Label
}
