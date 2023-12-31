package projectors

import "go.uber.org/fx"

var Module = fx.Module(
	"projectors.Module",
	fx.Provide(
		fx.Annotated{
			Group:  "cqrs_event_handlers",
			Target: NewUserCreatedProjector,
		},
	),
)
