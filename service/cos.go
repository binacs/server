package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"

	"github.com/binacsgo/log"

	pb "github.com/BinacsLee/server/api/cos"
	"github.com/BinacsLee/server/config"
)

// CosServiceImpl COS service implement
type CosServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	cli          *cos.Client
	timeTemplate string
}

// AfterInject do inject
func (cs *CosServiceImpl) AfterInject() error {
	u, err := url.Parse(cs.Config.CosConfig.BucketURL)
	if err != nil {
		return err
	}
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cs.Config.CosConfig.SecretID,
			SecretKey: cs.Config.CosConfig.SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})
	cs.cli = c
	cs.timeTemplate = "2006-01-02 15:04:05"
	return nil
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
	name := cs.generateFileName(req.GetFileName())
	f := bytes.NewReader(req.GetFileBytes())

	if _, err := cs.cli.Object.Put(ctx, name, f, nil); err != nil {
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
	name := req.GetCosURI()
	resp, err := cs.cli.Object.Get(ctx, name, nil)
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cs.Logger.Info("CosServiceImpl", "CosGet", string(bs))

	return &pb.CosGetResp{
		Data: &pb.CosGetResObj{
			FileBytes: bs,
		},
	}, nil
}