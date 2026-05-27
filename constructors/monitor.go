package constructors

import (
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/interfaces/monitor"
	"go.uber.org/fx"
)

const (
	// FxGroupMonitorContextExtractors defines group name
	FxGroupMonitorContextExtractors = "monitorContextExtractors"
)

type monitorDeps struct {
	fx.In

	LifeCycle         fx.Lifecycle
	Config            cfg.Config
	Logger            log.Logger
	MonitorBuilder    monitor.Builder
	ContextExtractors []monitor.ContextExtractor `group:"monitorContextExtractors"`
}

// DefaultMonitor is a constructor that will create a Metrics reporter based on values from the Config Map
// such as
//
//   - Tags: we will look for default tags using mortar.MonitorTagsKey within the configuration map
func DefaultMonitor(deps monitorDeps) monitor.Metrics {
	_ = "STUB: not implemented"
	return *new(monitor.Metrics)
}

// can be empty
