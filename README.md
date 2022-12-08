<h1 align="center">Go Cinch Layout</h1>

<div align="center">
Cinch Layout是一套轻量级微服务项目模版.
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/go-cinch/layout" alt="Go version"/>
<img src="https://img.shields.io/badge/Kratos-v2.5.3-brightgreen" alt="Kratos version"/>
<img src="https://img.shields.io/badge/MySQL-8.0-brightgreen" alt="MySQL version"/>
<img src="https://img.shields.io/badge/go--redis-v8.11.5-brightgreen" alt="Go redis version"/>
<img src="https://img.shields.io/badge/Gorm-1.24.0-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/badge/Wire-0.5.0-brightgreen" alt="Wire version"/>
<img src="https://img.shields.io/github/license/go-cinch/layout" alt="License"/>
</p>
</div>


# 起源


你的单体服务架构是否遇到一些问题, 不能满足业务需求? 那么微服务会是好的解决方案.

Cinch是一套轻量级微服务脚手架, 基于[Kratos], 节省基础服务搭建时间, 快速投入业务开发.

我们参考了Go的许多微服务架构, 结合实际需求, 最终选择简洁的[Kratos]作为基石(B站架构), 从架构的设计思路以及代码的书写格式和我们非常匹配.

> cinch意为简单的事, 小菜. 希望把复杂的事变得简单, 提升开发效率.

若你想深入学习微服务每个组件, 建议直接看Kratos官方文档. 本项目整合一些业务常用组件, 开箱即用, 并不会把每个组件都介绍那么详细.


# 特性


- `Proto` - proto协议同时开启gRPC & HTTP支持, 只需开发一次接口, 不用写两套
- `Jwt` - 认证, 用户登入登出一键搞定
- `Action` - 权限, 基于行为的权限校验
- `Redis` - 缓存, 内置防缓存穿透/缓存击穿/缓存雪崩示例
- `Gorm` - 数据库ORM管理框架, 可自行扩展多种数据库类型, 目前使用MySQL, 其他自行扩展
- `SqlMigrate` - 数据库迁移工具, 每次更新平滑迁移
- `Asynq` - 分布式定时任务(异步任务)
- `Log` - 日志, 在Kratos基础上增加一层包装, 无需每个方法传入
- `Embed` - go 1.16文件嵌入属性, 轻松将静态文件打包到编译后的二进制应用中
- `Opentelemetry` - 链路追踪, 跨服务调用可快速追踪完整链路
- `Idempotent` - 接口幂等性(解决重复点击或提交)
- `Pprof` - 内置性能分析开关, 对并发/性能测试友好
- `Wire` - 依赖注入, 编译时完成依赖注入
- `Swagger` - Api文档一键生成, 无需在代码里写注解
- `I18n` - 国际化支持, 简单切换多语言


# 项目结构

```
├── api                  // 各个微服务的proto/go文件, proto文件通过submodule管理包含proto后缀
│   ├── auth             // auth服务所需的go文件(通过命令生成)
│   ├── auth-proto       // auth服务proto文件
│   ├── reason
│   ├── reason-proto     // 公共reason(错误码建议统一管理, 不要不同服务搞不同的)
│   ├── xxx              // xxx服务所需的go文件(通过命令生成)
│   ├── xxx-proto        // xxx服务proto文件
│   └── ...
├── cmd                  // 程序主入口
│   └── order            // 项目名
│       ├── main.go
│       ├── wire.go      // wire依赖注入
│       └── wire_gen.go
├── configs              // 配置文件目录
│   ├── config.yaml      // 主配置文件
│   ├── client.yaml      // 配置grpc服务client, 如auth
│   └── ...              // 其他自定义配置文件以yaml结尾
├── internal             // 内部逻辑代码
│   ├── biz              // 业务逻辑的组装层, 类似 DDD 的 domain 层, data 类似 DDD 的 repo, 而 repo 接口在这里定义, 使用依赖倒置的原则. 
│   │   ├── biz.go
│   │   ├── greeter.go
│   │   └── reason.go    // 定义错误描述
│   ├── conf             // 内部使用的config的结构定义, 使用proto格式生成
│   │   ├── db           // sql文件目录, 每一次数据库变更都放在这里, 参考https://github.com/rubenv/sql-migrate
│   │   │   ├── xxx.sql  // sql文件
│   │   │   └── ...
│   │   ├── conf.pb.go
│   │   ├── conf.proto
│   │   └── migrate.go   // embed sql文件
│   ├── data             // 业务数据访问, 包含 cache、db 等封装, 实现了 biz 的 repo 接口. 我们可能会把 data 与 dao 混淆在一起, data 偏重业务的含义, 它所要做的是将领域对象重新拿出来, 我们去掉了 DDD 的 infra层. 
│   │   ├── cache.go     // cache层, 防缓存击穿/缓存穿透/缓存雪崩
│   │   ├── client.go    // 各个微服务client初始化
│   │   ├── ctx.go
│   │   ├── data.go
│   │   ├── greeter.go
│   │   └── tracer.go    // 链路追踪tracer初始化
│   ├── pkg              // 自定义扩展包
│   │   ├── idempotent   // 接口幂等性
│   │   └── task         // 异步任务, 内部调用asynq
│   ├── server           // http和grpc实例的创建和配置
│   │   ├── middleware   // 自定义中间件
│   │   │   └── locales  // i18n多语言map配置文件
│   │   ├── grpc.go
│   │   ├── http.go
│   │   └── server.go
│   └── service          // 实现了 api 定义的服务层, 类似 DDD 的 application 层, 处理 DTO 到 biz 领域实体的转换(DTO -> DO), 同时协同各类 biz 交互, 但是不应处理复杂逻辑
│       ├── greeter.go
│       └── service.go
├── third_party          // api依赖的第三方proto, 编译proto文件需要用到
│   ├── errors
│   ├── google
│   ├── openapi
│   │── validate
│   └── ...              //  其他自定义依赖
├─ .gitignore
├─ .gitmodules           // submodule配置文件
├─ Dockerfile
├─ LICENSE
├─ Makefile
└─ README.md
```


