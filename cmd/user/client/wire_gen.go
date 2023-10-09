// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/router"
	"github.com/zsmartex/zsmartex/internal/user/usecases"
	"github.com/zsmartex/zsmartex/pkg/mongodb"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(ctx context.Context, cfg *config.Config, opts []grpc.DialOption) (*App, func(), error) {
	commandRegistry := setup.NewCommandRegistry()
	eventBus := cfg.EventBus
	eventRegistry := setup.NewEventRegistry()
	bus, cleanup, err := setup.NewEventBus(ctx, eventBus, eventRegistry)
	if err != nil {
		return nil, nil, err
	}
	commandBus := setup.NewCommandBus(commandRegistry, bus, eventRegistry)
	mongoDB := cfg.MongoDB
	client, err := mongodb.NewMongoClient(ctx, mongoDB)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userUsecase := usecases.NewUserUsecase(ctx, commandBus, client)
	userServiceServer := router.NewUserServiceServer(userUsecase)
	server := NewGRPCServer(userServiceServer)
	configGRPC := cfg.GRPC
	serveMux, err := NewGRPCGateway(ctx, configGRPC, opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app := NewApp(server, serveMux, eventRegistry, commandRegistry)
	return app, func() {
		cleanup()
	}, nil
}