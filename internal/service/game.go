package service

import (
	"context"

	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/page"
	"github.com/go-cinch/common/proto/params"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/game"
	"github.com/go-cinch/layout/internal/biz"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GameService) CreateGame(ctx context.Context, req *game.CreateGameRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateGame")
	defer span.End()
	rp = &emptypb.Empty{}
	r := &biz.Game{}
	copierx.Copy(&r, req)
	err = s.game.Create(ctx, r)
	return
}

func (s *GameService) GetGame(ctx context.Context, req *game.GetGameRequest) (rp *game.GetGameReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGame")
	defer span.End()
	rp = &game.GetGameReply{}
	res, err := s.game.Get(ctx, req.Id)
	if err != nil {
		return
	}
	copierx.Copy(&rp, res)
	return
}

func (s *GameService) FindGame(ctx context.Context, req *game.FindGameRequest) (rp *game.FindGameReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "FindGame")
	defer span.End()
	rp = &game.FindGameReply{}
	rp.Page = &params.Page{}
	r := &biz.FindGame{}
	r.Page = page.Page{}
	copierx.Copy(&r, req)
	copierx.Copy(&r.Page, req.Page)
	res, err := s.game.Find(ctx, r)
	if err != nil {
		return
	}
	copierx.Copy(&rp.Page, r.Page)
	copierx.Copy(&rp.List, res)
	return
}

func (s *GameService) UpdateGame(ctx context.Context, req *game.UpdateGameRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateGame")
	defer span.End()
	rp = &emptypb.Empty{}
	r := &biz.UpdateGame{}
	copierx.Copy(&r, req)
	err = s.game.Update(ctx, r)
	return
}

func (s *GameService) DeleteGame(ctx context.Context, req *params.IdsRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteGame")
	defer span.End()
	rp = &emptypb.Empty{}
	err = s.game.Delete(ctx, utils.Str2Uint64Arr(req.Ids)...)
	return
}
