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
	ch := make(chan error, 64)
	s.errCh = ch
	return ch
}

func (s *service) Pipeline(fn func(startTrace func() opentracing.Span) error, this any) {
	name := runtime.FuncForPC(reflect.ValueOf(this).Pointer()).Name()
	name = name[strings.LastIndex(name, ".")+1:]
	name = name[:len(name)-3]

	var sp opentracing.Span
	sp = jaegertracing.CreateChildSpan(s.c, name)
	startTrace := func() opentracing.Span {
		sp = jaegertracing.CreateChildSpan(s.c, name)
		return sp
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				s.errCh <- r.(error)
				ext.LogError(sp, r.(error))
			}
		}()

		err := fn(startTrace)
		defer sp.Finish()
		if err != nil {
			s.errCh <- err
			ext.LogError(sp, err)
		}
	}()
}
