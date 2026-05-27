package handlers

import (
	"github.com/go-masonry/mortar/constructors/partial"
)

const (
	// the path of pprof must start with /debug/pprof because of https://github.com/golang/go/issues/14286
	profilePrefix = internalPatternPrefix + "/pprof"
)

// InternalProfileHandlerFunctions profile handlers
func InternalProfileHandlerFunctions() []partial.HTTPHandlerFuncPatternPair {
	_ = "STUB: not implemented"
	return nil
}
