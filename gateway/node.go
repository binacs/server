package gateway

import (
	"sync"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
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
	ns.Logger.Info("Node Service Onstart")

	// TODO catch error by channel
	var waiter sync.WaitGroup
	if ns.Config.Mode == config.ALL || ns.Config.Mode == config.WEB {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) error {
			defer wg.Done()
			if err := ns.WebSvc.Serve(); err != nil {
				ns.Logger.Error("Web Serve Done", "err", err)
				return err
			}
			ns.Logger.Info("Node Service WebSvc")
			return nil
		}(&waiter)
	}
	if ns.Config.Mode == config.ALL || ns.Config.Mode == config.GRPC {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) error {
			defer wg.Done()
			if err := ns.GRPCSvc.Serve(); err != nil {
				ns.Logger.Error("GRPC Serve Done", "err", err)
				return err
			}
			ns.Logger.Info("Node Service GRPCSvc")
			return nil
		}(&waiter)
	}

	waiter.Wait()

	ns.Logger.Info("Node Service End")
	return nil
}
