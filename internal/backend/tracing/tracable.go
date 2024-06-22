package tracing

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"reflect"
	"runtime"
	"strings"
)

type Tracable interface {
	StartSpan(ctx context.Context, name any) (opentracing.Span, context.Context)
}

type tracable struct {
	tracer opentracing.Tracer
}

func NewTracable(serviceName ...string) Tracable {
	if len(serviceName) == 0 {
		return &tracable{}
	}

	cfg := config.Configuration{
		ServiceName: serviceName[0],
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, _, err := cfg.NewTracer()
	if err != nil {
		log.Fatal(err)
	}

	return &tracable{tracer}
}

func (t *tracable) StartSpan(ctx context.Context, name any) (opentracing.Span, context.Context) {
	switch name.(type) {
	case string:
	default:
		name = runtime.FuncForPC(reflect.ValueOf(name).Pointer()).Name()
		name = name.(string)[strings.LastIndex(name.(string), ".")+1:]
		name = name.(string)[:len(name.(string))-3]
	}

	tracer, ok := ctx.Value("tracer").(opentracing.Tracer)
	if !ok {
		tracer = opentracing.GlobalTracer()
	}

	if t.tracer != nil {
		tracer = t.tracer
	}

	sp, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, name.(string))
	ctx = context.WithValue(ctx, "tracer", tracer)
	return sp, ctx
}
