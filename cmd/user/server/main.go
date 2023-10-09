package main

import (
	"fmt"
	"time"

	"github.com/modernice/goes/projection/schedule"
	"github.com/zsmartex/pkg/v2/log"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/events"
	"github.com/zsmartex/zsmartex/pkg/logging"
	"github.com/zsmartex/zsmartex/pkg/setup"
)

func init() {
	log.New("user")
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := setup.Context()
	defer cancel()

	app, disconnect, err := InitApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect()
	protectionErrors, err := setup.Project(ctx, app.userProjection, app.eventBus, app.eventStore, events.ListEvents, schedule.Debounce(1*time.Second))
	if err != nil {
		log.Panic(fmt.Errorf("projection: %w", err))
	}

	commandErrors := app.commandHandler.MustHandle(ctx)

	logging.LogErrors(ctx, protectionErrors, commandErrors)
}
