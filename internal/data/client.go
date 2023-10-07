package data

import (
	"context"
	"strings"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pkg/errors"
	g "google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func NewAuthClient(c *conf.Bootstrap) (auth.AuthClient, error) {
	return NewClient[auth.AuthClient](
		"auth",
		c.Client.Auth,
		c.Client.Health,
		c.Client.Timeout.AsDuration(),
		auth.NewAuthClient,
	)
}

func NewClient[T any](name, endpoint string, health bool, timeout time.Duration, newClient func(cc g.ClientConnInterface) T) (client T, err error) {
	ops := []grpc.ClientOption{
		grpc.WithEndpoint(endpoint),
		grpc.WithMiddleware(
			tracing.Client(),
			metadata.Client(),
			recovery.Recovery(),
		),
		grpc.WithOptions(g.WithDisableHealthCheck()),
		grpc.WithTimeout(timeout),
	}
	conn, err := grpc.DialInsecure(context.Background(), ops...)
	if err != nil {
		err = errors.WithMessage(err, strings.Join([]string{"initialize", name, "client failed"}, " "))
		return
	}
	if health {
		healthClient := healthpb.NewHealthClient(conn)
		_, err = healthClient.Check(context.Background(), &healthpb.HealthCheckRequest{})
		if err != nil {
			err = errors.WithMessage(err, strings.Join([]string{name, "healthcheck failed"}, " "))
			return
		}
	}
	client = newClient(conn)
	log.
		WithField("endpoint", endpoint).
		Info(strings.Join([]string{"initialize", name, "client success"}, " "))
	return
}
