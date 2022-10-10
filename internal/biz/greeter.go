package biz

import (
	"context"
	"fmt"
	"github.com/go-cinch/common/utils"
)

type Greeter struct {
	Id   uint64 `json:"id,string"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type GreeterRepo interface {
	Create(ctx context.Context, item *Greeter) error
	Update(ctx context.Context, id uint64, item *Greeter) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*Greeter, error)
	List(ctx context.Context, item *Greeter) ([]*Greeter, error)
}

type GreeterUseCase struct {
	repo GreeterRepo
	tx   Transaction
}

func NewGreeterUseCase(repo GreeterRepo, tx Transaction) *GreeterUseCase {
	return &GreeterUseCase{repo: repo, tx: tx}
}

func (uc *GreeterUseCase) Create(ctx context.Context, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.repo.Create(ctx, item)
	})
}

func (uc *GreeterUseCase) Update(ctx context.Context, id uint64, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.repo.Update(ctx, id, item)
	})
}

func (uc *GreeterUseCase) Delete(ctx context.Context, id uint64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.repo.Delete(ctx, id)
	})
}

func (uc *GreeterUseCase) Get(ctx context.Context, id uint64) (*Greeter, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *GreeterUseCase) List(ctx context.Context, item *Greeter) ([]*Greeter, error) {
	return uc.repo.List(ctx, item)
}

type GreeterWithCacheUseCase struct {
	repo  GreeterRepo
	tx    Transaction
	cache Cache
}

func NewGreeterWithCacheUseCase(repo GreeterRepo, tx Transaction, cache Cache) *GreeterWithCacheUseCase {
	cache.Register("greeter_data")
	return &GreeterWithCacheUseCase{repo: repo, tx: tx, cache: cache}
}

func (uc *GreeterWithCacheUseCase) Create(ctx context.Context, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Create(ctx, item)
		})
	})
}

func (uc *GreeterWithCacheUseCase) Update(ctx context.Context, id uint64, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Update(ctx, id, item)
		})
	})
}

func (uc *GreeterWithCacheUseCase) Delete(ctx context.Context, id uint64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Delete(ctx, id)
		})
	})
}

func (uc *GreeterWithCacheUseCase) Get(ctx context.Context, id uint64) (item *Greeter, err error) {
	item = &Greeter{}
	action := fmt.Sprintf("get_%d", id)
	str, ok, lock, _ := uc.cache.Get(ctx, action, func(ctx context.Context) (string, bool) {
		return uc.get(ctx, action, id)
	})
	if ok {
		utils.Json2Struct(item, str)
	} else if !lock {
		err = TooManyRequests
		return
	}
	return
}

func (uc *GreeterWithCacheUseCase) get(ctx context.Context, action string, id uint64) (res string, ok bool) {
	// read data from db and write to cache
	p := &Greeter{}
	var err error
	p, err = uc.repo.Get(ctx, id)
	if err != nil && err != ErrGreeterNotFound {
		return
	}
	res = utils.Struct2Json(p)
	uc.cache.Set(ctx, action, res, err == ErrGreeterNotFound)
	ok = true
	return
}

func (uc *GreeterWithCacheUseCase) List(ctx context.Context, item *Greeter) (list []*Greeter, err error) {
	list = make([]*Greeter, 0)
	action := fmt.Sprintf("list_%d_%s_%d", item.Id, item.Name, item.Age)
	str, ok, lock, _ := uc.cache.Get(ctx, action, func(ctx context.Context) (res string, ok bool) {
		return uc.list(ctx, action, item)
	})
	if ok {
		utils.Json2Struct(list, str)
	} else if !lock {
		err = TooManyRequests
		return
	}
	return
}

func (uc *GreeterWithCacheUseCase) list(ctx context.Context, action string, item *Greeter) (res string, ok bool) {
	// read data from db and write to cache
	list := make([]*Greeter, 0)
	var err error
	list, err = uc.repo.List(ctx, item)
	if err != nil {
		return
	}
	res = utils.Struct2Json(list)
	uc.cache.Set(ctx, action, res, len(list) == 0)
	ok = true
	return
}
