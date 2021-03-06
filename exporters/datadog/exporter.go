// Package datadog provides exporter to export collected metrics to DataDog agent.
package datadog

import (
	"context"
	"flag"
	"strings"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/zviadm/stats-go/metrics"
	"github.com/zviadm/zlog"
)

var flagPushFrequency = flag.Duration(
	"stats.datadog.push_frequency", 10*time.Second, "Frequency at which to push stats to DogStatsD daemon.")
var flagAgentAddr = flag.String(
	"stats.datadog.addr", "127.0.0.1:8125", "DogStatsD daemon address.")

// ExporterGo starts exporter go routine that exports stats to datadog agent. Exporter will run
// in the background until context is canceled.
// To match naming convention in DataDog, "/"-s are replaced with "."-s in metric names.
func ExporterGo(ctx context.Context) error {
	instanceName, nodeTags := metrics.InstanceNameAndNodeTags()
	if instanceName == "" {
		zlog.Info("not exporting stats, instance name not configured")
		return nil
	}
	c, err := statsd.New(*flagAgentAddr)
	if err != nil {
		return err
	}
	c.Tags = []string{"instance:" + instanceName}
	_ = nodeTags // TODO(zviad): figure out best way to expose node stats in Datadog.

	var exportPrev map[string]metrics.MetricData
	cachedTagMap := make(map[string]map[metrics.ValueList][]string)
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
