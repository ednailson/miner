package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup() {
	// This should be configurable in a production application
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	log, _ := config.Build()
	zap.ReplaceGlobals(log)
}
