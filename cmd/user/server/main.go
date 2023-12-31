package main

import (
	"go.uber.org/fx"

	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	"github.com/zsmartex/zsmartex/internal/user/projectors"
	"github.com/zsmartex/zsmartex/pkg/context_fx"
	"github.com/zsmartex/zsmartex/pkg/cqrs_fx"
	"github.com/zsmartex/zsmartex/pkg/logger"
	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
	"github.com/zsmartex/zsmartex/pkg/nats_fx"
)

func main() {

	app := fx.New(
		config.Module,
		nats_fx.Module,
		logger.Module,
		context_fx.Module,
		mongo_fx.Module,
		repo.CollectionModule,
		repo.ReadModule,
		repo.WriteModule,
		cqrs_fx.LoggerModule,
		cqrs_fx.PublisherModule,
		cqrs_fx.SubscriberModule,
		cqrs_fx.MarshalerModule,
		cqrs_fx.RouterModule,
		cqrs_fx.CommandProcessorModule,
		cqrs_fx.EventBusModule,
		cqrs_fx.EventProcessorModule,
		handlers.Module,
		projectors.Module,
	)

	app.Run()
}
