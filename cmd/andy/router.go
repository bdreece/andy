package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gorilla/sessions"
	echoSession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	slogEcho "github.com/samber/slog-echo"
	"go.uber.org/fx"

	"github.com/bdreece/andy/internal/auth"
	"github.com/bdreece/andy/internal/comment"
	"github.com/bdreece/andy/internal/content"
	"github.com/bdreece/andy/internal/passwordreset"
	"github.com/bdreece/andy/internal/post"
	"github.com/bdreece/andy/internal/profile"
	"github.com/bdreece/andy/internal/resource"
	"github.com/bdreece/andy/pkg/config"
	"github.com/bdreece/andy/pkg/session"
	"github.com/bdreece/andy/web"
)

func createRouter(p struct {
	fx.In

	Lifecycle fx.Lifecycle

	Auth          *auth.Controller
	Comments      *comment.Controller
	Content       *content.Controller
	PasswordReset *passwordreset.Controller
	Posts         *post.Controller
	Profile       *profile.Controller
	Resources     *resource.Controller

	Logger       *slog.Logger
	Renderer     echo.Renderer
	Validator    echo.Validator
	SessionStore sessions.Store
	Config       *config.Root
}) *echo.Echo {
	e := echo.New()
	e.Renderer = p.Renderer
	e.Validator = p.Validator
	e.HTTPErrorHandler = handleError

	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.BodyLimit("4M"),
		echoMiddleware.Decompress(),
		echoMiddleware.Gzip(),
		echoMiddleware.Secure(),
		echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
            CookieName: "_andy-csrf",
        }),
		echoSession.Middleware(p.SessionStore),
		echoMiddleware.StaticWithConfig(echoMiddleware.StaticConfig{
			Filesystem: http.FS(web.Static),
		}),
		echoMiddleware.StaticWithConfig(echoMiddleware.StaticConfig{
			Filesystem: http.FS(web.App),
		}),
		slogEcho.New(p.Logger),
        session.Middleware,
	)

	e.GET("/", p.Content.Home)

	e.GET("/login", p.Auth.LoginPage)
	e.POST("/login", p.Auth.Login)
	e.GET("/register", p.Auth.RegisterPage)
	e.POST("/register", p.Auth.Register)
	e.GET("/logout", p.Auth.Logout)

	e.POST("/password-reset/send", p.PasswordReset.Send)
	e.POST("/password-reset/verify", p.PasswordReset.Verify)

	e.GET("/user/:id", p.Profile.Get)
	e.PUT("/user/:id", p.Profile.Update)
	e.GET("/user/:id/resource", p.Resources.ListByUser)

	e.GET("/post", p.Posts.List)
	e.POST("/post", p.Posts.Create)
	e.GET("/post/:id", p.Posts.Get)
	e.PUT("/post/:id", p.Posts.Update)
	e.DELETE("/post/:id", p.Posts.Delete)
	e.POST("/post/:id/vote", p.Posts.Vote)
	e.GET("/post/:id/comment", p.Comments.ListByPost)
	e.GET("/post/:id/resource", p.Resources.ListByPost)

	e.POST("/comment", p.Comments.Create)
	e.GET("/comment/:id", p.Comments.Get)
	e.DELETE("/comment/:id", p.Comments.Delete)
	e.POST("/comment/:id/vote", p.Posts.Vote)
	e.GET("/comment/:id/comment", p.Comments.ListByComment)

	e.POST("/resource", p.Resources.Create)
	e.GET("/resource/:id", p.Resources.Get)
	e.PUT("/resource/:id", p.Resources.Update)
	e.DELETE("/resource/:id", p.Resources.Delete)

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go e.Start(p.Config.Addr())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	return e
}

func handleError(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	c.Logger().Error(err)

	_ = c.JSON(code, map[string]any{
		"message": err.Error(),
	})
}

func invokeRoute(*echo.Echo) {}
