package middleware

import (
	"context"

	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
)

const (
	WhitelistIdempotentCategory uint32 = 2
)

func Idempotent(authClient auth.AuthClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (rp interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = biz.ErrIdempotentMissingToken(ctx)
				return
			}
			var method, path string
			switch tr.Kind() {
			case transport.KindHTTP:
				if ht, ok3 := tr.(kratosHttp.Transporter); ok3 {
					method = ht.Request().Method
					path = ht.Request().URL.Path
				}
			}
			// check idempotent blacklist
			whitelist, err := authClient.HasWhitelist(ctx, &auth.HasWhitelistRequest{
				Category: WhitelistIdempotentCategory,
				Permission: &auth.HasWhitelistRequest_CheckPermission{
					Resource: tr.Operation(),
					Method:   method,
					Uri:      path,
				},
			})
			if err != nil {
				return
			}
			if !whitelist.Ok {
				return handler(ctx, req)
			}
			// check idempotent token
			token := tr.RequestHeader().Get("x-idempotent")
			if token == "" {
				err = biz.ErrIdempotentMissingToken(ctx)
				return
			}
			_, err = authClient.CheckIdempotent(ctx, &auth.CheckIdempotentRequest{Token: token})
			if err != nil {
				return
			}
			return handler(ctx, req)
		}
	}
}
