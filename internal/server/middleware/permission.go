package middleware

import (
	"context"
	jwtLocal "github.com/go-cinch/common/jwt"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/api/reason"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
				var reply metadata.MD
				res, err := authClient.Permission(ctx,
					&auth.PermissionRequest{
						Resource: resource,
					},
					grpc.Header(&reply),
				)
				if err != nil {
					return
				}
				if !res.Pass {
					err = reason.ErrorForbidden(i18n.FromContext(ctx).T(biz.NoPermission))
					return
				}
				ctx = jwtLocal.NewServerContextByReplyMD(ctx, reply)
				return handler(ctx, req)
			}
		},
	).Match(permissionWhitelist()).Build()
}
