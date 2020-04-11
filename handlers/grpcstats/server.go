package grpcstats

import (
	"context"

	"google.golang.org/grpc/stats"
)

type serverHandler struct{}

// NewServer creates new stats handler for GRPC server.
func NewServer() stats.Handler {
	return &serverHandler{}
}

func (h *serverHandler) TagRPC(
	ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (h *serverHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {

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
		connectionsServerG.Add(1)
		connectsServerG.Count(1)
	case *stats.ConnEnd:
		connectionsServerG.Add(-1)
		_ = s
	}
}
