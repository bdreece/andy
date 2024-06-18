package passwordreset

import (
	"github.com/bdreece/andy/pkg/database"
	"github.com/labstack/echo/v4"
)

type Controller struct {
    db database.Querier
}

func (ctrl *Controller) Send(c echo.Context) error {
    return nil
}

func (ctrl *Controller) Verify(c echo.Context) error {
    return nil
}

func NewController(db database.Querier) *Controller {
    return &Controller{db}
}
