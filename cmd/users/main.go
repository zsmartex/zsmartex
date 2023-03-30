package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/users/config"
	"github.com/zsmartex/zsmartex/internal/users/app"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		logger.Error("failed set max procs", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("failed get config", err)
	}

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	redisClient, err := redis.New(cfg.Redis.URI)
	if err != nil {
		logger.Error("failed connect to redis", err)
	}

	_, err = app.InitApp(cfg, server, redisClient)
	if err != nil {
		logger.Error("failed init app", err)
		cancel()
	}

	address := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		logger.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}

	logger.Info("🌏 start server...", "address", address)

	defer func() {
		if err1 := l.Close(); err != nil {
			logger.Error("failed to close", err1, "network", network, "address", address)
		}
	}()

	err = server.Serve(l)
	if err != nil {
		logger.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.Info("signal.Notify", v)
	case done := <-ctx.Done():
		logger.Info("ctx.Done", done)
	}
}
