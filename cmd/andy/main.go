package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/bdreece/andy/internal/auth"
	"github.com/bdreece/andy/internal/comment"
	"github.com/bdreece/andy/internal/content"
	"github.com/bdreece/andy/internal/passwordreset"
	"github.com/bdreece/andy/internal/post"
	"github.com/bdreece/andy/internal/profile"
	"github.com/bdreece/andy/internal/resource"
	"github.com/bdreece/andy/pkg/cli"
	"github.com/bdreece/andy/pkg/database"
	"github.com/bdreece/andy/pkg/logging"
	"github.com/bdreece/andy/pkg/renderer"
	"github.com/bdreece/andy/pkg/session"
	"github.com/bdreece/andy/pkg/validator"
)

func main() {
	fx.New(
        fx.Provide(cli.Parse),
		fx.Provide(loadConfig),
		fx.Provide(logging.NewLogger),
		fx.Provide(
            database.NewDBTX,
			fx.Annotate(
				database.New,
				fx.As(new(database.Querier)),
			),
		),
		fx.Provide(
			fx.Annotate(
				session.NewCookieStore,
				fx.As(new(sessions.Store)),
			),
            fx.Annotate(
                validator.New,
                fx.As(new(echo.Validator)),
            ),
			fx.Annotate(
				renderer.New,
				fx.As(new(echo.Renderer)),
			),
		),
		fx.Provide(
		    auth.NewController,
            comment.NewController,
            content.NewController,
		    passwordreset.NewController,
		    post.NewController,
		    profile.NewController,
		    resource.NewController,
        ),
		fx.Provide(createRouter),
		fx.Invoke(invokeRoute),
	).Run()
}
