package constructors

import (
	"context"

	jwtInt "github.com/go-masonry/mortar/interfaces/auth/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	authorizationHeader                = "authorization"
	grpcGatewayAuthorizationWithPrefix = runtime.MetadataPrefix + "authorization"
)

// DefaultJWTTokenExtractor simple TokenExtractor
func DefaultJWTTokenExtractor() jwtInt.TokenExtractor {
	_ = "STUB: not implemented"
	return *new(jwtInt.TokenExtractor)
}

// Handles use cases where 'authorization' header value is
//
//	bearer <token>
//	basic <token>
func contextExtractorAuthWithBearer(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
