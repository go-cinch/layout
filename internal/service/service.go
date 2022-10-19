package service

import (
	"github.com/go-cinch/layout/api/greeter"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/pkg/task"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

// GreeterService is a greeter service.
type GreeterService struct {
	greeter.UnimplementedGreeterServer

	task    *task.Task
	greeter *biz.GreeterUseCase
}

// NewGreeterService new a service.
func NewGreeterService(task *task.Task, greeter *biz.GreeterUseCase) *GreeterService {
	return &GreeterService{task: task, greeter: greeter}
}
