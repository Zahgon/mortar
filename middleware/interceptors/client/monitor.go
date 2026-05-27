package client

import (
	"github.com/go-masonry/mortar/interfaces/http/client"
	"github.com/go-masonry/mortar/interfaces/monitor"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Names
const (
	ClientTimerMetric            = "client_calls_duration"
	ClientTimerMetricDescription = "Monitor external HTTP client calls"
	TargetTag                    = "target"
	PathTag                      = "path"
	SuccessTag                   = "success"
	TypeTag                      = "ctype"
	TypeGRPC                     = "grpc"
	TypeREST                     = "rest"
)

type monitorDeps struct {
	fx.In

	Metrics monitor.Metrics `optional:"true"`
}

// MonitorGRPCClientCallsInterceptor create a new GRPC Unary Client interceptor that monitor all external client calls
func MonitorGRPCClientCallsInterceptor(deps monitorDeps) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// MonitorRESTClientCallsInterceptor create a new REST Client interceptor that monitor all external client calls
func MonitorRESTClientCallsInterceptor(deps monitorDeps) client.HTTPClientInterceptor {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientInterceptor)
}

func prepareTags(host, path, clientType, err string) monitor.Tags {
	_ = "STUB: not implemented"
	return *new(monitor.Tags)
}

// remove trailing port if exists
