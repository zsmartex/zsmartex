package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/zsmartex/zsmartex/pkg/config"
)

func NewNats(config config.Nats) (*nats.Conn, error) {
	return nats.Connect(fmt.Sprintf("nats://%s", config.Host))
}
