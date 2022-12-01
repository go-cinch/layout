package biz

import (
	"context"
	"fmt"
	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/common/page"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/reason"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/pkg/errors"
)

type Greeter struct {
	Id   uint64 `json:"id,string"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type FindGreeter struct {
	Page page.Page `json:"page"`
	Name *string   `json:"name"`
	Age  *int32    `json:"age"`
}

type FindGreeterCache struct {
	Page page.Page `json:"page"`
	List []Greeter `json:"list"`
}

type UpdateGreeter struct {
	Id   *uint64 `json:"id,string,omitempty"`
	Name *string `json:"name,omitempty"`
	Age  *int32  `json:"age,omitempty"`
}

type GreeterRepo interface {
	Create(ctx context.Context, item *Greeter) error
	Get(ctx context.Context, id uint64) (*Greeter, error)
	Find(ctx context.Context, condition *FindGreeter) []Greeter
	Update(ctx context.Context, item *UpdateGreeter) error
	Delete(ctx context.Context, ids ...uint64) error
}

type GreeterUseCase struct {
	c     *conf.Bootstrap
	repo  GreeterRepo
	tx    Transaction
	cache Cache
}

func NewGreeterUseCase(c *conf.Bootstrap, repo GreeterRepo, tx Transaction, cache Cache) *GreeterUseCase {
	// prefix rule = project name + _ + business name, example: layout_greeter
	return &GreeterUseCase{c: c, repo: repo, tx: tx, cache: cache.WithPrefix(fmt.Sprintf("%s_greeter", c.Name))}
}

func (uc *GreeterUseCase) Create(ctx context.Context, item *Greeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Create(ctx, item)
		})
	})
}

func (uc *GreeterUseCase) Get(ctx context.Context, id uint64) (rp *Greeter, err error) {
	rp = &Greeter{}
	action := fmt.Sprintf("get_%d", id)
	str, ok := uc.cache.Get(ctx, action, func(ctx context.Context) (string, bool) {
		return uc.get(ctx, action, id)
	})
	if ok {
		utils.Json2Struct(&rp, str)
		if rp.Id == constant.UI0 {
			err = reason.ErrorNotFound("%s Greeter.id: %d", i18n.FromContext(ctx).T(RecordNotFound), id)
		}
		return
	}
	err = reason.ErrorTooManyRequests(i18n.FromContext(ctx).T(TooManyRequests))
	return
}

func (uc *GreeterUseCase) get(ctx context.Context, action string, id uint64) (res string, ok bool) {
	// read data from db and write to cache
	rp := &Greeter{}
	item, err := uc.repo.Get(ctx, id)
	notFound := errors.Is(err, reason.ErrorNotFound(i18n.FromContext(ctx).T(RecordNotFound)))
	if err != nil && !notFound {
		return
	}
	copierx.Copy(&rp, item)
	res = utils.Struct2Json(rp)
	uc.cache.Set(ctx, action, res, notFound)
	ok = true
	return
}

func (uc *GreeterUseCase) Find(ctx context.Context, condition *FindGreeter) (rp []Greeter) {
	// use md5 string as cache replay json str, key is short
	action := fmt.Sprintf("find_%s", utils.StructMd5(condition))
	str, ok := uc.cache.Get(ctx, action, func(ctx context.Context) (string, bool) {
		return uc.find(ctx, action, condition)
	})
	if ok {
		var cache FindGreeterCache
		utils.Json2Struct(&cache, str)
		condition.Page = cache.Page
		rp = cache.List
	}
	return
}

func (uc *GreeterUseCase) find(ctx context.Context, action string, condition *FindGreeter) (res string, ok bool) {
	// read data from db and write to cache
	list := uc.repo.Find(ctx, condition)
	var cache FindGreeterCache
	cache.List = list
	cache.Page = condition.Page
	res = utils.Struct2Json(cache)
	uc.cache.Set(ctx, action, res, len(list) == 0)
	ok = true
	return
}

func (uc *GreeterUseCase) Update(ctx context.Context, item *UpdateGreeter) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Update(ctx, item)
			return
		})
	})
}

func (uc *GreeterUseCase) Delete(ctx context.Context, ids ...uint64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Delete(ctx, ids...)
			return
		})
	})
}
