package handlers

import (
	"github.com/google/uuid"
	"github.com/modernice/goes/aggregate"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
	"github.com/zsmartex/zsmartex/internal/code/commands"
	"github.com/zsmartex/zsmartex/internal/code/events"
)

const CodeAggregate = "code"

type Code struct {
	*handler.BaseHandler
	*aggregate.Base
}

func NewCode(id uuid.UUID) *Code {
	codeHandler := &Code{
		BaseHandler: handler.NewBase(),
		Base:        aggregate.New("code", id),
	}

	command.ApplyWith(codeHandler, codeHandler.GenerateCode, commands.GenerateCodeCmd)
	command.ApplyWith(codeHandler, codeHandler.GenerateCode, commands.ValidateCodeCmd)

	return codeHandler
}

func (c *Code) GenerateCode(data commands.GenerateCodeData) error {
	aggregate.Next(c, string(events.CodeGenerated), events.CodeGeneratedData{})

	return nil
}

func (c *Code) ValidateCodes(data commands.GenerateCodeData) error {
	aggregate.Next(c, string(events.CodeValidated), events.CodeValidatedData{})

	return nil
}
