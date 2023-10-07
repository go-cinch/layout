package middleware

import (
	"context"
	"strings"

	jwtLocal "github.com/go-cinch/common/jwt"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Permission(c *conf.Bootstrap, authClient auth.AuthClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (rp interface{}, err error) {
			var v auth.PermissionRequest
			var resource, method, path string
			if tr, ok := transport.FromServerContext(ctx); ok {
				switch tr.Kind() {
				case transport.KindHTTP:
					if ht, ok2 := tr.(kratosHttp.Transporter); ok2 {
						method = ht.Request().Method
						path = strings.Join([]string{c.Server.Http.Path, ht.Request().URL.Path}, "")
					}
				case transport.KindGRPC:
					resource = tr.Operation()
				}
			}
			v.Resource = &resource
			v.Method = &method
			v.Uri = &path
			ctx = jwtLocal.AppendToClientContext(ctx)
			var reply metadata.MD
			_, err = authClient.Permission(ctx, &v, grpc.Header(&reply))
			if err != nil {
				err = biz.ErrNoPermission(ctx)
				return
			}
			ctx = jwtLocal.NewServerContextByReplyMD(ctx, reply)
			return handler(ctx, req)
		}
	}
}
