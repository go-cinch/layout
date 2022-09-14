package data

import (
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data) (d *Data, cleanup func()) {
	d = &Data{}
	cleanup = func() {
		log.Info("closing the data resources")
	}
	return
}
