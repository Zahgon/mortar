package server

import (
	"context"
	"net"
	"sync"

	"github.com/go-masonry/mortar/interfaces/http/server"
	"google.golang.org/grpc"
)

type listenerMuxPair struct {
	m mux
	l net.Listener
}

type webService struct {
	sync.Mutex
	serviceConfig   *webServiceConfig
	grpcServer      *grpc.Server
	grpcAddr        string
	muxAndListeners []*listenerMuxPair
	close           bool
}

func newWebService(cfg *webServiceConfig) (instance server.WebService, err error) {
	_ = "STUB: not implemented"
	return *new(server.WebService), nil
}

// make sure to clean, since there might still be open listeners

func (ws *webService) Run(context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (ws *webService) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// already closed

func (ws *webService) Ports() (list []server.ListenInfo) { _ = "STUB: not implemented"; return nil }

func (ws *webService) setupGRPC(cfg *grpcConfig) (err error) { _ = "STUB: not implemented"; return nil }

// Listener

// Server

// save, since this should run first we have no problem with previous values

// we need this later for grpc gateway

func (ws *webService) setupREST(restConfigs []*restConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// indicate that we have some kind of handler here, grpcgateway or custom handler/handlerfunc

// Listener

// Server

// register handlers

// register handler functions

// GRPC Gateway

// Check if the root '/' pattern is taken

// register grpc gateway handlers

// check if we have configured anything

// Save

// Sanity

var _ server.WebService = (*webService)(nil)
