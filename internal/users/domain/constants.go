package domain

import eh "github.com/looplab/eventhorizon"

const UsersAggregateType eh.AggregateType = "users"

// Commands.
const (
	UserRegisterCommand          eh.CommandType = "register"
	UserUpdateInformationCommand eh.CommandType = "update_information"
)

// Events.
const (
	UserCreatedEvent eh.EventType = "user_created"
	UserUpdatedEvent eh.EventType = "user_updated"
)
