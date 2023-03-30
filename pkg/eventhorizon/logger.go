package eventhorizon

import (
	"context"
	"reflect"

	eh "github.com/looplab/eventhorizon"
	"go.uber.org/zap"
)

// LoggingMiddleware is a tiny command handle middleware for logging.
func LoggingMiddleware(logger *zap.Logger) eh.CommandHandlerMiddleware {
	return func(h eh.CommandHandler) eh.CommandHandler {
		return eh.CommandHandlerFunc(func(ctx context.Context, cmd eh.Command) error {
			logger.Debug("eh_command", zap.Any("command", cmd))
			return h.HandleCommand(ctx, cmd)
		})
	}
}

type Logger struct {
	*zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger,
	}
}

// HandlerType implements the HandlerType method of the eventhorizon.EventHandler interface.
func (l *Logger) HandlerType() eh.EventHandlerType {
	return eh.EventHandlerType(reflect.TypeOf(l.Logger).String())
}

// HandleEvent implements the HandleEvent method of the EventHandler interface.
func (l *Logger) HandleEvent(_ context.Context, event eh.Event) error {
	l.Debug("eh_logger", zap.Any("event", event))

	return nil
}
