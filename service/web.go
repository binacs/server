package service

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/binacsgo/log"

	"github.com/BinacsLee/server/config"
	web_middleware "github.com/BinacsLee/server/service/web/middleware"
	web_service "github.com/BinacsLee/server/service/web/service"
	"github.com/BinacsLee/server/types/table"
)

// WebService the web service
type WebService interface {
	Serve() error
}

// WebServiceImpl inplement of WebService
type WebServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`
	Trace  TraceService   `inject-name:"TraceService"`

	BasicSvc    *web_service.WebBasicServiceImpl    `inject-name:"WebBasicService`
	CryptoSvc   *web_service.WebCryptoServiceImpl   `inject-name:"WebCryptoService"`
	TinyURLSvc  *web_service.WebTinyURLServiceImpl  `inject-name:"WebTinyURLService"`
	PastebinSvc *web_service.WebPastebinServiceImpl `inject-name:"WebPastebin_Service"`

	r *gin.Engine
	s *http.Server
}

// AfterInject inject
func (ws *WebServiceImpl) AfterInject() error {
	ws.r = gin.New()
	ws.r.Use(gin.Recovery())
	ws.r.Use(web_middleware.TLSTransfer(ws.Config.WebConfig.Host + ":" + ws.Config.WebConfig.HttpsPort))
	ws.r.Use(web_middleware.JaegerTrace(ws.Trace.GetTracer()))
	ws.r.LoadHTMLGlob(ws.Config.WebConfig.TmplPath + "*.html")
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

// ------------------ Gin Router ------------------

// setRouter set all router
func (ws *WebServiceImpl) setRouter(r *gin.Engine) {
	ws.setBasicRouter(r)
	ws.setAPIRouter(r.Group("api"))
	ws.setRedirectRouter(r.Group("r"))
	ws.setPagesRouter(r.Group("p"))

	// ws.setManagerRouter(r.Group("manager"))
	// ws.setMonitorRouter(r.Group("monitor"))
}

func (ws *WebServiceImpl) setBasicRouter(r *gin.Engine) {
	r.GET("/", ws.BasicSvc.ServeHome)
	r.GET("/toys", ws.BasicSvc.ServeToys)
	r.GET("/toys/crypto", ws.BasicSvc.ServeCrypto)
	r.GET("/toys/tinyurl", ws.BasicSvc.ServeTinyURL)
	r.GET("/toys/pastebin", ws.BasicSvc.ServePastebin)
	r.GET("/about", ws.BasicSvc.ServeAbout)
}

func (ws *WebServiceImpl) setRedirectRouter(r *gin.RouterGroup) {
	r.GET("/:turl", ws.redirect)
}

func (ws *WebServiceImpl) setPagesRouter(r *gin.RouterGroup) {
	r.GET("/:turl", ws.pages)
}

func (ws *WebServiceImpl) setAPIRouter(r *gin.RouterGroup) {
	r.POST("/v1/crypto/encrypto", ws.apiV1CryptoEncrypto)
	r.POST("/v1/crypto/decrypto", ws.apiV1CryptoDecrypto)

	r.POST("/v1/tinyurl/encode", ws.apiV1TinyURLEncode)
	r.POST("/v1/tinyurl/decode", ws.apiV1TinyURLDecode)

	r.POST("/v1/pastebin/submit", ws.apiV1PastebinSubmit)
}

// ------------------ Gin Service ------------------

func (ws *WebServiceImpl) redirect(c *gin.Context) {
	span := ws.Trace.FromGinContext(c, "TinyURLSvc URLSearch")
	rsp, err := ws.TinyURLSvc.URLSearch(c.Param("turl"))
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl redirect", "URLSearch err", err, "turl", c.Param("turl"))
	}

	c.Redirect(http.StatusMovedPermanently, rsp)
}

func (ws *WebServiceImpl) pages(c *gin.Context) {
	span := ws.Trace.FromGinContext(c, "PastebinSvc URLSearch")
	rsp, err := ws.PastebinSvc.URLSearch(c.Param("turl"))
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl pages", "URLSearch err", err, "turl", c.Param("turl"))
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Pages",
			"Body":  template.HTML(err.Error()),
		})
	}

	span = ws.Trace.FromGinContext(c, "PastebinSvc Parse")
	body := ws.PastebinSvc.Parse(rsp.Content, rsp.Syntax)
	span.Finish()

	c.HTML(http.StatusOK, "pages", gin.H{
		"Title":    "binacs.cn - Pages",
		"TinyURL":  rsp.TinyURL,
		"Poster":   rsp.Poster,
		"Syntax":   rsp.Syntax,
		"CreateAt": rsp.CreatedAt,
		"Content":  template.HTML(body),
	})
}

