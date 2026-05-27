package server

import (
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type loggerInterceptorDeps struct {
	fx.In

	Config cfg.Config
	Logger log.Logger
}

// LoggerGRPCInterceptor logging interceptor, it will log grpc server call with request/response if configured
func LoggerGRPCInterceptor(deps loggerInterceptorDeps) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// log request if needed

// log response if needed

func addBodyToLogger(entry log.Fields, name string, i interface{}) log.Fields {
	_ = "STUB: not implemented"
	return *new(log.Fields)
}
