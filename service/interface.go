package service

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/binacsgo/trace"

	pb "github.com/BinacsLee/server/api/cos"
	crypto "github.com/BinacsLee/server/api/crypto"
	pastebin "github.com/BinacsLee/server/api/pastebin"
	tinyurl "github.com/BinacsLee/server/api/tinyurl"
	user "github.com/BinacsLee/server/api/user"
	"github.com/BinacsLee/server/types/table"
)

// ----------------- DB service -----------------

// MysqlService mysql service
type MysqlService interface {
	Sync2() error
	GetEngineG() *xorm.EngineGroup
}

// RedisService redis service
type RedisService interface {
	Ping() error
	Set(string, string, time.Duration) error
	Get(string) (string, error)
	Del(string) error
	GetExpireAt(string) (time.Time, error)
}

// ----------------- API service -----------------

// CryptoService crypto service which encrypt and decrypt text
type CryptoService interface {
	Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error
	CryptoEncrypt(ctx context.Context, req *crypto.CryptoEncryptReq) (*crypto.CryptoEncryptResp, error)
	CryptoDecrypt(ctx context.Context, req *crypto.CryptoDecryptReq) (*crypto.CryptoDecryptResp, error)
}

// PastebinService pastebin service
type PastebinService interface {
	Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error
	PastebinSubmit(ctx context.Context, req *pastebin.PastebinSubmitReq) (*pastebin.PastebinSubmitResp, error)
	URLSearch(turl string) (*table.Page, error)
	Parse(content, syntax string) string
}

// TinyURLService tinyurl service
type TinyURLService interface {
	Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error
	TinyURLEncode(ctx context.Context, req *tinyurl.TinyURLEncodeReq) (*tinyurl.TinyURLEncodeResp, error)
	TinyURLDecode(ctx context.Context, req *tinyurl.TinyURLDecodeReq) (*tinyurl.TinyURLDecodeResp, error)
	Encode(url string) string
	URLSearch(turl string) (string, error)
}

// CosService tinyurl service
type CosService interface {
	Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error
	CosBucketURL(ctx context.Context, req *pb.CosBucketURLReq) (*pb.CosBucketURLResp, error)
	CosPut(ctx context.Context, req *pb.CosPutReq) (*pb.CosPutResp, error)
	CosGet(ctx context.Context, req *pb.CosGetReq) (*pb.CosGetResp, error)
}

// UserService user service
type UserService interface {
	Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error
	UserTest(ctx context.Context, req *user.UserTestReq) (*user.UserTestResp, error)
	UserRegister(ctx context.Context, req *user.UserRegisterReq) (*user.UserRegisterResp, error)
	UserAuth(ctx context.Context, req *user.UserAuthReq) (*user.UserAuthResp, error)
	UserRefresh(ctx context.Context, req *user.UserRefreshReq) (*user.UserRefreshResp, error)
	UserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoResp, error)
}

// ----------------- other service -----------------

// TraceService trace service
type TraceService interface {
	StartSpan(operationName string) opentracing.Span
	Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error
	Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error)
	Close() error

	GetTracer() trace.Trace
	FromGinContext(c *gin.Context, serviceName string) opentracing.Span
}

// BasicService base service which retrun the static html
type BasicService interface {
	ServeHome() string
	ServeToys() string
	ServeCrypto() string
	ServeTinyURL() string
	ServePastebin() string
	ServeStorage() string
	ServeAbout() string
}
