package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	configs "github.com/zsmartex/zsmartex/pkg/config"
)

type (
	Config struct {
		configs.App   `yaml:"app"`
		configs.HTTP  `yaml:"http"`
		configs.Log   `yaml:"logger"`
		configs.Redis `yaml:"redis"`
		configs.Nats  `yaml:"nats"`
		GRPC          `yaml:"grpc"`
	}

	GRPC struct {
		UserHost string `env-required:"true" yaml:"user_host" env:"GRPC_USER_HOST"`
		UserPort int64  `env-required:"true" yaml:"user_port" env:"GRPC_USER_PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
