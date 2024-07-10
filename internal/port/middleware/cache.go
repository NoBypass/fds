package middleware

import (
	"bytes"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type CacheMiddleware struct {
	cache *mincache.Cache
}

func NewCacheMiddleware(cache *mincache.Cache) CacheMiddleware {
	return CacheMiddleware{cache}
}

func (cmw CacheMiddleware) Cache(d time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().RequestURI
			res, ok := cmw.cache.Get(key)
			if ok {
				res, ok := res.([]byte)
				if !ok {
					cmw.cache.Delete(key)
				} else {
					return c.JSONBlob(200, res)
				}
			}

			body := new(bytes.Buffer)
			c.Response().Writer = &responseCapture{
				ResponseWriter: c.Response().Writer,
				body:           body,
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			cmw.cache.Set(key, body.Bytes(), d)
			return nil
		}
	}
}

type responseCapture struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r *responseCapture) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
