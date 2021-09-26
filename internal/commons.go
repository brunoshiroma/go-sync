package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger  *zap.Logger
	LoggerS *zap.SugaredLogger
)

func init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	Logger, _ = loggerConfig.Build()
	LoggerS = Logger.Sugar()
}
