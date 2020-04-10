package datadog

import (
	"context"
	"errors"
	"flag"
	"strings"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/zviadm/stats-go/metrics"
)

var flagPushFrequency = flag.Duration("stats.datadog.push_frequency", 10*time.Second, "")
var flagAgentAddr = flag.String("stats.datadog.addr", "127.0.0.1:8125", "")

// Export starts exporter to datadog agent. Exporter will run in the background until
// context is canceled.
func Export(ctx context.Context) error {
	instanceName, nodeTags := metrics.InstanceNameAndNodeTags()
	if instanceName == "" {
		return errors.New("instance name must be configured to export stats")
	}
	c, err := statsd.New(*flagAgentAddr)
	if err != nil {
		return err
	}
	_ = instanceName
	_ = nodeTags

	var exportPrev map[string]metrics.MetricData
	var cachedTagMap map[string]map[metrics.ValueList][]string
	go func() {
		ticker := time.NewTicker(*flagPushFrequency)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			export := metrics.Export()
			tagMap := make(map[string]map[metrics.ValueList][]string, len(cachedTagMap))
			for name, mData := range export {
				mDataPrev := exportPrev[name]
				nameDD := strings.ReplaceAll(name, "/", ".")
				switch mData.Type {
				case metrics.CounterType, metrics.GaugeType:
					for vList, v := range mData.F64s {
						tags := cacheDatadogTags(
							tagMap, name, mData.Tags, vList, cachedTagMap[name])
						switch mData.Type {
						case metrics.CounterType:
							vDelta := v - mDataPrev.F64s[vList]
							_ = c.Count(nameDD, int64(vDelta), tags, 1)
						case metrics.GaugeType:
							_ = c.Gauge(nameDD, v, tags, 1)
						}
					}
				default:
					// Other types not supported.
				}
			}
			exportPrev = export
			cachedTagMap = tagMap
		}
	}()
	return nil
}

func cacheDatadogTags(
	tagMap map[string]map[metrics.ValueList][]string,
	name string,
	tagNames []string,
	vList metrics.ValueList,
	cachedMap map[metrics.ValueList][]string) []string {
	tags, ok := tagMap[name][vList]
	if !ok {
		tags, ok = cachedMap[vList]
		if !ok {
			tags = encodeDatadogTags(tagNames, vList.Decode())
		}

		if tagMap[name] == nil {
			tagMap[name] = make(map[metrics.ValueList][]string, len(cachedMap))
		}
		tagMap[name][vList] = tags
	}
	return tags
}

func encodeDatadogTags(names []string, values []string) []string {
	tags := make([]string, len(names))
	for idx := range tags {
		if values[idx] == "" {
			continue // Skip empty tags.
		}
		tags[idx] = names[idx] + ":" + values[idx]
	}
	return tags
}
