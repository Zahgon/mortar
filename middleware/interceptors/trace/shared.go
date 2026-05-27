package trace

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/fx"
)

type tracingDeps struct {
	fx.In

	Logger log.Logger
	Config cfg.Config
	Tracer opentracing.Tracer `optional:"true"`
}

func addBodyToSpan(span opentracing.Span, name string, msg interface{}) {
	_ = "STUB: not implemented"
	return
}

// TODO: can exceed length limit, introduce option

// If marshaling failed let's try to log msg.ToString()

func (d tracingDeps) extractIncomingCarrier(ctx context.Context) utils.MDTraceCarrier {
	_ = "STUB: not implemented"
	return *new(utils.MDTraceCarrier)
}

// make a copy since this map is not thread safe

func (d tracingDeps) extractOutgoingCarrier(ctx context.Context) utils.MDTraceCarrier {
	_ = "STUB: not implemented"
	return *new(utils.MDTraceCarrier)
}

// make a copy since this map is not thread safe

var grpcTag = opentracing.Tag{Key: string(ext.Component), Value: "gRPC"}
var restTag = opentracing.Tag{Key: string(ext.Component), Value: "REST"}
