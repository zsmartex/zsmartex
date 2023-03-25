package config

type (
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Postgres struct {
		Host     string `env-required:"true" yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
		Port     int    `env-required:"true" yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
		User     string `env-required:"true" yaml:"user" env:"POSTGRES_USER" env-default:"root"`
		Password string `env-required:"true" yaml:"password" env:"POSTGRES_PASSWORD" env-default:"changeme"`
		Database string `env-required:"true" yaml:"database" env:"POSTGRES_DATABASE" env-default:"zsmartex"`
	}

	Redis struct {
		URL string `env-required:"true" yaml:"url" env:"REDIS_URL" env-default:"redis://localhost:6379"`
	}
)
