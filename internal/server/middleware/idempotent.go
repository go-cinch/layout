package middleware

import (
	"context"
	"github.com/go-cinch/layout/internal/idempotent"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
)

func Idempotent(idt *idempotent.Idempotent) middleware.Middleware {
	return selector.Server(
		func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (rp interface{}, err error) {
				if tr, ok := transport.FromServerContext(ctx); ok {
					token := tr.RequestHeader().Get("X-IDEMPOTENT")
					if token != "" {
						if idt.Check(ctx, token) {
							return handler(ctx, req)
						}
						err = IdempotentTokenExpired
						return
					}
				}
				err = MissingIdempotentToken
				return
			}
		},
	).Match(idempotentBlacklist()).Build()
}
