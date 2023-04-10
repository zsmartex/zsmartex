package domain

import eh "github.com/looplab/eventhorizon"

const UserAggregateType eh.AggregateType = "user"

// Commands.
const (
	UserCreateCommand       eh.CommandType = "user_create"
	UserUpdateCommand       eh.CommandType = "user_update"
	UserLabelApplyCommand   eh.CommandType = "user_label_apply"
	UserLabelDestroyCommand eh.CommandType = "user_label_destroy"
	UserDataUpdateCommand   eh.CommandType = "user_data_update"
)

// Events.
const (
	UserCreatedEvent        eh.EventType = "user_created"
	UserUpdatedEvent        eh.EventType = "user_updated"
	UserLabelAppliedEvent   eh.EventType = "user_label_applied"
	UserLabelDestroyedEvent eh.EventType = "user_label_destroyed"
	UserDataUpdatedEvent    eh.EventType = "user_data_updated"
)
