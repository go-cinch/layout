package biz

import (
	"context"
	v1 "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrGreeterNotFound is greeter not found.
	ErrGreeterNotFound = errors.NotFound(v1.ErrorReason_GREETER_NOT_FOUND.String(), "greeter not found")
)

type Greeter struct {
	Id   int64
	Name string
	Age  int32
}

type GreeterRepo interface {
	Create(ctx context.Context, item *Greeter) error
	Update(ctx context.Context, id int64, item *Greeter) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*Greeter, error)
	List(ctx context.Context, item *Greeter) ([]*Greeter, error)
}

type GreeterUseCase struct {
	repo GreeterRepo
}

func NewGreeterUseCase(repo GreeterRepo) *GreeterUseCase {
	return &GreeterUseCase{repo: repo}
}

func (uc *GreeterUseCase) Create(ctx context.Context, item *Greeter) error {
	return uc.repo.Create(ctx, item)
}

func (uc *GreeterUseCase) Update(ctx context.Context, id int64, item *Greeter) error {
	return uc.repo.Update(ctx, id, item)
}

func (uc *GreeterUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *GreeterUseCase) Get(ctx context.Context, id int64) (p *Greeter, err error) {
	return uc.repo.Get(ctx, id)
}

func (uc *GreeterUseCase) List(ctx context.Context, item *Greeter) (ps []*Greeter, err error) {
	return uc.repo.List(ctx, item)
}
