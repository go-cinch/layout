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
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, idt *idempotent.Idempotent, authClient auth.AuthClient, svc *service.GreeterService) *grpc.Server {
	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		ratelimit.Server(),
	}
	if c.Tracer.Enable {
		middlewares = append(middlewares, tracing.Server(), commonMiddleware.TraceId())
	}
	middlewares = append(
		middlewares,
		logging.Server(log.DefaultWrapper.Options().Logger()),
		metadata.Server(),
		localMiddleware.Permission(authClient),
		validate.Validator(),
		localMiddleware.Idempotent(idt),
	)
	var opts = []grpc.ServerOption{grpc.Middleware(middlewares...)}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	greeter.RegisterGreeterServer(srv, svc)
	return srv
}
