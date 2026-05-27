package logger

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/log"
)

const (
	noAdditionalSkipFrames = 0
)

type contextAwareLogEntry struct {
	contextExtractors []log.ContextExtractor
	innerLogger       log.Fields
	fields            map[string]interface{}
	err               error
	withFields        bool
}

func newEntry(contextExtractors []log.ContextExtractor, logger log.Fields, withFields bool) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

func (c *contextAwareLogEntry) Trace(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) Debug(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) Info(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) Warn(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) Error(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) Custom(ctx context.Context, level log.Level, skipAdditionalFrames int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *contextAwareLogEntry) WithError(err error) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

func (c *contextAwareLogEntry) WithField(name string, value interface{}) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}

func (c *contextAwareLogEntry) log(ctx context.Context, level log.Level, skipAdditionalFrames int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// if no fields, we have one less layer to peel

func (c *contextAwareLogEntry) enrich(ctx context.Context) (logger log.Fields) {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}
