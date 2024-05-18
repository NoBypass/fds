package service

import (
	"fmt"
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
	Setup(echo.Context)
}

type service struct {
	errCh chan error
	c     echo.Context
}

func (s *service) Request(c echo.Context) <-chan error {
	s.c = c
	s.errCh = make(chan error, 32)
	return s.errCh
}

func (s *service) Setup(ctx echo.Context) {
	s.c = ctx
}

func (s *service) Trace(this any) (func(), opentracing.Span) {
	name := runtime.FuncForPC(reflect.ValueOf(this).Pointer()).Name()
	name = name[strings.LastIndex(name, ".")+1:]
	name = name[:len(name)-3]

	var sp opentracing.Span
	sp = jaegertracing.CreateChildSpan(s.c, name)
	endFn := func() {
		sp.Finish()

		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}

			stack := make([]byte, 1<<10)
			length := runtime.Stack(stack, false)
			stack = stack[:length]

			msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
			s.c.Logger().Error(msg)

			s.errCh <- echo.ErrInternalServerError
			ext.LogError(sp, r.(error))
		}
	}

	return endFn, sp
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
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				stack := make([]byte, 1<<10)
				length := runtime.Stack(stack, false)
				stack = stack[:length]

				msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
				s.c.Logger().Error(msg)

				s.errCh <- echo.ErrInternalServerError
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
