package aggregates

import (
	"context"
	"errors"
	"fmt"
	"time"

	eh "github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/commands"
	"github.com/zsmartex/zsmartex/internal/users/domain/events"
	"go.uber.org/zap"
)

var (
	ErrActivityAlreadyExists error = errors.New("user already exists")
)

var _ = eh.Aggregate(&ActivityAggregate{})

type ActivityAggregate struct {
	*ehEvents.AggregateBase
	logger *zap.Logger
}

func NewActivityAggregate(id uuid.UUID, logger *zap.Logger) *ActivityAggregate {
	return &ActivityAggregate{
		AggregateBase: ehEvents.NewAggregateBase(domain.UserAggregateType, id),
		logger:        logger,
	}
}

func (a *ActivityAggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case commands.ActivityCreateCommand:
		a.AppendEvent(domain.ActivityCreatedEvent, events.ActivityCreatedEvent{
			UserID:    cmd.UserID,
			Category:  cmd.Category,
			UserIP:    cmd.UserIP,
			UserAgent: cmd.UserAgent,
			Topic:     cmd.Topic,
			Action:    cmd.Action,
			Result:    cmd.Result,
			Device:    cmd.Device,
			Data:      cmd.Data,
		}, time.Now())
	default:
		a.logger.Error("invalid command", zap.Any("command", cmd))

		return fmt.Errorf("invalid command %v", cmd)
	}

	return nil
}
