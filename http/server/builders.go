package server

import (
	"container/list"
	"context"
	"net"
	"net/http"

	"github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// ******************************************************************************************************************************************************
// ***************************************************************************REST BUILDER***************************************************************
// ******************************************************************************************************************************************************

type restConfig struct {
	addr                    string
	server                  *http.Server
	listener                net.Listener
	handlers                map[string]http.Handler
	handlerFuncs            map[string]http.HandlerFunc
	grpcGatewayMux          *runtime.ServeMux
	grpcGatewayHandlers     []server.GRPCGatewayGeneratedHandlers
	grpcGatewayOptions      []runtime.ServeMuxOption
	grpcGatewayInterceptors []server.GRPCGatewayInterceptor
}

type restBuilder struct {
	parent server.GRPCWebServiceBuilder
	cfg    *restConfig
	ll     *list.List
}

func newRESTBuilder(cfg *restConfig, parent server.GRPCWebServiceBuilder) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) ListenOn(addr string) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) SetCustomServer(server *http.Server) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) SetCustomListener(listener net.Listener) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) AddHandler(pattern string, handler http.Handler) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) AddHandlerFunc(pattern string, handlerFunc http.HandlerFunc) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) SetCustomGRPCGatewayMux(mux *runtime.ServeMux) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) RegisterGRPCGatewayHandlers(handlers ...server.GRPCGatewayGeneratedHandlers) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) AddGRPCGatewayOptions(options ...runtime.ServeMuxOption) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) AddGRPCGatewayInterceptors(interceptors ...server.GRPCGatewayInterceptor) server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (r *restBuilder) BuildRESTPart() server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

// ******************************************************************************************************************************************************
// ***************************************************************************GRPC BUILDER***************************************************************
// ******************************************************************************************************************************************************

type grpcConfig struct {
	addr         string
	server       *grpc.Server
	listener     net.Listener
	registerAPI  []server.GRPCServerAPI
	options      []grpc.ServerOption
	panicHandler func(interface{}) error
}

type webServiceConfig struct {
	grpc   *grpcConfig
	rest   []*restConfig
	logger func(ctx context.Context, format string, args ...interface{})
}

type serviceBuilder struct {
	ll *list.List
}

// Builder creates a new gRPC web service builder, call it if you want to custom define your web services
func Builder() server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) ListenOn(addr string) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) SetCustomGRPCServer(server *grpc.Server) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) SetCustomListener(listener net.Listener) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) RegisterGRPCAPIs(apis ...server.GRPCServerAPI) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) AddGRPCServerOptions(options ...grpc.ServerOption) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) SetPanicHandler(handler func(interface{}) error) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) SetLogger(logger func(ctx context.Context, format string, args ...interface{})) server.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(server.GRPCWebServiceBuilder)
}

func (s *serviceBuilder) AddRESTServerConfiguration() server.RESTBuilder {
	_ = "STUB: not implemented"
	return *new(server.RESTBuilder)
}

func (s *serviceBuilder) Build() (server.WebService, error) {
	_ = "STUB: not implemented"
	return *new(server.WebService), nil
}

// no log

// make sure they are outer most

// Sanity
var _ server.GRPCWebServiceBuilder = (*serviceBuilder)(nil)
var _ server.RESTBuilder = (*restBuilder)(nil)
