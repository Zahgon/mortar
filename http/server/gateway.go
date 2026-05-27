package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

type grpcGatewayMuxOptionsDeps struct {
	fx.In

	Tracer opentracing.Tracer `optional:"true"`
}

// MetadataTraceCarrierOption is a nice trick to avoid creating an additional server span from REST to GRPC
// it will extract trace context from Headers and put that information into the Context without creating a new Span
// However if you would like to create a new Span on the REST layer, you should read how to do it here
// https://grpc-ecosystem.github.io/grpc-gateway/docs/customizingyourgateway.html scroll to "OpenTracing Support"
func MetadataTraceCarrierOption(deps grpcGatewayMuxOptionsDeps) runtime.ServeMuxOption {
	_ = "STUB: not implemented"
	return *new(runtime.ServeMuxOption)
}

// we ignore error here, since we assume that a new span will be open by gRPC Trace Interceptor anyway
