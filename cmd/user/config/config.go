package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	configs "github.com/zsmartex/zsmartex/pkg/config"
)

type (
	Config struct {
		configs.App      `yaml:"app"`
		configs.HTTP     `yaml:"http"`
		configs.Log      `yaml:"logger"`
		configs.Postgres `yaml:"postgres"`
		configs.Redis    `yaml:"redis"`
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
