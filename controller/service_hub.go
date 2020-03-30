package controller

import (
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service"
)

var (
	Services *ServiceHub
)

func init() {
	Services = &ServiceHub{}
}

type ServiceHub struct {
	ConfigService service.ConfigService `inject-name:"ConfigService"`
	RedisService  service.RedisService  `inject-name:"RedisService"`
	MysqlService  service.MysqlService  `inject-name:"MysqlService"`
	ServiceLogger log.Logger            `inject-name:"ServiceLogger"`
}
