package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/metruzanca/rss/db"
	"github.com/metruzanca/rss/views"
)

// Some example feeds to get started with
var feeds = []db.Feed{
	{
		Id:   0,
		Name: "Jack O'Shea",
		Url:  "https://www.youtube.com/feeds/videos.xml?channel_id=UCY-BxclFSs6BFNzb7InPQDw",
	},
	{
		Id:   1,
		Url:  "https://gleam.run/feed",
		Name: "Gleam",
	},
	{
		Id:   2,
		Url:  "https://odysee.com/$/rss/@Coldfusion:f",
		Name: "Cold Fusion",
	},
}

const PORT = ":3000"

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return views.Render(c, views.Feeds(feeds))
	})

	e.GET("/f", func(c echo.Context) error {
		return views.Render(c, views.Feeds(feeds))
	})

	e.GET("/f/:id", func(c echo.Context) error {
		id := c.Param("id")
		value, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return c.String(http.StatusBadRequest, "Error invalid feed ID")
		}
		return views.Render(c, views.FeedPage((feeds[value])))
	})

	fmt.Printf("Listening on http://localhost%s", PORT)
	e.Start(PORT)
}
