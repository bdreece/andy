package comment

import (
	"github.com/bdreece/andy/pkg/database"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	db database.Querier
}

func (ctrl *Controller) ListByPost(c echo.Context) error {
	return nil
}

func (ctrl *Controller) ListByComment(c echo.Context) error {
	return nil
}

func (ctrl *Controller) Get(c echo.Context) error {
	return nil
}

func (ctrl *Controller) Create(c echo.Context) error {
	return nil
}

func (ctrl *Controller) Update(c echo.Context) error {
	return nil
}

func (ctrl *Controller) Delete(c echo.Context) error {
	return nil
}

func (ctrl *Controller) Vote(c echo.Context) error {
	return nil
}

func NewController(db database.Querier) *Controller {
	return &Controller{db}
}
