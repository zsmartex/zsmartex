package setup

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"github.com/modernice/goes/aggregate/repository"
	"github.com/modernice/goes/backend/mongo"
	natsBackend "github.com/modernice/goes/backend/nats"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/cmdbus"
	"github.com/modernice/goes/event"
	"github.com/modernice/goes/event/eventstore"
	"github.com/modernice/goes/projection"
	"github.com/modernice/goes/projection/schedule"
	"github.com/zsmartex/pkg/v2/log"
	"github.com/zsmartex/zsmartex/pkg/config"
	"github.com/zsmartex/zsmartex/pkg/mongodb"
	natsPkg "github.com/zsmartex/zsmartex/pkg/nats"
)

var Set = wire.NewSet(
	NewEventStore,
	NewEventBus,
	NewEventRegistry,
	NewCommandRegistry,
	NewCommandBus,
	NewAggregate,
)

type EventRegistry interface {
	Map() map[string]func() any
	Marshal(data any) ([]byte, error)
	New(name string) (any, error)
	Register(name string, factory func() any)
	Unmarshal(b []byte, name string) (any, error)
}

type CommandRegistry interface {
	Map() map[string]func() any
	Marshal(data any) ([]byte, error)
	New(name string) (any, error)
	Register(name string, factory func() any)
	Unmarshal(b []byte, name string) (any, error)
}

// Context returns a context.Context that is cancelled when the program receives
// an interrupt signal (os.Interrupt, os.Kill, syscall.SIGTERM).
func Context() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
}

func NewEventStore(ctx context.Context, cfg config.EventStore, eventBus event.Bus, eventRegistry EventRegistry) (event.Store, error) {
	mongoClient, err := mongodb.NewMongoClient(ctx, config.MongoDB{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	})
	if err != nil {
		return nil, err
	}

	eventStore := mongo.NewEventStore(eventRegistry, mongo.Client(mongoClient), mongo.Database(cfg.Database))
	store := eventstore.WithBus(eventStore, eventBus)

	return store, nil
}

func NewEventBus(ctx context.Context, cfg config.EventBus, eventRegistry EventRegistry) (event.Bus, func(), error) {
	natsConn, err := natsPkg.NewNats(config.Nats{
		Host: cfg.Host,
	})
	if err != nil {
		return nil, nil, err
	}

	bus := natsBackend.NewEventBus(eventRegistry, natsBackend.Conn(natsConn))

	return bus, func() {
		log.Info("Disconnecting from NATS ...")

		if err := bus.Disconnect(ctx); err != nil {
			log.Panicf("Failed to disconnect from NATS: %v", err)
		}
	}, nil
}

func NewEventRegistry() EventRegistry {
	return event.NewRegistry()
}

func NewCommandRegistry() CommandRegistry {
	return command.NewRegistry()
}

func NewCommandBus(commandRegistry CommandRegistry, eventBus event.Bus, eventRegistry EventRegistry) command.Bus {
	cmdbus.RegisterEvents(eventRegistry)
	return cmdbus.New[int](commandRegistry, eventBus)
}

func NewAggregate(eventStore event.Store) *repository.Repository {
	return repository.New(eventStore)
}

func Project(
	ctx context.Context,
	pj projection.Target[any],
	bus event.Bus,
	store event.Store,
	events []string,
	opts ...schedule.ContinuousOption,
) (<-chan error, error) {
	s := schedule.Continuously(bus, store, events, opts...)

	errs, err := s.Subscribe(ctx, func(ctx projection.Job) error {
		start := time.Now()
		log.Debugf("Applying projection job ...")
		defer func() { log.Debugf("Applied projection job. (%s)", time.Since(start)) }()

		return ctx.Apply(ctx, pj)
	})
	if err != nil {
		return nil, fmt.Errorf("subscribe to projection schedule: %w", err)
	}

	return errs, nil
}
