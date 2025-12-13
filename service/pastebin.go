package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/binacsgo/log"
	"github.com/binacsgo/pastebin"

	pb "github.com/binacs/server/api/pastebin"
	"github.com/binacs/server/config"
	"github.com/binacs/server/types"
	"github.com/binacs/server/types/table"
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

// hashPassword hashes the password using SHA256
func (ps *PastebinServiceImpl) hashPassword(password string) string {
	if password == "" {
		return ""
	}
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// verifyPassword verifies if the provided password matches the stored hash
func (ps *PastebinServiceImpl) verifyPassword(password, hash string) bool {
	if hash == "" {
		return true // No password set, allow access
	}
	if password == "" {
		return false // Password required but not provided
	}
	computedHash := ps.hashPassword(password)
	return computedHash == hash
}

// PastebinSubmit the text to DB
func (ps *PastebinServiceImpl) PastebinSubmit(ctx context.Context, req *pb.PastebinSubmitReq) (*pb.PastebinSubmitResp, error) {
	turl := ps.TinyURLSvc.Encode(req.GetText() + strconv.FormatInt(time.Now().Unix(), 10))

	// Hash password if provided
	passwordHash := ps.hashPassword(req.GetPassword())

	page := &table.Page{
		Poster:   req.GetAuthor(),
		Syntax:   req.GetSyntax(),
		Content:  req.GetText(),
		TinyURL:  turl,
		Password: passwordHash,
	}
	affected, err := ps.MysqlSvc.GetEngineG().InsertOne(page)
	if err != nil || affected == 0 {
		return nil, err
	}

	ps.Logger.Info("PastebinServiceImpl PastebinSubmit Success", "turl", turl, "hasPassword", passwordHash != "")
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

// VerifyPassword verifies if the provided password is correct for the page
func (ps *PastebinServiceImpl) VerifyPassword(page *table.Page, password string) bool {
	return ps.verifyPassword(password, page.Password)
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

// RecentPosts show the recent posts
func (ps *PastebinServiceImpl) RecentPosts() ([]table.Page, error) {
	counts, _ := ps.MysqlSvc.GetEngineG().Count(&table.Page{})
	pages := make([]table.Page, counts)
	if err := ps.MysqlSvc.GetEngineG().Find(&pages); err != nil {
		return nil, err
	}
	return pages, nil
}
