package context

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type loggerContextExtractorDeps struct {
	fx.In

	Config cfg.Config
}

// LoggerGRPCIncomingContextExtractor creates a context extractor for logger
// This is useful if you want to add different fields from gRPC incoming metadata.MD to a log entry
func LoggerGRPCIncomingContextExtractor(deps loggerContextExtractorDeps) log.ContextExtractor {
	_ = "STUB: not implemented"
	return *new(log.ContextExtractor)
}

type headerPrefixes []string // if this slice will be very large it's better to build a trie map

func (h headerPrefixes) Extract(ctx context.Context) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
