package service

import (
	"context"
	helloword "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
)

func (s *HellowordService) CreateGreeter(ctx context.Context, req *helloword.CreateGreeterRequest) (rp *helloword.CreateGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateGreeter")
	defer span.End()
	s.log.WithContext(ctx).Infof("input data %v", req)
	rp = &helloword.CreateGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	err = s.greeter.Create(ctx, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, r)
	return
}

func (s *HellowordService) UpdateGreeter(ctx context.Context, req *helloword.UpdateGreeterRequest) (rp *helloword.UpdateGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateGreeter")
	defer span.End()
	s.log.WithContext(ctx).Infof("input data %v", req)
	rp = &helloword.UpdateGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	err = s.greeter.Update(ctx, req.Id, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, r)
	return
}

func (s *HellowordService) DeleteGreeter(ctx context.Context, req *helloword.DeleteGreeterRequest) (rp *helloword.DeleteGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteGreeter")
	defer span.End()
	s.log.WithContext(ctx).Infof("input data %v", req)
	rp = &helloword.DeleteGreeterReply{}
	err = s.greeter.Delete(ctx, req.Id)
	return
}

func (s *HellowordService) GetGreeter(ctx context.Context, req *helloword.GetGreeterRequest) (rp *helloword.GetGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGreeter")
	defer span.End()
	s.log.WithContext(ctx).Infof("input data %v", req)
	rp = &helloword.GetGreeterReply{}
	res, err := s.greeter.Get(ctx, req.Id)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, res)
	return
}

func (s *HellowordService) ListGreeter(ctx context.Context, req *helloword.ListGreeterRequest) (rp *helloword.ListGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "ListGreeter")
	defer span.End()
	s.log.WithContext(ctx).Infof("input data %v", req)
	rp = &helloword.ListGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	list, err := s.greeter.List(ctx, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.List, list)
	return
}
