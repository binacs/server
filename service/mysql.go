package service

import (
	"fmt"
	"time"

	// _ to run `init()`
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"xorm.io/core"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/types/table"
)

func newMysqlCli(cfg config.MysqlConfig, logger *zap.Logger) (*xorm.EngineGroup, error) {
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

// ----------------------------------------------------------------------

// MysqlServiceImpl inplement of MysqlService
type MysqlServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"MysqlLogger"`
	ZapLogger *zap.Logger    `inject-name:"ZapLogger"`
	EngineG   *xorm.EngineGroup
}

// AfterInject inject
func (ms *MysqlServiceImpl) AfterInject() error {
	// Ignore the Sync2 error in AfterInject
	_ = ms.buildClient()
	go ms.checkLoop()
	return nil
}

func (ms *MysqlServiceImpl) buildClient() (err error) {
	if ms.EngineG, err = newMysqlCli(ms.Config.MysqlConfig, ms.ZapLogger); err != nil {
		return err
	} else if err = ms.Sync2(); err != nil {
		return err
	}
	return nil
}

func (ms *MysqlServiceImpl) checkLoop() {
	timer := time.NewTimer(dbCheckInterval)
	defer timer.Stop()
	for {
		timer.Reset(dbCheckInterval)
		select {
		case <-timer.C:
			if err := ms.EngineG.Ping(); err != nil {
				ms.Logger.Error("MysqlServiceImpl checkLoop Ping", "err", err)
				if err := ms.buildClient(); err != nil {
					ms.Logger.Error("MysqlServiceImpl checkLoop buildClient", "err", err)
				} else {
					ms.Logger.Error("MysqlServiceImpl checkLoop buildClient success")
				}
			} else {
				ms.Logger.Info("MysqlServiceImpl checkLoop Ping success")
			}
			if err := ms.Sync2(); err != nil {
				ms.Logger.Error("MysqlServiceImpl checkLoop Sync2", "err", err)
			} else {
				ms.Logger.Error("MysqlServiceImpl checkLoop Sync2 success")
			}
		}
	}
}

// Sync2 sync the db
func (ms *MysqlServiceImpl) Sync2() error {
	if err := ms.EngineG.Master().Sync2(
		new(table.User),
		new(table.Page),
	); err != nil {
		return fmt.Errorf("MysqlServiceImpl Create table, err: %v", err)
	}
	ms.Logger.Info("MysqlService Sync2 Succeed")
	return nil
}

// GetEngineG return Engine Group
func (ms *MysqlServiceImpl) GetEngineG() *xorm.EngineGroup {
	return ms.EngineG
}

// 通过group可以实现对读写分离的支持
