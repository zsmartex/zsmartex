package handlers

import (
	"github.com/google/uuid"
	"github.com/modernice/goes/aggregate"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/domain"
	"github.com/zsmartex/zsmartex/internal/user/events"
)

const UserAggregate = "user"

type User struct {
	*handler.BaseHandler
	*aggregate.Base
}

func NewUser(id uuid.UUID) *User {
	userHandler := &User{
		BaseHandler: handler.NewBase(),
		Base:        aggregate.New("user", id),
	}

	command.ApplyWith(userHandler, userHandler.RegisterUser, string(commands.RegisterUserCmd))

	return userHandler
}

func (u *User) RegisterUser(data commands.RegisterUserData) error {
	aggregate.Next(u, string(events.UserCreated), events.UserCreatedData{
		UID:            data.UID,
		Email:          data.Email,
		PasswordDigest: data.PasswordDigest,
		Role:           domain.UserRoleMember,
		State:          domain.UserStatePending,
		Labels: []domain.Label{
			{
				Key:   "email",
				Value: "pending",
				Scope: domain.LabelScopePrivate,
			},
		},
	})

	return nil
}
