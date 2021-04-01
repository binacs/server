

计划使用go**优雅**地重构个人server

[![License](https://img.shields.io/github/license/BinacsLee/server?color=%2387CEFA)](https://github.com/BinacsLee/server/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/BinacsLee/server.svg?branch=master)](https://travis-ci.com/BinacsLee/server)
[![Go Report Card](https://goreportcard.com/badge/github.com/BinacsLee/server)](https://goreportcard.com/report/github.com/BinacsLee/server)

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
| Ice                        | 飞冰前端           |



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
│   ├── tinyurl.go
│   ├── trace.go
│   └── user.go
├── scripts        # 部署相关
│   ├── docker-compose
│   │   ├── config.toml
│   │   └── docker-compose.yml
│   └── kubernetes
│       ├── binacs-cn.yml
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
    └── types.go
└── version     # 版本
    └── version.go

```



### 如何熟悉

必备：

1. cobra
2. zap 和日志分割
3. gin
4. grpc & grpc-gateway
5. redis & mysql 基础 以及对象存储 xorm
6. protobuf & swagger

选读：

1. opentracing 协议
2. prometheus 组件
3. docker & k8s

### * 关于依赖注入

[这里](https://github.com/binacsgo/inject)



## Kubernetes 线上部署

Dockerfile 来自本项目。

镜像来自 [Docker Hub](https://hub.docker.com/r/binacslee/binacs-cn) 。

部署配置位于 [deployment-binacs-cn](https://github.com/OpenKikCoc/deployment-binacs-cn) 项目。



## 终端客户端

[BinacsLee/cli](https://github.com/BinacsLee/cli)

