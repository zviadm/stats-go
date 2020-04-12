package metrics

import (
	"go.uber.org/atomic"
)

// GaugeMetric represents definition for a gauge metric. Gauges hold
// last observation in a time range for point-in-time measurements.
type GaugeMetric struct {
	m *metric
}

// V creates instance of Gauge for given Key->Value tags.
func (m GaugeMetric) V(tags ...KV) Gauge {
	return Gauge{v: m.m.V(tags...)}
}

// DefineGauge defines new gauge metric. By convention, DefineGauge calls are
// expected to happen at import time, and all metrics exported by a package to be defined
// in a single `stats.go` file.
func DefineGauge(name string, opts ...MetricOption) GaugeMetric {
	c, err := registryGlobal.DefineGauge(name, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// Gauge is a single instance of GaugeMetric created for specific tag values.
type Gauge struct {
	v *atomic.Float64
}

// Add changes gauge by delta of `v`. Value can be negative too.
func (c Gauge) Add(v float64) {
	c.v.Add(v)
}

// Set sets the value of a gauge to a specific value.
func (c Gauge) Set(v float64) {
	c.v.Store(v)
}

// Get returns current gauge value.
func (c Gauge) Get() float64 {
	return c.v.Load()
}
