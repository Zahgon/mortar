package grpcgateway

import (
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
)

type grpcGatewayHeadersDeps struct {
	fx.In

	Config cfg.Config
}

// MapHTTPHeadersToClientMetadataMuxOption maps incoming HTTP Headers to gRPC Context by checking if they match a list of prefixes
func MapHTTPHeadersToClientMetadataMuxOption(deps grpcGatewayHeadersDeps) runtime.ServeMuxOption {
	_ = "STUB: not implemented"
	return *new(runtime.ServeMuxOption)
}

// `key` is already canonicalized by Grpc-Gateway before calling this code
