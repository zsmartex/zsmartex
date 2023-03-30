//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	"github.com/looplab/eventhorizon/eventbus/nats"
	"github.com/looplab/eventhorizon/eventstore/mongodb"
	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	"github.com/zsmartex/pkg/v3/infrastucture/database"
	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/users/config"
	"github.com/zsmartex/zsmartex/internal/users/app/router"
	"github.com/zsmartex/zsmartex/internal/users/infras/repo"
	usersUC "github.com/zsmartex/zsmartex/internal/users/usecases/users"
	"github.com/zsmartex/zsmartex/pkg/session"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitApp(
	ctx context.Context,
	cfg *config.Config,
	grpcServer *grpc.Server,
	redisClient *redis.RedisClient,
) (*App, error) {
	panic(wire.Build(
		New,
		postgresFunc,
		mongoFunc,
		eventStoreFunc,
		eventBusFunc,
		session.NewStore,
		repo.RepositorySet,
		usersUC.UseCaseSet,
		router.ProductGRPCServerSet,
	))
}

func postgresFunc(config *config.Config) (*gorm.DB, error) {
	return database.New(&database.Config{
		Host:     config.Postgres.Host,
		Port:     config.Postgres.Port,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		DBName:   config.Postgres.Database,
	})
}

func mongoFunc(config *config.Config) (*mongo.Client, error) {
	return mongo.NewClient()
}

func eventStoreFunc(client *mongo.Client) (*mongodb.EventStore, error) {
	return mongodb.NewEventStoreWithClient(client, "user_events")
}

func eventBusFunc(config *config.Config) (*nats.EventBus, error) {
	return nats.NewEventBus(config.Nats.URI, "user")
}

func outBoxFunc(client *mongo.Client) (*mongoOutbox.Outbox, error) {
	return mongoOutbox.NewOutboxWithClient(client, "user_events")
}
