package auth

import (
	"database/sql"
	"encoding/base64"
	"net/http"

	"github.com/bdreece/andy/pkg/database"
	"github.com/bdreece/andy/pkg/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginForm struct {
	EmailAddress string `form:"email-address" validate:"required,max=127"`
	Password     string `form:"password" validate:"required,min=8,max=127"`
}

func (ctrl *Controller) LoginPage(c echo.Context) error {
	const template string = "login.gotmpl"
	return c.Render(http.StatusOK, template, echo.Map{})
}

func (ctrl *Controller) Login(c echo.Context) error {
	var form loginForm
	if err := c.Bind(&form); err != nil {
		return err
	}
	if err := c.Validate(&form); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := ctrl.db.UserByEmail(ctx, database.UserByEmailParams{
		EmailAddress: form.EmailAddress,
	})

	if err != nil && err != sql.ErrNoRows {
		return err
	} else if err == sql.ErrNoRows {
		c.Logger().Debugf("user not found: %q", form.EmailAddress)
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
	}

	hash, err := base64.StdEncoding.DecodeString(user.PasswordHash)
	if err != nil {
		c.Logger().Warnf("failed to decode user password: %q", user.PasswordHash)
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
	}

	if err = bcrypt.CompareHashAndPassword(hash, []byte(form.Password)); err != nil {
		c.Logger().Debug("invalid user password")
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
	}

	sess, err := session.Get(c)
	if err != nil {
		return err
	}

	if err = session.Encode(sess, &session.Values{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}); err != nil {
		return err
	}

	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Render(http.StatusOK, "partials::alert", echo.Map{
		"Icon":    "tabler:info",
		"Variant": "succes",
		"Content": "Logged in successfully!",
	})
}
