package gateway

import (
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/binacsgo/log"

	"github.com/binacs/server/config"
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

	if ns.Config.PerfConfig.HttpPort != config.NoPerf {
		// Go Pprof. Execution traces are available on demand and time-bounded
		// via net/http/pprof at /debug/pprof/trace?seconds=N. We intentionally do
		// NOT call runtime/trace.Start to a file here: that captures the whole
		// program for the entire process lifetime, growing trace.out without bound
		// (it once filled the disk and stalled MySQL writes).
		ns.Logger.Info("Perf Pprof", "HttpPort", ns.Config.PerfConfig.HttpPort)
		go func() {
			if err := http.ListenAndServe("0.0.0.0:"+ns.Config.PerfConfig.HttpPort, nil); err != nil {
				ns.Logger.Error("Pprof server failed", "error", err)
			}
		}()
	}

	// TODO catch error by channel
	var waiter sync.WaitGroup
	if ns.Config.Mode == config.ALL || ns.Config.Mode == config.WEB {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if err := ns.WebSvc.Serve(); err != nil {
				ns.Logger.Error("Web Serve Done", "err", err)
			} else {
				ns.Logger.Info("Node Service WebSvc")
			}
		}(&waiter)
	}
	if ns.Config.Mode == config.ALL || ns.Config.Mode == config.GRPC {
		waiter.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if err := ns.GRPCSvc.Serve(); err != nil {
				ns.Logger.Error("GRPC Serve Done", "err", err)
			} else {
				ns.Logger.Info("Node Service GRPCSvc")
			}
		}(&waiter)
	}

	waiter.Wait()

	ns.Logger.Info("Node Service End")
	return nil
}
