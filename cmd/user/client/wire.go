//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/router"
	"github.com/zsmartex/zsmartex/internal/user/usecases"
	"github.com/zsmartex/zsmartex/pkg/mongodb"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"google.golang.org/grpc"
)

func InitApp(
	ctx context.Context,
	cfg *config.Config,
	opts []grpc.DialOption,
) (*App, func(), error) {
	panic(wire.Build(
		NewApp,
		wire.NewSet(
			mongodb.NewMongoClient,
			wire.FieldsOf(new(*config.Config), "MongoDB"),
		),
		usecases.Set,
		router.NewUserServiceServer,
		NewGRPCServer,
		wire.NewSet(
			NewGRPCGateway,
			wire.FieldsOf(new(*config.Config), "GRPC"),
		),
		wire.NewSet(
			setup.Set,
			wire.FieldsOf(new(*config.Config), "EventBus"),
		),
	))
}
