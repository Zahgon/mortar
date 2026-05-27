package server

import (
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/interfaces/monitor"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

const (
	gRPCCodeTagName = "code"
	grpcNamePrefix  = "grpc_"
)

type gRPCMetricInterceptorsDeps struct {
	fx.In

	Logger  log.Logger
	Metrics monitor.Metrics `optional:"true"`
}

// MonitorGRPCInterceptor sends gRPC method invocation metrics to the configured Metrics server (Prometheus, Datadog)
func MonitorGRPCInterceptor(deps gRPCMetricInterceptorsDeps) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// fetch one from registry or create new

func gRPCCodeTagValue(err error) string { _ = "STUB: not implemented"; return "" }
