package data

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/migrate"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewTransaction, NewGreeterRepo)

// Data .
type Data struct {
	db *gorm.DB
}

type contextTxKey struct{}

func (d *Data) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

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
func NewData(db *gorm.DB) (d *Data, cleanup func()) {
	d = &Data{
		db: db,
	}
	cleanup = func() {
		log.Info("clean data")
	}
	return
}

// NewDB gorm db without tx
func NewDB(c *conf.Data) (db *gorm.DB, err error) {
	err = migrate.Do(
		migrate.WithUri(c.Database.Dsn),
		migrate.WithFs(conf.SqlFiles),
		migrate.WithFsRoot("db"),
		migrate.WithBefore(func(ctx context.Context) (err error) {
			db, err = gorm.Open(m.Open(c.Database.Dsn), &gorm.Config{})
			return
		}),
	)
	var showDsn string
	cfg, e := mysql.ParseDSN(c.Database.Dsn)
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
