package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"

	"github.com/binacsgo/log"

	pb "github.com/binacs/server/api/cos"
	"github.com/binacs/server/config"
	"github.com/binacs/server/types"
)

// CosServiceImpl COS service implement
type CosServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	u *url.URL
	t *cos.AuthorizationTransport

	timeTemplate string
}

// AfterInject do inject
func (cs *CosServiceImpl) AfterInject() error {
	u, err := url.Parse(cs.Config.CosConfig.BucketURL)
	if err != nil {
		return err
	}
	cs.u = u

	cs.t = &cos.AuthorizationTransport{
		SecretID:  cs.Config.CosConfig.SecretID,
		SecretKey: cs.Config.CosConfig.SecretKey,
		Transport: &debug.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   false,
		},
	}

	cs.timeTemplate = "2006-01-02 15:04:05"
	return nil
}

func (cs *CosServiceImpl) makeCosClient() *cos.Client {
	return cos.NewClient(
		&cos.BaseURL{
			BucketURL: cs.u,
		},
		&http.Client{
			Transport: cs.t,
		})
}

func (cs *CosServiceImpl) generateFileName(name string) string {
	fileWithSuffix := path.Base(name)
	suffixOnly := path.Ext(fileWithSuffix)
	fileOnly := strings.TrimSuffix(fileWithSuffix, suffixOnly)
	fileOnly += "_" + time.Now().Format(cs.timeTemplate)
	return fileOnly + suffixOnly
}

func (cs *CosServiceImpl) generateCosURI(name string) string {
	return strings.ReplaceAll(cs.Config.CosConfig.BucketURL+"/"+name, " ", "%20")
}

func (cs *CosServiceImpl) processCosURI(name string) string {
	return strings.TrimPrefix(name, cs.Config.CosConfig.BucketURL+"/")
}

// checkPassKey verifies the bearer token extracted by the gRPC auth
// interceptor (types.AccessTokenContextKey) against the configured
// CosConfig.PassKey. This mirrors the check already enforced on the
// web upload form (gateway/web.go apiCosPut) so both entry points into
// COS storage require the same shared secret.
func (cs *CosServiceImpl) checkPassKey(ctx context.Context) error {
	token, _ := ctx.Value(types.AccessTokenContextKey).(string)
	if token == "" || token != cs.Config.CosConfig.PassKey {
		return status.Error(codes.Unauthenticated, "invalid or missing cos pass key")
	}
	return nil
}

// Register the service
func (cs *CosServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error {
	if gsrv != nil {
		pb.RegisterCosServer(gsrv, cs)
	}

	if gwmux != nil {
		if err := pb.RegisterCosHandlerFromEndpoint(ctx, gwmux, ":"+cs.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

// CosBucketURL return the BucketURL
func (cs *CosServiceImpl) CosBucketURL(ctx context.Context, req *pb.CosBucketURLReq) (*pb.CosBucketURLResp, error) {
	return &pb.CosBucketURLResp{
		Data: &pb.CosBucketURLResObj{
			BucketURL: cs.Config.CosConfig.BucketURL,
		},
	}, nil
}

// CosGet put file to COS storage
func (cs *CosServiceImpl) CosPut(ctx context.Context, req *pb.CosPutReq) (*pb.CosPutResp, error) {
	if err := cs.checkPassKey(ctx); err != nil {
		cs.Logger.Error("CosServiceImpl", "CosPut auth err", err)
		return nil, err
	}

	name := cs.generateFileName(req.GetFileName())
	f := bytes.NewReader(req.GetFileBytes())

	if _, err := cs.makeCosClient().Object.Put(ctx, name, f, nil); err != nil {
		cs.Logger.Error("CosServiceImpl", "CosPut err", err)
		return nil, err
	}

	return &pb.CosPutResp{
		Data: &pb.CosPutResObj{
			CosURI: cs.generateCosURI(name),
		},
	}, nil
}

// CosGet get file from COS storage
// There is no need for server to do this, we can download from URI directly.
func (cs *CosServiceImpl) CosGet(ctx context.Context, req *pb.CosGetReq) (*pb.CosGetResp, error) {
	if err := cs.checkPassKey(ctx); err != nil {
		cs.Logger.Error("CosServiceImpl", "CosGet auth err", err)
		return nil, err
	}

	name := cs.processCosURI(req.GetCosURI())
	resp, err := cs.makeCosClient().Object.Get(ctx, name, nil)
	// resp maybe nil in Tencent COS SDK
	if err != nil || resp == nil {
		cs.Logger.Error("CosServiceImpl", "CosGet err", err)
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		cs.Logger.Error("CosServiceImpl", "CosGet err", err)
		return nil, err
	}

	cs.Logger.Info("CosServiceImpl", "CosGet", string(bs))

	return &pb.CosGetResp{
		Data: &pb.CosGetResObj{
			FileBytes: bs,
		},
	}, nil
}
