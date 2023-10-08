module github.com/go-cinch/layout

go 1.20

require (
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/go-cinch/common/constant v1.0.4
	github.com/go-cinch/common/i18n v1.0.4
	github.com/go-cinch/common/id v1.0.4
	github.com/go-cinch/common/log v1.0.4
	github.com/go-cinch/common/middleware/i18n v1.0.1
	github.com/go-cinch/common/middleware/tenant v1.0.1
	github.com/go-cinch/common/middleware/trace v1.0.1
	github.com/go-cinch/common/plugins/gorm/filter v1.0.0
	github.com/go-cinch/common/plugins/gorm/log v1.0.4
	github.com/go-cinch/common/plugins/gorm/tenant v1.0.1
	github.com/go-cinch/common/plugins/k8s/pod v1.0.0
	github.com/go-cinch/common/plugins/kratos/config/env v1.0.0
	github.com/go-cinch/common/plugins/kratos/encoding/yml v1.0.1
	github.com/go-cinch/common/proto/params v1.0.0
	github.com/go-cinch/common/utils v1.0.4
	github.com/go-cinch/common/worker v1.0.4
	github.com/go-kratos/kratos/v2 v2.7.0
	github.com/google/wire v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.2.1
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	golang.org/x/text v0.11.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.31.0
	gorm.io/gorm v1.25.2
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-cinch/common/migrate v1.0.4 // indirect
	github.com/go-cinch/common/nx v1.0.4 // indirect
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang-module/carbon/v2 v2.2.8 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/hibiken/asynq v0.24.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/r3labs/diff/v3 v3.0.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rubenv/sql-migrate v1.5.1 // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sony/sonyflake v1.1.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/gen v0.3.22 // indirect
	gorm.io/hints v1.1.2 // indirect
	gorm.io/plugin/dbresolver v1.4.1 // indirect
)
