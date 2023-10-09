package main

import (
	"github.com/modernice/goes/command/handler"
	"github.com/modernice/goes/event"
	"github.com/zsmartex/zsmartex/internal/user/commands"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	"github.com/zsmartex/zsmartex/internal/user/projection"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

type App struct {
	userProjection *projection.User
	commandHandler *handler.Of[*handlers.User]
	eventBus       event.Bus
	eventStore     event.Store
}

func NewApp(
	userProjection *projection.User,
	commandHandler *handler.Of[*handlers.User],
	eventBus event.Bus,
	eventStore event.Store,
	eventRegistry setup.EventRegistry,
	commandRegistry setup.CommandRegistry,
) *App {
	events.RegisterEvents(eventRegistry)
	commands.RegisterCommands(commandRegistry)

	return &App{
		userProjection: userProjection,
		commandHandler: commandHandler,
		eventBus:       eventBus,
		eventStore:     eventStore,
	}
}
