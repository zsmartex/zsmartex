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

	Mongo struct {
		URI string `env-required:"true" yaml:"uri" env:"MONGO_URI" env-default:"mongodb://root:changeme@127.0.0.1:27017/"`
	}

	Redis struct {
		URI string `env-required:"true" yaml:"uri" env:"REDIS_URI" env-default:"redis://localhost:6379"`
	}

	Nats struct {
		URI string `env-required:"true" yaml:"uri" env:"NATS_URI" env-default:"localhost:8222"`
	}
)
