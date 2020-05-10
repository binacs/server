package service

import (
	"sync"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

// NodeService the node service
type NodeService interface {
	OnStart() error
}

// NodeServiceImpl the implement of node service
type NodeServiceImpl struct {
	Config  *config.Config `inject-name:"Config"`
	Logger  log.Logger     `inject-name:"NodeLogger"`
	WebSvc  WebService     `inject-name:"WebService"`
	GRPCSvc GRPCService    `inject-name:"GRPCService"`
}

// AfterInject do inject
func (ns *NodeServiceImpl) AfterInject() error {
	return nil
}

// OnStart start all the service
func (ns *NodeServiceImpl) OnStart() error {
	var waiter sync.WaitGroup

	if ns.Config.Mode == "all" || ns.Config.Mode == "web" {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) error {
			defer wg.Done()
			err := ns.WebSvc.Serve()
			if err != nil {
				ns.Logger.Error("Web Serve Done", "err", err)
				return err
			}
			return nil
		}(&waiter)
	}
	if ns.Config.Mode == "all" || ns.Config.Mode == "grpc" {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) error {
			defer wg.Done()
			err := ns.GRPCSvc.Serve()
			if err != nil {
				ns.Logger.Error("GRPC Serve Done", "err", err)
				return err
			}
			return nil
		}(&waiter)
	}

	waiter.Wait()

	return nil
}
