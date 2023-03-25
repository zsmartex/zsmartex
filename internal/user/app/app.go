package app

import (
	"github.com/zsmartex/zsmartex/cmd/user/config"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
)

type App struct {
	Cfg            *config.Config
	UserGRPCServer userv1.UserServiceServer
}

func New(cfg *config.Config, userGRPCServer userv1.UserServiceServer) *App {
	return &App{
		Cfg:            cfg,
		UserGRPCServer: userGRPCServer,
	}
}
