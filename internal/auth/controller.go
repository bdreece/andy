package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/bdreece/andy/pkg/database"
)

type Controller struct {
	db database.Querier
}

func (ctrl *Controller) Logout(c echo.Context) error {
	return nil
}

func NewController(db database.Querier) *Controller {
	return &Controller{db}
}
