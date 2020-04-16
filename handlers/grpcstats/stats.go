// Package grpcstats provides gRPC stats.Handler implementations.
package grpcstats

import "github.com/zviadm/stats-go/metrics"

var (
	serverConnectionsG = metrics.DefineGauge(
		"grpc/server/connetions",
		metrics.WithDesc("Currently established GRPC connections."),
	).V(nil)
	serverConnectsC = metrics.DefineCounter(
		"grpc/server/connects",
		metrics.WithDesc("Rate of GRPC connects."),
	).V(nil)

	serverRequestsCounter = metrics.DefineCounter(
		"grpc/server/requests",
		metrics.WithDesc("Rate of GRPC requests."),
		metrics.WithTags("method", "code"),
	)
	serverRequestsInflightGauge = metrics.DefineGauge(
		"grpc/server/requests_inflight",
		metrics.WithDesc("GRPC requests currently inflight."),
		metrics.WithTags("method"),
	)
)
