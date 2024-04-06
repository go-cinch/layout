package biz

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/page"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/pkg/errors"
)

type Game struct {
	Id   uint64 `json:"id,string"`
	Name string `json:"name"`
}

type FindGame struct {
	Page page.Page `json:"page"`
	Name *string   `json:"name"`
}

type FindGameCache struct {
	Page page.Page `json:"page"`
	List []Game    `json:"list"`
}

type UpdateGame struct {
	Id   uint64  `json:"id,string"`
	Name *string `json:"name,omitempty"`
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
	return &GameUseCase{
		c:    c,
		repo: repo,
		tx:   tx,
		cache: cache.WithPrefix(strings.Join([]string{
			c.Name, "game",
		}, "_")),
	}
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
	action := strings.Join([]string{"get", strconv.FormatUint(id, 10)}, "_")
	str, err := uc.cache.Get(ctx, action, func(ctx context.Context) (string, error) {
		return uc.get(ctx, action, id)
	})
	if err != nil {
		return
	}
	utils.Json2Struct(&rp, str)
	if rp.Id == constant.UI0 {
		err = ErrRecordNotFound(ctx)
		return
	}
	return
}

func (uc *GameUseCase) get(ctx context.Context, action string, id uint64) (res string, err error) {
	// read data from db and write to cache
	rp := &Game{}
	item, err := uc.repo.Get(ctx, id)
	notFound := errors.Is(err, ErrRecordNotFound(ctx))
	if err != nil && !notFound {
		return
	}
	copierx.Copy(&rp, item)
	res = utils.Struct2Json(rp)
	uc.cache.Set(ctx, action, res, notFound)
	return
}

func (uc *GameUseCase) Find(ctx context.Context, condition *FindGame) (rp []Game, err error) {
	// use md5 string as cache replay json str, key is short
	action := strings.Join([]string{"find", utils.StructMd5(condition)}, "_")
	str, err := uc.cache.Get(ctx, action, func(ctx context.Context) (string, error) {
		return uc.find(ctx, action, condition)
	})
	if err != nil {
		return
	}
	var cache FindGameCache
	utils.Json2Struct(&cache, str)
	condition.Page = cache.Page
	rp = cache.List
	return
}

func (uc *GameUseCase) find(ctx context.Context, action string, condition *FindGame) (res string, err error) {
	// read data from db and write to cache
	list := uc.repo.Find(ctx, condition)
	var cache FindGameCache
	cache.List = list
	cache.Page = condition.Page
	res = utils.Struct2Json(cache)
	uc.cache.Set(ctx, action, res, len(list) == 0)
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
