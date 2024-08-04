package main

import (
	"github.com/labstack/echo/v4"
	"github.com/metruzanca/rss/db"
)

func main() {
	db.Init()
	db.GetUsers()

	e := echo.New()
	// e.Static("/", "public")

	e.Start(":3000")
}
