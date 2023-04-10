package events

import (
	"encoding/json"

	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain/enums"
)

type ActivityCreatedEvent struct {
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
