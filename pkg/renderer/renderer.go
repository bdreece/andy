package renderer

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"

	"github.com/bdreece/andy/web"
)

var (
	ErrClone = errors.New("renderer: failed to clone templates")
	ErrParse = errors.New("renderer: failed to parse templates")
)

type Renderer struct {
	*template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := r.Clone()
	if err != nil {
		return errors.Join(ErrClone, err)
	}

	t.Funcs(funcMap(c))

	if strings.Contains(name, ".gotmpl") {
		_, err = t.ParseFS(web.Templates, name)
		if err != nil {
			return errors.Join(fmt.Errorf("failed to parse template: %q", name), err)
		}
	}

	if err = t.ExecuteTemplate(w, name, data); err != nil {
		return errors.Join(fmt.Errorf("failed to render template: %q", name), err)
	}

	return nil
}

func New(logger *slog.Logger) (*Renderer, error) {
	tmpl, err := template.New("").
		Funcs(sprig.FuncMap()).
		Funcs(funcMap(nil)).
		ParseFS(web.Templates,
			"layout/*.gotmpl",
			"partials/*.gotmpl")

	if err != nil {
		return nil, errors.Join(ErrParse, err)
	}

	return &Renderer{tmpl}, nil
}
