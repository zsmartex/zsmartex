package domain

import (
	"encoding/json"
	"time"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain/enums"
)

var _ = eh.Entity(&Activity{})

type Activity struct {
	ID            uuid.UUID              `json:"id" bson:"_id"`
	UserID        uuid.UUID              `json:"user_id" bson:"user_id"`
	TargetUID     string                 `json:"target_uid,omitempty" bson:"target_uid,omitempty"`
	Category      enums.ActivityCategory `json:"category" bson:"category"`
	UserIP        string                 `json:"user_ip" bson:"user_ip"`
	UserIPCountry string                 `json:"user_ip_country,omitempty" bson:"user_ip_countr,omitempty"`
	UserAgent     string                 `json:"user_agent" bson:"user_agent"`
	Topic         enums.ActivityTopic    `json:"topic" bson:"topic"`
	Action        enums.ActivityAction   `json:"action" bson:"action"`
	Result        enums.ActivityResult   `json:"result" bson:"result"`
	Device        string                 `json:"device" bson:"device"`
	Data          json.RawMessage        `json:"data,omitempty" bson:"data,omitempty"`
	Version       int                    `json:"version" bson:"version"`
	CreatedAt     time.Time              `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at" bson:"updated_at"`
}

// EntityID implements the EntityID method of the eventhorizon.Entity interface.
func (t *Activity) EntityID() uuid.UUID {
	return t.ID
}
