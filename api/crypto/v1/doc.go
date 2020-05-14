package binacs_api_crypto_v1

//generate grpc stub code
//go:generate protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/google/protobuf/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. crypto.proto

//generate grpc gateway reverse proxy
//go:generate protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/google/protobuf/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. crypto.proto

//generate grpc swagger definition
//go:generate protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/google/protobuf/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. crypto.proto
