package auth

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/bdreece/andy/pkg/database"
)

type registerForm struct {
	FirstName       string `form:"first-name" validate:"required,max=63"`
	LastName        string `form:"last-name" validate:"required,max=63"`
	EmailAddress    string `form:"email-address" validate:"required,max=127"`
	Password        string `form:"password" validate:"required,max=127,eqfield=ConfirmPassword"`
	ConfirmPassword string `form:"confirm-password" validate:"required,max=127"`
}

func (ctrl *Controller) RegisterPage(c echo.Context) error {
	const template string = "register.gotmpl"
	return c.Render(http.StatusOK, template, echo.Map{})
}

func (ctrl *Controller) Register(c echo.Context) error {
	var form registerForm
	if err := c.Bind(&form); err != nil {
		return err
	}

	if err := c.Validate(&form); err != nil {
		return err
	}

	ctx := c.Request().Context()
	_, err := ctrl.db.UserByEmail(ctx, database.UserByEmailParams{
		EmailAddress: form.EmailAddress,
	})
	if err == nil {
		return echo.NewHTTPError(http.StatusConflict, "user already exists!")
	} else if err != sql.ErrNoRows {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(form.Password),
		bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	now := time.Now()
	if err = ctrl.db.InsertUser(ctx, database.InsertUserParams{
		CreatedAt:     now,
		UpdatedAt:     now,
		FirstName:     form.FirstName,
		LastName:      form.LastName,
		EmailAddress:  form.EmailAddress,
		EmailVerified: false,
		PasswordHash:  base64.StdEncoding.EncodeToString(hash),
	}); err != nil {
		return err
	}

	return c.Render(http.StatusOK, "partials::alert", echo.Map{
		"Icon":    "tabler:info",
		"Variant": "success",
		"Content": "Registered successfully!",
	})
}
