package middleware

import (
	"context"
	jwtLocal "github.com/go-cinch/common/jwt"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
)

func Permission(authClient auth.AuthClient) middleware.Middleware {
	return selector.Server(
		func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (rp interface{}, err error) {
				var resource string
				if tr, ok := transport.FromServerContext(ctx); ok {
					resource = tr.Operation()
				}
				ctx = jwtLocal.AppendToClientContext(ctx)
				res, err := authClient.Permission(ctx, &auth.PermissionRequest{
					Resource: resource,
				})
				if err != nil {
					return
				}
				if !res.Pass {
					err = biz.NoPermission
					return
				}
				return handler(ctx, req)
			}
		},
	).Match(permissionWhitelist()).Build()
}
