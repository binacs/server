package gateway

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/binacsgo/log"

	api_cos "github.com/BinacsLee/server/api/cos"
	api_crypto "github.com/BinacsLee/server/api/crypto"
	api_pastebin "github.com/BinacsLee/server/api/pastebin"
	api_tinyurl "github.com/BinacsLee/server/api/tinyurl"
	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/middleware"
	"github.com/BinacsLee/server/service"
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

	TraceSvc    service.TraceService    `inject-name:"TraceService"`
	BasicSvc    service.BasicService    `inject-name:"BasicService"`
	CryptoSvc   service.CryptoService   `inject-name:"CryptoService"`
	TinyURLSvc  service.TinyURLService  `inject-name:"TinyURLService"`
	PastebinSvc service.PastebinService `inject-name:"PastebinService"`
	CosSvc      service.CosService      `inject-name:"CosService"`

	r *gin.Engine
	s *http.Server
}

// AfterInject inject
func (ws *WebServiceImpl) AfterInject() error {
	ws.r = gin.New()

	ws.r.Use(gin.Recovery())
	ws.r.Use(middleware.JaegerTrace(ws.TraceSvc.GetTracer()))
	if ws.Config.WebConfig.SSLRedirect {
		ws.r.Use(middleware.TLSTransfer(ws.Config.WebConfig.Host + ":" + ws.Config.WebConfig.HttpsPort))
	}

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

	if ws.Config.WebConfig.SSLRedirect {
		ws.Logger.Info("WebService Serve ", "ListenAndServeTLS", true)
		go func() {
			if err := ws.r.RunTLS(":"+ws.Config.WebConfig.HttpsPort, ws.Config.WebConfig.CertPath, ws.Config.WebConfig.KeyPath); err != nil {
				ws.Logger.Error("WebService Serve", "ListenAndServeTLS err", err)
			}
		}()
	}

	// In fact, there is no need for `ws.s` (http server),
	// `Kubernetes Ingress` will handle the requests.
	if err := ws.s.ListenAndServe(); err != nil {
		ws.Logger.Error("WebService Serve", "ListenAndServe err", err)
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
	ws.setRecentPosts(r)
}

func (ws *WebServiceImpl) setBasicRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Home",
			"Body":  template.HTML(ws.BasicSvc.ServeHome()),
		})
	})
	r.GET("/toys", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Toys",
			"Body":  template.HTML(ws.BasicSvc.ServeToys()),
		})
	})
	r.GET("/toys/crypto", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Crypto",
			"Body":  template.HTML(ws.BasicSvc.ServeCrypto()),
		})
	})
	r.GET("/toys/tinyurl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - TinyURL",
			"Body":  template.HTML(ws.BasicSvc.ServeTinyURL()),
		})
	})
	r.GET("/toys/pastebin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Pastebin",
			"Body":  template.HTML(ws.BasicSvc.ServePastebin()),
		})
	})
	r.GET("/toys/storage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Storage",
			"Body":  template.HTML(ws.BasicSvc.ServeStorage()),
		})
	})
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - About",
			"Body":  template.HTML(ws.BasicSvc.ServeAbout()),
		})
	})
}

func (ws *WebServiceImpl) setRedirectRouter(r *gin.RouterGroup) {
	r.GET("/:turl", ws.redirect)
}

func (ws *WebServiceImpl) setPagesRouter(r *gin.RouterGroup) {
	r.GET("/:turl", ws.pages)
}

func (ws *WebServiceImpl) setRecentPosts(r *gin.Engine) {
	r.GET("/posts", ws.recentPosts)
}

// ------------------ Gin Service ------------------
func (ws *WebServiceImpl) redirect(c *gin.Context) {
	span := ws.TraceSvc.FromGinContext(c, "TinyURLSvc URLSearch")
	rsp, err := ws.TinyURLSvc.URLSearch(c.Param("turl"))
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl redirect", "URLSearch err", err, "turl", c.Param("turl"))
	}

	c.Redirect(http.StatusMovedPermanently, rsp)
}

func (ws *WebServiceImpl) pages(c *gin.Context) {
	span := ws.TraceSvc.FromGinContext(c, "PastebinSvc URLSearch")
	rsp, err := ws.PastebinSvc.URLSearch(c.Param("turl"))
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl pages", "URLSearch err", err, "turl", c.Param("turl"))
		c.HTML(http.StatusOK, "global", gin.H{
			"Title": "binacs.cn - Pages",
			"Body":  template.HTML(err.Error()),
		})
	}

	span = ws.TraceSvc.FromGinContext(c, "PastebinSvc Parse")
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

func (ws *WebServiceImpl) recentPosts(c *gin.Context) {
	span := ws.TraceSvc.FromGinContext(c, "Recent Posts")
	pages, err := ws.PastebinSvc.RecentPosts()
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl recent posts", "err", err)
	}

	f := func(page table.Page) string {
		link := ws.Config.WebConfig.GetDomain() + "/p/" + page.TinyURL
		a := fmt.Sprintf("<a href=\"%s\" target=\"_blank\">%s</a>", link, link)
		return fmt.Sprintf("Poster [ %s ]:  %s  at  %s<br>", page.Poster, a, page.CreatedAt)
	}

	var ret string
	for i := len(pages) - 1; i >= 0; i-- {
		if len(pages[i].Content) > 0 {
			ret += f(pages[i])
		}
	}
	c.HTML(http.StatusOK, "posts", template.HTML(ret))
}

