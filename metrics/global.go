package metrics

import "flag"

var flagInstanceName = flag.String("stats.instance_name", "", "")
var flagNodeTags = flag.String("stats.node_tags", "", "")

var registryGlobal *registry

func InstanceNameAndNodeTags() (name string, tags map[string]string) {
	return registryGlobal.InstanceNameAndNodeTags()
}

func SetInstanceNameAndNodeTags(name string, tags map[string]string) {
	registryGlobal.SetInstanceNameAndNodeTags(name, tags)
}

func Export() map[string]MetricData {
	return registryGlobal.Export()
}

func init() {
	registryGlobal = newRegistry()
}
