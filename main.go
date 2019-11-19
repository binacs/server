package main

import (
	//"log"
	//"os"
	"github.com/BinacsLee/server/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := iris.New()
	// set the static file path
	app.StaticWeb("/", "static")

	// package middleware

	// package route

	app.Any("/debug/pprof/{action:path}", pprof.New())

	conf := config.GetConfig()

	var addr = conf.Host + conf.Port
	app.Run(iris.Addr(addr))
}

func init() {
	/*
		log_file, err := os.OpenFile("log/Debug.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalln("fail to create log/Debug.log file")
		}
		defer log_file.Close()
		// format
	*/
}
