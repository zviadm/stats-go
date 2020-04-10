package metrics

import (
	"sync"

	"go.uber.org/atomic"
)

type metric struct {
	opts metricOptions
	m    sync.Map
	mLen atomic.Int64 // len(m), but can be temporarily out of sync, since there are no locks.
}

func (m *metric) V(tagsAll ...map[string]string) *atomic.Float64 {
	values := make([]string, 0, len(m.opts.Tags))
	for _, tag := range m.opts.Tags {
		var ok bool
		var v string
		for idx := len(tagsAll) - 1; idx >= 0; idx-- {
			v, ok = tagsAll[idx][tag]
			if ok {
				break
			}
		}
		values = append(values, v)
	}
	// TODO(zviad): Check for missing tags?
	key := encodeValueList(values)
	v, ok := m.m.Load(key)
	if !ok {
		v = atomic.NewFloat64(0)
		v, ok = m.m.LoadOrStore(key, v)
		if !ok {
			m.mLen.Add(1)
		}
	}
	return v.(*atomic.Float64)
}

func (m *metric) Export() MetricData {
	r := MetricData{
		Type: m.opts.Type,
		Tags: m.opts.Tags,
	}
	r.F64s = make(map[string]float64, int(m.mLen.Load()))
	m.m.Range(func(k, v interface{}) bool {
		kk := k.(string)
		vv := v.(*atomic.Float64).Load()
		r.F64s[kk] = vv
		return true
	})
	return r
}
