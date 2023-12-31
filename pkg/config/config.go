package config

type (
	MongoDB struct {
		URL      string `env:"MONGODB_URL" envDefault:"mongodb://mongodb:password@localhost:27017/"`
		Database string `env:"MONGODB_DATABASE" envDefault:"zsmartex"`
	}

	Nats struct {
		URL string `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	}

	HTTP struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port int    `env:"PORT" envDefault:"8080"`
	}

	GRPC struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port int    `env:"PORT" envDefault:"9090"`
	}
)
