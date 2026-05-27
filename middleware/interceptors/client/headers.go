package client

import (
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/http/client"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type copyHeadersDeps struct {
	fx.In

	Config cfg.Config
}

// CopyGRPCHeadersClientInterceptor copies filtered Headers found in the Incoming metadata.MD to the Outgoing one.
//
// # This is useful if you want to propagate them to the next service when using grpc Client
//
// For Example: "authorization" header containing user token
func CopyGRPCHeadersClientInterceptor(deps copyHeadersDeps) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// CopyGRPCHeadersHTTPClientInterceptor copies filtered Headers found in the Incoming metadata.MD to the Outgoing Request Headers.
//
// This is useful if you want to propagate them to the next service when using `http.Client`
//
// For Example: "authorization" header containing user token
func CopyGRPCHeadersHTTPClientInterceptor(deps copyHeadersDeps) client.HTTPClientInterceptor {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientInterceptor)
}

// Remember the key will be canonicalized by `http.CanonicalHeaderKey`
