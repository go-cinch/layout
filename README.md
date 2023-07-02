<h1 align="center">Go Cinch Layout</h1>

<div align="center">
Layout是一套轻量级微服务项目模版.
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/go-cinch/layout" alt="Go version"/>
<img src="https://img.shields.io/badge/Kratos-v2.6.2-brightgreen" alt="Kratos version"/>
<img src="https://img.shields.io/badge/MySQL-8.0-brightgreen" alt="MySQL version"/>
<img src="https://img.shields.io/badge/Go--Redis-v9.0.5-brightgreen" alt="Go redis version"/>
<img src="https://img.shields.io/badge/Gorm-v1.25.2-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/badge/Wire-0.5.0-brightgreen" alt="Wire version"/>
<img src="https://img.shields.io/github/license/go-cinch/layout" alt="License"/>
</p>
</div>

# 简介

## 起源

你的单体服务架构是否遇到一些问题, 不能满足业务需求? 那么微服务会是好的解决方案.  
Cinch是一套轻量级微服务脚手架, 基于[Kratos], 节省基础服务搭建时间, 快速投入业务开发.  
我们参考了Go的许多微服务架构, 结合实际需求, 最终选择简洁的[Kratos]作为基石(B站架构), 从架构的设计思路以及代码的书写格式和我们非常匹配.
> cinch意为简单的事, 小菜. 希望把复杂的事变得简单, 提升开发效率.

> 若你想深入学习微服务每个组件, 建议直接看Kratos官方文档. 本项目整合一些业务常用组件, 开箱即用, 并不会把每个组件都介绍那么详细.

## 特性

- `Go` - 当前Go版本直接上`v1.20`
- `Proto` - proto协议同时开启gRPC & HTTP支持, 只需开发一次接口, 不用写两套
- `Jwt` - 认证, 用户登入登出一键搞定
- `Action` - 权限, 基于行为的权限校验
- `Redis` - 缓存, 内置防缓存穿透/缓存击穿/缓存雪崩示例
- `Gorm` - 数据库ORM管理框架, 可自行扩展多种数据库类型, 目前使用MySQL, 其他自行扩展
- `Gorm Gen` - Gorm实体自动生成, 无需再写数据库实体映射
- `Tenant` - 基于Gorm的schema层面的多租户(一个租户一个数据库schema)
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

# 当前版本 <img src="https://img.shields.io/badge/Layout-v1.0.3-brightgreen" alt="Current version"/>

