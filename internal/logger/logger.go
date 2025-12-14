package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		// Fallback to development logger if production fails
		logger, _ = zap.NewDevelopment()
	}

	return logger
}
