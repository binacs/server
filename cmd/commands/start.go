package commands

import (
	"github.com/spf13/cobra"

	"github.com/binacsgo/inject"

	"github.com/BinacsLee/server/gateway"
	"github.com/BinacsLee/server/service"
)

var (
	// StartCmd the start command
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Command",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var node gateway.NodeService
			if node, err = initService(); err != nil {
				return err
			} else if err = node.OnStart(); err != nil {
				return err
			}
			return nil
		},
	}
)

func initService() (gateway.NodeService, error) {
	nodeSvc := gateway.NodeServiceImpl{}

	// Config
	inject.Regist(Inject_Config, cfg)

	// Loggers
	inject.Regist(Inject_ZAPLOGGER, zaplogger)
	inject.Regist(Inject_LOGGER, logger)
	inject.Regist(Inject_Node_LOGGER, logger.With("module", "node"))
	inject.Regist(Inject_Web_LOGGER, logger.With("module", "web"))
	inject.Regist(Inject_GRPC_LOGGER, logger.With("module", "grpc"))
	inject.Regist(Inject_Redis_LOGGER, logger.With("module", "redis"))
	inject.Regist(Inject_Mysql_LOGGER, logger.With("module", "mysql"))

	// Services
	inject.Regist(Inject_Redis_Service, &service.RedisServiceImpl{})
	inject.Regist(Inject_Mysql_Service, &service.MysqlServiceImpl{})

	inject.Regist(Inject_Trace_Service, &service.TraceServiceImpl{})
	inject.Regist(Inject_Basic_Service, &service.BasicServiceImpl{})
	inject.Regist(Inject_User_Service, &service.UserServiceImpl{})
	inject.Regist(Inject_Crypto_Service, &service.CryptoServiceImpl{})
	inject.Regist(Inject_TinyURL_Service, &service.TinyURLServiceImpl{})
	inject.Regist(Inject_Pastebin_Service, &service.PastebinServiceImpl{})
	inject.Regist(Inject_Cos_Service, &service.CosServiceImpl{})

	inject.Regist(Inject_Web_Service, &gateway.WebServiceImpl{})
	inject.Regist(Inject_GRPC_Service, &gateway.GRPCServiceImpl{})
	inject.Regist(Inject_Node_Service, &nodeSvc)

	if err := inject.DoInject(); err != nil {
		return nil, err
	}
	return &nodeSvc, nil
}
