package nats_fx

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/pkg/config"
)

var Module = fx.Module(
	"nats_fx.Module",
	fx.Provide(
		NewNATSConn,
	),
)

func NewNATSConn(config config.Nats) (*nats.Conn, error) {
	return nats.Connect(config.URL, nats.Compression(true))
}
