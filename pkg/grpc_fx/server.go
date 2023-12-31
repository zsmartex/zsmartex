package grpc_fx

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zsmartex/zsmartex/pkg/config"
	"github.com/zsmartex/zsmartex/pkg/logger"
)

var ServerModule = fx.Module(
	"grpc_fx.ServerModule",
	fx.Provide(
		NewGRPCServer,
	),
	fx.Invoke(registerGRPCServerHooks),
)

func NewGRPCServer(logger *logger.Logger) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Call the gRPC handler function
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Error("error", zap.Error(err))
		}

		// Check for an error and convert it to a custom error message
		if err != nil {
			message := "An error occurred: " + err.Error()
			return nil, status.Errorf(codes.Internal, message)
		}

		// Return the response
		return resp, nil
	}))

	return grpcServer
}

func registerGRPCServerHooks(config config.GRPC, lc fx.Lifecycle, grpcServer *grpc.Server, mux *http.ServeMux, logger *logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
			if err != nil {
				logger.Fatal("failed to listen", zap.Error(err))
			}
			go func() {
				logger.Info("starting gRPC server", zap.Int("port", config.Port))
				if err := grpcServer.Serve(lis); err != nil {
					logger.Fatal("failed to serve", zap.Error(err))
				}
			}()

			return nil
		},
	})
}
