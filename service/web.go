package service

import (
	"net/http"
	"time"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service/web/controller"
	"github.com/gin-gonic/gin"
)

type WebService interface {
	Serve() error
}

type WebServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`
	r      *gin.Engine
	s      *http.Server
}

func (ws *WebServiceImpl) AfterInject() error {
	ws.r = gin.New()
	ws.r.Use(gin.Recovery())
	controller.SetRouter(ws.r)
	ws.s = &http.Server{
		Addr:           ":" + ws.Config.WebConfig.HttpPort,
		Handler:        ws.r,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return nil
}

func (ws *WebServiceImpl) Serve() error {
	ws.Logger.Info("GRPCService Serve", "HttpPort", ws.Config.WebConfig.HttpPort)
	err := ws.s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
