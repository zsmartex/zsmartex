package grpc_fx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zsmartex/zsmartex/pkg/config"
	"github.com/zsmartex/zsmartex/pkg/logger"
)

var GatewayModule = fx.Module(
	"grpc_fx.GatewayModule",
	fx.Provide(
		http.NewServeMux,
		NewGRPCGateway,
	),
	fx.Invoke(registerGRPCGatewayHooks),
)

type grpcGatewayParams struct {
	fx.In

	GRPCServerConfig config.GRPC
}

func NewGRPCGateway(ctx context.Context, params grpcGatewayParams, logger *logger.Logger) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			logger.Error("error", zap.Error(err))
			code := codes.Internal
			msg := "Internal Server Error"
			s, ok := status.FromError(err)
			if ok {
				code = s.Code()
				msg = s.Message()
				spew.Dump(s.WithDetails())
			}

			response := map[string][]string{"errors": {msg}}
			body, _ := json.Marshal(response)

			// Write the HTTP response
			writer.Header().Set("Content-Type", marshaler.ContentType(response))
			status := runtime.HTTPStatusFromCode(code)
			writer.WriteHeader(status)
			writer.Write(body)
		}),
	)

	return mux, nil
}

func registerGRPCGatewayHooks(ctx context.Context, lc fx.Lifecycle, config config.HTTP, mux *http.ServeMux, gwMux *runtime.ServeMux, logger *logger.Logger) {
	mux.Handle("/", gwMux)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				<-ctx.Done()
				logger.Info("Shutting down the http server")
				if err := httpServer.Shutdown(context.Background()); err != nil {
					logger.Error("Failed to shutdown http server", zap.Error(err))
				}
			}()

			go func() {
				logger.Info("Starting the http server", zap.Int("port", config.Port))
				// start server
				err := httpServer.ListenAndServe()
				if err != nil {
					logger.Fatal("failed to start http server", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
}