// -------- Crypto Service --------
func (ws *WebServiceImpl) apiV1CryptoEncrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")

	span := ws.Trace.FromGinContext(c, "CryptoSvc CryptoEncrypt")
	rsp, err := ws.CryptoSvc.CryptoEncrypt(text, tp)
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1CryptoEncrypto", "err", err, "text", text, "type", tp)
	}

	c.String(http.StatusOK, rsp)
}

func (ws *WebServiceImpl) apiV1CryptoDecrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")

	span := ws.Trace.FromGinContext(c, "CryptoSvc CryptoDecrypto")
	rsp, err := ws.CryptoSvc.CryptoDecrypt(text, tp)
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1CryptoDecrypto", "err", err, "text", text, "type", tp)
	}

	c.String(http.StatusOK, rsp)
}

// -------- TinyURL Service --------
func (ws *WebServiceImpl) apiV1TinyURLEncode(c *gin.Context) {
	text := c.Request.FormValue("text")

	span := ws.Trace.FromGinContext(c, "TinyURLSvc TinyURLEncode")
	rsp, err := ws.TinyURLSvc.URLEncode(text)
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1TinyURLEncode", "err", err, "text", text)
	}

	c.String(http.StatusOK, rsp)
}

func (ws *WebServiceImpl) apiV1TinyURLDecode(c *gin.Context) {
	text := c.Request.FormValue("text")

	span := ws.Trace.FromGinContext(c, "TinyURLSvc TinyURLDecode")
	rsp, err := ws.TinyURLSvc.URLDecode(text)
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiV1TinyURLDecode", "err", err, "text", text)
	}

	c.String(http.StatusOK, rsp)
}

// -------- PasteBin Service --------
func (ws *WebServiceImpl) apiV1PastebinSubmit(c *gin.Context) {
	content := c.Request.FormValue("content")

	span := ws.Trace.FromGinContext(c, "TinyURLSvc Encode")
	turl := ws.TinyURLSvc.Encode(content + string(time.Now().Unix()))
	span.Finish()

	p := &table.Page{
		Poster:  c.Request.FormValue("poster"),
		Syntax:  c.Request.FormValue("syntax"),
		Content: content,
		TinyURL: turl,
	}

	span = ws.Trace.FromGinContext(c, "PastebinSvc Submit")
	err := ws.PastebinSvc.Submit(p)
	span.Finish()
	if err != nil {
		ws.Logger.Error("apiV1PastebinSubmit Submit", "err", err)
		c.String(http.StatusNotFound, "Submit to database failed!")
		return
	}

	c.String(http.StatusOK, "/p/"+turl)
	// c.Redirect(http.StatusMovedPermanently, ws.Config.WebConfig.Host+"/p/"+turl)
}

/* ------------------ NO USE FOR NOW ------------------
func (ws *WebServiceImpl) setBasicRouter(r *gin.Engine) {
	r.StaticFile("/", ws.Config.WebConfig.StaticPath+"index")
	r.StaticFile("/toys", ws.Config.WebConfig.StaticPath+"toys")
	r.StaticFile("/toys/crypto", ws.Config.WebConfig.StaticPath+"crypto")
	r.StaticFile("/toys/tinyurl", ws.Config.WebConfig.StaticPath+"tinyurl")
	r.StaticFile("/about", ws.Config.WebConfig.StaticPath+"about")
}
func (ws *WebServiceImpl) setManagerRouter(r *gin.RouterGroup) {
	// r.POST("/reload", Reload)
}
func (ws *WebServiceImpl) setMonitorRouter(r *gin.RouterGroup) {
	// r.Any("/prometheus/*path", ws.prometheusReverseProxy())
	// r.GET("/grafana/*path", ws.grafanaReverseProxy())
}
 ------------------ ReverseProxy ------------------
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
 ------------------ ReverseProxy ------------------ */
