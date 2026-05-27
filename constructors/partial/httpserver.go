package partial

import (
	"net/http"

	"github.com/go-masonry/mortar/interfaces/cfg"
	serverInt "github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/interfaces/monitor"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// PanicHandlerCounter is the metric name to count all recovered panics
const PanicHandlerCounter = "panic_handler_total"

// Group order is not guaranteed, if it's important then add them manually
const (
	// FxGroupGRPCServerAPIs defines group name
	FxGroupGRPCServerAPIs = "grpcServerAPIs"
	// FxGroupGRPCGatewayGeneratedHandlers defines group name
	FxGroupGRPCGatewayGeneratedHandlers = "grpcGatewayGeneratedHandlers"
	// FxGroupGRPCGatewayMuxOptions defines group name
	FxGroupGRPCGatewayMuxOptions = "grpcGatewayMuxOptions"
	// FxGroupExternalHTTPHandlers defines group name
	FxGroupExternalHTTPHandlers = "externalHttpHandlers"
	// FxGroupExternalHTTPHandlerFunctions defines group name
	FxGroupExternalHTTPHandlerFunctions = "externalHttpHandlerFunctions"
	// FxGroupExternalHTTPInterceptors defines group name
	FxGroupExternalHTTPInterceptors = "externalHttpInterceptors"
	// FxGroupUnaryServerInterceptors defines group name
	FxGroupUnaryServerInterceptors = "unaryServerInterceptors"
	// FxGroupUnaryServerInterceptors defines group name
	FxGroupStreamServerInterceptors = "streamServerInterceptors"
	// FxGroupInternalHTTPHandlers defines group name
	FxGroupInternalHTTPHandlers = "internalHttpHandlers"
	// FxGroupInternalHTTPHandlerFunctions defines group name
	FxGroupInternalHTTPHandlerFunctions = "internalHttpHandlerFunctions"
	// FxGroupInternalHTTPInterceptors defines group name
	FxGroupInternalHTTPInterceptors = "internalHttpInterceptors"
)

// HTTPHandlerPatternPair defines pattern -> handler pair
type HTTPHandlerPatternPair struct {
	Pattern string
	Handler http.Handler
}

// HTTPHandlerFuncPatternPair defines patter -> handler func pair
type HTTPHandlerFuncPatternPair struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type httpServerDeps struct {
	fx.In

	Config  cfg.Config
	Logger  log.Logger
	Metrics monitor.Metrics `optional:"true"`
	// GRPC
	GRPCServerAPIs     []serverInt.GRPCServerAPI      `group:"grpcServerAPIs"`
	UnaryInterceptors  []grpc.UnaryServerInterceptor  `group:"unaryServerInterceptors"`
	StreamInterceptors []grpc.StreamServerInterceptor `group:"streamServerInterceptors"`
	// External REST
	GRPCGatewayGeneratedHandlers []serverInt.GRPCGatewayGeneratedHandlers `group:"grpcGatewayGeneratedHandlers"`
	GRPCGatewayMuxOptions        []runtime.ServeMuxOption                 `group:"grpcGatewayMuxOptions"`
	ExternalHTTPHandlers         []HTTPHandlerPatternPair                 `group:"externalHttpHandlers"`
	ExternalHTTPHandlerFunctions []HTTPHandlerFuncPatternPair             `group:"externalHttpHandlerFunctions"`
	ExternalHTTPInterceptors     []serverInt.GRPCGatewayInterceptor       `group:"externalHttpInterceptors"`
	// Internal REST
	InternalHTTPHandlers         []HTTPHandlerPatternPair           `group:"internalHttpHandlers"`
	InternalHTTPHandlerFunctions []HTTPHandlerFuncPatternPair       `group:"internalHttpHandlerFunctions"`
	InternalHTTPInterceptors     []serverInt.GRPCGatewayInterceptor `group:"internalHttpInterceptors"`
}

// HTTPServerBuilder true to it's name, it is partially initialized builder.
//
// It uses some default assumptions and configurations, which are mostly good.
// However, if you need to customize your configuration it's better to build yours from scratch
func HTTPServerBuilder(deps httpServerDeps) serverInt.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(serverInt.GRPCWebServiceBuilder)
}

// GRPC port

// GRPC unary server interceptors

// GRPC stream server interceptors

func (deps httpServerDeps) buildExternalAPI(builder serverInt.GRPCWebServiceBuilder) serverInt.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(serverInt.GRPCWebServiceBuilder)
}

// register grpc APIs

// add GRPC Gateway on top and expose on external REST Port

func (deps httpServerDeps) buildInternalAPI(builder serverInt.GRPCWebServiceBuilder) serverInt.GRPCWebServiceBuilder {
	_ = "STUB: not implemented"
	return *new(serverInt.GRPCWebServiceBuilder)
}

// add internal GRPC health endpoint
// Internal

// Health

func (deps httpServerDeps) panicHandler(r interface{}) error { _ = "STUB: not implemented"; return nil }
