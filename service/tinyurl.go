package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/binacsgo/log"
	"github.com/binacsgo/tinyurl"

	pb "github.com/BinacsLee/server/api/tinyurl"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/types"
)

// TinyURLServiceImpl Web TinyURL service implement
type TinyURLServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	RedisSvc RedisService `inject-name:"RedisService"`
	MysqlSvc MysqlService `inject-name:"MysqlService"`

	instance tinyurl.TinyURL
	prefix   string
}

// AfterInject do inject
func (ts *TinyURLServiceImpl) AfterInject() error {
	// Select a impl of TinyURL interface
	// choose md5 for now
	ts.instance = tinyurl.GetMD5Impl()
	ts.prefix = ts.Config.WebConfig.GetDomain() + "/r/"
	return nil
}

// Register the service
func (ts *TinyURLServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error {
	if gsrv != nil {
		pb.RegisterTinyURLServer(gsrv, ts)
	}

	if gwmux != nil {
		if err := pb.RegisterTinyURLHandlerFromEndpoint(ctx, gwmux, ":"+ts.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

// TinyURLEncode return the tiny-url of origin-url
func (ts *TinyURLServiceImpl) TinyURLEncode(ctx context.Context, req *pb.TinyURLEncodeReq) (*pb.TinyURLEncodeResp, error) {
	ts.Logger.Info("TinyURLServiceImpl TinyURLEncode Start", "url", req.GetUrl())

	url := req.GetUrl()
	if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
		ts.Logger.Error("TinyURLServiceImpl TinyURLEncode, url illegal", "text", url)
		return nil, nil
	}

	turl := ts.Encode(url)
	if err := ts.RedisSvc.Set(turl, url, types.TinyURLExpireDuration()); err != nil {
		ts.Logger.Error("TinyURLServiceImpl TinyURLEncode RedisSvc.Set", "error", err, "url", url, "turl", turl)
		return nil, err
	}

	ts.Logger.Info("TinyURLServiceImpl TinyURLEncode Succeed", "turl", turl)
	return &pb.TinyURLEncodeResp{
		Data: &pb.TinyURLEncodeResObj{
			Turl: ts.prefix + turl,
		},
	}, nil
}

// TinyURLDecode return the origin-url of tiny-url
func (ts *TinyURLServiceImpl) TinyURLDecode(ctx context.Context, req *pb.TinyURLDecodeReq) (*pb.TinyURLDecodeResp, error) {
	ts.Logger.Info("TinyURLServiceImpl TinyURLDecode Start", "tiny-url", req.GetTurl())

	uri := req.GetTurl()
	if !strings.HasPrefix(uri, ts.prefix) {
		ts.Logger.Error("TinyURLServiceImpl TinyURLDecode, tiny-url illegal", "uri", uri)
		return nil, nil
	}

	url, err := ts.RedisSvc.Get(strings.TrimPrefix(uri, ts.prefix))
	if err != nil {
		ts.Logger.Error("TinyURLServiceImpl TinyURLDecode RedisSvc.Get", "error", err, "uri", uri)
		return nil, err
	}

	ts.Logger.Info("TinyURLServiceImpl TinyURLDecode Succeed", "url", url)
	return &pb.TinyURLDecodeResp{
		Data: &pb.TinyURLDecodeResObj{
			Url: url,
		},
	}, nil
}

// Encode return the encode string directly
func (ts *TinyURLServiceImpl) Encode(url string) string {
	return ts.instance.EncodeURL(url)
}

// URLSearch return the origin-url of tiny-url suffix
func (ts *TinyURLServiceImpl) URLSearch(turl string) (string, error) {
	ts.Logger.Info("TinyURLServiceImpl URLSearch Start", "tiny-url suffix", turl)
	url, err := ts.RedisSvc.Get(turl)
	if err != nil {
		ts.Logger.Error("TinyURLServiceImpl URLSearch RedisSvc.Get", "error", err, "url", url, "turl", turl)
		return "", err
	}
	return url, nil
}
