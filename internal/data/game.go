package data

import (
	"context"
	"strings"

	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/data/model"
	"github.com/go-cinch/layout/internal/data/query"
	"gorm.io/gen"
)

type gameRepo struct {
	data *Data
}

func NewGameRepo(data *Data) biz.GameRepo {
	return &gameRepo{
		data: data,
	}
}

func (ro gameRepo) Create(ctx context.Context, item *biz.CreateGame) (err error) {
	err = ro.NameExists(ctx, item.Name)
	if err == nil {
		err = biz.ErrDuplicateField(ctx, "name", item.Name)
		return
	}
	var m model.Game
	copierx.Copy(&m, item)
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	m.ID = ro.data.ID(ctx)
	err = db.Create(&m)
	return
}

func (ro gameRepo) Get(ctx context.Context, id uint64) (item *biz.Game, err error) {
	item = &biz.Game{}
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	m := db.GetByID(id)
	if m.ID == constant.UI0 {
		err = biz.ErrRecordNotFound(ctx)
		return
	}
	copierx.Copy(&item, m)
	return
}

func (ro gameRepo) Find(ctx context.Context, condition *biz.FindGame) (rp []biz.Game) {
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	rp = make([]biz.Game, 0)
	list := make([]model.Game, 0)
	conditions := make([]gen.Condition, 0, 2)
	if condition.Name != nil {
		conditions = append(conditions, p.Name.Like(strings.Join([]string{"%", *condition.Name, "%"}, "")))
	}
	condition.Page.Primary = "id"
	condition.Page.
		WithContext(ctx).
		Query(
			db.
				Order(p.ID.Desc()).
				Where(conditions...).
				UnderlyingDB(),
		).
		Find(&list)
	copierx.Copy(&rp, list)
	return
}

func (ro gameRepo) Update(ctx context.Context, item *biz.UpdateGame) (err error) {
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	m := db.GetByID(item.ID)
	if m.ID == constant.UI0 {
		err = biz.ErrRecordNotFound(ctx)
		return
	}
	change := make(map[string]interface{})
	utils.CompareDiff(m, item, &change)
	if len(change) == 0 {
		err = biz.ErrDataNotChange(ctx)
		return
	}
	if item.Name != nil && *item.Name != m.Name {
		err = ro.NameExists(ctx, *item.Name)
		if err == nil {
			err = biz.ErrDuplicateField(ctx, "name", *item.Name)
			return
		}
	}
	_, err = db.
		Where(p.ID.Eq(item.ID)).
		Updates(&change)
	return
}

func (ro gameRepo) Delete(ctx context.Context, ids ...uint64) (err error) {
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	_, err = db.
		Where(p.ID.In(ids...)).
		Delete()
	return
}

func (ro gameRepo) NameExists(ctx context.Context, name string) (err error) {
	p := query.Use(ro.data.DB(ctx)).Game
	db := p.WithContext(ctx)
	arr := strings.Split(name, ",")
	for _, item := range arr {
		res := db.GetByCol("name", item)
		if res.ID == constant.UI0 {
			err = biz.ErrRecordNotFound(ctx)
			log.
				WithError(err).
				Warn("invalid `name`: %s", name)
			return
		}
	}
	return
}
