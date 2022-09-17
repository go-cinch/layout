package service

import (
	helloword "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHellowordService)

// HellowordService is a greeter service.
type HellowordService struct {
	helloword.UnimplementedHelloworldServer

	greeter          *biz.GreeterUseCase
	greeterWithCache *biz.GreeterWithCacheUseCase
}

// NewHellowordService new a helloword service.
func NewHellowordService(greeter *biz.GreeterUseCase, greeterWithCache *biz.GreeterWithCacheUseCase) *HellowordService {
	return &HellowordService{greeter: greeter, greeterWithCache: greeterWithCache}
}
