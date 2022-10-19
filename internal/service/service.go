package service

import (
	helloword "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/pkg/task"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHellowordService)

// HellowordService is a greeter service.
type HellowordService struct {
	helloword.UnimplementedHelloworldServer

	task    *task.Task
	greeter *biz.GreeterUseCase
}

// NewHellowordService new a helloword service.
func NewHellowordService(task *task.Task, greeter *biz.GreeterUseCase) *HellowordService {
	return &HellowordService{task: task, greeter: greeter}
}
