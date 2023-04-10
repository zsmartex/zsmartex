package eventhorizon

import (
	"context"

	"github.com/looplab/eventhorizon"
	eh "github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	"go.uber.org/zap"
)

func New(
	ctx context.Context,
	eventStore eh.EventStore,
	eventBus eh.EventBus,
	outBox eh.Outbox,
	repo eh.ReadWriteRepo,
	projector *projector.EventHandler,
	eventTypes []eh.EventType,
	commandTypes []eh.CommandType,
	aggregateType eh.AggregateType,
	logger *zap.Logger,
) error {
	commandBus := bus.NewCommandHandler()

	// Add the EventBus as the last handler of the outbox.
	if err := outBox.AddHandler(ctx, eh.MatchAll{}, eventBus); err != nil {
		return err
	}

	// Add a logger as an observer.
	if err := eventBus.AddHandler(ctx, eh.MatchAll{}, eh.UseEventHandlerMiddleware(NewLogger(logger))); err != nil {
		return err
	}

	// Create the aggregate store.
	aggregateStore, err := ehEvents.NewAggregateStore(eventStore)
	if err != nil {
		return err
	}

	usersHandler, err := aggregate.NewCommandHandler(aggregateType, aggregateStore)
	if err != nil {
		return err
	}

	commandHandler := eventhorizon.UseCommandHandlerMiddleware(usersHandler, LoggingMiddleware(logger))
	for _, cmd := range commandTypes {
		err := commandBus.SetHandler(commandHandler, cmd)
		if err != nil {
			return err
		}
	}

	if err = eventBus.AddHandler(ctx, eh.MatchEvents(eventTypes), projector); err != nil {
		return nil
	}

	outBox.Start()

	logger.Debug("eventhorizon setup has done!")

	return nil
}
