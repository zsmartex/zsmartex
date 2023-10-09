package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/zsmartex/zsmartex/pkg/config"
)

type (
	Config struct {
		config.App
		config.HTTP
		config.GRPC
		MongoDB    config.MongoDB    `envPrefix:"MONGODB_"`
		EventStore config.EventStore `envPrefix:"EVENTSTORE_"`
		EventBus   config.EventBus   `envPrefix:"EVENTBUS_"`
	}
)

func New() (*Config, error) {
	cfg := new(Config)

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	return cfg, nil
}
