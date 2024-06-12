package renderer

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Index struct {
	Name string
}

func (t *Template) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", Index{"NoBypass"})
}
