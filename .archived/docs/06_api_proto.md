# proto



## 1. 概述

略



## 2. 重要概念

message & service

用以支持 grpc-gateway 的 google/api/annotations.proto



## 3. 使用示例

### 3.1 准备

```shell
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u google.golang.org/protobuf/protoc-gen-go
```

对于mac：

```shell
brew install protobuf
```

出现：

```
/Users/xxx/go/src/github.com/protocolbuffers/protobuf/src: warning: directory does not exist.

/Users/xxx/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis: warning: directory does not exist.
```

git clone https://github.com/protocolbuffers/protobuf.git
git clone https://github.com/grpc-ecosystem/grpc-gateway.git

手动下载即可。

### 3.2 使用

编写 proto 服务和消息定义文件，同目录下编写注释方便使用  `go generate` 。

*详情请参考 `api` 目录下的 Makefile 以及 doc.go 文件。*



## 4. 关于 go generate

使用以下注释快速生成支持 grpc、grpc-gateway、swagger 所需文件。

```go
package package_name

//generate grpc stub code
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go-grpc_out=. file-name.proto

//generate grpc gateway reverse proxy
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. service.proto

//generate grpc swagger definition
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. service.proto

```

