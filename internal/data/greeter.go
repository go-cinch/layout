package data

import (
	"context"
	"fmt"
	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/reason"
	"github.com/go-cinch/layout/internal/biz"
	"strings"
)

type greeterRepo struct {
	data *Data
}

// Greeter is database fields map
type Greeter struct {
	Id   uint64 `json:"id,string"` // auto increment id
	Name string `json:"name"`      // name
	Age  int32  `json:"age"`       // age
}

func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}

func (ro greeterRepo) Create(ctx context.Context, item *biz.Greeter) (err error) {
	var m Greeter
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

func (ro greeterRepo) Get(ctx context.Context, id uint64) (item *biz.Greeter, err error) {
	item = &biz.Greeter{}
	var m Greeter
	ro.data.DB(ctx).
		Where("`id` = ?", id).
		First(&m)
	if m.Id == constant.UI0 {
		err = reason.ErrorNotFound("%s Greeter.id: %d", i18n.FromContext(ctx).T(biz.RecordNotFound), id)
		return
	}
	copierx.Copy(&item, m)
	return
}

func (ro greeterRepo) Find(ctx context.Context, condition *biz.FindGreeter) (rp []biz.Greeter) {
	db := ro.data.DB(ctx)
	db = db.
		Model(&Greeter{}).
		Order("id DESC")
	rp = make([]biz.Greeter, 0)
	list := make([]Greeter, 0)
	if condition.Name != nil {
		db.Where("`name` LIKE ?", fmt.Sprintf("%%%s%%", *condition.Name))
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

func (ro greeterRepo) Update(ctx context.Context, item *biz.UpdateGreeter) (err error) {
	var m Greeter
	db := ro.data.DB(ctx)
	db.
		Where("`id` = ?", item.Id).
		First(&m)
	if m.Id == constant.UI0 {
		err = reason.ErrorNotFound("%s Greeter.id: %d", i18n.FromContext(ctx).T(biz.RecordNotFound), item.Id)
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

func (ro greeterRepo) Delete(ctx context.Context, ids ...uint64) (err error) {
	db := ro.data.DB(ctx)
	err = db.
		Where("`id` IN (?)", ids).
		Delete(&Greeter{}).Error
	return
}

func (ro greeterRepo) NameExists(ctx context.Context, name string) (err error) {
	var m Greeter
	db := ro.data.DB(ctx)
	arr := strings.Split(name, ",")
	for _, item := range arr {
		db.
			Where("`name` = ?", item).
			First(&m)
		if m.Id == constant.UI0 {
			err = reason.ErrorNotFound("%s Greeter.name: %s", i18n.FromContext(ctx).T(biz.RecordNotFound), item)
			return
		}
	}
	return
}
