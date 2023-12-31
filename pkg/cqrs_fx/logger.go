package cqrs_fx

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/zsmartex/zsmartex/pkg/logger"
)

var LoggerModule = fx.Provide(
	NewLogger,
)

type Logger struct {
	*logger.Logger
}

func NewLogger(logger *logger.Logger) *Logger {
	return &Logger{logger}
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	return &Logger{
		l.Logger.With(logFields...),
	}
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Debug(msg, logFields...)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Debug(msg, logFields...)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Info(msg, logFields...)
}

func (l *Logger) Warn(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Warn(msg, logFields...)
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	logFields = append(logFields, zap.Error(err))

	l.Logger.Error(msg, logFields...)
}

func (l *Logger) Fatal(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Fatal(msg, logFields...)
}

func (l *Logger) Panic(msg string, fields watermill.LogFields) {
	logFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		logFields = append(logFields, zap.Any(k, v))
	}

	l.Logger.Panic(msg, logFields...)
}
