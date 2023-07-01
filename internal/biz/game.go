package biz

import (
	"github.com/go-cinch/layout/internal/conf"
)

type GameRepo interface{}

type GameUseCase struct {
	c     *conf.Bootstrap
	repo  GameRepo
	tx    Transaction
	cache Cache
}

func NewGameUseCase(c *conf.Bootstrap, repo GameRepo, tx Transaction, cache Cache) *GameUseCase {
	return &GameUseCase{
		c:     c,
		repo:  repo,
		tx:    tx,
		cache: cache.WithPrefix("game"),
	}
}
