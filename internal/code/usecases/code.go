package usecases

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/modernice/goes/command"
	"github.com/zsmartex/pkg/v2/utils"
	"github.com/zsmartex/zsmartex/internal/code/commands"
	"github.com/zsmartex/zsmartex/internal/code/domain"
	"github.com/zsmartex/zsmartex/internal/code/infras/repo"
	"github.com/zsmartex/zsmartex/internal/user/handlers"
	servicesv1 "github.com/zsmartex/zsmartex/proto/services/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ CodeUsecase = (*codeUsecase)(nil)

type codeUsecase struct {
	commandBus        command.Bus
	userServiceClient servicesv1.UserServiceClient

	codeRepo repo.CodeRepo
}

func NewUserUsecase(
	ctx context.Context,
	commandBus command.Bus,
	userServiceClient servicesv1.UserServiceClient,
	codeRepo repo.CodeRepo,
) CodeUsecase {
	return &codeUsecase{
		commandBus:        commandBus,
		userServiceClient: userServiceClient,
		codeRepo:          codeRepo,
	}
}

func (c *codeUsecase) GenerateCode(ctx context.Context, userId primitive.ObjectID, t domain.CodeType, category domain.CodeCategory, data json.RawMessage) error {
	user, err := c.userServiceClient.GetUser(ctx, &servicesv1.GetUserRequest{
		QueryBy:    "id",
		QueryValue: userId.Hex(),
	})
	if err != nil {
		return err
	}

	cmd := command.New(commands.GenerateCodeCmd, commands.GenerateCodeData{
		UserID:   userId,
		Email:    user.Email,
		Code:     utils.RandomNumber(6),
		Type:     t,
		Category: category,
		Data:     data,
	}, command.Aggregate(handlers.UserAggregate, uuid.New()))

	return c.commandBus.Dispatch(ctx, cmd.Any())
}
