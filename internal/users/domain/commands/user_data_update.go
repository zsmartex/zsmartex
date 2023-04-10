package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
)

var _ = eh.Command(UserDataUpdateCommand{})

type UserDataUpdateCommand struct {
	ID uuid.UUID `json:"id"`
}

func (c UserDataUpdateCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c UserDataUpdateCommand) AggregateType() eh.AggregateType {
	return domain.UserAggregateType
}

func (c UserDataUpdateCommand) CommandType() eh.CommandType {
	return domain.UserDataUpdateCommand
}
