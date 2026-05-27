package monitoring

import (
	"time"

	"github.com/go-masonry/mortar/interfaces/monitor"
)

type noop struct {
	name, desc string
	err        error
	onError    func(error)
}

type noopCounter struct {
	*noop
}

func (n *noopCounter) WithTags(tags map[string]string) (monitor.Counter, error) {
	_ = "STUB: not implemented"
	return *new(monitor.Counter), nil
}

func newNoopCounter(name, desc string, err error, onError func(error)) monitor.BricksCounter {
	_ = "STUB: not implemented"
	return *new(monitor.BricksCounter)
}

type noopGauge struct {
	*noop
}

func (n *noopGauge) WithTags(tags map[string]string) (monitor.Gauge, error) {
	_ = "STUB: not implemented"
	return *new(monitor.Gauge), nil
}

func newNoopGauge(name, desc string, err error, onError func(error)) monitor.BricksGauge {
	_ = "STUB: not implemented"
	return *new(monitor.BricksGauge)
}

type noopHistogram struct {
	*noop
}

func (n *noopHistogram) WithTags(tags map[string]string) (monitor.Histogram, error) {
	_ = "STUB: not implemented"
	return *new(monitor.Histogram), nil
}

func newNoopHistogram(name, desc string, err error, onError func(error)) monitor.BricksHistogram {
	_ = "STUB: not implemented"
	return *new(monitor.BricksHistogram)
}

type noopTimer struct {
	*noop
}

func (n *noopTimer) WithTags(tags map[string]string) (monitor.Timer, error) {
	_ = "STUB: not implemented"
	return *new(monitor.Timer), nil
}

func (n *noopTimer) Record(d time.Duration) { _ = "STUB: not implemented"; return }

func newNoopTimer(name, desc string, err error, onError func(error)) monitor.BricksTimer {
	_ = "STUB: not implemented"
	return *new(monitor.BricksTimer)
}

// Inc increments the counter by 1
func (n *noop) Inc() {
	_ = "STUB: not implemented"

	// Add adds the given value to the counter, negative values are not advised
	return
}

func (n *noop) Add(v float64) {
	_ = "STUB: not implemented"

	// Record value
	return
}

func (n *noop) Record(v float64) {
	_ = "STUB: not implemented"

	// Set sets Gauge value
	return
}

func (n *noop) Set(v float64) {
	_ = "STUB: not implemented"

	// Dec adds -1
	return
}

func (n *noop) Dec() { _ = "STUB: not implemented"; return }

func (n *noop) do() { _ = "STUB: not implemented"; return }
