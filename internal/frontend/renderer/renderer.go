package renderer

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, fmt.Sprintf("%s.gohtml", name), data)
}

func New() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("internal/frontend/templates/*.gohtml")),
	}
}
