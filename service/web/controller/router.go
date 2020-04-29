package controller

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	SetBasicRouter(r)
	SetApiRouter(r.Group("api"))
	SetManagerRouter(r.Group("manager"))
	SetMonitorRouter(r.Group("monitor"))
}

func SetBasicRouter(r *gin.Engine) {
	r.StaticFile("/", "./test/static/index")
	r.StaticFile("/toys", "./test/static/toys")
	r.StaticFile("/toys/crypto", "./test/static/crypto")
	r.StaticFile("/about", "./test/static/about")
}

func SetApiRouter(r *gin.RouterGroup) {
}

func SetManagerRouter(r *gin.RouterGroup) {
	//r.POST("/reload", Reload)
}

func SetMonitorRouter(r *gin.RouterGroup) {
	//r.Get("/monitor", Monitor)
}
