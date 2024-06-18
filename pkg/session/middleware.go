package session

import "github.com/labstack/echo/v4"

const ContextKey string = "session-context"

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var values *Values
		sess, err := Get(c)
		if err != nil {
			goto done
		}

		values, err = Decode(sess)
		if err != nil {
			goto done
		}

		c.Set(ContextKey, values)

	done:
		return next(c)
	}
}
