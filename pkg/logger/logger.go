package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(
	NewLogger,
)

type Logger struct {
	*zap.Logger
}

func NewLogger() (*Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	zapLogger.Named("cqrs")

	return &Logger{
		Logger: zapLogger,
	}, nil
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}
