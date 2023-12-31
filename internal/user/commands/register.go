package commands

import "github.com/google/uuid"

type RegisterCmd struct {
	ID             uuid.UUID `json:"id"`
	UID            string    `json:"uid"`
	Email          string    `json:"email"`
	PasswordDigest string    `json:"password"`
}

func NewRegisterCommand(id uuid.UUID, uid, email, passwordDigest string) *RegisterCmd {
	return &RegisterCmd{
		ID:             id,
		UID:            uid,
		Email:          email,
		PasswordDigest: passwordDigest,
	}
}
