package renderer

import (
	"html/template"

	"github.com/bdreece/andy/pkg/session"
	"github.com/labstack/echo/v4"
)

func funcMap(c echo.Context) template.FuncMap {
	return template.FuncMap{
		"session": func() *session.Values {
			return c.Get(session.ContextKey).(*session.Values)
		},
	}
}