# 创建服务

## 环境准备


启动项目前, 我默认你已准备好:
- [go](https://golang.org/dl/)
- 开启go modules
- [mysql](https://www.mysql.com)(本地测试建议使用docker-compose搭建)
- [redis](https://redis.io)(本地测试建议使用docker-compose搭建)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)
- [git](https://git-scm.com)
- [kratos cli工具](https://go-kratos.dev/docs/getting-started/usage)
    ```
    go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
    ```


## [Auth服务](https://github.com/go-cinch/auth)


权限认证服务无需再开发, 下载开箱即用

```
git clone https://github.com/go-cinch/auth
# 可以指定tag
# git clone -b v1.0.1 https://github.com/go-cinch/auth
```


## 创建Game服务


```
# 1.通过模板创建项目 -r 指定仓库 -b 指定分支
kratos new game -r https://github.com/go-cinch/layout.git -b v1.0.1

# 2. 进入项目
cd game
git init
# 一般的我们建议以dev作为开发分支
git checkout -b dev

# 3. 初始化submodule
rm -rf api/auth-proto
# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/auth-proto https://github.com/go-cinch/auth-proto.git ./api/auth-proto

rm -rf api/reason-proto
# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/reason-proto https://github.com/go-cinch/reason-proto.git ./api/reason-proto

# 这里用greeter作为示例
rm -rf api/greeter-proto
# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/greeter-proto https://github.com/go-cinch/greeter-proto.git ./api/greeter-proto

# 4. 下载依赖
go mod download

# 5. 编译项目
make all
```


## 启动


```
# 启动auth
cd auth
export AUTH_DATA_DATABASE_DSN=root:root@tcp(127.0.0.1:3306)/auth?parseTime=True
export AUTH_DATA_REDIS_DSN=redis://127.0.0.1:6379/0
kratos run

# 启动game
cd game
export CINCH_DATA_DATABASE_DSN=root:root@tcp(127.0.0.1:3306)/game?parseTime=True
export CINCH_DATA_REDIS_DSN=redis://127.0.0.1:6379/0
export CINCH_CLIENT_AUTH=127.0.0.1:6160
kratos run
```


## 测试访问


auth服务:
```
curl http://127.0.0.1:6060/idempotent
# 输出如下说明服务通了只是没有权限, 出现其他说明配置有误
# {"code":401, "reason":"UNAUTHORIZED", "message":"token is missing", "metadata":{}}
```

game服务:
```
curl http://127.0.0.1:8080/greeter
# 输出如下说明服务通了只是没有权限, 出现其他说明配置有误
# {"code":401, "reason":"UNAUTHORIZED", "message":"token is missing", "metadata":{}}
```

> 至此, 微服务已启动完毕, auth以及game, 接下来可以自定义你的game啦~


# 文档

以下几个较为常用, 当然你也可以按顺序查看[Docs](https://go-cinch.github.io/docs)
- [第一个接口](https://go-cinch.github.io/docs/#/started.1.first-api)
- [配置](https://go-cinch.github.io/docs/#/base.0.config)
- [数据库](https://go-cinch.github.io/docs/#/base.2.db)
- [Make命令](https://go-cinch.github.io/docs/#/base.4.make)
- [错误管理](https://go-cinch.github.io/docs/#/base.5.reason)


[Kratos]: (https://go-kratos.dev/docs/)

