package data

import (
	"context"
	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/reason"
	"github.com/go-cinch/layout/internal/biz"
	"strings"
)

type gameRepo struct {
	data *Data
}

// Game is database fields map
type Game struct {
	Id   uint64 `json:"id,string"` // auto increment id
	Name string `json:"name"`      // name
	Age  int32  `json:"age"`       // age
}

func NewGameRepo(data *Data) biz.GameRepo {
	return &gameRepo{
		data: data,
	}
}

func (ro gameRepo) Create(ctx context.Context, item *biz.Game) (err error) {
	var m Game
	err = ro.NameExists(ctx, item.Name)
	if err == nil {
		err = reason.ErrorIllegalParameter("%s `name`: %s", i18n.FromContext(ctx).T(biz.DuplicateField), item.Name)
		return
	}
	copierx.Copy(&m, item)
	db := ro.data.DB(ctx)
	m.Id = ro.data.Id(ctx)
	err = db.Create(&m).Error
	return
}

func (ro gameRepo) Get(ctx context.Context, id uint64) (item *biz.Game, err error) {
	item = &biz.Game{}
	var m Game
	ro.data.DB(ctx).
		Where("`id` = ?", id).
		First(&m)
	if m.Id == constant.UI0 {
		err = reason.ErrorNotFound("%s Game.id: %d", i18n.FromContext(ctx).T(biz.RecordNotFound), id)
		return
	}
	copierx.Copy(&item, m)
	return
}

func (ro gameRepo) Find(ctx context.Context, condition *biz.FindGame) (rp []biz.Game) {
	db := ro.data.DB(ctx)
	db = db.
		Model(&Game{}).
		Order("id DESC")
	rp = make([]biz.Game, 0)
	list := make([]Game, 0)
	if condition.Name != nil {
		db.Where("`name` LIKE ?", strings.Join([]string{"%", *condition.Name, "%"}, ""))
	}
	if condition.Age != nil {
		db.Where("`age` = ?", condition.Age)
	}
	condition.Page.Primary = "id"
	condition.Page.
		WithContext(ctx).
		Query(db).
		Find(&list)
	copierx.Copy(&rp, list)
	return
}

func (ro gameRepo) Update(ctx context.Context, item *biz.UpdateGame) (err error) {
	var m Game
	db := ro.data.DB(ctx)
	db.
		Where("`id` = ?", item.Id).
		First(&m)
	if m.Id == constant.UI0 {
		err = reason.ErrorNotFound("%s Game.id: %d", i18n.FromContext(ctx).T(biz.RecordNotFound), item.Id)
		return
	}
	change := make(map[string]interface{})
	utils.CompareDiff(m, item, &change)
	if len(change) == 0 {
		err = reason.ErrorIllegalParameter(i18n.FromContext(ctx).T(biz.DataNotChange))
		return
	}
	if item.Name != nil && *item.Name != m.Name {
		err = ro.NameExists(ctx, *item.Name)
		if err == nil {
			err = reason.ErrorIllegalParameter("%s `name`: %s", i18n.FromContext(ctx).T(biz.DuplicateField), *item.Name)
			return
		}
	}
	err = db.
		Model(&m).
		Updates(&change).Error
	return
}

func (ro gameRepo) Delete(ctx context.Context, ids ...uint64) (err error) {
	db := ro.data.DB(ctx)
	err = db.
		Where("`id` IN (?)", ids).
		Delete(&Game{}).Error
	return
}

func (ro gameRepo) NameExists(ctx context.Context, name string) (err error) {
	var m Game
	db := ro.data.DB(ctx)
	arr := strings.Split(name, ",")
	for _, item := range arr {
		db.
			Where("`name` = ?", item).
			First(&m)
		if m.Id == constant.UI0 {
			err = reason.ErrorNotFound("%s Game.name: %s", i18n.FromContext(ctx).T(biz.RecordNotFound), item)
			return
		}
	}
	return
}
