package middleware

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/BinacsLee/server/types"
)

func GlobalAuthFunc() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		return Auth(ctx)
	}
}

func Auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, types.TokenType_Bearer)
	if err != nil {
		return ctx, err
	}
	newCtx := context.WithValue(ctx, types.TokenContextKey, token)
	return newCtx, nil
}
