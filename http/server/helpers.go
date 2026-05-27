package server

import (
	"context"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

// Interfaces
type mux interface {
	Serve(listener net.Listener) error
}

type grpcServerStopper interface {
	GracefulStop()
	Stop()
}

type restServerShutdown interface {
	Shutdown(ctx context.Context) error
	Close() error
}

func defaultPanicHandler(r interface{}) error { _ = "STUB: not implemented"; return nil }

func panicHandlerUnaryInterceptor(panicHandler func(interface{}) error) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func panicHandlerStreamInterceptor(panicHandler func(interface{}) error) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func createListener(network, addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

type muxHandler interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

func extractPort(addr string) int { _ = "STUB: not implemented"; return 0 }
