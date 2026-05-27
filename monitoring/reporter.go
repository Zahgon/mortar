package monitoring

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/monitor"
)

type mortarReporter struct {
	externalMetrics monitor.BricksMetrics
	cfg             *monitorConfig
	registry        *externalRegistry
}

// NewMortarReporter creates a new mortar monitoring reporter which is a wrapper to support
//   - ContextExtractors
//   - Default Tags, for example: {"version":"v1.0.1", "service":"awesome"}
//
// Meaning, it is possible to also extract tag values from the context, this is useful when the value is set per request/call within the context.Context:
//   - Canary release https://martinfowler.com/bliki/CanaryRelease.html identifier
//   - Authentication Token values, but avoid using high cardinality values such as UserID
func newMortarReporter(cfg *monitorConfig) monitor.Reporter {
	_ = "STUB: not implemented"
	return *new(monitor.Reporter)
}

func (r *mortarReporter) Connect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *mortarReporter) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *mortarReporter) Metrics() monitor.Metrics {
	_ = "STUB: not implemented"

	// Counter creates a counter with possible predefined tags
	return *new(monitor.Metrics)
}

func (r *mortarReporter) Counter(name string, desc string) monitor.TagsAwareCounter {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareCounter)
}

// Gauge creates a gauge with possible predefined tags
func (r *mortarReporter) Gauge(name string, desc string) monitor.TagsAwareGauge {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareGauge)
}

// Histogram creates a histogram with possible predefined tags
func (r *mortarReporter) Histogram(name string, desc string, buckets monitor.Buckets) monitor.TagsAwareHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareHistogram)
}

// Timer creates a timer with possible predefined tags
func (r *mortarReporter) Timer(name string, desc string) monitor.TagsAwareTimer {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareTimer)
}

// WithTags sets custom tags to be included if possible in every Metric
func (r *mortarReporter) WithTags(tags monitor.Tags) monitor.Metrics {
	_ = "STUB: not implemented"
	return *new(monitor.Metrics)
}

// first apply default tags
// then apply custom ones
