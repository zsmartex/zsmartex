package context_fx

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("context_fx", fx.Provide(New), fx.Invoke(registerHooks))

func New() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func registerHooks(lc fx.Lifecycle, cancel context.CancelFunc) {
	lc.Append(fx.StopHook(func() {
		cancel()
	}))
}
