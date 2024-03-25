package service

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"reflect"
	"runtime"
	"strings"
)

type Service interface {
	Request(echo.Context) <-chan error
}

type service struct {
	errCh chan<- error
	c     echo.Context
}

func (s *service) Request(c echo.Context) <-chan error {
	s.c = c
	ch := make(chan error, 1)
	s.errCh = ch
	return ch
}

func (s *service) Pipeline(fn func(startTrace func(), span opentracing.Span) error, this any) {
	name := runtime.FuncForPC(reflect.ValueOf(this).Pointer()).Name()
	name = name[strings.LastIndex(name, ".")+1:]
	name = name[:len(name)-3]

	var sp opentracing.Span
	startTrace := func() {
		sp = jaegertracing.CreateChildSpan(s.c, name)
	}

	go func() {
		err := fn(startTrace, sp)
		if err != nil {
			s.errCh <- err
			ext.LogError(sp, err)
		}
		sp.Finish()
	}()
}
