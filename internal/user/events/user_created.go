package events

import "github.com/google/uuid"

type UserCreated struct {
	ID             uuid.UUID `json:"id"`
	UID            string    `json:"uid"`
	Email          string    `json:"email"`
	PasswordDigest string    `json:"password"`
}
