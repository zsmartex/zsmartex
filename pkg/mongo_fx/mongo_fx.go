package mongo_fx

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/pkg/config"
	"github.com/zsmartex/zsmartex/pkg/logger"
)

var Module = fx.Module(
	"mongo_fx.Module",
	fx.Provide(
		New,
		NewDatabase,
	),
)

func New(ctx context.Context, config config.MongoDB, logger *logger.Logger) (*mongo.Client, error) {
	sink := NewLogger(logger)

	loggerOptions := options.
		Logger().
		SetSink(sink).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	options := options.Client().
		ApplyURI(config.URL).
		SetLoggerOptions(loggerOptions)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewDatabase(client *mongo.Client, config config.MongoDB) *mongo.Database {
	return client.Database(config.Database)
}
