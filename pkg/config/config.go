package config

type (
	App struct {
		Name string `env:"APP_NAME"`
	}

	HTTP struct {
		Host string `env:"HTTP_HOST" envDefault:"localhost"`
		Port int    `env:"HTTP_PORT" envDefault:"8080"`
	}

	GRPC struct {
		Host string `env:"GRPC_HOST" envDefault:"localhost"`
		Port int    `env:"GRPC_PORT" envDefault:"9090"`
	}

	// should not use directly
	// must use with envPrefix
	Postgres struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     int    `env:"PORT" envDefault:"5432"`
		Username string `env:"USERNAME" envDefault:"postgres"`
		Password string `env:"PASSWORD" envDefault:"password"`
		Database string `env:"DATABASE" envDefault:"postgres"`
	}

	// should not use directly
	// must use with envPrefix
	MongoDB struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     int    `env:"PORT" envDefault:"27017"`
		Username string `env:"USERNAME" envDefault:"mongodb"`
		Password string `env:"PASSWORD" envDefault:"password"`
		Database string `env:"DATABASE" envDefault:"mongo"`
	}

	// should not use directly
	// must use with envPrefix
	Nats struct {
		Host string `env:"HOST" envDefault:"localhost:4222"`
	}

	EventStore MongoDB
	EventBus   Nats
)
