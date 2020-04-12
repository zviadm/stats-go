// Package grpcstats provides gRPC stats.Handler implementations.
package grpcstats

import "github.com/zviadm/stats-go/metrics"

var (
	serverConnectionsG = metrics.DefineGauge(
		"grpc/server/connetions",
		metrics.WithDesc("Currently established GRPC connections."),
	).V()
	serverConnectsC = metrics.DefineCounter(
		"grpc/server/connects",
		metrics.WithDesc("Rate of GRPC connects."),
	).V()

	serverRequestsCounter = metrics.DefineCounter(
		"grpc/server/requests",
		metrics.WithDesc("Rate of GRPC requests."),
		metrics.WithTags("method", "code"),
	)
)
