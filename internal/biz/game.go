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

type Game struct {
	Id   uint64 `json:"id,string"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type FindGame struct {
	Page page.Page `json:"page"`
	Name *string   `json:"name"`
	Age  *int32    `json:"age"`
}

type FindGameCache struct {
	Page page.Page `json:"page"`
	List []Game    `json:"list"`
}

type UpdateGame struct {
	Id   *uint64 `json:"id,string,omitempty"`
	Name *string `json:"name,omitempty"`
	Age  *int32  `json:"age,omitempty"`
}

type GameRepo interface {
	Create(ctx context.Context, item *Game) error
	Get(ctx context.Context, id uint64) (*Game, error)
	Find(ctx context.Context, condition *FindGame) []Game
	Update(ctx context.Context, item *UpdateGame) error
	Delete(ctx context.Context, ids ...uint64) error
}

type GameUseCase struct {
	c     *conf.Bootstrap
	repo  GameRepo
	tx    Transaction
	cache Cache
}

func NewGameUseCase(c *conf.Bootstrap, repo GameRepo, tx Transaction, cache Cache) *GameUseCase {
	// prefix rule = project name + _ + business name, example: layout_game
	return &GameUseCase{c: c, repo: repo, tx: tx, cache: cache.WithPrefix(fmt.Sprintf("%s_game", c.Name))}
}

func (uc *GameUseCase) Create(ctx context.Context, item *Game) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Create(ctx, item)
		})
	})
}

func (uc *GameUseCase) Get(ctx context.Context, id uint64) (rp *Game, err error) {
	rp = &Game{}
	action := fmt.Sprintf("get_%d", id)
	str, ok := uc.cache.Get(ctx, action, func(ctx context.Context) (string, bool) {
		return uc.get(ctx, action, id)
	})
	if ok {
		utils.Json2Struct(&rp, str)
		if rp.Id == constant.UI0 {
			err = reason.ErrorNotFound("%s Game.id: %d", i18n.FromContext(ctx).T(RecordNotFound), id)
		}
		return
	}
	err = reason.ErrorTooManyRequests(i18n.FromContext(ctx).T(TooManyRequests))
	return
}

func (uc *GameUseCase) get(ctx context.Context, action string, id uint64) (res string, ok bool) {
	// read data from db and write to cache
	rp := &Game{}
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

func (uc *GameUseCase) Find(ctx context.Context, condition *FindGame) (rp []Game) {
	// use md5 string as cache replay json str, key is short
	action := fmt.Sprintf("find_%s", utils.StructMd5(condition))
	str, ok := uc.cache.Get(ctx, action, func(ctx context.Context) (string, bool) {
		return uc.find(ctx, action, condition)
	})
	if ok {
		var cache FindGameCache
		utils.Json2Struct(&cache, str)
		condition.Page = cache.Page
		rp = cache.List
	}
	return
}

func (uc *GameUseCase) find(ctx context.Context, action string, condition *FindGame) (res string, ok bool) {
	// read data from db and write to cache
	list := uc.repo.Find(ctx, condition)
	var cache FindGameCache
	cache.List = list
	cache.Page = condition.Page
	res = utils.Struct2Json(cache)
	uc.cache.Set(ctx, action, res, len(list) == 0)
	ok = true
	return
}

func (uc *GameUseCase) Update(ctx context.Context, item *UpdateGame) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Update(ctx, item)
			return
		})
	})
}

func (uc *GameUseCase) Delete(ctx context.Context, ids ...uint64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Delete(ctx, ids...)
			return
		})
	})
}
