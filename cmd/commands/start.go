package commands

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/controller"
	"github.com/BinacsLee/server/libs/base"
	"github.com/BinacsLee/server/libs/inject"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Command",
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			fmt.Println(*cfg)
			node := initService(logger, cfg)
			//node.OnStart()
			fmt.Println("node = ", node)
			startHttpService(logger, cfg)
			base.Forever()

			return nil
		},
	}
)

func initService(logger log.Logger, cfg *config.Config) service.NodeService {
	nodeSvc := service.NodeServiceImpl{}

	inject.Regist(Inject_Config, cfg)

	inject.Regist(Inject_LOGGER, logger)
	inject.Regist(Inject_Node_LOGGER, logger.With("module", "node"))
	inject.Regist(Inject_Config_LOGGER, logger.With("module", "config"))
	inject.Regist(Inject_Redis_LOGGER, logger.With("module", "redis"))
	inject.Regist(Inject_Mysql_LOGGER, logger.With("module", "mysql"))
	inject.Regist(Inject_Service_LOGGER, logger.With("module", "service"))

	inject.Regist(Inject_Node_Service, &nodeSvc)
	inject.Regist(Inject_Redis_Service, &service.RedisServiceImpl{})
	inject.Regist(Inject_Mysql_Service, &service.MysqlServiceImpl{})
	inject.Regist(Inject_Config_Service, &service.ConfigServiceImpl{})
	inject.Regist(Inject_ServiceHub, controller.Services)

	err := inject.DoInject()
	if err != nil {
		panic(err.Error())
	}
	return nodeSvc
}

func startHttpService(logger log.Logger, cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Recovery())
	controller.SetRouter(r)
	s := &http.Server{
		Addr:           ":" + cfg.GinConfig.HttpPort,
		Handler:        r,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
	logger.Info("http start finish", "port", cfg.GinConfig.HttpPort)
}
