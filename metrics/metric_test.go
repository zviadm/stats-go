package metrics

import (
	"testing"
)

var metricsVCounter = DefineCounter("benchmark/metrics_v", WithTags("tag1", "tag2", "tag3"))

// BenchmarkMetricsVWithTags-4 - 1643517 - 707 ns/op - 8 B/op - 1 allocs/op
// TODO(zviadm): ideally this should do `0` allocations.
func BenchmarkMetricsVWithTags(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		metricsVCounter.V(KV{"tag1": "t1", "tag2": "t2", "tag3": "t3"}).Count(1)
	}
}

// BenchmarkMetricsVCountOnly-4 - 78687260 - 14.1 ns/op - 0 B/op - 0 allocs/op
func BenchmarkMetricsVCountOnly(b *testing.B) {
	c := metricsVCounter.V(KV{"tag1": "t1", "tag2": "t2", "tag3": "t3"})
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Count(1)
	}
}
