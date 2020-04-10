package metrics

import "strings"

type MetricData struct {
	Type MetricType
	Tags []string
	F64s map[ValueList]float64
}

type ValueList string

const (
	valueListSep = "|"
)

func (v ValueList) Decode() []string {
	return strings.Split(string(v), valueListSep)
}

func encodeValues(v ...string) ValueList {
	return ValueList(strings.Join(v, valueListSep))
}

func (r *registry) Export() map[string]MetricData {
	r.mx.Lock()
	defer r.mx.Unlock()
	export := make(map[string]MetricData, len(r.metrics))
	for name, m := range r.metrics {
		export[name] = m.Export()
	}
	return export
}
