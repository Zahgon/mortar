package client

import (
	"github.com/go-masonry/mortar/interfaces/http/client"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type dumpHTTPDeps struct {
	fx.In

	Logger log.Logger
}

// DumpRESTClientInterceptor usefull when you want to log what is actually sent to the external HTTP server
// and was returned.
func DumpRESTClientInterceptor(deps dumpHTTPDeps) client.HTTPClientInterceptor {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientInterceptor)
}
