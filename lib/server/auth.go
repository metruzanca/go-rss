package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO OAuth See https://github.com/markbates/goth/blob/master/examples/main.go
// TODO Cookie Session for dev

func RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	g.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusTeapot, "TODO: Not Implemented")
	})

	g.POST("/logout", func(c echo.Context) error {
		return c.String(http.StatusTeapot, "TODO: Not Implemented")
	})
}
