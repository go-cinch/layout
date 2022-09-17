package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUseCase, NewGreeterWithCacheUseCase)

type Transaction interface {
	Tx(ctx context.Context, handler func(ctx context.Context) error) error
}

type Cache interface {
	// Register is register cache key prefix
	Register(prefix string)
	// Get is get cache data by key from redis
	Get(ctx context.Context, action string) (string, bool)
	// Set is set data to redis
	Set(ctx context.Context, action, data string)
	// SetWithExpiration is set data to redis with custom expiration
	SetWithExpiration(ctx context.Context, action, data string, seconds int64)
	// Flush is clean association cache if handler err=nil
	Flush(ctx context.Context, handler func(ctx context.Context) error) error
	// Lock is get lock from redis, false means don't get lock
	Lock(ctx context.Context, action string) bool
	// Unlock is release lock from redis
	Unlock(ctx context.Context, action string)
}
