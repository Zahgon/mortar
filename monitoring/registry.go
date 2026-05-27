package monitoring

import (
	"sync"

	"github.com/go-masonry/mortar/interfaces/monitor"
)

type externalRegistry struct {
	cm       sync.RWMutex
	gm       sync.RWMutex
	tm       sync.RWMutex
	hm       sync.RWMutex
	external monitor.BricksMetrics
	// TODO perhaps change this to self evicting cache that will remove metrics if unused for a long time to save space
	counters   *sync.Map
	gauges     *sync.Map
	histograms *sync.Map
	timers     *sync.Map
}

func newRegistry(externalMetrics monitor.BricksMetrics) *externalRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (r *externalRegistry) loadOrStoreCounter(name, desc string, keys ...string) (bricksCounter monitor.BricksCounter, err error) {
	_ = "STUB: not implemented"
	return *new(monitor.BricksCounter), nil
}

// if a previous duplicate is already there (was created by other go routine)

// perhaps it's already there (was created by other go routine) and the underlying impl have a dup check

func (r *externalRegistry) loadOrStoreGauge(name, desc string, keys ...string) (bricksGauge monitor.BricksGauge, err error) {
	_ = "STUB: not implemented"
	return *new(monitor.BricksGauge), nil
}

// if a previous duplicate is already there (was created by other go routine)

// perhaps it's already there (was created by other go routine) and the underlying impl have a dup check

func (r *externalRegistry) loadOrStoreHistogram(name, desc string, buckets monitor.Buckets, keys ...string) (bricksHistogram monitor.BricksHistogram, err error) {
	_ = "STUB: not implemented"
	return *new(monitor.BricksHistogram), nil
}

// if a previous duplicate is already there (was created by other go routine)

// perhaps it's already there (was created by other go routine) and the underlying impl have a dup check

func (r *externalRegistry) loadOrStoreTimer(name, desc string, keys ...string) (bricksTimer monitor.BricksTimer, err error) {
	_ = "STUB: not implemented"
	return *new(monitor.BricksTimer), nil
}

// if a previous duplicate is already there (was created by other go routine)

// perhaps it's already there (was created by other go routine) and the underlying impl have a dup check

func calcID(name string, keys ...string) (ID string) { _ = "STUB: not implemented"; return "" }

// preallocate slice with extra space

// avoid allocation and prepend name
// add empty string to the end -> len++
// shift
// name is the first string now
