package session

import (
	"github.com/gorilla/sessions"
	echoSession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type Values struct {
	ID        int64  `mapstructure:"id"`
	FirstName string `mapstructure:"first_name"`
	LastName  string `mapstructure:"last_name"`
}

func Get(c echo.Context) (*sessions.Session, error) {
	return echoSession.Get("_andy-session", c)
}

func Decode(sess *sessions.Session) (*Values, error) {
	values := new(Values)
	if err := mapstructure.Decode(sess.Values, values); err != nil {
		return nil, err
	}

	return values, nil
}

func Encode(sess *sessions.Session, values *Values) error {
	return mapstructure.Decode(values, &sess.Values)
}
