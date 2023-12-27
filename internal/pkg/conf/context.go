package conf

import (
	"github.com/labstack/echo/v4"
	"server/internal/core/repository"
)

type Context struct {
	echo.Context
	Config *Config
	Repo   *repository.Repository
}

func Construct(conf *Config) *Context {
	repo := conf.ConnectDB()

	return &Context{
		Repo:   repo,
		Config: conf,
	}
}

func (c *Context) Populate(ctx echo.Context) *Context {
	c.Context = ctx
	return c
}

func Preload(e echo.Context) *Context {
	return e.(*Context)
}
