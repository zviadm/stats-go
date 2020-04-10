package metrics

import (
	"flag"
	"fmt"
	"sync"
)

var flagInstanceName = flag.String("stats.instance_name", "", "")
var flagNodeTags = flag.String("stats.node_tags", "", "")

type registry struct {
	mx           sync.Mutex
	initialized  bool
	instanceName string
	nodeTags     map[string]string
	metrics      map[string]*metric
}

var registryGlobal *registry

func newRegistry() *registry {
	return &registry{metrics: make(map[string]*metric)}
}

func (r *registry) DefineCounter(name string, opts ...MetricOption) (CounterMetric, error) {
	m, err := r.defineMetric(name, counterType, opts...)
	return CounterMetric{m}, err
}
func (r *registry) DefineGauge(name string, opts ...MetricOption) (GaugeMetric, error) {
	m, err := r.defineMetric(name, gaugeType, opts...)
	return GaugeMetric{m}, err
}

func (r *registry) defineMetric(name string, t metricType, opts ...MetricOption) (*metric, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	if _, ok := r.metrics[name]; ok {
		return nil, fmt.Errorf("metric: %s already defined", name)
	}
	o := metricOptions{Type: t}
	for _, opt := range opts {
		opt(&o)
	}
	m := &metric{opts: o}
	r.metrics[name] = m
	return m, nil
}

func (r *registry) InstanceNameAndNodeTags() (name string, tags map[string]string) {
	r.mx.Lock()
	defer r.mx.Unlock()
	if !r.initialized {
		if r.instanceName == "" {
			r.instanceName = *flagInstanceName
			r.nodeTags = parseNodeTags(*flagNodeTags)
		}
		r.initialized = true
	}
	return r.instanceName, r.nodeTags
}

func (r *registry) SetInstanceNameAndNodeTags(name string, tags map[string]string) {
	r.mx.Lock()
	defer r.mx.Unlock()
	if r.initialized {
		panic("registry is already initialized!")
	}
	if name == "" {
		panic("name must be set!")
	}
	r.instanceName = name
	r.nodeTags = tags
}

func parseNodeTags(nodeTags string) map[string]string {
	r := make(map[string]string)
	return r
}

func init() {
	registryGlobal = newRegistry()
}
