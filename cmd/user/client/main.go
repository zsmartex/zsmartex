package main

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	"github.com/zsmartex/zsmartex/internal/user/router"
	"github.com/zsmartex/zsmartex/pkg/context_fx"
	"github.com/zsmartex/zsmartex/pkg/cqrs_fx"
	"github.com/zsmartex/zsmartex/pkg/grpc_fx"
	"github.com/zsmartex/zsmartex/pkg/logger"
	"github.com/zsmartex/zsmartex/pkg/mongo_fx"
)

func main() {

	app := fx.New(
		config.Module,
		logger.Module,
		context_fx.Module,
		mongo_fx.Module,
		repo.CollectionModule,
		repo.ReadModule,
		cqrs_fx.LoggerModule,
		cqrs_fx.MarshalerModule,
		cqrs_fx.PublisherModule,
		cqrs_fx.CommandBusModule,
		grpc_fx.ServerModule,
		grpc_fx.GatewayModule,
		router.Module,
		fx.Provide(func() []grpc.DialOption {
			return []grpc.DialOption{grpc.WithInsecure()}
		}),
		fx.Invoke(
			registerGatewayHooks,
		),
	)

	app.Run()
}
