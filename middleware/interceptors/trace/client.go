package trace

import (
	"context"
	"net/http"

	"github.com/go-masonry/mortar/interfaces/http/client"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// TracerGRPCClientInterceptor is a grpc tracing client interceptor, it can log req/resp if needed
func TracerGRPCClientInterceptor(deps tracingDeps) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// log request if needed

// log response if needed

// TracerRESTClientInterceptor is a REST tracing client interceptor, it can log req/resp if needed
func TracerRESTClientInterceptor(deps tracingDeps) client.HTTPClientInterceptor {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientInterceptor)
}

func (d tracingDeps) newClientSpanForGRPC(ctx context.Context, methodName string) (opentracing.Span, context.Context) {
	_ = "STUB: not implemented"
	return *new(opentracing.Span), *new(context.Context)
}

func (d tracingDeps) newClientSpanForREST(req *http.Request) (opentracing.Span, context.Context) {
	_ = "STUB: not implemented"
	return *new(opentracing.Span), *new(context.Context)
}
