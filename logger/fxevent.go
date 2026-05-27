package logger

import (
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx/fxevent"
)

// CreateFxEventLogger is a constructor to create fxevent.Logger
// This one is used by fx itself to output its events
func CreateFxEventLogger(logger log.Logger, cfg cfg.Config) fxevent.Logger {
	_ = "STUB: not implemented"
	return *new(fxevent.Logger)
}

type logWrapper struct {
	log.Logger
	startStopLogLevel log.Level
}

func (zw *logWrapper) LogEvent(event fxevent.Event) { _ = "STUB: not implemented"; return }

// Do nothing. Will log on Invoked.
