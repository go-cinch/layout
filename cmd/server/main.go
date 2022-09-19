package main

import (
	"flag"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	kratosLog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "cinch-layout"
	// EnvPrefix is the prefix of the env params
	EnvPrefix = "CINCH_"
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yml")
}

func newApp(gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.DefaultWrapper.Options().Logger()),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	log.DefaultWrapper = log.NewWrapper(
		log.WithLogger(
			kratosLog.With(
				kratosLog.NewStdLogger(os.Stdout),
				"ts", kratosLog.DefaultTimestamp,
				"caller", kratosLog.DefaultCaller,
				"service.id", id,
				"service.name", Name,
				"service.version", Version,
				"trace.id", tracing.TraceID(),
				"span.id", tracing.SpanID(),
			),
		),
	)
	c := config.New(
		config.WithSource(
			env.NewSource(EnvPrefix),
			file.NewSource(flagConf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	bc.Name = Name
	bc.Version = Version

	app, cleanup, err := wireApp(&bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
