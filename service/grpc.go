package service

import (
	//"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	//grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	//grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
	//grpc_service "github.com/BinacsLee/server/service/grpc/service"
)

type GRPCService interface {
	Serve() error
}

type GRPCServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"GRPCLogger"`
	ZapLogger *zap.Logger    `inject-name:"ZapLogger"`

	//XXXSvc *grpc_service.GRPCXXXServiceImpl `inject-name:"GRPCXXXService"`

	tlsCfg *tls.Config
	creds  credentials.TransportCredentials

	gsrv  *grpc.Server
	gwmux *runtime.ServeMux
	srv   *http.Server
}

func (gs *GRPCServiceImpl) AfterInject() error {
	var err error
	gs.tlsCfg, err = tlsConfig(gs.Config.GRPCConfig.CertPath, gs.Config.GRPCConfig.KeyPath)
	if err != nil {
		return err
	}
	gs.creds, err = credentials.NewServerTLSFromFile(gs.Config.GRPCConfig.CertPath, gs.Config.GRPCConfig.KeyPath)
	if err != nil {
		return err
	}

	grpc_zap.ReplaceGrpcLoggerV2(gs.ZapLogger)
	gs.gsrv = grpc.NewServer(grpc.Creds(gs.creds))
	/*
		opts := []grpc_zap.Option{
			//grpc_zap.WithLevels(customFunc),
		}
		gs.gsrv = grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_zap.UnaryServerInterceptor(gs.ZapLogger, opts...),
				grpc_auth.UnaryServerInterceptor(middleware.GlobalAuthFunc()),
			),
			grpc_middleware.WithStreamServerChain(
				grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_zap.StreamServerInterceptor(gs.ZapLogger, opts...),
				grpc_auth.StreamServerInterceptor(middleware.GlobalAuthFunc()),
			),
			grpc.Creds(gs.creds))
	*/
	//ctx := context.Background()
	gs.gwmux = runtime.NewServeMux()

	return nil
}

func (gs *GRPCServiceImpl) Serve() error {
	gs.Logger.Info("GRPCService service register")
	//ctx := context.Background()
	//if err := gs.XXXSvc.Register(ctx, gs.gsrv, gs.gwmux); err != nil {
	//	return err
	//}
	mux := http.NewServeMux()
	mux.Handle("/", gs.gwmux)
	gs.srv = &http.Server{
		Addr:      ":" + gs.Config.GRPCConfig.HttpPort,
		Handler:   HandlerFunc(gs.gsrv, mux),
		TLSConfig: gs.tlsCfg,
	}
	gs.Logger.Info("GRPCService Serve", "HttpPort", gs.Config.GRPCConfig.HttpPort)
	listener, err := net.Listen("tcp", ":"+gs.Config.GRPCConfig.HttpPort)
	if err != nil {
		return err
	}
	//reflection.Register(gs.gsrv)
	if err := gs.srv.Serve(tls.NewListener(listener, gs.tlsCfg)); err != nil {
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

func tlsConfig(certPath, keyPath string) (*tls.Config, error) {
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("Read TLS cert file %s, err: %v", certPath, err)
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("Read TLS key file %s, err: %v", keyPath, err)
	}

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, fmt.Errorf("TLS KeyPair err: %v", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{pair},
		NextProtos:   []string{http2.NextProtoTLS},
	}, nil
}
