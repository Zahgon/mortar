package client

import (
	"container/list"
	"net/http"

	"github.com/go-masonry/mortar/interfaces/http/client"
)

type restBuilderConfig struct {
	predefinedClient *http.Client
	interceptors     []client.HTTPClientInterceptor
}

type builderImpl struct {
	ll *list.List
}

// HTTPClientBuilder creates a fresh http.Client builder
func HTTPClientBuilder() client.HTTPClientBuilder {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientBuilder)
}

func (impl *builderImpl) AddInterceptors(interceptors ...client.HTTPClientInterceptor) client.HTTPClientBuilder {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientBuilder)
}

func (impl *builderImpl) WithPreconfiguredClient(client *http.Client) client.HTTPClientBuilder {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientBuilder)
}

func (impl *builderImpl) Build() *http.Client { _ = "STUB: not implemented"; return nil }

type customRoundTripper struct {
	inner             http.RoundTripper
	unitedInterceptor client.HTTPClientInterceptor
}

func prepareCustomRoundTripper(actual http.RoundTripper, interceptors ...client.HTTPClientInterceptor) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (crt *customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uniteInterceptors(interceptors []client.HTTPClientInterceptor) client.HTTPClientInterceptor {
	_ = "STUB: not implemented"
	return *new(client.HTTPClientInterceptor)
}

// That's why we needed an alias to http.RoundTripper.RoundTrip

var _ client.HTTPClientBuilder = (*builderImpl)(nil)
