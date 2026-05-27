package handlers

import (
	"net/http"

	"github.com/go-masonry/mortar/constructors/partial"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

const (
	obfuscationEdgeLength = 4
	selfHandlerPrefix     = "/self"
)

type selfHandlerDeps struct {
	fx.In

	Logger log.Logger
	Config cfg.Config
}

// SelfHandlers this service information handlers
func SelfHandlers(deps selfHandlerDeps) []partial.HTTPHandlerPatternPair {
	_ = "STUB: not implemented"
	return nil
}

func (s *selfHandlerDeps) BuildInfo() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (s *selfHandlerDeps) ConfigMap() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (s *selfHandlerDeps) getConfigVariables() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *selfHandlerDeps) obfuscateMapWhereNeeded(prefix string, confMap map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *selfHandlerDeps) getEnvVariables() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (s *selfHandlerDeps) obfuscateIfNeeded(key string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// if none exist slice will be empty
