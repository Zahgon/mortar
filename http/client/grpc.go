package client

import (
	"container/list"
	"context"

	"github.com/go-masonry/mortar/interfaces/http/client"
	"google.golang.org/grpc"
)

type grpcClientConnOptions struct {
	options []grpc.DialOption
}

type grpcClientConnBuilder struct {
	ll *list.List
}

// GRPCClientConnBuilder creates a fresh gRPC connection for client builder
func GRPCClientConnBuilder() client.GRPCClientConnectionBuilder {
	_ = "STUB: not implemented"
	return *new(client.GRPCClientConnectionBuilder)
}

func (g *grpcClientConnBuilder) AddOptions(opts ...grpc.DialOption) client.GRPCClientConnectionBuilder {
	_ = "STUB: not implemented"
	return *new(client.GRPCClientConnectionBuilder)
}

func (g *grpcClientConnBuilder) Build() client.GRPCClientConnectionWrapper {
	_ = "STUB: not implemented"
	return *new(client.GRPCClientConnectionWrapper)
}

type grpcClientConnImpl struct {
	options *grpcClientConnOptions
}

func (g *grpcClientConnImpl) Dial(ctx context.Context, target string, extraOptions ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientConnInterface), nil
}
