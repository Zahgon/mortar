package naive

import (
	"container/list"
	"io"

	logInt "github.com/go-masonry/mortar/interfaces/log"
)

const defaultSkipDepth = 4

type defaultConfig struct {
	writer        io.Writer
	level         logInt.Level
	depth         int
	excludeTime   bool
	includeCaller bool
}

type defaultBuilder struct {
	ll *list.List
}

// NativeLogBuilder is a helper interface to configure native log.Logger instance.
type NativeLogBuilder interface {
	logInt.Builder
	// SetWriter set where output should be printed
	SetWriter(writer io.Writer) NativeLogBuilder
	// ExcludeTime configures standard Logger to exclude any time field
	ExcludeTime() NativeLogBuilder
	// IncludeCaller adds caller:line to the output
	IncludeCaller() NativeLogBuilder
}

// Builder creates a fresh default Logger builder, this will eventually build a std logger wrapper without structured logging
func Builder() NativeLogBuilder { _ = "STUB: not implemented"; return *new(NativeLogBuilder) }

func (d *defaultBuilder) SetWriter(writer io.Writer) NativeLogBuilder {
	_ = "STUB: not implemented"
	return *new(NativeLogBuilder)
}

func (d *defaultBuilder) ExcludeTime() NativeLogBuilder {
	_ = "STUB: not implemented"
	return *new(NativeLogBuilder)
}

func (d *defaultBuilder) IncludeCaller() NativeLogBuilder {
	_ = "STUB: not implemented"
	return *new(NativeLogBuilder)
}

func (d *defaultBuilder) IncrementSkipFrames(inc int) logInt.Builder {
	_ = "STUB: not implemented"
	return *new(logInt.Builder)
}

func (d *defaultBuilder) SetLevel(level logInt.Level) logInt.Builder {
	_ = "STUB: not implemented"
	return *new(logInt.Builder)
}

func (d *defaultBuilder) Build() logInt.Logger {
	_ = "STUB: not implemented"
	return *new(logInt.Logger)
}

// 2 is used within the log package

var _ logInt.Builder = (*defaultBuilder)(nil)
