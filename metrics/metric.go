package metrics

import (
	"strings"
	"sync"

	"go.uber.org/atomic"
)

type metric struct {
	opts metricOptions
	m    sync.Map
	mLen atomic.Int64 // len(m), but can be temporarily out of sync, since there are no locks.
}

// KV represents tag key->value mapping.
type KV map[string]string

func (m *metric) V(tags KV) *atomic.Float64 {
	maxKeyLen := len(m.opts.Tags)
	for _, v := range tags {
		maxKeyLen += len(v)
	}
	b := strings.Builder{}
	b.Grow(maxKeyLen)
	for idx, tag := range m.opts.Tags {
		v := tags[tag]
		v = sanitizeTagValue(v)
		b.WriteString(v)
		if idx < len(m.opts.Tags)-1 {
			b.WriteByte(valueListSep)
		}
	}
	key := ValueList(b.String())
	// TODO(zviad): Check for missing tags?

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
	r.F64s = make(map[ValueList]float64, int(m.mLen.Load()))
	m.m.Range(func(k, v interface{}) bool {
		kk := k.(ValueList)
		vv := v.(*atomic.Float64).Load()
		r.F64s[kk] = vv
		return true
	})
	return r
}
