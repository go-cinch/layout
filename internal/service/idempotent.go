package service

import (
	"context"
	"github.com/go-cinch/layout/api/game"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GameService) Idempotent(ctx context.Context, req *emptypb.Empty) (rp *game.IdempotentReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Idempotent")
	defer span.End()
	rp = &game.IdempotentReply{}
	rp.Token = s.idempotent.Token(ctx)
	return
}
