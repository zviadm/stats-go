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

// ExporterGo starts exporter go routine that exports stats to datadog agent. Exporter will run
// in the background until context is canceled.
func ExporterGo(ctx context.Context) error {
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
			for name, mData := range export {
				mDataPrev := exportPrev[name]
				if cachedTagMap[name] == nil {
					cachedTagMap[name] = make(
						map[metrics.ValueList][]string, len(mData.F64s))
				}
				nameDD := strings.ReplaceAll(name, "/", ".")
				switch mData.Type {
				case metrics.CounterType, metrics.GaugeType:
					for vList, v := range mData.F64s {
						tags := cacheDatadogTags(
							cachedTagMap[name], mData.Tags, vList)
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

			// Clear out not longer reporting entries from cachedTagMap.
			for name, cachedMap := range cachedTagMap {
				mData, ok := export[name]
				if !ok {
					delete(cachedTagMap, name)
					continue
				}
				for vList := range cachedMap {
					_, ok := mData.F64s[vList]
					if !ok {
						delete(cachedMap, vList)
						continue
					}
				}
			}
		}
	}()
	return nil
}

func cacheDatadogTags(
	cachedMap map[metrics.ValueList][]string,
	tagNames []string,
	vList metrics.ValueList) []string {
	tags, ok := cachedMap[vList]
	if !ok {
		tags = encodeDatadogTags(tagNames, vList.Decode())
		cachedMap[vList] = tags
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
