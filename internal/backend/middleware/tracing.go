package middleware

import (
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config2 "github.com/uber/jaeger-client-go/config"
	"io"
)

func StartTracer(v string) (opentracing.Tracer, io.Closer) {
	cfg := config2.Configuration{
		ServiceName: fmt.Sprintf("FDS backend %s", v),
		Reporter: &config2.ReporterConfig{
			LogSpans: true,
		},
		Sampler: &config2.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}
	tracer, closer, err := cfg.NewTracer(config2.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func Trace(tracer opentracing.Tracer) echo.MiddlewareFunc {
	return jaegertracing.TraceWithConfig(jaegertracing.TraceConfig{
		Skipper:       jaegertracing.DefaultTraceConfig.Skipper,
		IsBodyDump:    true,
		ComponentName: "backend",
		OperationNameFunc: func(c echo.Context) string {
			return fmt.Sprintf("%s %s", c.Request().Method, c.Path())
		},
		Tracer: tracer,
	})
}
