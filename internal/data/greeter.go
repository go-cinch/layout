package data

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/pkg/errors"

	"github.com/go-cinch/layout/internal/biz"
)

type greeterRepo struct {
	data *Data
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}

func (ro greeterRepo) Create(ctx context.Context, item *biz.Greeter) (err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Create, name: %s, age: %d", item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Update(ctx context.Context, id int64, item *biz.Greeter) (err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Update, id: %d, name: %s, age: %d", item.Id, item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Delete(ctx context.Context, id int64) (err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Delete, id: %d", id)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Get(ctx context.Context, id int64) (item *biz.Greeter, err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Get, id: %d", id)
	err = biz.ErrGreeterNotFound
	return
}

func (ro greeterRepo) List(ctx context.Context, item *biz.Greeter) (list []*biz.Greeter, err error) {
	// TODO implement me
	list = make([]*biz.Greeter, 0)
	log.WithContext(ctx).Info("List, name: %s, age: %d", item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}
