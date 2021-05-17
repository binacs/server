package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/binacsgo/log"
	"github.com/binacsgo/pastebin"

	pb "github.com/BinacsLee/server/api/pastebin"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/types"
	"github.com/BinacsLee/server/types/table"
)

// PastebinServiceImpl Web PasteBin service implement
type PastebinServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	RedisSvc RedisService `inject-name:"RedisService"`
	MysqlSvc MysqlService `inject-name:"MysqlService"`

	TinyURLSvc TinyURLService `inject-name:"TinyURLService"`

	instance pastebin.PasteBin
	prefix   string
}

// AfterInject do inject
func (ps *PastebinServiceImpl) AfterInject() error {
	// select a impl of PasteBin interface
	// choose markdown for now
	ps.instance = new(pastebin.MarkDownImpl)
	ps.prefix = ps.Config.WebConfig.GetDomain() + "/p/"
	return nil
}

// Register the service
func (ps *PastebinServiceImpl) Register(ctx context.Context, gsrv *grpc.Server, gwmux *runtime.ServeMux, opts []grpc.DialOption) error {
	if gsrv != nil {
		pb.RegisterPastebinServer(gsrv, ps)
	}

	if gwmux != nil {
		if err := pb.RegisterPastebinHandlerFromEndpoint(ctx, gwmux, ":"+ps.Config.GRPCConfig.HttpPort, opts); err != nil {
			return fmt.Errorf("RegisterUserHandlerFromEndpoint: %v", err)
		}
	}
	return nil
}

// PastebinSubmit the text to DB
func (ps *PastebinServiceImpl) PastebinSubmit(ctx context.Context, req *pb.PastebinSubmitReq) (*pb.PastebinSubmitResp, error) {
	turl := ps.TinyURLSvc.Encode(req.GetText() + strconv.FormatInt(time.Now().Unix(), 10))

	page := &table.Page{
		Poster:  req.GetAuthor(),
		Syntax:  req.GetSyntax(),
		Content: req.GetText(),
		TinyURL: turl,
	}
	affected, err := ps.MysqlSvc.GetEngineG().InsertOne(page)
	if err != nil || affected == 0 {
		return nil, err
	}

	ps.Logger.Info("PastebinServiceImpl PastebinSubmit Success", "turl", turl)
	return &pb.PastebinSubmitResp{
		Data: &pb.PastebinSubmitResObj{
			Purl: ps.prefix + turl,
		},
	}, nil
}

// URLSearch the origin content from DB by turl
func (ps *PastebinServiceImpl) URLSearch(turl string) (*table.Page, error) {
	var page table.Page
	if exsit, err := ps.MysqlSvc.GetEngineG().Where("tinyurl = ?", turl).Get(&page); err != nil || !exsit {
		return nil, err
	}
	return &page, nil
}

// Parse the content to markdown... etc.
func (ps *PastebinServiceImpl) Parse(content, syntax string) string {
	if syntax != types.PastebinSyntaxRaw {
		res, err := ps.instance.ParseContent(content)
		if err != nil {
			return err.Error()
		}
		return res
	}
	return content
}
