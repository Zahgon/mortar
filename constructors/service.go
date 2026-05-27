package constructors

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type webServiceDependencies struct {
	fx.In

	LifeCycle         fx.Lifecycle
	Logger            log.Logger
	WebServiceBuilder server.GRPCWebServiceBuilder
}

// Service should be invoked by FX, it will build the entire dependencies graph and add lifecycle hooks
func Service(deps webServiceDependencies) (server.WebService, error) {
	_ = "STUB: not implemented"
	return *new(server.WebService), nil
}

// this should exit only when service was shutdown

func (deps webServiceDependencies) pingService(ctx context.Context, service server.WebService) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (deps webServiceDependencies) getGRPCAddress(ports []server.ListenInfo) string {
	_ = "STUB: not implemented"
	return ""
}
