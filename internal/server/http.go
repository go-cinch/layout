package server

import (
	"github.com/go-cinch/common/i18n"
	"github.com/go-cinch/common/log"
	i18nMiddleware "github.com/go-cinch/common/middleware/i18n"
	traceMiddleware "github.com/go-cinch/common/middleware/trace"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/api/game"
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
	"golang.org/x/text/language"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Bootstrap, idt *idempotent.Idempotent, authClient auth.AuthClient, svc *service.GameService) *http.Server {
	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		ratelimit.Server(),
		localMiddleware.Header(),
	}
	if c.Tracer.Enable {
		middlewares = append(middlewares, tracing.Server(), traceMiddleware.Id())
	}
	middlewares = append(
		middlewares,
		logging.Server(log.DefaultWrapper.Options().Logger()),
		i18nMiddleware.Translator(i18n.WithLanguage(language.Make(c.Server.Language)), i18n.WithFs(locales)),
	)
	if c.Server.Permission {
		middlewares = append(middlewares, localMiddleware.Permission(authClient))
	}
	if c.Server.Idempotent {
		middlewares = append(middlewares, localMiddleware.Idempotent(idt))
	}
	if c.Server.Validate {
		middlewares = append(middlewares, validate.Validator())
	}
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
	game.RegisterGameHTTPServer(srv, svc)
	return srv
}
