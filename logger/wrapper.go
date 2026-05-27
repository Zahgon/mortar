package logger

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/log"
)

const (
	compensateMortarLoggerWrapper = 1
)

type loggerWrapper struct {
	contextExtractors []log.ContextExtractor
	logger            log.Logger
}

// CreateMortarLogger creates a new mortar logger which is a wrapper to support
//   - ContextExtractors
//
// **Important**
//
//	This constructor will call builder.IncrementSkipFrames to peel additional layer of itself.
func CreateMortarLogger(builder log.Builder, contextExtractors ...log.ContextExtractor) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// add 1

func (l *loggerWrapper) Trace(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) Debug(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) Info(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) Warn(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) Error(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) Custom(ctx context.Context, level log.Level, skipAdditionalFrames int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *loggerWrapper) WithError(err error) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

func (l *loggerWrapper) WithField(name string, value interface{}) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

func (l *loggerWrapper) Configuration() log.LoggerConfiguration {
	_ = "STUB: not implemented"
	return *new(log.LoggerConfiguration)
}
