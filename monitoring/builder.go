package monitoring

import (
	"container/list"

	"github.com/go-masonry/mortar/interfaces/monitor"
)

type monitorConfig struct {
	tags       monitor.Tags
	extractors []monitor.ContextExtractor
	onError    func(err error)
	reporter   monitor.BricksReporter
}

// WrapperBuilder is a helper builder to define internal Mortar monitoring wrapper
type WrapperBuilder interface {
	// Build builds monitor.Reporter
	Build(bricksBuilder monitor.Builder) monitor.Reporter
	// DoOnError is a helper function to act when receiving an error during Metric creation
	DoOnError(onError func(error)) WrapperBuilder
	// AddExtractors adds ContextExtractors that might override tag values when calling metric functions
	AddExtractors(extractors ...monitor.ContextExtractor) WrapperBuilder
	// SetTags saves defaults tags, these tags will always be included in every metric
	SetTags(tags monitor.Tags) WrapperBuilder
}

type wrapperBuilder struct {
	ll *list.List
}

// Builder creates a WrapperBuilder
func Builder() WrapperBuilder { _ = "STUB: not implemented"; return *new(WrapperBuilder) }

func (b *wrapperBuilder) SetTags(tags monitor.Tags) WrapperBuilder {
	_ = "STUB: not implemented"
	return *new(WrapperBuilder)
}

// make sure tags are always empty, not nil

func (b *wrapperBuilder) AddExtractors(extractors ...monitor.ContextExtractor) WrapperBuilder {
	_ = "STUB: not implemented"
	return *new(WrapperBuilder)
}

func (b *wrapperBuilder) DoOnError(onError func(error)) WrapperBuilder {
	_ = "STUB: not implemented"
	return *new(WrapperBuilder)
}

func (b *wrapperBuilder) Build(bricksBuilder monitor.Builder) monitor.Reporter {
	_ = "STUB: not implemented"
	return *new(monitor.Reporter)
}

var _ WrapperBuilder = (*wrapperBuilder)(nil)
