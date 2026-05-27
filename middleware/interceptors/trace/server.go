package trace

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// GRPCTracingUnaryServerInterceptor is a grpc unary server interceptor that adds trace information of the invoked grpc method and starts a new span
func GRPCTracingUnaryServerInterceptor(deps tracingDeps) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// log request if needed

// call handler

// log response if needed

func (d tracingDeps) newServerSpan(ctx context.Context, methodName string) (opentracing.Span, context.Context) {
	_ = "STUB: not implemented"
	return *new(opentracing.Span), *new(context.Context)
}

// really low level information in my opinion
