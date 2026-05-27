package constructors

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/cfg"
	logInt "github.com/go-masonry/mortar/interfaces/log"

	"go.uber.org/fx"
)

// FxGroupLoggerContextExtractors defines group name
const FxGroupLoggerContextExtractors = "loggerContextExtractors"
const (
	application = "app"
	hostname    = "host"
	gitCommit   = "git"
)
const compensateDefaultLogger = 1

type loggerDeps struct {
	fx.In

	Config            cfg.Config
	LoggerBuilder     logInt.Builder            `optional:"true"`
	ContextExtractors []logInt.ContextExtractor `group:"loggerContextExtractors"`
}

// DefaultLogger is a constructor that will create a logger with some default values on top of provided ones
func DefaultLogger(deps loggerDeps) logInt.Logger {
	_ = "STUB: not implemented"
	return *new(logInt.Logger)
}

func (d loggerDeps) selfStaticFieldsContextExtractor(_ context.Context) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (d loggerDeps) getLogBuilder() logInt.Builder {
	_ = "STUB: not implemented"
	return *new(logInt.Builder)
}
