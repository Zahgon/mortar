package naive

import (
	"context"
	"log"

	logInt "github.com/go-masonry/mortar/interfaces/log"
)

const (
	noAdditionalFramesToSkip = 0
)

type defaultLogger struct {
	cfg    *defaultConfig
	logger *log.Logger
}

func (d *defaultLogger) Level() logInt.Level { _ = "STUB: not implemented"; return *new(logInt.Level) }

func (d *defaultLogger) Implementation() interface{} { _ = "STUB: not implemented"; return nil }

func (d *defaultLogger) Trace(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultLogger) Info(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultLogger) Error(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultLogger) Custom(ctx context.Context, level logInt.Level, skipAdditionalFrames int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// WithError not supported
func (d *defaultLogger) WithError(err error) logInt.Fields {
	_ = "STUB: not implemented"

	// WithField not supported
	return *new(logInt.Fields)
}

func (d *defaultLogger) WithField(name string, value interface{}) logInt.Fields {
	_ = "STUB: not implemented"
	return *new(logInt.Fields)
}

func (d *defaultLogger) Configuration() logInt.LoggerConfiguration {
	_ = "STUB: not implemented"
	return *new(logInt.LoggerConfiguration)
}

func (d *defaultLogger) log(skipAdditionalFrames int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func newDefaultLogger(cfg *defaultConfig) logInt.Logger {
	_ = "STUB: not implemented"
	return *new(logInt.Logger)
}

var _ logInt.Logger = (*defaultLogger)(nil)
