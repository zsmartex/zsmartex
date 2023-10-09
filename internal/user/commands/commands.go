package commands

import (
	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

type CommandType string

const (
	RegisterUserCmd CommandType = "register.user"
)

// RegisterCommands registers commands into a registry.
func RegisterCommands(r setup.CommandRegistry) {
	codec.Register[RegisterUserData](r, string(RegisterUserCmd))
}

type RegisterUserData struct {
	UID            string
	Email          string
	PasswordDigest string
}
