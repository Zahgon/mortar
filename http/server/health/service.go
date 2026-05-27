package health

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type healthService struct {
	UnimplementedHealthServer
}

// RegisterInternalGRPCGatewayHandler grpc-gateway health handler
func RegisterInternalGRPCGatewayHandler(mux *runtime.ServeMux, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterInternalHealthService grpc server health api registration
func RegisterInternalHealthService(srv *grpc.Server) { _ = "STUB: not implemented"; return }

// ImplementedHealthService internal health service
func ImplementedHealthService() HealthServer { _ = "STUB: not implemented"; return *new(HealthServer) }

func (*healthService) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
