package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/binacsgo/trace"
)

const (
	NameOfGinCtxTracer    = "tracer"
	NameOfGinCtxTracerCtx = "tracerCtx"
)

// JaegerTrace the jaegre trace middleware
func JaegerTrace(tracer trace.Trace) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span

		spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
		} else {
			parentSpan = tracer.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
		}
		defer parentSpan.Finish()
		c.Set(NameOfGinCtxTracer, tracer)
		c.Set(NameOfGinCtxTracerCtx, opentracing.ContextWithSpan(context.Background(), parentSpan))
		c.Next()
	}
}
