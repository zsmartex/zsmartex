package eventhorizon

import (
	"context"
	"fmt"

	eh "github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	"github.com/looplab/eventhorizon/eventbus/nats"
	ehProjector "github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/looplab/eventhorizon/eventstore/mongodb"
	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	mongoRepo "github.com/looplab/eventhorizon/repo/mongodb"
	"github.com/looplab/eventhorizon/repo/version"
	"github.com/looplab/eventhorizon/tracing"
	"github.com/zsmartex/zsmartex/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type EventHorizonRepoType int

const (
	EventHorizonRepoTypeVersion EventHorizonRepoType = iota
	EventHorizonRepoTypeTracing
)

type EventHorizon struct {
	eventStore *mongodb.EventStore
	eventBus   *nats.EventBus
	outBox     *mongoOutbox.Outbox
	repo       eh.ReadWriteRepo
	logger     *zap.Logger
}

func New(
	ctx context.Context,
	mongoConfig config.Mongo,
	natsConfig config.Nats,
	projector ehProjector.Projector,
	repoType EventHorizonRepoType,
	aggregateType eh.AggregateType,
	eventTypes []eh.EventType,
	commandTypes []eh.CommandType,
	logger *zap.Logger,
) (*EventHorizon, error) {
	mongoClient, err := mongo.NewClient()
	if err != nil {
		return nil, err
	}

	eventStore, err := mongodb.NewEventStoreWithClient(mongoClient, fmt.Sprintf("%s_store", aggregateType))
	if err != nil {
		return nil, err
	}

	eventBus, err := nats.NewEventBus(natsConfig.URI, fmt.Sprintf("%s_events", aggregateType))
	if err != nil {
		return nil, err
	}

	outBox, err := mongoOutbox.NewOutboxWithClient(mongoClient, fmt.Sprintf("%s_outbox", aggregateType))
	if err != nil {
		return nil, err
	}

	mongoRepo, err := mongoRepo.NewRepoWithClient(mongoClient, string(aggregateType), string(aggregateType))
	if err != nil {
		return nil, err
	}

	var repo eh.ReadWriteRepo
	switch repoType {
	case EventHorizonRepoTypeVersion:
		repo = version.NewRepo(mongoRepo)
	case EventHorizonRepoTypeTracing:
		repo = tracing.NewRepo(mongoRepo)
	}

	commandBus := bus.NewCommandHandler()

	// Add the EventBus as the last handler of the outbox.
	if err := outBox.AddHandler(ctx, eh.MatchAll{}, eventBus); err != nil {
		return nil, err
	}

	// Add a logger as an observer.
	if err := eventBus.AddHandler(ctx, eh.MatchAll{}, eh.UseEventHandlerMiddleware(NewLogger(logger))); err != nil {
		return nil, err
	}

	// Create the aggregate store.
	aggregateStore, err := ehEvents.NewAggregateStore(eventStore)
	if err != nil {
		return nil, err
	}

	usersHandler, err := aggregate.NewCommandHandler(aggregateType, aggregateStore)
	if err != nil {
		return nil, err
	}

	commandHandler := eh.UseCommandHandlerMiddleware(usersHandler, LoggingMiddleware(logger))
	for _, cmd := range commandTypes {
		err := commandBus.SetHandler(commandHandler, cmd)
		if err != nil {
			return nil, err
		}
	}

	projectorEventHandler := ehProjector.NewEventHandler(projector, repo)

	if err = eventBus.AddHandler(ctx, eh.MatchEvents(eventTypes), projectorEventHandler); err != nil {
		return nil, err
	}

	outBox.Start()

	logger.Debug("eventhorizon setup has done!")

	return &EventHorizon{
		eventStore: eventStore,
		eventBus:   eventBus,
		outBox:     outBox,
		repo:       repo,
		logger:     logger,
	}, nil
}

func (e *EventHorizon) Close() error {
	err := e.eventBus.Close()
	if err != nil {
		e.logger.Sugar().Errorw("error closing event bus", "error", err)
		return err
	}

	err = e.eventStore.Close()
	if err != nil {
		e.logger.Sugar().Errorw("error closing event store", "error", err)
		return err
	}

	err = e.outBox.Close()
	if err != nil {
		e.logger.Sugar().Errorw("error closing outbox", "error", err)
		return err
	}

	err = e.repo.Close()
	if err != nil {
		e.logger.Sugar().Errorw("error closing repo", "error", err)
		return err
	}

	return nil
}
