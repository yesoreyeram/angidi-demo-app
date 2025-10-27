package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.Logger for structured logging
type Logger struct {
	*zap.SugaredLogger
}

// New creates a new logger instance
func New(level, format string) (*Logger, error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	var config zap.Config
	if format == "json" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(zapLevel)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return &Logger{
		SugaredLogger: logger.Sugar(),
	}, nil
}

// WithField returns a new logger with an additional field
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		SugaredLogger: l.SugaredLogger.With(key, value),
	}
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	if l.SugaredLogger != nil {
		return l.SugaredLogger.Sync()
	}
	return nil
}

// Default returns a default logger for testing or when config isn't available
func Default() *Logger {
	logger, _ := zap.NewDevelopment()
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// init sets up default logger
func init() {
	// Ensure we have a default logger for package-level logging
	if os.Getenv("DEBUG") != "" {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	}
}
