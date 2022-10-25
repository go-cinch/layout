package data

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pkg/errors"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func NewAuthClient(c *conf.Bootstrap) (client auth.AuthClient, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	ops := []grpc.ClientOption{
		grpc.WithEndpoint(c.Client.Auth),
		grpc.WithMiddleware(
			tracing.Client(),
			metadata.Client(),
			recovery.Recovery(),
		),
	}
	conn, err := grpc.DialInsecure(context.Background(), ops...)
	if err != nil {
		err = errors.WithMessage(err, "initialize auth client failed")
		return
	}
	health := healthpb.NewHealthClient(conn)
	_, err = health.Check(context.Background(), &healthpb.HealthCheckRequest{})
	if err != nil {
		err = errors.WithMessage(err, "initialize auth client failed")
		return
	}
	client = auth.NewAuthClient(conn)
	log.
		WithField("auth.endpoint", c.Client.Auth).
		Info("initialize auth client success")
	return
}
