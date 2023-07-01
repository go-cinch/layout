package data

import (
	"github.com/go-cinch/layout/internal/biz"
)

type gameRepo struct {
	data *Data
}

func NewGameRepo(data *Data) biz.GameRepo {
	return &gameRepo{
		data: data,
	}
}
