package service

import (
	helloword "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHellowordService)

// HellowordService is a greeter service.
type HellowordService struct {
	helloword.UnimplementedHelloworldServer

	log     *log.Helper
	greeter *biz.GreeterUseCase
}

// NewHellowordService new a helloword service.
func NewHellowordService(greeter *biz.GreeterUseCase, logger log.Logger) *HellowordService {
	return &HellowordService{greeter: greeter, log: log.NewHelper(logger)}
}
