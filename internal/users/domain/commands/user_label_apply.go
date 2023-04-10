package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
)

var _ = eh.Command(UserLabelApplyCommand{})

type UserLabelApplyCommand struct {
	ID          uuid.UUID `json:"id"`
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
}

func (c UserLabelApplyCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c UserLabelApplyCommand) AggregateType() eh.AggregateType {
	return domain.UserAggregateType
}

func (c UserLabelApplyCommand) CommandType() eh.CommandType {
	return domain.UserLabelApplyCommand
}
