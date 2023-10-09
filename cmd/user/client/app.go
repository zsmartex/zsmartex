package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"google.golang.org/grpc"
)

type App struct {
	grpcServer *grpc.Server
	gwMux      *runtime.ServeMux
}

func NewApp(
	grpcServer *grpc.Server,
	gwMux *runtime.ServeMux,
	eventRegistry setup.EventRegistry,
	commandRegistry setup.CommandRegistry,
) *App {
	events.RegisterEvents(eventRegistry)
	commands.RegisterCommands(commandRegistry)

	return &App{
		grpcServer: grpcServer,
		gwMux:      gwMux,
	}
}
