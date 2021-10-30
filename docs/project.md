## 项目结构

目录结构如下:

```
/
├── api  // 下面维护了大仓中微服务使用的proto文件以及根据它们所生成的go文件
│    └── user
│        └── service
│            └── v1
│                ├── error_reason_errors.pb.go
│                ├── error_reason.pb.go
│                ├── error_reason.proto
│                ├── error_reason.swagger.json
│                ├── user_grpc.pb.go
│                ├── user_http.pb.go
│                ├── user.pb.go
│                ├── user.proto
│                └── user.swagger.json
├── app  // 大仓中所有微服务的集合
│    ├── logic  // 逻辑服务
│    │    └── service  // 应用类型[interface、service、job、admin]
│    │        ├── bin
│    │        │    └── logic  // 项目编译后的启动文件
│    │        ├── cmd  // 项目启动的入口文件
│    │        │    └── logic
│    │        │        ├── main.go
│    │        │        ├── wire_gen.go
│    │        │        └── wire.go  // 我们使用wire来维护依赖注入
│    │        ├── configs // 这里通常维护一些本地调试用的样例配置文件
│    │        │    └── config.yaml
│    │        ├── Dockerfile
│    │        ├── generate.go
│    │        ├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│    │        │    ├── biz  // 业务逻辑的组装层  repo 接口在这里定义，使用依赖倒置的原则
│    │        │    │    ├── biz.go
│    │        │    │    ├── greeter.go
│    │        │    │    └── README.md
│    │        │    ├── conf  // 内部使用的config的结构定义，使用proto格式生成
│    │        │    │    ├── conf.pb.go
│    │        │    │    └── conf.proto
│    │        │    ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口
│    │        │    │    ├── data.go
│    │        │    │    ├── greeter.go
│    │        │    │    └── README.md
│    │        │    ├── server // http和grpc实例的创建和配置
│    │        │    │    ├── grpc.go
│    │        │    │    ├── http.go
│    │        │    │    └── server.go
│    │        │    └── service // 实现了 api 定义的服务层，
│    │        │        ├── greeter.go
│    │        │        ├── README.md
│    │        │        └── service.go
│    │        ├── LICENSE
│    │        └── Makefile
│    └── user  //用户服务
│        └── service
│            ├── bin
│            │    └── suim
│            ├── cmd
│            │    └── suim
│            │        ├── main.go
│            │        ├── wire_gen.go
│            │        └── wire.go
│            ├── configs
│            │    └── config.yaml
│            ├── Dockerfile
│            ├── generate.go
│            ├── internal
│            │    ├── biz
│            │    │    ├── biz.go
│            │    │    ├── README.md
│            │    │    └── user.go
│            │    ├── conf
│            │    │    ├── conf.pb.go
│            │    │    └── conf.proto
│            │    ├── data
│            │    │    ├── data.go
│            │    │    ├── README.md
│            │    │    └── user.go
│            │    ├── server
│            │    │    ├── grpc.go
│            │    │    ├── http.go
│            │    │    └── server.go
│            │    └── service
│            │        ├── README.md
│            │        ├── service.go
│            │        └── user.go
│            ├── LICENSE
│            └── Makefile
├── app_makefile  // 大仓的主要 makefile
├── deploy // 使用docker-compose 快速搭建测试环境,亦可线上部署
│    ├── docker-compose.yaml
│    ├── etcd
│    ├── mariadb
│    │    ├── conf // 配置文件,映射到docker内部
│    │    │    ├── conf.d
│    │    │    │    └── docker.cnf
│    │    │    ├── debian.cnf
│    │    │    ├── debian-start
│    │    │    ├── mariadb.cnf
│    │    │    ├── mariadb.conf.d
│    │    │    │    ├── 50-client.cnf
│    │    │    │    ├── 50-mysql-clients.cnf
│    │    │    │    ├── 50-mysqld_safe.cnf
│    │    │    │    ├── 50-server.cnf
│    │    │    │    ├── 60-galera.cnf
│    │    │    │    └── 99-enable-encryption.cnf.preset
│    │    │    │        └── enable_encryption.preset
│    │    │    └── my.cnf -> /etc/alternatives/my.cnf
│    │    ├── data  // mariadb 数据文件
│    │    └── log
│    ├── nats
│    └── redis
│        ├── conf
│        │    └── redis.conf
│        └── data // redis数据文件
├── docs //项目文档
│    ├── im
│    │    ├── im_proto.proto
│    │    └── proto.md
│    └── project.md
├── go.mod
├── go.sum
├── LICENSE
├── MakeFile
├── README.md
└── third_party // api 依赖的第三方proto
    ├── errors
    │    └── errors.proto
    ├── google
    │    └── api
    │        ├── annotations.proto
    │        ├── httpbody.proto
    │        └── http.proto
    ├── protoc-gen-openapiv2
    │    └── options
    │        ├── annotations.proto
    │        └── openapiv2.proto
    ├── README.md
    └── validate
        ├── README.md
        └── validate.proto
```