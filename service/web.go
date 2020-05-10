package service

import (
	"net/http"
	"time"

	"github.com/unrolled/secure"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
	"github.com/BinacsLee/server/service/web/controller"
	"github.com/gin-gonic/gin"
)

// WebService the web service
type WebService interface {
	Serve() error
}

// WebServiceImpl inplement of WebService
type WebServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`
	r      *gin.Engine
	s      *http.Server
}

// AfterInject inject
func (ws *WebServiceImpl) AfterInject() error {
	ws.r = gin.New()
	ws.r.Use(gin.Recovery())
	ws.r.Use(ws.tlsTransfer())
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

// // Serve start web serve
func (ws *WebServiceImpl) Serve() error {
	ws.Logger.Info("WebService Serve", "HttpPort", ws.Config.WebConfig.HttpPort, "HttpsPort", ws.Config.WebConfig.HttpsPort)
	go func() {
		if err := ws.s.ListenAndServe(); err != nil {
			ws.Logger.Error("WebService Serve", "ListenAndServe err", err)
		}
	}()
	err := ws.r.RunTLS(":"+ws.Config.WebConfig.HttpsPort, ws.Config.WebConfig.CertPath, ws.Config.WebConfig.KeyPath)
	if err != nil {
		ws.Logger.Error("WebService Serve", "ListenAndServeTLS err", err)
		return err
	}
	return nil
}

func (ws *WebServiceImpl) tlsTransfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ws.Config.WebConfig.Host + ":" + ws.Config.WebConfig.HttpsPort,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			ws.Logger.Error("WebService tlsTransfer", "Process err", err)
			c.Abort()
			return
		}
		// Avoid header rewrite if response is a redirection.
		//if status := c.Writer.Status(); status > 300 && status < 399 {
		//	c.Abort()
		//}
		c.Next()
	}
}
