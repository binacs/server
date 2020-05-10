package types

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	GrpcMatadataResult_Key = "debug"
)

var (
	grpcMatadataResult_Debug = false
)

// GrpcMatadataDebugOpen set true
func GrpcMatadataDebugOpen() {
	grpcMatadataResult_Debug = true
}

// GrpcMatadataSetResultFail set fail
func GrpcMatadataSetResultFail(ctx context.Context, rsp interface{}) {
	if !grpcMatadataResult_Debug {
		return
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		return
	}
	grpc.SendHeader(ctx, metadata.Pairs(GrpcMatadataResult_Key, string(data)))
}
