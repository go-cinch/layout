package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/go-cinch/common/log"
	_ "github.com/go-cinch/common/plugins/gorm/filter"
	"github.com/go-cinch/common/plugins/k8s/pod"
	"github.com/go-cinch/common/plugins/kratos/config/env"
	_ "github.com/go-cinch/common/plugins/kratos/encoding/yml"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	kratosLog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "layout"
	// EnvPrefix is the prefix of the env params
	EnvPrefix = "SERVICE"
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "c", "../../configs", "config path, eg: -c config.yml")
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
	// set default log before read config
	logOps := []func(*log.Options){
		log.WithLogger(
			kratosLog.With(
				kratosLog.NewStdLogger(os.Stdout),
				"ts", kratosLog.DefaultTimestamp,
				"service.id", id,
				"service.name", Name,
				"service.version", Version,
				"trace.id", tracing.TraceID(),
				"span.id", tracing.SpanID(),
			),
		),
		log.WithLevel(log.InfoLevel),
	}
	log.DefaultWrapper = log.NewWrapper(logOps...)
	c := config.New(
		config.WithSource(file.NewSource(flagConf)),
		config.WithResolver(
			env.NewRevolver(
				env.WithPrefix(EnvPrefix),
				env.WithLoaded(func(k string, v interface{}) {
					log.Info("env loaded: %s=%v", k, v)
				}),
			),
		),
	)
	defer c.Close()

	fields := log.Fields{
		"conf": flagConf,
	}
	if err := c.Load(); err != nil {
		log.
			WithError(err).
			WithFields(fields).
			Fatal("load conf failed")
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.
			WithError(err).
			WithFields(fields).
			Fatal("scan conf failed")
	}
	bc.Name = Name
	bc.Version = Version
	// override log level after read config
	logOps = append(logOps, log.WithLevel(log.NewLevel(bc.Server.LogLevel)))
	log.DefaultWrapper = log.NewWrapper(logOps...)
	if bc.Server.MachineId == "" {
		// if machine id not set, gen from pod ip
		machineId, err := pod.MachineId()
		if err == nil {
			bc.Server.MachineId = strconv.FormatUint(uint64(machineId), 10)
		} else {
			bc.Server.MachineId = "0"
		}
	}

	app, cleanup, err := wireApp(&bc)
	if err != nil {
		str := utils.Struct2Json(&bc)
		log.
			WithError(err).
			Error("wire app failed")
		// env str maybe very long, log with another line
		log.
			WithFields(fields).
			Fatal(str)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		log.
			WithError(err).
			WithFields(fields).
			Fatal("run app failed")
	}
}
