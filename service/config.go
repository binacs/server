package service

import (
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

// ConfigService config service
type ConfigService interface {
	Reload() error
}

// ConfigServiceImpl implement of ConfigService
type ConfigServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"ConfigLogger"`
}

// AfterInject do inject
func (cs *ConfigServiceImpl) AfterInject() error {
	return nil
}

// Reload the config
func (cs *ConfigServiceImpl) Reload() error {
	return cs.Config.Reload()
}
