package command

import (
	"context"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/infras/repo"
	"go.uber.org/zap"
)

type CreateUserHandler eh.CommandHandler

type createUserHandler struct {
	userRepository repo.UserRepository
}

func NewCreateUserHandler(
	userRepository repo.UserRepository,
	logger *zap.SugaredLogger,
) CreateUserHandler {
	return createUserHandler{userRepository}
}

func (h createUserHandler) HandleCommand(ctx context.Context, cmd CreateUserCommand) error {
	return nil
}

type CreateUserCommand struct {
	ID       uuid.UUID
	Email    string
	Phone    string
	Password string
}

func (c CreateUserCommand) AggregateID() uuid.UUID {
	return c.ID
}

func (c CreateUserCommand) AggregateType() eh.AggregateType {
	return domain.UsersAggregateType
}

func (c CreateUserCommand) CommandType() eh.CommandType {
	return domain.UserRegisterCommand
}
