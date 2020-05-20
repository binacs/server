package service

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
	web_service "github.com/BinacsLee/server/service/web/service"
)

// WebService the web service
type WebService interface {
	Serve() error
}

// WebServiceImpl inplement of WebService
type WebServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	CryptoSvc *web_service.WebCryptoServiceImpl `inject-name:"WebCryptoService"`

	r *gin.Engine
	s *http.Server
}

// AfterInject inject
func (ws *WebServiceImpl) AfterInject() error {
	ws.r = gin.New()
	ws.r.Use(gin.Recovery())
	ws.r.Use(ws.tlsTransfer())
	ws.setRouter(ws.r)
	ws.s = &http.Server{
		Addr:           ":" + ws.Config.WebConfig.HttpPort,
		Handler:        ws.r,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return nil
}

// Serve start web serve
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

// ------------------ Gin Router ------------------

// setRouter set all router
func (ws *WebServiceImpl) setRouter(r *gin.Engine) {
	ws.setBasicRouter(r)
	ws.setApiRouter(r.Group("api"))
	ws.setManagerRouter(r.Group("manager"))
	ws.setMonitorRouter(r.Group("monitor"))
}

// setBasicRouter set basic router
func (ws *WebServiceImpl) setBasicRouter(r *gin.Engine) {
	r.StaticFile("/", ws.Config.WebConfig.StaticPath+"index")
	r.StaticFile("/toys", ws.Config.WebConfig.StaticPath+"toys")
	r.StaticFile("/toys/crypto", ws.Config.WebConfig.StaticPath+"crypto")
	r.StaticFile("/about", ws.Config.WebConfig.StaticPath+"about")
}

// setApiRouter set RESTful api router
func (ws *WebServiceImpl) setApiRouter(r *gin.RouterGroup) {
	r.POST("/v1/crypto/encrypto", ws.apiV1CryptoEncrypto)
	r.POST("/v1/crypto/decrypto", ws.apiV1CryptoDecrypto)
}

// setManagerRouter set manager router
func (ws *WebServiceImpl) setManagerRouter(r *gin.RouterGroup) {
	//r.POST("/reload", Reload)
}

// setMonitorRouter set monitor router
func (ws *WebServiceImpl) setMonitorRouter(r *gin.RouterGroup) {
	r.Any("/prometheus/*path", ws.prometheusReverseProxy())
	//r.GET("/grafana/*path", ws.grafanaReverseProxy())
}

// ------------------ Gin Service ------------------

func (ws *WebServiceImpl) apiV1CryptoEncrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")
	rsp, err := ws.CryptoSvc.CryptoEncrypt(c, text, tp)
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1CryptoEncrypto", "err", err, "text", text, "type", tp)
	}
	c.String(http.StatusOK, rsp)
}

func (ws *WebServiceImpl) apiV1CryptoDecrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")
	rsp, err := ws.CryptoSvc.CryptoDecrypt(c, text, tp)
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1CryptoDecrypto", "err", err, "text", text, "type", tp)
	}
	c.String(http.StatusOK, rsp)
}

// ------------------ ReverseProxy ------------------

func (ws *WebServiceImpl) prometheusReverseProxy() gin.HandlerFunc {
	target := ws.Config.WebConfig.ReverseProxy["prometheus"]
	//target := "http://127.0.0.1:9000" //转向的host
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			ws.Logger.Error("WebService prometheusReverseProxy", "error", err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		c.Request.URL.Path = "monitor/prometheus" + c.Param("path") //请求API
		ws.Logger.Info("WebService prometheusReverseProxy ready to serve", "path", c.Param("path"), "url.path", c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

/*
func (ws *WebServiceImpl) grafanaReverseProxy() gin.HandlerFunc {
	target := ws.Config.WebConfig.ReverseProxy["grafana"]
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			ws.Logger.Error("WebService grafanaReverseProxy", "error", err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		c.Request.URL.Path = c.Param("path") //请求API
		ws.Logger.Info("WebService grafanaReverseProxy ready to serve", "path", c.Param("path"), "url.path", c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
*/
