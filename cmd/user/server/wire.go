//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/modernice/goes/aggregate/repository"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	"github.com/zsmartex/zsmartex/internal/user/projection"
	"github.com/zsmartex/zsmartex/pkg/mongodb"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

func InitApp(
	ctx context.Context,
	cfg *config.Config,
) (*App, func(), error) {
	panic(wire.Build(
		NewApp,
		wire.NewSet(
			projection.NewUser,
			mongodb.NewMongoClient,
			wire.FieldsOf(new(*config.Config), "MongoDB"),
		),
		NewCommandHandler,
		wire.NewSet(
			setup.Set,
			wire.FieldsOf(new(*config.Config), "EventStore", "EventBus"),
		),
	))
}

func NewCommandHandler(repo *repository.Repository, commandBus command.Bus) *handler.Of[*handlers.User] {
	return handler.New(handlers.NewUser, repo, commandBus)
}
