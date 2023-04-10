//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	eh "github.com/looplab/eventhorizon"
	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/users/config"
	"github.com/zsmartex/zsmartex/internal/users/app/router"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/projections"
	"github.com/zsmartex/zsmartex/pkg/eventhorizon"
	"github.com/zsmartex/zsmartex/pkg/session"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitApp(
	ctx context.Context,
	cfg *config.Config,
	grpcServer *grpc.Server,
	redisClient *redis.RedisClient,
) (*App, error) {
	panic(wire.Build(
		New,
		session.NewStore,
		eventhorizonUserSetup,
		eventhorizonActivitySetup,
		router.ProductGRPCServerSet,
	))
}

type UserEventHorizon struct {
}

func eventhorizonUserSetup(
	ctx context.Context,
	config *config.Config,
	logger *zap.Logger,
) (*eventhorizon.EventHorizon, error) {
	return eventhorizon.New(
		ctx,
		config.Mongo,
		config.Nats,
		projections.NewUserProjector(logger),
		eventhorizon.EventHorizonRepoTypeVersion,
		domain.UserAggregateType,
		[]eh.EventType{
			domain.UserCreatedEvent,
			domain.UserUpdatedEvent,
			domain.UserLabelAppliedEvent,
			domain.UserLabelDestroyedEvent,
			domain.UserDataUpdatedEvent,
		},
		[]eh.CommandType{
			domain.UserCreateCommand,
			domain.UserUpdateCommand,
			domain.UserLabelApplyCommand,
			domain.UserLabelDestroyCommand,
			domain.UserDataUpdateCommand,
		},
		logger,
	)
}

func eventhorizonActivitySetup(
	ctx context.Context,
	config *config.Config,
	logger *zap.Logger,
) (*eventhorizon.EventHorizon, error) {
	return eventhorizon.New(
		ctx,
		config.Mongo,
		config.Nats,
		projections.NewUserProjector(logger),
		eventhorizon.EventHorizonRepoTypeTracing,
		domain.ActivityAggregateType,
		[]eh.EventType{
			domain.ActivityCreatedEvent,
		},
		[]eh.CommandType{
			domain.ActivityCreateCommand,
		},
		logger,
	)
}
