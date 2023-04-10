package types

import "github.com/looplab/eventhorizon/uuid"

func (u UUID) ToUUID() (uuid.UUID, error) {
	return uuid.Parse(u.String())
}

func UUIDToPB(uuid uuid.UUID) string {
	return uuid.String()
}
