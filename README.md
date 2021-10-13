# Suim 即时通讯项目

suim 是一个即时通讯服务器，代码全部使用golang完成,采用 Kratos 搭建大仓,可以很方便的扩充业务。

## 主要功能包括

1.支持tcp，websocket接入  
2.离线消息同步    
3.单用户多设备同时在线    
4.单聊，群聊，频道,以及房间聊天场景  
5.支持服务水平扩展

### 使用技术:
- 数据库:MySQL+Redis  
- 消息队列:Nats  
- 通讯框架:GRPC  
- 项目框架:Kratos  
- 长连接通讯协议:Protocol Buffers
- ORM框架:GORM

## Create a service

```
# 创建模板项目
kratos new server

cd server
# 添加一个 proto 模板
kratos proto add api/server/server.proto
# 生成 proto 代码
kratos proto client api/server/server.proto
# 通过proto文件生成服务的源码
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## 通过Makefile生成其他辅助文件

```bash

Usage:
 make [target]

Targets:
init                   init env
grpc                   generate grpc code
http                   generate http code
errors                 generate errors code
swagger                generate swagger
proto                  generate internal proto struct
generate               generate client code
build                  build
test                   test
wire                   generate wire
api                    generate api proto
all                    generate all
help                   show help
```

例如:生成User 服务的 http grcp swagger proto 代码

```bash
cd app/xxx/service
make http grpc swagger
#or
make api
```

## 自动初始化 (wire)

```bash
# install wire
go install github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## docker-compsoe

启动项目的底层服务,例如MariaDb,Redis,Etcd,Nats

```bash
# run
docker-compose up -d 

# stop
docker-compose stop

# Stop and remove containers, networks
docker-compose down
```

## 应用类型

微服务中的 app 服务类型主要分为4类:interface、service、job、admin，应用 cmd 目录负责程序的:启动、关闭、配置初始化等。

- interface: 对外的 BFF 服务(英文全称:Backend For Frontend)，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口。
- service: 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。
- admin:区别于 service，更多是面向运营测的服务，通常数据权限更高，隔离带来更好的代码级别安全。
- job: 流式任务处理的服务，上游一般依赖 message broker。
- task: 定时任务，类似 cronjob，部署到 task 托管平台中。