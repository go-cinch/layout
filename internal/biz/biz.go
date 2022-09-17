package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUseCase)

type Transaction interface {
	Tx(ctx context.Context, handler func(ctx context.Context) error) error
}
