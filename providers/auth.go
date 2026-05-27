package providers

import (
	"github.com/go-masonry/mortar/constructors"
	"go.uber.org/fx"
)

// JWTExtractorFxOption adds default JWT extractor from context.Context to the graph
func JWTExtractorFxOption() fx.Option { _ = "STUB: not implemented"; return *new(fx.Option) }

// JWTExtractor is a constructor for Default JWT Token Extractor
//
// Consider using JWTExtractorFxOption if you only want to provide it.
var JWTExtractor = constructors.DefaultJWTTokenExtractor
