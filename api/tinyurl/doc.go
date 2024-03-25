package api_tinyurl

//generate go code
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. tinyurl.proto

//generate grpc stub code
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go-grpc_out=require_unimplemented_servers=false:. tinyurl.proto

//generate grpc gateway reverse proxy
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. tinyurl.proto

//generate grpc swagger definition
//go:generate protoc -I. -I$GOPATH/pkg/mod -I$GOPATH/pkg/mod/github.com/protocolbuffers/protobuf/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. tinyurl.proto
