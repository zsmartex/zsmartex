package handlers

import "go.uber.org/fx"

var Module = fx.Module(
	"handlers.Module",
	fx.Provide(
		fx.Annotated{
			Group:  "cqrs_command_handlers",
			Target: NewRegisterHandler,
		},
	),
)
