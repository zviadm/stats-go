package metrics

type metricOptions struct {
	Type MetricType
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
		opts.Tags = tags
	}
}
