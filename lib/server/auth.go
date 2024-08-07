package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/metruzanca/rss/lib/views"
)

// TODO OAuth See https://github.com/markbates/goth/blob/master/examples/main.go
// TODO Cookie Session for dev

func RegisterAuthRoutes(e *echo.Echo) {
	e.GET("/login", func(c echo.Context) error {
		return views.Render(c, views.LoginForm())
	})

	e.POST("/login", func(c echo.Context) error {
		// TODO login
		return c.Redirect(http.StatusOK, "/")
	})

	e.POST("/register", func(c echo.Context) error {
		// TODO register
		return c.Redirect(http.StatusOK, "/")
	})

	e.POST("/logout", func(c echo.Context) error {
		// TODO logout
		return c.Redirect(http.StatusOK, "/")
	})
}
