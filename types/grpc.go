package types

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	GrpcMatadataResult_Key = "debug"

	GrpcMsgSize = 1024 * 1024 * 1024
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
	if err := grpc.SendHeader(ctx, metadata.Pairs(GrpcMatadataResult_Key, string(data))); err != nil {
		// Log error if needed, but continue execution
		// This is a debug function, so we don't want to fail the main flow
		_ = err // Suppress unused variable warning
	}
}
