package cqrs_fx

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"go.uber.org/fx"
)

var RouterModule = fx.Module(
	"cqrs_fx.RouterModule",
	fx.Provide(
		NewRouter,
	),
	fx.Invoke(RegisterRouterHooks),
)

func NewRouter(logger *Logger) (*message.Router, error) {
	return message.NewRouter(message.RouterConfig{}, logger)
}

type RouterHooksParams struct {
	fx.In

	Router *message.Router
}

func RegisterRouterHooks(lc fx.Lifecycle, params RouterHooksParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Router.AddMiddleware(middleware.Recoverer)

			go params.Router.Run(context.Background())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			go params.Router.Close()
			return nil
		},
	})
}
