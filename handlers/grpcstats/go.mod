module github.com/zviadm/stats-go/handlers/grpcstats

go 1.14

require (
	github.com/zviadm/stats-go/metrics v0.0.0
	google.golang.org/grpc v1.28.1
)

replace github.com/zviadm/stats-go/metrics => ../../metrics
