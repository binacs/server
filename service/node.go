package service

import (
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

type NodeService interface {
	OnStart() error
}

type NodeServiceImpl struct {
	Config   *config.Config  `inject-name:"Config"`
	Logger   log.Logger      `inject-name:"NodeLogger"`
	WebSvc   WebService      `inject-name:"WebService"`
	GRPCSvc  GRPCService     `inject-name:"GRPCService"`
}

func (ns *NodeServiceImpl) AfterInject() error {
	return nil
}

func (ns *NodeServiceImpl) OnStart() error {
	err := ns.GRPCSvc.Serve()
	if err != nil {
		return err
	}

	return nil
}