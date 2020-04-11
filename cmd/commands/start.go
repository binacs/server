package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/base"
	"github.com/BinacsLee/server/libs/inject"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service"
	"github.com/BinacsLee/server/service/db"
	grpc_service "github.com/BinacsLee/server/service/grpc/service"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Command",
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			fmt.Println(*cfg)
			node := initService(logger, cfg)
			fmt.Println("node = ", node)
			if err = node.OnStart(); err != nil {
				fmt.Println(err)
			}
			base.Forever()

			return nil
		},
	}
)

func initService(logger log.Logger, cfg *config.Config) *service.NodeServiceImpl {
	nodeSvc := service.NodeServiceImpl{}

	inject.Regist(Inject_Config, cfg)

	inject.Regist(Inject_ZAPLOGGER, zaplogger)
	inject.Regist(Inject_LOGGER, logger)
	inject.Regist(Inject_Node_LOGGER, logger.With("module", "node"))
	inject.Regist(Inject_Web_LOGGER, logger.With("module", "web"))
	inject.Regist(Inject_GRPC_LOGGER, logger.With("module", "grpc"))
	inject.Regist(Inject_Config_LOGGER, logger.With("module", "config"))
	inject.Regist(Inject_Redis_LOGGER, logger.With("module", "redis"))
	inject.Regist(Inject_Mysql_LOGGER, logger.With("module", "mysql"))
	//inject.Regist(Inject_Service_LOGGER, logger.With("module", "service"))

	inject.Regist(Inject_Node_Service, &nodeSvc)
	inject.Regist(Inject_Web_Service, &service.WebServiceImpl{})
	inject.Regist(Inject_GRPC_Service, &service.GRPCServiceImpl{})
	inject.Regist(Inject_GRPCUser_Service, &grpc_service.GRPCUserServiceImpl{})

	inject.Regist(Inject_Redis_Service, &db.RedisServiceImpl{})
	inject.Regist(Inject_Mysql_Service, &db.MysqlServiceImpl{})
	inject.Regist(Inject_Config_Service, &service.ConfigServiceImpl{})
	//inject.Regist(Inject_ServiceHub, controller.Services)

	err := inject.DoInject()
	if err != nil {
		panic(err.Error())
	}
	return &nodeSvc
}
