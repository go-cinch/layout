package service

import (
	"context"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/page"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/greeter"
	"github.com/go-cinch/layout/internal/biz"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GreeterService) CreateGreeter(ctx context.Context, req *greeter.CreateGreeterRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateGreeter")
	defer span.End()
	rp = &emptypb.Empty{}
	r := &biz.Greeter{}
	copierx.Copy(&r, req)
	err = s.greeter.Create(ctx, r)
	return
}

func (s *GreeterService) GetGreeter(ctx context.Context, req *greeter.GetGreeterRequest) (rp *greeter.GetGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGreeter")
	defer span.End()
	rp = &greeter.GetGreeterReply{}
	res, err := s.greeter.Get(ctx, req.Id)
	if err != nil {
		return
	}
	copierx.Copy(&rp, res)
	return
}

func (s *GreeterService) FindGreeter(ctx context.Context, req *greeter.FindGreeterRequest) (rp *greeter.FindGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "FindGreeter")
	defer span.End()
	rp = &greeter.FindGreeterReply{}
	rp.Page = &greeter.Page{}
	r := &biz.FindGreeter{}
	r.Page = page.Page{}
	copierx.Copy(&r, req)
	copierx.Copy(&r.Page, req.Page)
	res := s.greeter.Find(ctx, r)
	copierx.Copy(&rp.Page, r.Page)
	copierx.Copy(&rp.List, res)
	return
}

func (s *GreeterService) UpdateGreeter(ctx context.Context, req *greeter.UpdateGreeterRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateGreeter")
	defer span.End()
	rp = &emptypb.Empty{}
	r := &biz.UpdateGreeter{}
	copierx.Copy(&r, req)
	err = s.greeter.Update(ctx, r)
	return
}

func (s *GreeterService) DeleteGreeter(ctx context.Context, req *greeter.IdsRequest) (rp *emptypb.Empty, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteGreeter")
	defer span.End()
	rp = &emptypb.Empty{}
	err = s.greeter.Delete(ctx, utils.Str2Uint64Arr(req.Ids)...)
	return
}