使用cinch工具快速[创建Game服务](https://github.com/go-cinch/layout#%E5%88%9B%E5%BB%BAgame%E6%9C%8D%E5%8A%A1)
```bash
go install github.com/go-cinch/cinch/cmd/cinch@latest

cinch new game
```

## 在线演示

[Vue3入口](https://vue3.go-cinch.top/)  
[React入口](https://react.go-cinch.top/)

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
├── cmd                  
│   └── game             // 项目名
│       ├── main.go      // 程序主入口
│       ├── wire.go      // wire依赖注入
│       └── wire_gen.go
├── configs              // 配置文件目录
│   ├── config.yml       // 主配置文件
│   ├── client.yml       // 配置grpc服务client, 如auth
│   ├── gen.yml          // gen gorm或migrate会用到的配置文件
│   └── ...              // 其他自定义配置文件以yml/yaml结尾均可
├── internal             // 内部逻辑代码
│   ├── biz              // 业务逻辑的组装层, 类似 DDD 的 domain 层, data 类似 DDD 的 repo, 而 repo 接口在这里定义, 使用依赖倒置的原则. 
│   │   ├── biz.go
│   │   ├── reason.go    // 定义错误描述
│   │   ├── game.go      // game业务
│   │   └── xxx.go       // 其他业务
│   ├── conf
│   │   ├── conf.pb.go
│   │   └── conf.proto   // 内部使用的config的结构定义, 使用proto格式生成
│   ├── data             // 业务数据访问, 包含 cache、db 等封装, 实现了 biz 的 repo 接口. 我们可能会把 data 与 dao 混淆在一起, data 偏重业务的含义, 它所要做的是将领域对象重新拿出来, 我们去掉了 DDD 的 infra层. 
│   │   ├── model        // gorm gen生成model目录
│   │   ├── query        // gorm gen生成query目录
│   │   ├── cache.go     // cache层, 防缓存击穿/缓存穿透/缓存雪崩
│   │   ├── client.go    // 各个微服务client初始化
│   │   ├── ctx.go
│   │   ├── data.go      // 数据初始化, 如DB/Redis
│   │   ├── game.go
│   │   └── tracer.go    // 链路追踪tracer初始化
│   ├── db
│   │   ├── migrations   // sql迁移文件目录, 每一次数据库变更都放在这里, 参考https://github.com/rubenv/sql-migrate
│   │   │   ├── xxx.sql  // sql文件
│   │   │   └── ...
│   │   └── migrate.go   // embed sql文件
│   ├── pkg              // 自定义扩展包
│   │   ├── idempotent   // 接口幂等性
│   │   ├── task         // 异步任务, 内部调用asynq
│   │   └── xxx          // 其他扩展
│   ├── server           // http和grpc实例的创建和配置
│   │   ├── middleware   // 自定义中间件
│   │   │   ├── locales  // i18n多语言map配置文件
│   │   │   └── xxx.go   // 一些中间件
│   │   ├── grpc.go
│   │   ├── http.go
│   │   └── server.go
│   └── service          // 实现了 api 定义的服务层, 类似 DDD 的 application 层, 处理 DTO 到 biz 领域实体的转换(DTO -> DO), 同时协同各类 biz 交互, 但是不应处理复杂逻辑
│       ├── service.go
│       ├── game.go      // game接口入口
│       └── xxx.go       // 其他接口入口
├── third_party          // api依赖的第三方proto, 编译proto文件需要用到
│   ├── cinch            // cinch公共依赖
│   ├── errors
│   ├── google
│   ├── openapi
│   │── validate
│   └── ...              //  其他自定义依赖
├─ .gitignore
├─ .gitmodules           // submodule配置文件
├─ .golangci.yml         // golangci-lint
├─ Dockerfile
├─ go.mod
├─ go.sum
├─ LICENSE
├─ Makefile
└─ README.md
```

# 环境准备

启动项目前, 我默认你已准备好(部分软件按建议方式安装即可):

- [go](https://golang.org/dl)1.18+(建议使用[g](https://github.com/voidint/g))
  ```bash
  # sudo apt update
  # sudo apt install -y curl
  curl -sSL https://raw.githubusercontent.com/voidint/g/master/install.sh | bash
  source "$HOME/.g/env"
  # g --version
  # g version 1.5.0
  
  g install 1.20
  # go version
  # go version go1.20 linux/amd64
  
  echo "export GOPATH=/home/ubuntu/go" >> ~/.bashrc 
  # 设置go/bin目录到PATH, 若不设置, go安装的一些文件无法识别
  echo "export PATH=$PATH:/home/ubuntu/.g/go/bin:/home/ubuntu/go/bin" >> ~/.bashrc
  source ~/.bashrc
  ```
- 开启go modules
- [mysql](https://www.mysql.com)(本地测试建议使用docker-compose搭建)
- [redis](https://redis.io)(本地测试建议使用docker-compose搭建)
  ```bash
  # 安装docker
  sudo apt-get update
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  sudo apt-key fingerprint 0EBFCD88
  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  sudo apt-get update
  sudo apt-get install -y docker-ce
  # sudo docker -v
  # Docker version 23.0.1, build a5ee5b1

  # 国内加速安装
  sudo apt-get update
  sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
  curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
  sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
  sudo apt-get update
  sudo apt-get install -y docker-ce
  # sudo docker -v
  # Docker version 23.0.1, build a5ee5b1

  # 去除docker sudo
  sudo groupadd docker
  sudo gpasswd -a ${USER} docker
  sudo systemctl restart docker
  sudo chmod a+rw /var/run/docker.sock
  # docker -v
  # Docker version 23.0.1, build a5ee5b1

  # docker-compose
  sudo curl -L https://github.com/docker/compose/releases/download/v2.10.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  # docker-compose -v
  # Docker Compose version v2.10.2

  
  # 简单启动一个单机版mysql和redis
  git clone https://github.com/go-cinch/compose
  cd compose/single
  # 修改默认密码
  source myenv
  docker-compose -f docker-compose.db.yml up -d redis mysql
  # docker ps
  # CONTAINER ID  IMAGE         COMMAND                 CREATED            STATUS         PORTS                                                 NAMES
  # 918328d0aae1  mysql:8.0.19  "docker-entrypoint.s…"  About an hour ago  Up 59 minutes  0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp  mysql
  # 918b2cfcd72e  redis:7.0     "docker-entrypoint.s…"  About an hour ago  Up 59 minutes  0.0.0.0:6379->6379/tcp, :::6379->6379/tcp             redis
  ```
- [protoc](https://github.com/protocolbuffers/protobuf)
  ```bash
  curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protoc-3.20.3-`uname -s`-`uname -m`.zip
  # apt install -y unzip
  sudo unzip protoc-3.20.3-`uname -s`-`uname -m`.zip -d /usr
  # protoc --version
  # libprotoc 3.20.3
  ```
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
  # protoc-gen-go --version
  # protoc-gen-go v1.30.0
  ```
- [git](https://git-scm.com)
  ```bash
  sudo apt update
  sudo apt install -y git
  ```
- cinch cli工具
  ```bash
  go install github.com/go-cinch/cinch/cmd/cinch@latest
  # cinch -v
  cinch version v1.0.0
  ```

# [Auth服务](https://github.com/go-cinch/auth)

权限认证服务无需再开发, 下载开箱即用

```bash
git clone https://github.com/go-cinch/auth
# 可以指定tag
# git clone -b v1.0.3 https://github.com/go-cinch/auth
```

# 创建Game服务

```bash
# 1.通过模板创建项目
cinch new game
# -r 指定仓库 -b 指定分支
# cinch new game -r https://github.com/go-cinch/layout.git -b dev

# 2. 进入项目
cd game
# 创建个人开发分支
git init -b mark/dev
# 如果你的git版本较低
# git init
# git checkout -b mark/dev

# 3. 初始化submodule(不想使用可忽略此步骤)
# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/auth-proto https://github.com/go-cinch/auth-proto.git ./api/auth-proto

# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/reason-proto https://github.com/go-cinch/reason-proto.git ./api/reason-proto

# 这里用game作为示例, 按需修改
# -b指定分支 --name指定submodule名称
git submodule add -b master --name api/game-proto https://github.com/go-cinch/game-proto.git ./api/game-proto

# 删除一个已经存在的submodule
# git submodule deinit api/game-proto
# git rm --cached api/game-proto
# rm -rf .git/modules/api/game-proto
# rm -rf api/game-proto

# 4. 初始化依赖项(需确保已经安装make)
# sudo apt install -y make
make init

# 5. 编译项目
make all
```

# 启动

## 配置文件

```bash
# 修改auth项目配置
cd auth
# 将mysql/redis的配置修改成你本地配置
vim configs/config.yml

# 修改game项目配置
cd game
# 将mysql/redis的配置修改成你本地配置
vim game/configs/config.yml
# 将auth服务host和端口修改成你本地配置
vim game/configs/client.yml

# 启动auth
cd auth
cinch run

# 启动game
cd game
cinch run
```

## 环境变量

```bash
# 启动auth
# 如果你用的是compose/single
export AUTH_DATA_DATABASE_DSN='root:mysqlrootpwd@tcp(127.0.0.1:3306)/auth?parseTime=True'
export AUTH_DATA_REDIS_DSN='redis://:redispwd@127.0.0.1:6379/0'
# 其他按需修改
export AUTH_DATA_DATABASE_DSN='root:root@tcp(127.0.0.1:3306)/auth?parseTime=True'
export AUTH_DATA_REDIS_DSN='redis://127.0.0.1:6379/0'
cd auth
cinch run

# 启动game
# 如果你用的是compose/single
export GAME_DATA_DATABASE_DSN='root:mysqlrootpwd@tcp(127.0.0.1:3306)/game?parseTime=True'
export GAME_DATA_REDIS_DSN='redis://:redispwd@127.0.0.1:6379/0'
# 其他按需修改
export GAME_DATA_DATABASE_DSN='root:root@tcp(127.0.0.1:3306)/game?parseTime=True'
export GAME_DATA_REDIS_DSN='redis://127.0.0.1:6379/0'
# 设置auth服务
export GAME_CLIENT_AUTH='127.0.0.1:6160'
cd game
cinch run
```

> Tip: 环境变量前缀可在cmd/xxx/main.go中修改,
参见[环境变量前缀](/base/0.config?id=%e7%8e%af%e5%a2%83%e5%8f%98%e9%87%8f%e5%89%8d%e7%bc%80)

## 测试访问

auth服务:

```bash
curl http://127.0.0.1:6060/idempotent
# 输出如下说明服务通了只是没有权限, 出现其他说明配置有误
# {"code":401, "reason":"UNAUTHORIZED", "message":"token is missing", "metadata":{}}
```

game服务:

```bash
curl http://127.0.0.1:8080/hello/world
# 输出如下说明服务通了只是没有权限, 出现其他说明配置有误
# {"code":401, "reason":"UNAUTHORIZED", "message":"token is missing", "metadata":{}}
```

> 至此, 微服务已启动完毕, auth以及game, 接下来可以自定义你的game啦~

# 文档

以下几个较为常用, 当然你也可以按顺序查看[Docs](https://go-cinch.github.io/docs)

- [第一个CRUD](https://go-cinch.github.io/docs/#/started/1.first-api)
- [配置](https://go-cinch.github.io/docs/#/base/0.config)
- [数据库](https://go-cinch.github.io/docs/#/base/2.db)
- [Make命令](https://go-cinch.github.io/docs/#/base/4.make)
- [错误管理](https://go-cinch.github.io/docs/#/base/5.reason)

[Kratos]: (https://go-kratos.dev/docs/)

