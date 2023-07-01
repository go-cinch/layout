package main

import (
	"flag"
	"os"
	"strings"

	"github.com/go-cinch/common/log"
	_ "github.com/go-cinch/common/plugins/gorm/filter"
	"github.com/go-cinch/common/plugins/gorm/tenant"
	_ "github.com/go-cinch/common/plugins/kratos/encoding/yml"
	"github.com/go-cinch/common/utils"
	"github.com/go-cinch/layout/internal/conf"
	k8sConfig "github.com/go-kratos/kratos/contrib/config/kubernetes/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
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
	EnvPrefix = "LAYOUT_"
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string
	// flagK8sNamespace is read config from k8s's configmap which namespace is xxx
	flagK8sNamespace string
	// flagK8sNamespace is read config from k8s's configmap which label is app=xxx
	flagK8sLabel string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "c", "../../configs", "config path, eg: -c config.yml")
	flag.StringVar(&flagK8sNamespace, "n", "", "k8s namespace, eg: -n cinch")
	flag.StringVar(&flagK8sLabel, "l", "", "k8s configmap label, eg: -l layout")
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
				"service.id", id,
				"service.name", Name,
				"service.version", Version,
				"trace.id", tracing.TraceID(),
				"span.id", tracing.SpanID(),
				"tenant.id", tenant.ID(),
			),
		),
	)
	sources := []config.Source{
		env.NewSource(EnvPrefix),
	}
	if flagK8sNamespace != "" || flagK8sLabel != "" {
		namespace := "default"
		if flagK8sNamespace != "" {
			namespace = flagK8sNamespace
		}
		opts := []k8sConfig.Option{
			k8sConfig.Namespace(namespace),
		}
		if flagK8sLabel != "" {
			opts = append(opts, k8sConfig.LabelSelector(strings.Join([]string{"app", flagK8sLabel}, "=")))
		}
		sources = append(sources, k8sConfig.NewSource(opts...))
	} else {
		sources = append(sources, file.NewSource(flagConf))
	}
	c := config.New(config.WithSource(sources...))
	defer c.Close()

	if err := c.Load(); err != nil {
		log.
			WithError(err).
			WithFields(log.Fields{
				"flag.c": flagConf,
				"flag.n": flagK8sNamespace,
				"flag.l": flagK8sLabel,
			}).
			Fatal("load conf failed")
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.
			WithError(err).
			WithFields(log.Fields{
				"flag.c": flagConf,
				"flag.n": flagK8sNamespace,
				"flag.l": flagK8sLabel,
			}).
			Fatal("scan conf failed")
	}
	bc.Name = Name
	bc.Version = Version

	app, cleanup, err := wireApp(&bc)
	if err != nil {
		str := utils.Struct2Json(bc)
		log.
			WithError(err).
			WithFields(log.Fields{
				"flag.c": flagConf,
				"flag.n": flagK8sNamespace,
				"flag.l": flagK8sLabel,
				"json":   str,
			}).
			Fatal("wire app failed")
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		log.
			WithError(err).
			WithFields(log.Fields{
				"flag.c": flagConf,
				"flag.n": flagK8sNamespace,
				"flag.l": flagK8sLabel,
			}).
			Fatal("run app failed")
	}
}
