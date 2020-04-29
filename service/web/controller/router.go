package controller

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	SetBasicRouter(r)
	SetApiRouter(r.Group("api"))
	SetManagerRouter(r.Group("manager"))
	SetMonitorRouter(r.Group("monitor"))
}

func SetApiRouter(r *gin.RouterGroup) {
	r.StaticFile("/", "./test/static/index.html")
}

func SetApiRouter(r *gin.RouterGroup) {
}

func SetManagerRouter(r *gin.RouterGroup) {
	//r.POST("/reload", Reload)
}

func SetMonitorRouter(r *gin.RouterGroup) {
	//r.Get("/monitor", Monitor)
}
