package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/app"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	redisClient, err := redis.New(cfg.Redis.URL)
	if err != nil {
		slog.Error("failed connect to redis", err)
	}

	_, err = app.InitApp(cfg, server, redisClient)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
	}

	address := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}

	slog.Info("🌏 start server...", "address", address)

	defer func() {
		if err1 := l.Close(); err != nil {
			slog.Error("failed to close", err1, "network", network, "address", address)
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", done)
	}
}
