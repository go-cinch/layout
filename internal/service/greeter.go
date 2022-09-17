package service

import (
	"context"
	"github.com/go-cinch/common/log"
	helloword "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
)

func (s *HellowordService) CreateGreeter(ctx context.Context, req *helloword.CreateGreeterRequest) (rp *helloword.CreateGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateGreeter")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
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
	log.WithContext(ctx).Info("input data %v", req)
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
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.DeleteGreeterReply{}
	err = s.greeter.Delete(ctx, req.Id)
	return
}

func (s *HellowordService) GetGreeter(ctx context.Context, req *helloword.GetGreeterRequest) (rp *helloword.GetGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGreeter")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.GetGreeterReply{}
	rp.Item = &helloword.Greeter{}
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
	log.WithContext(ctx).Info("input data %v", req)
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

func (s *HellowordService) CreateGreeterWithCache(ctx context.Context, req *helloword.CreateGreeterRequest) (rp *helloword.CreateGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateGreeterWithCache")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.CreateGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	err = s.greeterWithCache.Create(ctx, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, r)
	return
}

func (s *HellowordService) UpdateGreeterWithCache(ctx context.Context, req *helloword.UpdateGreeterRequest) (rp *helloword.UpdateGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateGreeterWithCache")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.UpdateGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	err = s.greeterWithCache.Update(ctx, req.Id, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, r)
	return
}

func (s *HellowordService) DeleteGreeterWithCache(ctx context.Context, req *helloword.DeleteGreeterRequest) (rp *helloword.DeleteGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteGreeterWithCache")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.DeleteGreeterReply{}
	err = s.greeterWithCache.Delete(ctx, req.Id)
	return
}

func (s *HellowordService) GetGreeterWithCache(ctx context.Context, req *helloword.GetGreeterRequest) (rp *helloword.GetGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetGreeterWithCache")
	defer span.End()
	// log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.GetGreeterReply{}
	rp.Item = &helloword.Greeter{}
	res, err := s.greeterWithCache.Get(ctx, req.Id)
	if err != nil {
		return
	}
	copier.Copy(&rp.Item, res)
	return
}

func (s *HellowordService) ListGreeterWithCache(ctx context.Context, req *helloword.ListGreeterRequest) (rp *helloword.ListGreeterReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "ListGreeterWithCache")
	defer span.End()
	log.WithContext(ctx).Info("input data %v", req)
	rp = &helloword.ListGreeterReply{}
	r := &biz.Greeter{}
	copier.Copy(&r, req)
	list, err := s.greeterWithCache.List(ctx, r)
	if err != nil {
		return
	}
	copier.Copy(&rp.List, list)
	return
}
