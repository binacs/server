package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/service"
	"github.com/BinacsLee/server/types"
)

// GRPCService interface
type GRPCService interface {
	Serve() error
}

// GRPCServiceImpl inplement of GRPCService
type GRPCServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"GRPCLogger"`
	ZapLogger *zap.Logger    `inject-name:"ZapLogger"`

	RedisSvc service.RedisService `inject-name:"RedisService"`
	MysqlSvc service.MysqlService `inject-name:"MysqlService"`

	UserSvc     service.UserService     `inject-name:"UserService"`
	CryptoSvc   service.CryptoService   `inject-name:"CryptoService"`
	TinyURLSvc  service.TinyURLService  `inject-name:"TinyURLService"`
	PastebinSvc service.PastebinService `inject-name:"PastebinService"`
	CosSvc      service.CosService      `inject-name:"CosService"`

	tlsCfg *tls.Config
	creds  credentials.TransportCredentials

	gsrv  *grpc.Server
	gwmux *runtime.ServeMux
	srv   *http.Server
}

// AfterInject inject
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
	opts := []grpc_zap.Option{
		// grpc_zap.WithLevels(customFunc),
	}
	gs.gsrv = grpc.NewServer(
		grpc.Creds(gs.creds),
		grpc.MaxRecvMsgSize(types.GrpcMsgSize),
		grpc.MaxSendMsgSize(types.GrpcMsgSize),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(gs.ZapLogger, opts...),
			grpc_auth.UnaryServerInterceptor(gs.auth),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(gs.ZapLogger, opts...),
			grpc_auth.StreamServerInterceptor(gs.auth),
		),
	)
	gs.gwmux = runtime.NewServeMux()

	return nil
}

// Serve start grpc serve
func (gs *GRPCServiceImpl) Serve() error {
	gs.Logger.Info("GRPCService service register")

	creds, err := credentials.NewClientTLSFromFile(gs.Config.GRPCConfig.CertPath, gs.Config.GRPCConfig.Host)
	if err != nil {
		return fmt.Errorf("NewClientTLSFromFile: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	ctx := context.Background()
	if err := gs.UserSvc.Register(ctx, gs.gsrv, gs.gwmux, opts); err != nil {
		return err
	}
	if err := gs.CryptoSvc.Register(ctx, gs.gsrv, gs.gwmux, opts); err != nil {
		return err
	}
	if err := gs.TinyURLSvc.Register(ctx, gs.gsrv, gs.gwmux, opts); err != nil {
		return err
	}
	if err := gs.PastebinSvc.Register(ctx, gs.gsrv, gs.gwmux, opts); err != nil {
		return err
	}
	if err := gs.CosSvc.Register(ctx, gs.gsrv, gs.gwmux, opts); err != nil {
		return err
	}

	gs.Logger.Info("ServeMux build")
	mux := http.NewServeMux()
	mux.Handle("/", gs.gwmux)
	gs.srv = &http.Server{
		Addr:      ":" + gs.Config.GRPCConfig.HttpPort,
		Handler:   handlerFunc(gs.gsrv, mux),
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

func (gs *GRPCServiceImpl) auth(ctx context.Context) (context.Context, error) {
	fmt.Println(ctx)
	token, err := grpc_auth.AuthFromMD(ctx, types.TokenType_Bearer)
	if err != nil {
		fmt.Println("err=", err)
		return ctx, err
	}
	fmt.Println("token before check: ", token)
	// check token
	/*
		if types.IsBase64(token) {
			tokenDecBytes, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				return ctx, grpc.Errorf(codes.Unauthenticated, "Request unauthenticated because base64 decode failed")
			}
			token = string(tokenDecBytes)
		}
	*/

	newCtx := context.WithValue(ctx, types.AccessTokenContextKey, token)
	return newCtx, nil
}

func handlerFunc(gsrv *grpc.Server, gwmux http.Handler) http.Handler {
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
