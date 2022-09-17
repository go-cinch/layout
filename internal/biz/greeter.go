package biz

import (
	"context"
	"encoding/json"
	"fmt"
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

func (uc *GreeterUseCase) Update(ctx context.Context, id int64, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.repo.Update(ctx, id, item)
	})
}

func (uc *GreeterUseCase) Delete(ctx context.Context, id int64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.repo.Delete(ctx, id)
	})
}

func (uc *GreeterUseCase) Get(ctx context.Context, id int64) (p *Greeter, err error) {
	return uc.repo.Get(ctx, id)
}

func (uc *GreeterUseCase) List(ctx context.Context, item *Greeter) (ps []*Greeter, err error) {
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

func (uc *GreeterWithCacheUseCase) Update(ctx context.Context, id int64, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Update(ctx, id, item)
		})
	})
}

func (uc *GreeterWithCacheUseCase) Delete(ctx context.Context, id int64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Delete(ctx, id)
		})
	})
}

func (uc *GreeterWithCacheUseCase) Get(ctx context.Context, id int64) (p *Greeter, err error) {
	p = &Greeter{}
	action := fmt.Sprintf("get_%d", id)
	// first get cache
	str, ok := uc.cache.Get(ctx, action)
	if ok {
		json.Unmarshal([]byte(str), p)
		return
	}
	// get lock before read db
	ok = uc.cache.Lock(ctx, action)
	if !ok {
		err = TooManyRequests
		return
	}
	defer uc.cache.Unlock(ctx, action)
	// double check cache(avoid concurrency first get cache=false)
	str, ok = uc.cache.Get(ctx, action)
	if ok {
		json.Unmarshal([]byte(str), p)
		return
	}
	p, err = uc.repo.Get(ctx, id)
	if err != nil {
		if err == ErrGreeterNotFound {
			// if record not found, set a short expiration cache
			res, _ := json.Marshal(p)
			uc.cache.SetWithExpiration(ctx, action, string(res), 60)
		}
		return
	}
	res, _ := json.Marshal(p)
	uc.cache.Set(ctx, action, string(res))
	return
}

func (uc *GreeterWithCacheUseCase) List(ctx context.Context, item *Greeter) (list []*Greeter, err error) {
	list = make([]*Greeter, 0)
	action := fmt.Sprintf("list_%d_%s_%d", item.Id, item.Name, item.Age)
	// first get cache
	str, ok := uc.cache.Get(ctx, action)
	if ok {
		json.Unmarshal([]byte(str), list)
		return
	}
	// get lock before read db
	ok = uc.cache.Lock(ctx, action)
	if !ok {
		err = TooManyRequests
		return
	}
	defer uc.cache.Unlock(ctx, action)
	// double check cache(avoid concurrency first get cache=false)
	str, ok = uc.cache.Get(ctx, action)
	if ok {
		json.Unmarshal([]byte(str), list)
		return
	}
	list, err = uc.repo.List(ctx, item)
	if err != nil {
		if len(list) == 0 {
			// if list is empty, set a short expiration cache
			res, _ := json.Marshal(list)
			uc.cache.SetWithExpiration(ctx, action, string(res), 60)
		}
		return
	}
	res, _ := json.Marshal(list)
	uc.cache.Set(ctx, action, string(res))
	return
}
