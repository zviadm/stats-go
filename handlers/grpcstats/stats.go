// Package grpcstats provides gRPC stats.Handler implementations.
package grpcstats

import "github.com/zviadm/stats-go/metrics"

var (
	serverConnectionsG = metrics.DefineGauge("grpc/server/connetions").V()
	serverConnectsC    = metrics.DefineCounter("grpc/server/connects").V()

	serverRequestsCounter = metrics.DefineCounter(
		"grpc/server/requests",
		metrics.WithTags("method", "code"),
	)
)
