package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
)

var _ = eh.Command(UserLabelDestroyCommand{})

type UserLabelDestroyCommand struct {
	ID  uuid.UUID `json:"id"`
	Key string    `json:"key"`
}

func (c UserLabelDestroyCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c UserLabelDestroyCommand) AggregateType() eh.AggregateType {
	return domain.UserAggregateType
}

func (c UserLabelDestroyCommand) CommandType() eh.CommandType {
	return domain.UserLabelDestroyCommand
}
