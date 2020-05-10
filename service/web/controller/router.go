package controller

import "github.com/gin-gonic/gin"

// SetRouter set all router
func SetRouter(r *gin.Engine) {
	SetBasicRouter(r)
	SetApiRouter(r.Group("api"))
	SetManagerRouter(r.Group("manager"))
	SetMonitorRouter(r.Group("monitor"))
}

// SetBasicRouter set basic router
func SetBasicRouter(r *gin.Engine) {
	r.StaticFile("/", "./test/static/index")
	r.StaticFile("/toys", "./test/static/toys")
	r.StaticFile("/toys/crypto", "./test/static/crypto")
	r.StaticFile("/about", "./test/static/about")
}

// SetApiRouter set RESTful api router
func SetApiRouter(r *gin.RouterGroup) {
}

// SetManagerRouter set manager router
func SetManagerRouter(r *gin.RouterGroup) {
	//r.POST("/reload", Reload)
}

// SetMonitorRouter set monitor router
func SetMonitorRouter(r *gin.RouterGroup) {
	//r.Get("/monitor", Monitor)
}
