package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	userv1 "github.com/zsmartex/zsmartex/proto/common/user/v1"
)

var _ = eh.Command(UserUpdateCommand{})

type UserUpdateCommand struct {
	ID             uuid.UUID        `json:"id"`
	Email          string           `json:"email"`
	Phone          string           `json:"phone"`
	PasswordDigest string           `json:"password_digest"`
	Role           userv1.UserRole  `json:"role"`
	State          userv1.UserState `json:"state"`
}

func NewUserUpdateCommand(
	id uuid.UUID,
	email string,
	phone string,
	password_digest string,
	role userv1.UserRole,
	state userv1.UserState,
) UserUpdateCommand {
	return UserUpdateCommand{
		ID:             id,
		Email:          email,
		Phone:          phone,
		PasswordDigest: password_digest,
		Role:           role,
		State:          state,
	}
}

func (c UserUpdateCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c UserUpdateCommand) AggregateType() eh.AggregateType {
	return domain.UserAggregateType
}

func (c UserUpdateCommand) CommandType() eh.CommandType {
	return domain.UserUpdateCommand
}
