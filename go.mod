module github.com/go-cinch/layout

go 1.20

require (
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/go-cinch/common/constant v1.0.5
	github.com/go-cinch/common/copierx v1.0.3
	github.com/go-cinch/common/i18n v1.0.6
	github.com/go-cinch/common/id v1.0.4
	github.com/go-cinch/common/log v1.1.0
	github.com/go-cinch/common/middleware/i18n v1.0.5
	github.com/go-cinch/common/middleware/tenant v1.0.1
	github.com/go-cinch/common/middleware/trace v1.0.1
	github.com/go-cinch/common/mock v1.0.0
	github.com/go-cinch/common/page v1.0.4
	github.com/go-cinch/common/plugins/gorm/filter v1.0.0
	github.com/go-cinch/common/plugins/gorm/log v1.0.4
	github.com/go-cinch/common/plugins/gorm/tenant v1.0.1
	github.com/go-cinch/common/plugins/k8s/pod v1.0.0
	github.com/go-cinch/common/plugins/kratos/config/env v1.0.0
	github.com/go-cinch/common/plugins/kratos/encoding/yml v1.0.1
	github.com/go-cinch/common/proto/params v1.0.0
	github.com/go-cinch/common/utils v1.0.4
	github.com/go-cinch/common/worker v1.0.6
	github.com/go-kratos/kratos/v2 v2.7.0
	github.com/golang-module/carbon/v2 v2.2.8
	github.com/google/gnostic v0.7.0
	github.com/google/uuid v1.3.1
	github.com/google/wire v0.6.0
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.2.1
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	golang.org/x/text v0.14.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.31.0
	gorm.io/gen v0.3.22
	gorm.io/gorm v1.25.2
	gorm.io/plugin/dbresolver v1.4.1
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/alicebob/miniredis/v2 v2.31.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dolthub/flatbuffers/v23 v23.3.3-dh.2 // indirect
	github.com/dolthub/go-icu-regex v0.0.0-20230524105445-af7e7991c97e // indirect
	github.com/dolthub/go-mysql-server v0.17.0 // indirect
	github.com/dolthub/jsonpath v0.0.2-0.20230525180605-8dc13778fd72 // indirect
	github.com/dolthub/vitess v0.0.0-20230823204737-4a21a94e90c3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-cinch/common/migrate v1.0.4 // indirect
	github.com/go-cinch/common/nx v1.0.4 // indirect
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gocraft/dbr/v2 v2.7.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hibiken/asynq v0.24.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jinzhu/copier v0.4.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/r3labs/diff/v3 v3.0.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rubenv/sql-migrate v1.5.1 // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sony/sonyflake v1.1.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/tetratelabs/wazero v1.1.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yuin/gopher-lua v1.1.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	gopkg.in/src-d/go-errors.v1 v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/hints v1.1.2 // indirect
)
