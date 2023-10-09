package commands

import (
	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

const (
	RegisterUserCmd = "register.user"
)

func RegisterCommands(r setup.CommandRegistry) {
	codec.Register[RegisterUserData](r, string(RegisterUserCmd))
}

type RegisterUserData struct {
	UID            string
	Email          string
	PasswordDigest string
}
