package providers

import (
	"github.com/go-masonry/mortar/constructors"
	"github.com/go-masonry/mortar/constructors/partial"
	"github.com/go-masonry/mortar/middleware/grpcgateway"
	"github.com/go-masonry/mortar/middleware/interceptors/client"
	"go.uber.org/fx"
)

// BuildMortarWebServiceFxOption creates the entire dependency graph
// and registers all provided fx.LifeCycle hooks
func BuildMortarWebServiceFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

// BuildMortarWebService is a constructor that creates and registers fx.LifeCycle hooks for Mortar web services
//
// Consider using BuildMortarWebServiceFxOption if you only want to invoke it.
var BuildMortarWebService = constructors.Service

// HTTPServerBuilderFxOption adds Default Http Server builder which later injected to the Service Invoke option
// by calling BuildMortarWebServiceFxOption fx.Invoke option to the graph
func HTTPServerBuilderFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

// HTTPServerBuilder is a constructor that creates a partial Mortars HTTP Server Builder
//
// Consider using HTTPServerBuilderFxOption if you only want to provide it.
var HTTPServerBuilder = partial.HTTPServerBuilder

// HTTPClientBuildersFxOption adds both (GRPC, REST) partial http clients to the graph
func HTTPClientBuildersFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

// HTTPClientBuilder is a constructor that creates HTTP Client builder
//
// Consider using HTTPClientBuildersFxOption if you only want to provide it.
var HTTPClientBuilder = partial.HTTPClientBuilder

// GRPCClientConnectionBuilder is a constructor that creates gRPC Connection for client builder
//
// Consider using HTTPClientBuildersFxOption if you only want to provide it.
var GRPCClientConnectionBuilder = partial.GRPCClientConnectionBuilder

// CopyGRPCHeadersClientInterceptorFxOption adds grpc Client Interceptor that copies values from grpc Incoming to Outgoing metadata
func CopyGRPCHeadersClientInterceptorFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// CopyGRPCHeadersClientInterceptor is a constructor that creates gRPC Unary Client Interceptor
//
//	*Note* normally this dependency is part of a group. If you want to create it as a standalone
//	dependency, remember that there can be only one of this kind in the graph.
//
// Consider using CopyGRPCHeadersClientInterceptorFxOption if you only want to provide it.
var CopyGRPCHeadersClientInterceptor = client.CopyGRPCHeadersClientInterceptor

// CopyGRPCHeadersHTTPClientInterceptorFxOption copies filtered Headers found in the Incoming GRPC metadata.MD to the Outgoing HTTP Request Headers.
//
// This is useful if you want to propagate them to the next service when using `http.Client`
//
// For Example: "authorization" header containing user token
func CopyGRPCHeadersHTTPClientInterceptorFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// CopyGRPCHeadersHTTPClientInterceptor is a constructor that creates REST Client Interceptor
//
//	*Note* normally this dependency is part of a group. If you want to create it as a standalone
//	dependency, remember that there can be only one of this kind in the graph.
//
// Consider using CopyGRPCHeadersHTTPClientInterceptorFxOption if you only want to provide it.
var CopyGRPCHeadersHTTPClientInterceptor = client.CopyGRPCHeadersHTTPClientInterceptor

// DumpRESTClientInterceptorFxOption usefull when you want to log what is actually sent to the external HTTP server
// and what was returned.
func DumpRESTClientInterceptorFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// DumpRESTClientInterceptor is a constructor that creates REST Client Interceptor
//
//	*Note* normally this dependency is part of a group. If you want to create it as a standalone
//	dependency, remember that there can be only one of this kind in the graph.
//
// Consider using DumpRESTClientInterceptorFxOption if you only want to provide it.
var DumpRESTClientInterceptor = client.DumpRESTClientInterceptor

// MonitorGRPCClientCallsInterceptorFxOption usefull when you want to monitor all your unary gRPC Client calls
func MonitorGRPCClientCallsInterceptorFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// MonitorGRPCClientCallsInterceptor is a constructor that creates Unary gRPC Client Interceptor
//
//	*Note* normally this dependency is part of a group. If you want to create it as a standalone
//	dependency, remember that there can be only one of this kind in the graph.
//
// Consider using MonitorGRPCClientCallsInterceptorFxOption if you only want to provide it.
var MonitorGRPCClientCallsInterceptor = client.MonitorGRPCClientCallsInterceptor

// MonitorRESTClientCallsInterceptorFxOption usefull when you want to monitor all your REST Client calls
func MonitorRESTClientCallsInterceptorFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// MonitorRESTClientCallsInterceptor is a constructor that creates REST Client Interceptor
//
//	*Note* normally this dependency is part of a group. If you want to create it as a standalone
//	dependency, remember that there can be only one of this kind in the graph.
//
// Consider using MonitorRESTClientCallsInterceptorFxOption if you only want to provide it.
var MonitorRESTClientCallsInterceptor = client.MonitorRESTClientCallsInterceptor

// MapHTTPHeadersToClientMetadataMuxOptionFxOption adds a Grpc Gateway server mux option
// that maps incoming HTTP Headers to gRPC Context by checking if they match a list of prefixes.
// List of prefixes is controlled by config key: `mortar.middleware.map.httpHeaders`
func MapHTTPHeadersToClientMetadataMuxOptionFxOption() fx.Option {
	_ = "STUB: not implemented"
	return *new(fx.Option)
}

// MapHTTPHeadersToClientMetadataMuxOption is a Grpc Gateway server mux option
// that maps incoming HTTP Headers to gRPC Context by checking if they match a list of prefixes.
// List of prefixes is controlled by config key: `mortar.middleware.map.httpHeaders`
//
// Consider using MapHTTPHeadersToClientMetadataMuxOptionFxOption if you only want to provide it.
var MapHTTPHeadersToClientMetadataMuxOption = grpcgateway.MapHTTPHeadersToClientMetadataMuxOption
