module github.com/zviadm/stats-go/exporters/datadog

go 1.14

require (
	github.com/DataDog/datadog-go v3.5.0+incompatible
	github.com/stretchr/testify v1.5.1
	github.com/zviadm/stats-go/metrics v0.0.0-20200412122026-72fc1da5a98f
	github.com/zviadm/zlog v0.0.0-20200326214804-bea93fc07ffa
)

replace github.com/zviadm/stats-go/metrics => ../../metrics