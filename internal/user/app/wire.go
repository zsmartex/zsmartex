//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/zsmartex/pkg/v3/infrastucture/database"
	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/app/router"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	usersUC "github.com/zsmartex/zsmartex/internal/user/usecases/users"
	"github.com/zsmartex/zsmartex/pkg/session"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
	redisClient *redis.RedisClient,
) (*App, error) {
	panic(wire.Build(
		New,
		postgresFunc,
		session.NewStore,
		repo.RepositorySet,
		usersUC.UseCaseSet,
		router.ProductGRPCServerSet,
	))
}

func postgresFunc(config *config.Config) (*gorm.DB, error) {
	return database.New(&database.Config{
		Host:     config.Postgres.Host,
		Port:     config.Postgres.Port,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		DBName:   config.Postgres.Database,
	})
}
