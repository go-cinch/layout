package data

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-cinch/common/id"
	"github.com/go-cinch/common/log"
	glog "github.com/go-cinch/common/plugins/gorm/log"
	"github.com/go-cinch/common/plugins/gorm/tenant"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-cinch/layout/internal/db"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewRedis, NewDB, NewSonyflake, NewTracer, NewData, NewTransaction, NewCache,
	NewAuthClient,
	NewGameRepo,
)

// Data .
type Data struct {
	tenant    *tenant.Tenant
	redis     redis.UniversalClient
	sonyflake *id.Sonyflake

	auth auth.AuthClient
}

// NewData .
func NewData(
	redis redis.UniversalClient,
	gormTenant *tenant.Tenant,
	sonyflake *id.Sonyflake,
	tp *trace.TracerProvider,
	auth auth.AuthClient,
) (d *Data, cleanup func()) {
	d = &Data{
		redis:     redis,
		tenant:    gormTenant,
		sonyflake: sonyflake,
		auth:      auth,
	}
	cleanup = func() {
		if tp != nil {
			tp.Shutdown(context.Background())
		}
		log.Info("clean data")
	}
	return
}

type contextTxKey struct{}

// Tx is transaction wrapper
func (d *Data) Tx(ctx context.Context, handler func(ctx context.Context) error) error {
	return d.tenant.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return handler(ctx)
	})
}

// DB can get tx from ctx, if not exist return db
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.tenant.DB(ctx)
}

// HiddenSQL return a hidden sql ctx
func (*Data) HiddenSQL(ctx context.Context) context.Context {
	ctx = glog.NewHiddenSqlContext(ctx)
	return ctx
}

// Cache can get cache instance
func (d *Data) Cache() redis.UniversalClient {
	return d.redis
}

// Id can get unique id
func (d *Data) Id(ctx context.Context) uint64 {
	return d.sonyflake.Id(ctx)
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewRedis is initialize redis connection from config
func NewRedis(c *conf.Bootstrap) (client redis.UniversalClient, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var u *url.URL
	u, err = url.Parse(c.Data.Redis.Dsn)
	if err != nil {
		err = errors.WithMessage(err, "initialize redis failed")
		return
	}
	u.User = url.UserPassword(u.User.Username(), "***")
	showDsn, _ := url.PathUnescape(u.String())
	client, err = utils.ParseRedisURI(c.Data.Redis.Dsn)
	if err != nil {
		err = errors.WithMessage(err, "initialize redis failed")
		return
	}
	err = client.Ping(ctx).Err()
	if err != nil {
		err = errors.WithMessage(err, "initialize redis failed")
		return
	}
	log.
		WithField("redis.dsn", showDsn).
		Info("initialize redis success")
	return
}

// NewDB is initialize db connection from config
func NewDB(c *conf.Bootstrap) (gormTenant *tenant.Tenant, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()

	ops := make([]func(*tenant.Options), 0, len(c.Data.Database.Tenants)+3)
	if len(c.Data.Database.Tenants) > 0 {
		for k, v := range c.Data.Database.Tenants {
			ops = append(ops, tenant.WithDSN(k, v))
		}
	} else {
		dsn := c.Data.Database.Dsn
		if dsn == "" {
			dsn = fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?%s",
				c.Data.Database.Username,
				c.Data.Database.Password,
				c.Data.Database.Endpoint,
				c.Data.Database.Schema,
				c.Data.Database.Query,
			)
		}
		ops = append(ops, tenant.WithDSN("", dsn))
	}
	ops = append(ops, tenant.WithSQLFile(db.SQLFiles))
	ops = append(ops, tenant.WithSQLRoot(db.SQLRoot))

	gormTenant, err = tenant.New(ops...)
	if err != nil {
		err = errors.WithMessage(err, "initialize db failed")
		return
	}
	err = gormTenant.Migrate()
	if err != nil {
		err = errors.WithMessage(err, "initialize db failed")
		return
	}
	log.Info("initialize db success")
	return
}

// NewSonyflake is initialize sonyflake id generator
func NewSonyflake(c *conf.Bootstrap) (sf *id.Sonyflake, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	machineId, _ := strconv.ParseUint(c.Server.MachineId, 10, 16)
	sf = id.NewSonyflake(
		id.WithSonyflakeMachineId(uint16(machineId)),
		id.WithSonyflakeStartTime(time.Date(100, 10, 10, 0, 0, 0, 0, time.UTC)),
	)
	if sf.Error != nil {
		err = errors.WithMessage(sf.Error, "initialize sonyflake failed")
		return
	}
	log.
		WithField("machine.id", machineId).
		Info("initialize sonyflake success")
	return
}
