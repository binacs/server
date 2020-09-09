package service

import (
	"fmt"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/service/db"
	"github.com/BinacsLee/server/types/table"

	"github.com/binacsgo/log"
	"github.com/binacsgo/pastebin"
)

// WebPastebinServiceImpl Web PasteBin service implement
type WebPastebinServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	RedisSvc db.RedisService `inject-name:"RedisService"`
	MysqlSvc db.MysqlService `inject-name:"MysqlService"`

	instance pastebin.PasteBin
}

// AfterInject do inject
func (ps *WebPastebinServiceImpl) AfterInject() error {
	// select a impl of PasteBin interface
	// choose markdown for now
	ps.instance = new(pastebin.MarkDownImpl)
	return nil
}

// Submit the text to DB
func (ps *WebPastebinServiceImpl) Submit(page *table.Page) error {
	affected, err := ps.MysqlSvc.GetEngineG().InsertOne(page)
	if err != nil || affected == 0 {
		return fmt.Errorf("WebPasteBinServiceImpl Submit | InsertOne Get err: %+v", err)
	}
	return nil
}

// URLSearch the origin content from DB by turl
func (ps *WebPastebinServiceImpl) URLSearch(turl string) (*table.Page, error) {
	var page table.Page
	exsit, err := ps.MysqlSvc.GetEngineG().Where("tinyurl = ?", turl).Get(&page)
	if err != nil {
		return nil, fmt.Errorf("WebPasteBinServiceImpl URLSearch | Mysql Get err: %+v", err)
	} else if !exsit {
		return nil, fmt.Errorf("WebPasteBinServiceImpl URLSearch | tinyurl not exist")
	}
	// TODO 返回到渲染模版
	return &page, nil
}

// Parse the content to markdown... etc.
func (ps *WebPastebinServiceImpl) Parse(content, syntax string) string {
	// TODO syntax
	res, err := ps.instance.ParseContent(content)
	if err != nil {
		return err.Error()
	}
	return res
}
