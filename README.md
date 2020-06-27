

计划使用go**优雅**地重构个人server

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
├── api	    # 服务定义 此目录下执行make以依据proto文件生成go和swagger文档
│   ├── Makefile
│   ├── ...
│   └── ...
├── cmd	    # 程序入口 使用cobra以及自写 依赖注入* 组件
│   ├── commands
│   └── main.go
├── config	# 配置相关 支持热加载
│   ├── config.go
│   ├── grpc.go
│   ├── web.go
│   ├── mysql.go
│   ├── redis.go
│   └── log.go
├── config.toml	# 配置文件示例
├── libs	# 基础库
│   ├── base
│   ├── errcode
│   ├── inject
│   ├── log
│   ├── mycrypto
│   └── treemap
├── service	# 抽象为服务的各项功能
│   ├── db		# 存储服务
│   ├── grpc	# grpc服务
│   ├── web     # web服务
│   ├── grpc.go
│   ├── web.go
│   └── node.go
├── test	# 测试使用
│   ├── client	# grpc测试client端
│   ├── static
│   └── tls	# TLS证书
├── types	# 数据结构定义 常量定义
│   ├── table	# 数据库表定义
│   ├── grpc.go
│   ├── redis.go
│   └── types.go
└── version # 版本
    └── version.go

```



### 如何熟悉

必备：

1. cobra
2. zap和日志分割
3. gin
4. grpc & grpc-gateway
5. redis & mysql基础 以及对象存储xorm
6. protobuf & swagger

选读：

1. opentracing协议
2. prometheus组件
3. docker & k8s

### * 关于依赖注入

[这里](https://github.com/binacsgo/inject)





# 快速开始

## 1. 一键构建

在当前路径执行：

```shell
docker-compose up
```

借助docker-compose.yml配置快速搭建全套服务环境。

随后在浏览器中访问本机443端口即可。



## 2. 分块构建

分块构建server, redis 与 mysql, 需要注意的是, server构建需在redis与mysql构建完成之后。

### 2.1 redis

```shell
docker pull redis:latest
docker run -itd --name redis-test -p 6379:6379 redis --requirepass "password"

```

如果你需要在docker中直接操作redis, 请执行

```shell
$ docker exec -it redis-test /bin/bash
[redis容器] $ redis-cli
> auth password
```



### 2.2 mysql

```shell
docker pull mysql
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=testdb mysql 
```

在启动server之前, 你需要在mysql中先行建立DB, 使用collation : utf8mb4的配置并新建testdb。需要注意的是，前文 `docker run` 指令中的 `-e MYSQL_DATABASE=testdb` 与下列mysql指令等效：

```mysql
CREATE SCHEMA `testdb` DEFAULT CHARACTER SET utf8mb4 ;
```

如果你需要在docker中直接操作mysql, 请执行

```shell
$ docker exec -it mysql-test /bin/bash
[mysql容器] $ mysql -uroot -ppassword
```

1. linux可视化管理工具: Workbench

```
sudo apt-get install mysql-workbench
```

2. 对于windows macOS: Navicat

此时在host项需配置docker-machine ip而非localhost or 127.0.0.1



### 2.3 server

#### 2.3.1 本地构建

宿主机需要有Go语言环境。

请依次执行：

```shell
make

./start.sh
```

#### 2.3.2 镜像构建

在项目目录下执行：

```shell
docker build -t binacs/server:v1 .

docker run -itd --name server -p 443:443 -p 9500:9500 binacs/server:v1
```

随后在浏览器中访问本机443端口即可。