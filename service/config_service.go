package service

import (
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type ConfigService interface {
	Reload() error
}

type ConfigServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"ConfigLogger"`
}

func (cs *ConfigServiceImpl) AfterInject() error {
	return nil
}

func (cs *ConfigServiceImpl) Reload() error {
	return cs.Config.Reload()
}
