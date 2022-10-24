package data

import (
	"context"
	"github.com/go-cinch/common/id"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/migrate"
	glog "github.com/go-cinch/common/plugins/gorm/log"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/api/auth"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/sdk/trace"
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/url"
	"strconv"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewRedis, NewDB, NewSonyflake, NewTracer, NewData, NewTransaction, NewCache,
	NewAuthClient,
	NewGreeterRepo,
)

// Data .
type Data struct {
	db        *gorm.DB
	redis     redis.UniversalClient
	sonyflake *id.Sonyflake

	auth auth.AuthClient
}

type contextTxKey struct{}

// Tx is transaction wrapper
func (d *Data) Tx(ctx context.Context, handler func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
	return d.db.WithContext(ctx)
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

// NewData .
func NewData(
	redis redis.UniversalClient, db *gorm.DB, sonyflake *id.Sonyflake, tp *trace.TracerProvider,
	auth auth.AuthClient,
) (d *Data, cleanup func()) {
	d = &Data{
		redis:     redis,
		db:        db,
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
func NewDB(c *conf.Bootstrap) (db *gorm.DB, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = migrate.Do(
		migrate.WithCtx(ctx),
		migrate.WithUri(c.Data.Database.Dsn),
		migrate.WithFs(conf.SqlFiles),
		migrate.WithFsRoot("db"),
		migrate.WithBefore(func(ctx context.Context) (err error) {
			l := glog.New(
				glog.WithColorful(true),
				glog.WithSlow(200),
			)
			db, err = gorm.Open(m.Open(c.Data.Database.Dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
				QueryFields: true,
				Logger:      l,
			})
			return
		}),
	)
	var showDsn string
	cfg, e := mysql.ParseDSN(c.Data.Database.Dsn)
	if e == nil {
		// hidden password
		cfg.Passwd = "***"
		showDsn = cfg.FormatDSN()
	}
	if err != nil {
		err = errors.WithMessage(err, "initialize mysql failed")
		return
	}
	log.
		WithField("db.dsn", showDsn).
		Info("initialize mysql success")
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
	sf = id.NewSonyflake(id.WithSonyflakeMachineId(uint16(machineId)))
	if sf.Error != nil {
		err = errors.WithMessage(sf.Error, "initialize sonyflake failed")
		return
	}
	log.
		WithField("sonyflake.id", machineId).
		Info("initialize sonyflake success")
	return
}
