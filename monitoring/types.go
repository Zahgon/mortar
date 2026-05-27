package monitoring

import (
	"context"
	"sync"
	"time"

	"github.com/go-masonry/mortar/interfaces/monitor"
)

// *******************************************************************
// *                             Counter                             *
// *******************************************************************
type counter struct {
	*tagsMetric
	bricksCounter monitor.BricksCounter
	extractors    []monitor.ContextExtractor
}

func (c *counter) Inc() { _ = "STUB: not implemented"; return }

func (c *counter) Add(v float64) { _ = "STUB: not implemented"; return }

func (c *counter) WithTags(tags monitor.Tags) monitor.TagsAwareCounter {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareCounter)
}

func (c *counter) WithContext(ctx context.Context) monitor.TagsAwareCounter {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareCounter)
}

// *******************************************************************
// *                             Gauge                               *
// *******************************************************************
type gauge struct {
	*tagsMetric
	bricksGauge monitor.BricksGauge
	extractors  []monitor.ContextExtractor
}

// Set sets Gauge value
func (g *gauge) Set(v float64) { _ = "STUB: not implemented"; return }

// Add adds (or subtracts if negative) from previously set value
func (g *gauge) Add(v float64) { _ = "STUB: not implemented"; return }

// Inc adds 1
func (g *gauge) Inc() { _ = "STUB: not implemented"; return }

// Dec adds -1
func (g *gauge) Dec() { _ = "STUB: not implemented"; return }

func (g *gauge) WithTags(tags monitor.Tags) monitor.TagsAwareGauge {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareGauge)
}

func (g *gauge) WithContext(ctx context.Context) monitor.TagsAwareGauge {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareGauge)
}

// *******************************************************************
// *                             histogram                           *
// *******************************************************************
type histogram struct {
	*tagsMetric
	bricksHistogram monitor.BricksHistogram
	extractors      []monitor.ContextExtractor
}

// Record value
func (h *histogram) Record(v float64) { _ = "STUB: not implemented"; return }

func (h *histogram) WithTags(tags monitor.Tags) monitor.TagsAwareHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareHistogram)
}

func (h *histogram) WithContext(ctx context.Context) monitor.TagsAwareHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareHistogram)
}

// *******************************************************************
// *                             timer                               *
// *******************************************************************
type timer struct {
	*tagsMetric
	bricksTimer monitor.BricksTimer
	extractors  []monitor.ContextExtractor
}

// Record uses Histogram to record timed duration
// Since Histogram accepts float64 we will take the d.Seconds() which returns float64
func (t *timer) Record(d time.Duration) { _ = "STUB: not implemented"; return }

func (t *timer) WithTags(tags monitor.Tags) monitor.TagsAwareTimer {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareTimer)
}

func (t *timer) WithContext(ctx context.Context) monitor.TagsAwareTimer {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareTimer)
}

// *******************************************************************
// *                             tags helper                         *
// *******************************************************************
type tagsMetric struct {
	sync.Mutex
	tags    monitor.Tags
	onError func(error)
	copied  bool
}

func (tm *tagsMetric) withTags(tags monitor.Tags) { _ = "STUB: not implemented"; return }

func (tm *tagsMetric) withContext(ctx context.Context, extractors []monitor.ContextExtractor) {
	_ = "STUB: not implemented"
	return
}

func (tm *tagsMetric) shouldLogMetric(err error) bool { _ = "STUB: not implemented"; return false }

// Metric Constructors

func newCounterWithTags(bricksCounter monitor.BricksCounter, predefinedTags monitor.Tags, extractors []monitor.ContextExtractor, onError func(error)) monitor.TagsAwareCounter {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareCounter)
}

func newGaugeWithTags(bricksGauge monitor.BricksGauge, predefinedTags monitor.Tags, extractors []monitor.ContextExtractor, onError func(error)) monitor.TagsAwareGauge {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareGauge)
}

func newHistogramWithTags(bricksHistogram monitor.BricksHistogram, predefinedTags monitor.Tags, extractors []monitor.ContextExtractor, onError func(error)) monitor.TagsAwareHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareHistogram)
}

func newTimerWithTags(bricksTimer monitor.BricksTimer, predefinedTags monitor.Tags, extractors []monitor.ContextExtractor, onError func(error)) monitor.TagsAwareTimer {
	_ = "STUB: not implemented"
	return *new(monitor.TagsAwareTimer)
}
