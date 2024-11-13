package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	LoggerKey = "logger"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type logger struct {
	logger *zap.Logger
}

func (l logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func New() Logger {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	return &logger{
		logger: zapLogger,
	}
}

func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}
