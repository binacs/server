package service

import (
	"net"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service/grpc/middleware"
)

type GRPCService interface {
	Serve() error
}

type GRPCServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"GRPCLogger"`
	ZapLogger *zap.Logger    `inject-name:"ZapLogger"`

	srv *http.Server
}

func (gs *GRPCServiceImpl) AfterInject() error {
	opts := []grpc_zap.Option{
		//grpc_zap.WithLevels(customFunc),
	}

	grpc_zap.ReplaceGrpcLoggerV2(gs.ZapLogger)
	gsrv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(gs.ZapLogger, opts...),
			grpc_auth.UnaryServerInterceptor(middleware.GlobalAuthFunc()),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(gs.ZapLogger, opts...),
			grpc_auth.StreamServerInterceptor(middleware.GlobalAuthFunc()),
		))
	gwmux := runtime.NewServeMux()
	// service

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	gs.srv = &http.Server{
		Addr:    ":" + gs.Config.GRPCConfig.HttpPort,
		Handler: HandlerFunc(gsrv, mux),
	}
	return nil
}

func (gs *GRPCServiceImpl) Serve() error {
	gs.Logger.Info("GRPCService Serve", "HttpPort", gs.Config.GRPCConfig.HttpPort)
	listener, err := net.Listen("tcp", ":"+gs.Config.GRPCConfig.HttpPort)
	if err != nil {
		return err
	}
	if err := gs.srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func HandlerFunc(gsrv *grpc.Server, gwmux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			gsrv.ServeHTTP(w, r)
		} else {
			gwmux.ServeHTTP(w, r)
		}
	})
}
