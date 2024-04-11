package middleware

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	config2 "github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
)

func StartTracer(v string) io.Closer {
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
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal(err)
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}

func Trace() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			tracer := opentracing.GlobalTracer()

			var sp opentracing.Span
			ctx, err := tracer.Extract(
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(req.Header),
			)
			op := fmt.Sprintf("HTTP %s %s", req.Method, req.URL.Path)
			if err != nil {
				sp = tracer.StartSpan(op)
			} else {
				sp = tracer.StartSpan(op, ext.RPCServerOption(ctx))
			}
			defer sp.Finish()

			if jaegerSpanContext, ok := sp.Context().(jaeger.SpanContext); ok {
				c.Set("traceID", jaegerSpanContext.TraceID().String())
			}

			reqBody := []byte{}
			if c.Request().Body != nil {
				reqBody, _ = io.ReadAll(c.Request().Body)
				sp.LogKV("http.req.body", string(reqBody))
			}
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			reqSpan := req.WithContext(opentracing.ContextWithSpan(req.Context(), sp))
			c.SetRequest(reqSpan)

			buf := new(bytes.Buffer)
			resp := c.Response()
			respDumper := &responseDumper{
				ResponseWriter: resp.Writer,
				mw:             io.MultiWriter(resp.Writer, buf),
				buf:            buf,
			}
			c.Response().Writer = respDumper

			err = next(c)
			if err != nil {
				c.Error(err)
				ext.LogError(sp, err)
			}

			sp.LogKV("http.resp.body", respDumper.GetResponse())
			return nil
		}
	}
}

type responseDumper struct {
	http.ResponseWriter
	mw  io.Writer
	buf *bytes.Buffer
}

func (d *responseDumper) Write(b []byte) (int, error) {
	return d.mw.Write(b)
}

func (d *responseDumper) GetResponse() string {
	return d.buf.String()
}
