package events

import (
	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

type EventType string

const (
	UserCreated EventType = "user.created"
)

// ListEvents are all events of a todo list.
var ListEvents = []string{
	string(UserCreated),
}

// RegisterEvents registers events into a registry.
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
