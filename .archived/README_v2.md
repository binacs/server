

计划使用 Go **优雅**地重构个人 server

[![License](https://img.shields.io/github/license/binacs/server?color=%2387CEFA)](https://github.com/binacs/server/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/binacs/server.svg?branch=master)](https://travis-ci.com/binacs/server)
[![Go Report Card](https://goreportcard.com/badge/github.com/binacs/server)](https://goreportcard.com/report/github.com/binacs/server)
[![Test Coverage](https://codecov.io/gh/binacs/server/branch/master/graph/badge.svg)](https://codecov.io/github/binacs/server?branch=master)

### tools

| 工具族                     | 功能               |
| -------------------------- | ------------------ |
| spf13 cobra                | 解析命令行启动参数 |
| zap lumberjack             | 日志功能系统       |
| gin                        | web服务框架        |
| grpc grpc-gateway          | grpc服务框架       |
| redis mysql(xorm)          | 分布式存储         |
| proto                      | 服务消息结构       |
| swagger                    | RESTful接口        |
| opentracing jaeger         | 分布式追踪         |
| prometheus 全家桶+自研组件 | 监控系统           |
| Dockfile Deployment        | 容器化             |
| Materialize CSS            | CSS样式            |



### 项目结构

```
server
├── api            # 服务定义 此目录下执行make以依据proto文件生成go和swagger文档
│   ├── Makefile
│   ├── ...
│   └── ...
├── cmd            # 程序入口 使用cobra以及自写 依赖注入* 组件
│   ├── commands
│   └── main.go
├── config         # 配置相关 支持热加载
│   ├── config.go
│   ├── grpc.go
│   ├── log.go
│   ├── web.go
│   ├── mysql.go
│   ├── redis.go
│   ├── trace.go
│   └── web.go
├── config.toml    # 配置文件示例
├── gateway        # 路由
│   ├── grpc.go
│   ├── web.go
│   └── node.go
├── middleware     # 中间件
│   ├── auth.go
│   ├── tls.go
│   └── trace.go
├── service        # 服务
│   ├── basic.go
│   ├── const.go
│   ├── crypto.go
│   ├── interface.go
│   ├── mysql.go
│   ├── pastebin.go
│   ├── redis.go
│   ├── storage.go
│   ├── tinyurl.go
│   ├── trace.go
│   └── user.go
├── scripts        # 部署相关
│   ├── docker-compose
│   │   ├── config.toml
│   │   └── docker-compose.yml
│   └── kubernetes
│       ├── binacs-space.yml
│       └── config.toml
├── test           # 测试
│   ├── client     # grpc 测试 client 端
│   ├── template
│   └── tls        # TLS证书
├── types          # 数据结构定义 常量定义
│   ├── table      # 数据库表定义
│   ├── grpc.go
│   ├── redis.go
│   ├── regexp.go
│   └── types.go
└── version     # 版本
    └── version.go

```

### * Injection

[binacsgo/inject](https://github.com/binacsgo/inject)

## Terminal Client

[binacs/cli](https://github.com/binacs/cli)

