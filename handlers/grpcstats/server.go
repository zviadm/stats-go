package grpcstats

import (
	"context"
	"strconv"

	"github.com/zviadm/stats-go/metrics"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

type serverHandler struct{}

// NewServer creates new stats handler for GRPC server.
func NewServer() stats.Handler {
	return &serverHandler{}
}

type ctxRpcTagInfo struct{}

var rpcTagInfoKey = ctxRpcTagInfo{}

type rpcInfo struct {
	FullMethodName string
	inflightG      metrics.Gauge
}

func (h *serverHandler) TagRPC(
	ctx context.Context, info *stats.RPCTagInfo) context.Context {
	i := &rpcInfo{
		FullMethodName: info.FullMethodName,
		inflightG:      serverRequestsInflightGauge.V(metrics.KV{"method": info.FullMethodName}),
	}
	return context.WithValue(ctx, rpcTagInfoKey, i)
}

func (h *serverHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {
	if s.IsClient() {
		return
	}
	switch s := s.(type) {
	case *stats.Begin:
		info := ctx.Value(rpcTagInfoKey).(*rpcInfo)
		info.inflightG.Add(1)
	case *stats.End:
		info := ctx.Value(rpcTagInfoKey).(*rpcInfo)
		info.inflightG.Add(-1)
		serverRequestsCounter.V(metrics.KV{
			"method": info.FullMethodName,
			"code":   strconv.Itoa(int(status.Convert(s.Error).Code())),
		}).Count(1)
	}
}

func (h *serverHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *serverHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	if s.IsClient() {
		return
	}
	switch s := s.(type) {
	case *stats.ConnBegin:
		serverConnectionsG.Add(1)
		serverConnectsC.Count(1)
	case *stats.ConnEnd:
		serverConnectionsG.Add(-1)
		_ = s
	}
}
