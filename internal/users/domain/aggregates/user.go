package domain

import (
	"context"
	"errors"
	"time"

	eh "github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/commands"
	"github.com/zsmartex/zsmartex/internal/users/domain/events"
	"github.com/zsmartex/zsmartex/internal/users/infras/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
)

var _ = eh.Aggregate(&User{})

type User struct {
	*ehEvents.AggregateBase
	userRepository repo.UserRepository
	logger         *zap.Logger
}

func NewUserAggregate(id uuid.UUID, userRepository repo.UserRepository, logger *zap.Logger) *User {
	return &User{
		AggregateBase:  ehEvents.NewAggregateBase(domain.UserAggregateType, id),
		userRepository: userRepository,
		logger:         logger,
	}
}

func (a *User) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case commands.UserCreateCommand:
		user, err := a.userRepository.GetUser(ctx, repo.GetUserParams{
			Email: cmd.Email,
			Phone: cmd.Phone,
		})
		if user != nil {
			return ErrUserAlreadyExists
		}
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

		a.AppendEvent(domain.UserCreatedEvent, events.UserCreatedEvent{
			Email:          cmd.Email,
			Phone:          cmd.Phone,
			PasswordDigest: cmd.PasswordDigest,
		}, time.Now())
		return nil
	case commands.UserUpdateCommand:
		a.AppendEvent(domain.UserUpdatedEvent, events.UserUpdatedEvent{
			Email:          cmd.Email,
			Phone:          cmd.Phone,
			PasswordDigest: cmd.PasswordDigest,
			Role:           cmd.Role,
			State:          cmd.State,
		}, time.Now())
	case commands.UserDataUpdateCommand:
		a.AppendEvent(domain.UserDataUpdatedEvent, events.UserDataUpdatedEvent{}, time.Now())
	case commands.UserLabelApplyCommand:
		a.AppendEvent(domain.UserLabelAppliedEvent, events.UserLabelAppliedEvent{
			Key:         cmd.Key,
			Value:       cmd.Value,
			Description: cmd.Description,
		}, time.Now())
	case commands.UserLabelDestroyCommand:
		a.AppendEvent(domain.UserLabelDestroyedEvent, events.UserLabelDestroyedEvent{
			Key: cmd.Key,
		}, time.Now())
	default:
		a.logger.Error("invalid command", zap.Any("command", cmd))

		return errors.New("")
	}

	return nil
}
