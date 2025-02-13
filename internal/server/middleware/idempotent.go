package middleware

import (
	"context"

	"github.com/go-cinch/common/idempotent"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/redis/go-redis/v9"
)

func Idempotent(rds redis.UniversalClient) middleware.Middleware {
	idt := idempotent.New(
		idempotent.WithPrefix("idempotent"),
		idempotent.WithRedis(rds),
	)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (rp interface{}, err error) {
			tr, _ := transport.FromServerContext(ctx)
			// check idempotent token if it has header
			token := tr.RequestHeader().Get("x-idempotent")
			if token == "" {
				return handler(ctx, req)
			}
			pass := idt.Check(ctx, token)
			if !pass {
				err = biz.ErrIdempotentTokenExpired(ctx)
				return
			}
			return handler(ctx, req)
		}
	}
}
