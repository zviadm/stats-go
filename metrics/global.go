// Package metrics provides APIs for instrumenting code.
package metrics

import "flag"

var flagInstanceName = flag.String("stats.instance_name", "", "")
var flagNodeTags = flag.String("stats.node_tags", "", "")

var registryGlobal *registry

// InstanceNameAndNodeTags returns instance name and node tags. Instance name and
// node tags can not change after first call to this function. Thus after first call
// they remain same until process exits.
func InstanceNameAndNodeTags() (name string, tags map[string]string) {
	return registryGlobal.InstanceNameAndNodeTags()
}

// SetInstanceName sets instance name. Instance name can also be set using a flag.
// If flag is set, flag takes the priority and overrides whatever application sets.
func SetInstanceName(name string) {
	registryGlobal.SetInstanceName(name)
}

// Export reads all current metric data. Expected to be used by exporters.
func Export() map[string]MetricData {
	return registryGlobal.Export()
}

func init() {
	registryGlobal = newRegistry()
}
