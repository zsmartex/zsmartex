package commands

import (
	"encoding/json"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/enums"
)

var _ = eh.Command(ActivityCreateCommand{})

type ActivityCreateCommand struct {
	ID        uuid.UUID              `json:"id"`
	UserID    uuid.UUID              `json:"user_id"`
	Category  enums.ActivityCategory `json:"category"`
	UserIP    string                 `json:"user_ip"`
	UserAgent string                 `json:"user_agent"`
	Topic     enums.ActivityTopic    `json:"topic"`
	Action    enums.ActivityAction   `json:"action"`
	Result    enums.ActivityResult   `json:"result"`
	Device    string                 `json:"device"`
	Data      json.RawMessage        `json:"data"`
}

func NewActivityCreateCommand(
	id uuid.UUID,
	userID uuid.UUID,
	category enums.ActivityCategory,
	userIP string,
	userAgent string,
	topic enums.ActivityTopic,
	action enums.ActivityAction,
	result enums.ActivityResult,
	device string,
	data json.RawMessage,
) ActivityCreateCommand {
	return ActivityCreateCommand{
		ID:        id,
		UserID:    userID,
		Category:  category,
		UserIP:    userIP,
		UserAgent: userAgent,
		Topic:     topic,
		Action:    action,
		Result:    result,
		Device:    device,
		Data:      data,
	}
}

func (c ActivityCreateCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c ActivityCreateCommand) AggregateType() eh.AggregateType {
	return domain.ActivityAggregateType
}

func (c ActivityCreateCommand) CommandType() eh.CommandType {
	return domain.ActivityCreateCommand
}
