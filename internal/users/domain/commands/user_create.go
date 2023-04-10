package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
)

var _ = eh.Command(UserCreateCommand{})

type UserCreateCommand struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	PasswordDigest string    `json:"password_digest"`
}

func NewCreateUserCommand(id uuid.UUID) UserCreateCommand {
	return UserCreateCommand{}
}

func (c UserCreateCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c UserCreateCommand) AggregateType() eh.AggregateType {
	return domain.UserAggregateType
}

func (c UserCreateCommand) CommandType() eh.CommandType {
	return domain.UserCreateCommand
}
