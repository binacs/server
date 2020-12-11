package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// TLSTransfer tls transfer
func TLSTransfer(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     host,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// Logger.Error("WebService TLSTransfer", "Process err", err)
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
