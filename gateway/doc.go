package gateway

//generate gateway mock code

//go:generate  mockgen -destination ../mock/gateway/mock_node.go -package mock_gateway -source node.go
//go:generate  mockgen -destination ../mock/gateway/mock_web.go -package mock_gateway -source web.go
//go:generate  mockgen -destination ../mock/gateway/mock_grpc.go -package mock_gateway -source grpc.go
