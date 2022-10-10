package data

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/pkg/errors"
)

type greeterRepo struct {
	data *Data
}

// Greeter is database fields map
type Greeter struct {
	Id   uint64 `json:"id,string"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
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
	// u can do create like this:
	// var m Greeter
	// copier.Copy(&m, item)
	// m.Id = ro.data.Id(ctx)
	// err = ro.data.DB(ctx).Create(&m).Error
	return
}

func (ro greeterRepo) Update(ctx context.Context, id uint64, item *biz.Greeter) (err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Update, id: %d, name: %s, age: %d", item.Id, item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Delete(ctx context.Context, id uint64) (err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Delete, id: %d", id)
	err = errors.Errorf("implement me")
	return
}

func (ro greeterRepo) Get(ctx context.Context, id uint64) (item *biz.Greeter, err error) {
	// TODO implement me
	log.WithContext(ctx).Info("Get, id: %d", id)
	item = &biz.Greeter{
		Id: id,
	}
	if id <= 0 {
		item.Id = 0
		err = biz.ErrGreeterNotFound
	}
	return
}

func (ro greeterRepo) List(ctx context.Context, item *biz.Greeter) (list []*biz.Greeter, err error) {
	// TODO implement me
	list = make([]*biz.Greeter, 0)
	log.WithContext(ctx).Info("List, name: %s, age: %d", item.Name, item.Age)
	err = errors.Errorf("implement me")
	return
}
