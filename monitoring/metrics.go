package monitoring

import (
	"github.com/go-masonry/mortar/interfaces/monitor"
)

type mortarMetric struct {
	*tagsMetric
	registry *externalRegistry
	cfg      *monitorConfig
}

func newMetric(registry *externalRegistry, cfg *monitorConfig) monitor.Metrics {
	_ = "STUB: not implemented"
	return *new(monitor.Metrics)
}

// Counter creates a counter with possible predefined tags
func (mm *mortarMetric) Counter(name, desc string) monitor.TagsAwareCounter {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareCounter)
}

// Gauge creates a gauge with possible predefined tags
func (mm *mortarMetric) Gauge(name, desc string) monitor.TagsAwareGauge {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareGauge)
}

// Histogram creates a histogram with possible predefined tags
func (mm *mortarMetric) Histogram(name, desc string, buckets monitor.Buckets) monitor.TagsAwareHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareHistogram)
}

// Timer creates a timer with possible predefined tags
func (mm *mortarMetric) Timer(name, desc string) monitor.TagsAwareTimer {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareTimer)
}

// WithTags sets custom tags to be included if possible in every Metric
func (mm *mortarMetric) WithTags(tags monitor.Tags) monitor.Metrics {
	_ = "STUB: not implemented"
	return *new(monitor.Metrics)
}

func (mm *mortarMetric) extractTagKeys() (keys []string) { _ = "STUB: not implemented"; return nil }
