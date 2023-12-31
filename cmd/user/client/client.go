package main

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/zsmartex/zsmartex/pkg/config"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
)

func registerGatewayHooks(ctx context.Context, gwMux *runtime.ServeMux, config config.GRPC, opts []grpc.DialOption) error {
	err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, gwMux, fmt.Sprintf("localhost:%d", config.Port), opts)
	if err != nil {
		return err
	}

	return nil
}
