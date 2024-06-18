package comment

import (
	"github.com/bdreece/andy/pkg/database"
	"github.com/labstack/echo/v4"
)

func QueryMiddleware(db database.Querier) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return next(c)
        }
    }
}
