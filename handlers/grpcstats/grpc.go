package grpcstats

import (
	"context"

	"google.golang.org/grpc/stats"
)

type handler struct{}

// New creates new stats handler.
func New() stats.Handler {
	return &handler{}
}

func (h *handler) TagRPC(
	ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (h *handler) HandleRPC(ctx context.Context, s stats.RPCStats) {

}

func (h *handler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *handler) HandleConn(ctx context.Context, s stats.ConnStats) {
	if !s.IsClient() {
		switch s := s.(type) {
		case *stats.ConnBegin:
			connectionsServerG.Add(1)
			connectsServerG.Count(1)
		case *stats.ConnEnd:
			connectionsServerG.Add(-1)
			_ = s
		}
	} else {
		switch s := s.(type) {
		case *stats.ConnBegin:
			connectionsClientG.Add(1)
			connectsClientG.Count(1)
		case *stats.ConnEnd:
			connectionsClientG.Add(-1)
			_ = s
		}
	}
}
