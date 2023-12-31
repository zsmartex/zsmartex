package config

import (
	"github.com/caarlos0/env/v10"
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/pkg/config"
)

var Module = fx.Module(
	"config.Module",
	fx.Provide(
		New,
	),
)

type Config struct {
	fx.Out

	config.GRPC `envPrefix:"GRPC_"`
	config.HTTP `envPrefix:"HTTP_"`
	config.Nats
	config.MongoDB
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
