package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/binacsgo/trace"

	"github.com/BinacsLee/server/config"
	web_middleware "github.com/BinacsLee/server/service/web/middleware"
)

const (
	FatalPscNilInGinContext = "fatal-psc-nil-in-gin-context"
	// ...
)

// TraceService trace service
type TraceService interface {
	StartSpan(operationName string) opentracing.Span
	Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error
	Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error)
	Close() error

	GetTracer() trace.Trace
	FromGinContext(c *gin.Context, serviceName string) opentracing.Span
}

// TraceServiceImpl inplement of TraceService
type TraceServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	tracer trace.Trace
}

// AfterInject inject
func (ts *TraceServiceImpl) AfterInject() error {
	var err error
	jaegerCfg := jaegercfg.Configuration{
		ServiceName: ts.Config.TraceConfig.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: ts.Config.TraceConfig.AgentHostPort,
		},
	}
	if ts.tracer, err = trace.MakeTrace(jaegerCfg); err != nil {
		return err
	}
	opentracing.SetGlobalTracer(ts.tracer)
	return nil
}

// StartSpan return the tracer's impl
func (ts *TraceServiceImpl) StartSpan(operationName string) opentracing.Span {
	return ts.tracer.StartSpan(operationName, nil)
}

// Inject return the tracer's impl
func (ts *TraceServiceImpl) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return ts.tracer.Inject(sm, format, carrier)
}

// Extract return the tracer's impl
func (ts *TraceServiceImpl) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return ts.tracer.Extract(format, carrier)
}

// Close the closer
func (ts *TraceServiceImpl) Close() error {
	return ts.tracer.Close()
}

// GetTracer return the tracer
func (ts *TraceServiceImpl) GetTracer() trace.Trace {
	return ts.tracer
}

// FromGinContext start a new span from gin context
func (ts *TraceServiceImpl) FromGinContext(c *gin.Context, serviceName string) opentracing.Span {
	psc, _ := c.Get(web_middleware.NameOfGinCtxTracerCtx)
	ctx, ok := psc.(context.Context)
	if ok {
		span, _ := opentracing.StartSpanFromContext(ctx, serviceName)
		return span
	}
	span := ts.StartSpan(FatalPscNilInGinContext)
	return span
}
