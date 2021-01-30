# ChangeLog Of [binacs.cn](binacs.cn) Server




## Version 2.0.0 [ 2021 - 01 - 31]

**主要变动**

1.  将原有 `service` 全部整理分类，形成：

    `gateway` & `middleware` & `service`

2.  将所有 `service` 进一步抽象、解耦
  
    对外仅暴露 `interface` ，便于 `mock` 和 `unit test`

3.  将所有 `service` 的与 `API` 相关的接口 `proto` 化

    便于使用 `GRPC` 跨语言前后端交互

4.  清除暂时无用的部分代码




**TODO**

- `service mock` & `unit test`

- `server` 划分为两个子程序 便于横向扩展
  1. `gin` : 仅 `router` 和 `static web`
  2. `grpc` : 仅 `API`

- 使用 `typescript` 重写 `js` 模块
    - 独立在 `ui` 目录







## Version 1.0 [ Before 2021 - 01 - 31 ]

项目结构

```
server
├── api	    	# 服务定义 此目录下执行make以依据proto文件生成go和swagger文档
│   ├── Makefile
│   ├── ...
│   └── ...
├── cmd	    	# 程序入口 使用cobra以及自写 依赖注入* 组件
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
│   ├── db	# 存储服务
│   ├── grpc	# grpc服务 包括 middleware 和 grpc_service
│   ├── web     # web服务  包括 middleware 和 web_service
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
└── version 	# 版本
    └── version.go

```