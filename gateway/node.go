package gateway

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
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
	Config     *config.Config    `inject-name:"Config"`
	Logger     log.Logger        `inject-name:"NodeLogger"`
	WebSvc     WebService        `inject-name:"WebService"`
	GRPCSvc    GRPCService       `inject-name:"GRPCService"`
	LogCleaner LogCleanerService `inject-name:"LogCleanerService"`
}

// AfterInject do inject
func (ns *NodeServiceImpl) AfterInject() error {
	return nil
}

// OnStart start all the service
func (ns *NodeServiceImpl) OnStart() error {
	ns.Logger.Info("Node Service Onstart")

	// Start log cleanup service
	if err := ns.LogCleaner.Start(); err != nil {
		ns.Logger.Error("Failed to start log cleanup service", "error", err)
	} else {
		ns.Logger.Info("Log cleanup service started successfully")
	}

	if ns.Config.PerfConfig.HttpPort != config.NoPerf {
		// Go Pprof
		{
			ns.Logger.Info("Perf Pprof", "HttpPort", ns.Config.PerfConfig.HttpPort)
			go func() {
				if err := http.ListenAndServe("0.0.0.0:"+ns.Config.PerfConfig.HttpPort, nil); err != nil {
					ns.Logger.Error("Pprof server failed", "error", err)
				}
			}()
		}
		// Go Trace
		{
			ns.Logger.Info("Perf Trace")
			if f, err := os.Create("trace.out"); err != nil {
				ns.Logger.Error("Perf Trace", "err", err)
			} else {
				if err := trace.Start(f); err != nil {
					ns.Logger.Error("Trace start failed", "error", err)
				} else {
					defer trace.Stop()
				}
			}
		}
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
