package metrics

import (
	"go.uber.org/atomic"
)

// CounterMetric represents definition for a counter metric. Counters measure
// total count of events/operations/throughput/etc... in a time range. Counters can
// measure total negative count too.
type CounterMetric struct {
	m *metric
}

// V creates instance of Counter for given Key->Value tags.
func (m CounterMetric) V(tags ...KV) Counter {
	return Counter{v: m.m.V(tags...)}
}

// DefineCounter defines new counter metric.
func DefineCounter(name string, opts ...MetricOption) CounterMetric {
	c, err := registryGlobal.DefineCounter(name, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// Counter is a single instance of CounterMetric created for specific tag values.
type Counter struct {
	v *atomic.Float64
}

// Count adds given value to total count in current time range. Value can be negative too.
func (c Counter) Count(v float64) {
	c.v.Add(v)
}
