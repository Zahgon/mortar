package providers

import (
	"go.uber.org/fx"
)

// TODO Revisit Monitoring

// MonitorFxOption adds default metric client to the graph
func MonitorFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

// MonitorGRPCInterceptorFxOption adds Unary Server Interceptor that will notify metric provider of every call
func MonitorGRPCInterceptorFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }
