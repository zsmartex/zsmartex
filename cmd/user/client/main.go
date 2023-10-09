package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/zsmartex/pkg/v2/log"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"google.golang.org/grpc"
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

	app, disconnect, err := InitApp(ctx, config, []grpc.DialOption{grpc.WithInsecure()})
	defer disconnect()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Infof("starting gRPC server on :%v", config.GRPC.Port)
		if err := app.grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	mux := http.NewServeMux()

	mux.Handle("/", app.gwMux)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Infof("Shutting down the http server")
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	log.Infof("starting HTTP server on :%v", config.HTTP.Port)
	// start server
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
