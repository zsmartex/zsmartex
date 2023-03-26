//go:build migrate

package app

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/zsmartex/pkg/v3/infrastucture/database"
	"github.com/zsmartex/zsmartex/cmd/user/config"
	"github.com/zsmartex/zsmartex/internal/user/infras/repo"
	"github.com/zsmartex/zsmartex/internal/user/migrations"
	usersUC "github.com/zsmartex/zsmartex/internal/user/usecases/users"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Seeds struct {
	Users []UserSeed `yaml:"users"`
}

type UserSeed struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
	State    string `yaml:"state"`
	Level    int32  `yaml:"level"`
}

func init() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
		return
	}

	postgres, err := database.New(&database.Config{
		Host:     config.Postgres.Host,
		Port:     config.Postgres.Port,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		DBName:   config.Postgres.Database,
	})

	err = migrate(postgres)
	if err != nil {
		panic(err)
	}

	err = seed(ctx, postgres)
	if err != nil {
		panic(err)
	}
}

func migrate(postgres *gorm.DB) error {
	migrator := gormigrate.New(postgres, gormigrate.DefaultOptions, migrations.ModelSchemaList)

	return migrator.Migrate()
}

func seed(ctx context.Context, postgres *gorm.DB) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(dir + "/seeds.yml")
	if err != nil {
		return err
	}

	var seeds Seeds

	err = yaml.Unmarshal(bytes, &seeds)
	if err != nil {
		return err
	}

	userCredentialsRepository := repo.NewUserCredentialsRepository(postgres)
	userRepository := repo.NewUserRepository(postgres, userCredentialsRepository)
	userUsecase := usersUC.NewUserUseCase(userRepository, userCredentialsRepository)

	for _, user := range seeds.Users {
		userExist, _ := userUsecase.GetUser(ctx, usersUC.GetUserParams{
			Email: user.Email,
		})

		if userExist != nil {
			continue
		}

		_, err = userUsecase.CreateUser(ctx, usersUC.CreateUserParams{
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
			State:    user.State,
			Level:    user.Level,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
