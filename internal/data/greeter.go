package data

import (
	"context"
	"github.com/pkg/errors"

	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ro greeterRepo) Create(ctx context.Context, item *biz.Greeter) (err error) {
	// TODO implement me
	ro.log.WithContext(ctx).Infof("Create, name: %s, age: %d", item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Update(ctx context.Context, id int64, item *biz.Greeter) (err error) {
	// TODO implement me
	ro.log.WithContext(ctx).Infof("Update, id: %d, name: %s, age: %d", item.Id, item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Delete(ctx context.Context, id int64) (err error) {
	// TODO implement me
	ro.log.WithContext(ctx).Infof("Delete, id: %d", id)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Get(ctx context.Context, id int64) (item *biz.Greeter, err error) {
	// TODO implement me
	ro.log.WithContext(ctx).Infof("Get, id: %d", id)
	err = biz.ErrGreeterNotFound
	return
}

func (ro greeterRepo) List(ctx context.Context, item *biz.Greeter) (list []*biz.Greeter, err error) {
	// TODO implement me
	list = make([]*biz.Greeter, 0)
	ro.log.WithContext(ctx).Infof("List, name: %s, age: %d", item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}
