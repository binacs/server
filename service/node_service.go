package service

import (
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type NodeService interface {
}

type NodeServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"NodeLogger"`
}

func (ns *NodeServiceImpl) AfterInject() error {
	return nil
}

//func
