package utils

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// ErrorMapper is a map function that maps HTTP Status Code into its gRPC counter part.
type ErrorMapper func(statusCode int) *status.Status

// ProtobufHTTPClient is a helper util in situations where you want to call a REST API, but you have all the definitions as Protobuf.
type ProtobufHTTPClient interface {
	// Example:
	//	var response *pbpkg.ResponseMessage
	//	var request = &pbpkg.RequestMessage{Name: "test"}
	//	err := Do(ctx, http.MethodPost, "http://host/path", request, &response)
	//
	// Error returned will be of type `(*status.Status).Err()` meaning it will be a gRPC type error.
	//
	// Note:
	// you can pass `nil` as the last parameter if you don't want to unmarshal HTTP response body.
	// or you know it's going to be empty == EOF
	//	err := Do(ctx, http.MethodPost, "http://host/path", request, nil)
	Do(ctx context.Context, method, url string, in proto.Message, out interface{}) error
}

// DefaultProtobufHTTPClient uses default http Client, error mapper and marshaller
var DefaultProtobufHTTPClient = CreateProtobufHTTPClient(nil, nil, nil)

// CreateProtobufHTTPClient Creates a custom Protobuf aware HTTP client
func CreateProtobufHTTPClient(client *http.Client, errorMapper ErrorMapper, marshaller runtime.Marshaler) ProtobufHTTPClient {
	_ = "STUB: not implemented"
	return *new(ProtobufHTTPClient)
}

type protobufHTTPClientImpl struct {
	client      *http.Client
	errorMapper ErrorMapper
	marshaller  runtime.Marshaler
}

// ProtoToHTTPRequest is a helper to convert proto Message into an HTTP Request
func (impl *protobufHTTPClientImpl) Do(ctx context.Context, method, url string, in proto.Message, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to unmarshal the body if it's a "nil interface"
// https://golang.org/doc/faq#nil_error

func defaultErrorMapper(httpStatus int) *status.Status { _ = "STUB: not implemented"; return nil }