// TODO @binacs DELETE EVERYTHING BELOW when grpc-web starts to work
func (ws *WebServiceImpl) setAPIRouter(r *gin.RouterGroup) {
	r.POST("/crypto/encrypto", ws.apiCryptoEncrypto)
	r.POST("/crypto/decrypto", ws.apiCryptoDecrypto)

	r.POST("/tinyurl/encode", ws.apiTinyURLEncode)
	r.POST("/tinyurl/decode", ws.apiTinyURLDecode)

	r.POST("/pastebin/submit", ws.apiPastebinSubmit)

	r.POST("/cos/put", ws.apiCosPut)
}

func (ws *WebServiceImpl) apiCryptoEncrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")

	span := ws.TraceSvc.FromGinContext(c, "CryptoSvc CryptoEncrypt")
	rsp, err := ws.CryptoSvc.CryptoEncrypt(nil, &api_crypto.CryptoEncryptReq{
		Algorithm: tp,
		PlainText: text,
	})
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiCryptoEncrypto", "err", err, "text", text, "type", tp)
	}

	c.String(http.StatusOK, rsp.Data.EncryptText)
}

func (ws *WebServiceImpl) apiCryptoDecrypto(c *gin.Context) {
	text := c.Request.FormValue("text")
	tp := c.Request.FormValue("type")

	span := ws.TraceSvc.FromGinContext(c, "CryptoSvc CryptoDecrypto")
	rsp, err := ws.CryptoSvc.CryptoDecrypt(nil, &api_crypto.CryptoDecryptReq{
		Algorithm:   tp,
		EncryptText: text,
	})
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiCryptoDecrypto", "err", err, "text", text, "type", tp)
	}

	c.String(http.StatusOK, rsp.Data.PlainText)
}

// -------- TinyURL Service --------
func (ws *WebServiceImpl) apiTinyURLEncode(c *gin.Context) {
	text := c.Request.FormValue("text")

	span := ws.TraceSvc.FromGinContext(c, "TinyURLSvc TinyURLEncode")
	rsp, err := ws.TinyURLSvc.TinyURLEncode(nil, &api_tinyurl.TinyURLEncodeReq{
		Url: text,
	})
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiTinyURLEncode", "err", err, "text", text)
	}

	c.String(http.StatusOK, rsp.Data.Turl)
}

func (ws *WebServiceImpl) apiTinyURLDecode(c *gin.Context) {
	text := c.Request.FormValue("text")

	span := ws.TraceSvc.FromGinContext(c, "TinyURLSvc TinyURLDecode")
	rsp, err := ws.TinyURLSvc.TinyURLDecode(nil, &api_tinyurl.TinyURLDecodeReq{
		Turl: text,
	})
	span.Finish()
	if err != nil {
		ws.Logger.Error("WebServiceImpl apiTinyURLDecode", "err", err, "text", text)
	}

	c.String(http.StatusOK, rsp.Data.Url)
}

// -------- PasteBin Service --------
func (ws *WebServiceImpl) apiPastebinSubmit(c *gin.Context) {
	span := ws.TraceSvc.FromGinContext(c, "PastebinSvc Submit")
	rsp, err := ws.PastebinSvc.PastebinSubmit(nil, &api_pastebin.PastebinSubmitReq{
		Author: c.Request.FormValue("poster"),
		Syntax: c.Request.FormValue("syntax"),
		Text:   c.Request.FormValue("content"),
	})
	span.Finish()
	if err != nil {
		ws.Logger.Error("apiPastebinSubmit Submit", "err", err)
		c.String(http.StatusNotFound, "Submit to database failed!")
		return
	}

	c.String(http.StatusOK, rsp.Data.GetPurl())
	// c.Redirect(http.StatusMovedPermanently, rsp.Data.GetPurl())
}

// -------- Cos Service --------
func (ws *WebServiceImpl) apiCosPut(c *gin.Context) {
	span := ws.TraceSvc.FromGinContext(c, "CosSvc Put")
	form, err := c.MultipartForm()
	if err != nil {
		ws.Logger.Error("apiCosPut MultipartForm", "err", err)
		c.String(http.StatusNotFound, "Get MultipartForm failed!")
		span.Finish()
		return
	}

	key := form.Value["key"]
	if len(key) != 1 || key[0] != ws.Config.CosConfig.PassKey {
		ws.Logger.Error("apiCosPut PassKey incorrect", "key", key)
		c.String(http.StatusOK, "PassKey incorrect!")
		span.Finish()
		return
	}

	if len(form.File) == 0 {
		ws.Logger.Error("apiCosPut form.File empty")
		c.String(http.StatusOK, "No file selected!")
		span.Finish()
		return
	}

	var URIS string
	files := form.File["file"]
	for i, f := range files {
		ws.Logger.Info("apiCosPut range files", "index", i)
		file, err := f.Open()
		if err != nil {
			ws.Logger.Error("apiCosPut f.Open", "err", err)
			c.String(http.StatusNotFound, "f.Open failed!")
			continue
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			ws.Logger.Error("apiCosPut io.Copy", "err", err)
			c.String(http.StatusNotFound, "io.Copy failed!")
			continue
		}

		resp, err := ws.CosSvc.CosPut(context.Background(), &api_cos.CosPutReq{
			FileName:  f.Filename,
			FileBytes: buf.Bytes(),
		})
		if err != nil {
			ws.Logger.Error("apiCosPut CosPut", "err", err)
			continue
		}
		URIS += resp.Data.CosURI + "\n\n"
	}
	span.Finish()
	ws.Logger.Info("apiCosPut", "URIS", URIS)

	c.String(http.StatusOK, URIS)
}

/* ------------------ NO USE FOR NOW ------------------
func (ws *WebServiceImpl) setManagerRouter(r *gin.RouterGroup) {
	// r.POST("/reload", Reload)
}
*/
