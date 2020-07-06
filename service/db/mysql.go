package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"xorm.io/core"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/types/table"
	"github.com/binacsgo/log"
)

// MysqlService mysql service
type MysqlService interface {
	Sync2() error
	GetEngineG() *xorm.EngineGroup
}

// MysqlServiceImpl inplement of MysqlService
type MysqlServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"MysqlLogger"`
	ZapLogger *zap.Logger    `inject-name:"ZapLogger"`
	EngineG   *xorm.EngineGroup
}

// AfterInject inject
func (ms *MysqlServiceImpl) AfterInject() error {
	var err error
	ms.EngineG, err = NewMysqlCli(ms.Config.MysqlConfig, ms.ZapLogger)
	if err != nil {
		return err
	}
	return nil
}

// NewMysqlCli return a (mysql client) Engine Group
func NewMysqlCli(cfg config.MysqlConfig, logger *zap.Logger) (*xorm.EngineGroup, error) {
	DSN := cfg.GenerateDSN()
	engine, err := xorm.NewEngineGroup("mysql", DSN)
	if err != nil {
		return nil, err
	}
	tableMapper := core.NewPrefixMapper(core.SameMapper{}, "t_")
	engine.SetTableMapper(tableMapper)
	engine.ShowExecTime(true)
	engine.SetLogger(log.NewMysqlLogger(logger))
	if cfg.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		engine.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	return engine, nil
}

// Sync2 sync the db
func (ms *MysqlServiceImpl) Sync2() error {
	if err := ms.EngineG.Master().Sync2(
		new(table.User),
	); err != nil {
		return fmt.Errorf("MysqlService Create table, err: %v\n", err)
	}
	ms.Logger.Info("MysqlService Sync2 Succeed")
	return nil
}

// GetEngineG return Engine Group
func (ms *MysqlServiceImpl) GetEngineG() *xorm.EngineGroup {
	return ms.EngineG
}

// 通过group可以实现对读写分离的支持
