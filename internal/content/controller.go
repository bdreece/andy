package content

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
    cfg *Config
}

func (ctrl *Controller) Home(c echo.Context) error {
    const template string = "home.gotmpl"

    return c.Render(http.StatusOK, template, echo.Map{})
}

func NewController(cfg *Config) *Controller {
    return &Controller{cfg}
}
