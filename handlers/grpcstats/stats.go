package grpcstats

import "github.com/zviadm/stats-go/metrics"

var (
	connectionsServerG = metrics.DefineGauge("grpc/server/connetions").V()
	connectsServerG    = metrics.DefineCounter("grpc/server/connects").V()

	// connectionsClientG = metrics.DefineGauge("grpc/client/connetions").V()
	// connectsClientG    = metrics.DefineCounter("grpc/client/connects").V()
)
