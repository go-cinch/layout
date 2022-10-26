package server

import (
	"github.com/go-cinch/common/log"
	commonMiddleware "github.com/go-cinch/common/middleware"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/api/greeter"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-cinch/layout/internal/pkg/idempotent"
	localMiddleware "github.com/go-cinch/layout/internal/server/middleware"
	"github.com/go-cinch/layout/internal/service"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Bootstrap, idt *idempotent.Idempotent, authClient auth.AuthClient, svc *service.GreeterService) *http.Server {
	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		ratelimit.Server(),
		localMiddleware.Header(),
	}
	if c.Tracer.Enable {
		middlewares = append(middlewares, tracing.Server(), commonMiddleware.TraceId())
	}
	middlewares = append(
		middlewares,
		logging.Server(log.DefaultWrapper.Options().Logger()),
		localMiddleware.Permission(authClient),
		validate.Validator(),
		localMiddleware.Idempotent(idt),
	)
	var opts = []http.ServerOption{http.Middleware(middlewares...)}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	greeter.RegisterGreeterHTTPServer(srv, svc)
	return srv
}
