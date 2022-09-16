package data

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/migrate"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewTracer, NewData, NewDB, NewTransaction, NewGreeterRepo)

// Data .
type Data struct {
	db *gorm.DB
}

type contextTxKey struct{}

// Tx is transaction wrapper
func (d *Data) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB can get tx from ctx, if not exist return db
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(db *gorm.DB, tp *trace.TracerProvider) (d *Data, cleanup func()) {
	d = &Data{
		db: db,
	}
	cleanup = func() {
		tp.Shutdown(context.Background())
		log.Info("clean data")
	}
	return
}

// NewDB gorm db without tx
func NewDB(c *conf.Bootstrap) (db *gorm.DB, err error) {
	err = migrate.Do(
		migrate.WithUri(c.Data.Database.Dsn),
		migrate.WithFs(conf.SqlFiles),
		migrate.WithFsRoot("db"),
		migrate.WithBefore(func(ctx context.Context) (err error) {
			db, err = gorm.Open(m.Open(c.Data.Database.Dsn), &gorm.Config{})
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
		log.
			WithError(err).
			WithField("db.dsn", showDsn).
			Error("initialize mysql failed")
		return
	}
	log.
		WithField("db.dsn", showDsn).
		Info("initialize mysql success")
	return
}
