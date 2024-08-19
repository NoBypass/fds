package trace

import (
	"fmt"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/common/version"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func SetupTracer(c *env.Env) io.Closer {
	cfg := config.Configuration{
		ServiceName: fmt.Sprintf("FDS backend %s", version.VERSION),
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: c.JaegerEndpoint,
		},
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("failed setting up the tracer: %s", err)
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}
