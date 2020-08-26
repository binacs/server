# Gin



[GoDoc](https://godoc.org/github.com/gin-gonic/gin)



## 1. 概述

Web 框架。

资料很多，略。



## 2. 关键概念

### 2.1 gin.Context

简单来说：gin的context实现了的context.Context `Interface` ；

封装 http ；

快速获取 Metadata（Set / Get 的键值对）、 Param（URL中RESTful风格的部分）、Query（URL中传递的参数）、PostForm（Body数据）等等。



### 2.2 gin.Engine

```go
func init() {
    router = gin.New()            // return a *gin.Engine
    router.Use(gin.Recovery())
    router.StaticFile("/", index.html)
    
    server = &http.Server{
        Addr:           ":" + ws.Config.WebConfig.HttpPort,
        Handler:        router,
        ReadTimeout:    time.Second,
        WriteTimeout:   time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    if err := server.ListenAndServe(); err != nil {
        // xxx
    }
}
```

Engine 可注册各类中间件以及路由，更多请参考 godoc ；

路由组（RouterGroup）及相关使用方法也请参考 godoc 。

这里需要提到：gin.Recovery 将返回一个 HandlerFunc ，其功能是 500 情况下自动从 panic 中恢复。



## 3. 使用示例

较简单，参考 2 部分以及项目代码。

