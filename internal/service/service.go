package service

import (
	"github.com/go-cinch/layout/api/game"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/pkg/idempotent"
	"github.com/go-cinch/layout/internal/pkg/task"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGameService)

// GameService is a game service.
type GameService struct {
	game.UnimplementedGameServer

	task       *task.Task
	idempotent *idempotent.Idempotent
	game       *biz.GameUseCase
}

// NewGameService new a service.
func NewGameService(task *task.Task, idempotent *idempotent.Idempotent, game *biz.GameUseCase) *GameService {
	return &GameService{task: task, idempotent: idempotent, game: game}
}
