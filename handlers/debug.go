package handlers

import (
	"net/http"
	"runtime"

	"github.com/go-masonry/mortar/constructors/partial"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

const (
	internalPatternPrefix = "/debug"
)

// StatsInfo some statistics information
type StatsInfo struct {
	Memory          *runtime.MemStats `json:"memory"`
	NumOfCPU        int               `json:"num_of_cpu"`
	NumOfGoRoutines int               `json:"num_of_go_routines"`
}

// DebugHandlers different debug handlers
type DebugHandlers interface {
	DebugVars() http.Handler
	Stats() http.HandlerFunc
	DumpFunc() http.HandlerFunc
}

type debugHandlersDeps struct {
	fx.In

	Logger log.Logger
}

// InternalDebugHandlers defines internal debug handlers
//   - dump heap
//   - expvar
//   - running stats
func InternalDebugHandlers(deps debugHandlersDeps) []partial.HTTPHandlerPatternPair {
	_ = "STUB: not implemented"
	return nil
}

func (d *debugHandlersDeps) DebugVars() http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (d *debugHandlersDeps) DumpFunc() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// remove garbage

func (d *debugHandlersDeps) Stats() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
