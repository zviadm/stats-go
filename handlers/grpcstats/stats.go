package grpcstats

import "github.com/zviadm/stats-go/metrics"

var (
	serverConnectionsG = metrics.DefineGauge("grpc/server/connetions").V()
	serverConnectsC    = metrics.DefineCounter("grpc/server/connects").V()

	serverRequestsCounter = metrics.DefineCounter(
		"grpc/server/requests",
		metrics.WithTags("method", "code"),
	)

	// connectionsClientG = metrics.DefineGauge("grpc/client/connetions").V()
	// connectsClientG    = metrics.DefineCounter("grpc/client/connects").V()
)
