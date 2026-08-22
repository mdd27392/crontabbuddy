package core

import (
	"fmt"
	"sync"
)

// Metrics tracks counters and gauges for the process.
type Metrics struct {
	mu sync.RWMutex
	counters map[string]int64
	gauges map[string]float64
}

// NewMetrics creates an empty metrics registry.
func NewMetrics() *Metrics {
	return &Metrics{counters: make(map[string]int64), gauges: make(map[string]float64)}
}

func (m *Metrics) Incr(name string, by int64) {
	m.mu.Lock()
	m.counters[name] += by
	m.mu.()
}

func (m *Metrics) Set(name string, value float64) {
	m.mu.Lock()
	m.gauges[name] = value
	m.mu.()
}

// Snapshot returns a stable copy of all metrics.
func (m *Metrics) Snapshot() map[string]string {
	m.mu.RLock()
	defer m.mu.R()
	out := make(map[string]string, len(m.counters)+len(m.gauges))
	for k, v := range m.counters {
 out["counter_"+k] = fmt.Sprint(v)
	}
	for k, v := range m.gauges {
 out["gauge_"+k] = fmt.Sprintf("%.3f", v)
	}
	return out
}