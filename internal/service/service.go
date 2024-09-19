package service

import (
	"github.com/go-cinch/common/worker"
	"github.com/go-cinch/layout/api/game"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGameService)

// GameService is a game service.
type GameService struct {
	game.UnimplementedGameServer

	task *worker.Worker
	game *biz.GameUseCase
}

// NewGameService new a service.
func NewGameService(task *worker.Worker, game *biz.GameUseCase) *GameService {
	return &GameService{task: task, game: game}
}
