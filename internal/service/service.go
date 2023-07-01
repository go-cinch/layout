package service

import (
	"github.com/go-cinch/layout/api/game"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/pkg/task"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGameService)

// GameService is a game service.
type GameService struct {
	game.UnimplementedGameServer

	task *task.Task
	game *biz.GameUseCase
}

// NewGameService new a service.
func NewGameService(task *task.Task, game *biz.GameUseCase) *GameService {
	return &GameService{task: task, game: game}
}
