package projections

import (
	"context"
	"errors"
	"fmt"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/zsmartex/zsmartex/internal/users/domain"
	"github.com/zsmartex/zsmartex/internal/users/domain/events"
	"go.uber.org/zap"
)

var _ = projector.Projector(&UserProjector{})

const activityProjectorType projector.Type = "user"

type ActivityProjector struct {
	logger *zap.Logger
}

func NewActivityProjector(
	logger *zap.Logger,
) *ActivityProjector {
	return &ActivityProjector{
		logger,
	}
}

func (p *ActivityProjector) ProjectorType() projector.Type {
	return userProjectorType
}

func (p *ActivityProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	activity, ok := entity.(*domain.Activity)
	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case domain.ActivityCreatedEvent:
		data, ok := event.Data().(*events.ActivityCreatedEvent)
		if !ok {
			return nil, errors.New("invalid event data")
		}

		activity.UserID = data.UserID
		activity.Category = data.Category
		activity.UserIP = data.UserIP
		// TODO: add support user ip country
		activity.UserAgent = data.UserAgent
		activity.Topic = data.Topic
		activity.Action = data.Action
		activity.Result = data.Result
		activity.Device = data.Device
		activity.Data = data.Data
	default:
		p.logger.Error("invalid event", zap.Any("event", event))

		return nil, fmt.Errorf("invalid event %v", event)
	}

	activity.Version++
	return activity, nil
}
