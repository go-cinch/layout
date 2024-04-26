package main

import (
	"flag"
	"github.com/go-cinch/common/log/caller"
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
	// beforeReadConfigLogLevel is log level before read config.
	beforeReadConfigLogLevel = log.DebugLevel

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
		log.WithJSON(false),
		log.WithLevel(beforeReadConfigLogLevel),
		log.WithValuer("service.id", id),
		log.WithValuer("service.name", Name),
		log.WithValuer("service.version", Version),
		log.WithValuer("trace.id", tracing.TraceID()),
		log.WithValuer("span.id", tracing.SpanID()),
		log.WithCallerOptions(
			caller.WithSource(false),
			caller.WithLevel(2),
			caller.WithVersion(true),
		),
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
	logOps = append(logOps,
		[]func(*log.Options){
			log.WithLevel(log.NewLevel(bc.Log.Level)),
			log.WithJSON(bc.Log.JSON),
		}...,
	)
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
	// os.Setenv("COPIERX_UTC", "true")

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
