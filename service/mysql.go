package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type MysqlService interface {

}

type MysqlServiceImpl struct {
	Config  *config.Config `inject-name:"Config"`
	Logger  log.Logger     `inject-name:"MysqlLogger"`
	EngineG *xorm.EngineGroup
}

func (ms *MysqlServiceImpl) AfterInject() error {
	var err error
	ms.EngineG, err = NewMysqlCli(ms.Config.MysqlConfig, ms.Logger)
	if err != nil {
		return err
	}
	return nil
}

func NewMysqlCli(cfg config.MysqlConfig, logger log.Logger) (*xorm.EngineGroup, error) {
	DSN := cfg.GenerateDSN()
	engine, err := xorm.NewEngineGroup("mysql", DSN)
	if err != nil {
		return nil, err
	}
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	engine.SetTableMapper(tableMapper)
	//engine.ShowExecTime(true)
	//engine.SetLogger(logger)
	if cfg.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		engine.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	return engine, nil
}