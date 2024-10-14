package logutil

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLogger = zap.NewNop()

func SetDefaultZapLogger(logger *zap.Logger) {
	DefaultLogger = logger
}

func NewZapLogger(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	cfg := zap.Config{
		Level:         lvl,
		Development:   true,
		Encoding:      "console",
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		OutputPaths:   []string{"stdout"},
	}
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg.Build()
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}
