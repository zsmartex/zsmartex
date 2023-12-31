package mongo_fx

import (
	"go.uber.org/zap"

	"github.com/zsmartex/zsmartex/pkg/logger"
)

type Logger struct {
	*logger.Logger
}

func NewLogger(logger *logger.Logger) *Logger {
	return &Logger{
		logger,
	}
}

func (logger *Logger) Info(level int, msg string, keyandvalues ...interface{}) {
	fields := []zap.Field{}
	for i := 0; i < len(keyandvalues); i += 2 {
		key, ok := keyandvalues[i].(string)
		if !ok {
			continue
		}
		value := keyandvalues[i+1]
		fields = append(fields, zap.Any(key, value))
	}

	logger.Logger.Debug(msg, fields...)
}

func (logger *Logger) Error(err error, msg string, keyandvalues ...interface{}) {
	fields := []zap.Field{}

	fields = append(fields, zap.Error(err))
	for i := 0; i < len(keyandvalues); i += 2 {
		key, ok := keyandvalues[i].(string)
		if !ok {
			continue
		}
		value := keyandvalues[i+1]
		fields = append(fields, zap.Any(key, value))
	}

	logger.Logger.Error(msg, fields...)
}
