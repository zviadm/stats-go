package metrics

import "sort"

type metricOptions struct {
	Type MetricType
	Desc string
	Tags []string
}

// MetricType enumerates type of a metric.
type MetricType string

// Supported metric types.
const (
	CounterType = "c"
	GaugeType   = "g"
)

// MetricOption defines options for metric definitions.
type MetricOption func(*metricOptions)

// WithTags adds list of tags to metric definition.
func WithTags(tags ...string) MetricOption {
	return func(opts *metricOptions) {
		tagsC := make(sort.StringSlice, len(tags))
		copy(tagsC, tags)
		tagsC.Sort()
		opts.Tags = tagsC
	}
}

// WithDesc adds description to metrics definition.
func WithDesc(desc string) MetricOption {
	return func(opts *metricOptions) {
		opts.Desc = desc
	}
}
