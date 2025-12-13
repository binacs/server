package api_user

//generate go code
//go:generate protoc -I. -I$$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=. user.proto

//generate grpc stub code
//go:generate protoc -I. -I$$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go-grpc_out=require_unimplemented_servers=false:. user.proto

//generate grpc gateway reverse proxy (v2)
//go:generate protoc -I. -I$$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. user.proto

//generate openapi definition (v2)
//go:generate protoc -I. -I$$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --openapiv2_out=logtostderr=true:. user.proto
