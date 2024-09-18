package mock

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/go-cinch/common/mock"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-cinch/layout/internal/data"
	"github.com/go-cinch/layout/internal/service"
	"github.com/go-kratos/kratos/v2/transport"
)

func GameService() (gameService *service.GameService) {
	gameUseCase := GameUseCase()
	gameService = service.NewGameService(nil, gameUseCase)
	return
}

func GameUseCase() (gameUseCase *biz.GameUseCase) {
	c, dataData, cache := Data()
	gameRepo := GameRepo()
	transaction := data.NewTransaction(dataData)
	gameUseCase = biz.NewGameUseCase(c, gameRepo, transaction, cache)
	return
}

func GameRepo() biz.GameRepo {
	_, d, _ := Data()
	return data.NewGameRepo(d)
}

type headerCarrier http.Header

func (hc headerCarrier) Get(key string) string { return http.Header(hc).Get(key) }

func (hc headerCarrier) Set(key string, value string) { http.Header(hc).Set(key, value) }

func (hc headerCarrier) Add(key string, value string) { http.Header(hc).Add(key, value) }

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice value associated with the passed key.
func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}

func newUserHeader(k, v string) *headerCarrier {
	header := &headerCarrier{}
	header.Set(k, v)
	return header
}

type Transport struct {
	kind      transport.Kind
	endpoint  string
	operation string
	reqHeader transport.Header
}

func (tr *Transport) Kind() transport.Kind {
	return tr.kind
}

func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

func (tr *Transport) Operation() string {
	return tr.operation
}

func (tr *Transport) RequestHeader() transport.Header {
	return tr.reqHeader
}

func (*Transport) ReplyHeader() transport.Header {
	return nil
}

func NewContextWithUserId(ctx context.Context, u string) context.Context {
	tr := &Transport{
		reqHeader: newUserHeader("X-Md-Global-Code", u),
	}
	return transport.NewServerContext(ctx, tr)
}

var (
	onceC *conf.Bootstrap
	once  sync.Once
)

func Data() (c *conf.Bootstrap, dataData *data.Data, cache biz.Cache) {
	debug.SetGCPercent(-1)
	once.Do(func() {
		onceC = MySQLAndRedis()
	})
	c = onceC
	// os.Setenv("COPIERX_UTC", "true")

	universalClient, err := data.NewRedis(c)
	if err != nil {
		panic(err)
	}
	tenant, err := data.NewDB(c)
	if err != nil {
		panic(err)
	}
	sonyflake, err := data.NewSonyflake(c)
	if err != nil {
		panic(err)
	}
	tracerProvider, err := data.NewTracer(c)
	if err != nil {
		panic(err)
	}
	dataData, _ = data.NewData(universalClient, tenant, sonyflake, tracerProvider, nil)
	cache = data.NewCache(c, universalClient)
	return c, dataData, cache
}

func MySQLAndRedis() *conf.Bootstrap {
	host1, port1, err := mock.NewMySQL()
	if err != nil {
		panic(err)
	}
	host2, port2, err := mock.NewRedis()
	if err != nil {
		panic(err)
	}
	// host1 = "127.0.0.1"
	// port1 = 3306
	return &conf.Bootstrap{
		Server: &conf.Server{
			MachineId: "123",
		},
		Log: &conf.Log{
			Level:   "debug",
			JSON:    false,
			ShowSQL: true,
		},
		Data: &conf.Data{
			Database: &conf.Data_Database{
				Endpoint: fmt.Sprintf("%s:%d", host1, port1),
				Username: "root",
				Password: "passwd",
				Schema:   "game",
				Query:    "charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=10000ms",
			},
			Redis: &conf.Data_Redis{
				Dsn: fmt.Sprintf("redis://%s:%d", host2, port2),
			},
		},
		Tracer: &conf.Tracer{
			Enable: true,
			Otlp:   &conf.Tracer_Otlp{},
			Stdout: &conf.Tracer_Stdout{},
		},
	}
}
