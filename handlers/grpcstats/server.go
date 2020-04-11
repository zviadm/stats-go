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

func (h *serverHandler) TagRPC(
	ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return context.WithValue(ctx, rpcTagInfoKey, info)
}

func (h *serverHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {
	if s.IsClient() {
		return
	}
	switch s := s.(type) {
	case *stats.End:
		info := ctx.Value(rpcTagInfoKey).(*stats.RPCTagInfo)
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
