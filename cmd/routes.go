package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/metruzanca/rss/lib/db"
	"github.com/metruzanca/rss/lib/rss"
	"github.com/metruzanca/rss/lib/views"
)

// TODO move to these, and save some data in db
// var feedUrls = []string{
// 	"https://www.youtube.com/feeds/videos.xml?channel_id=UCY-BxclFSs6BFNzb7InPQDw", // Jack O'Shea
// 	"https://gleam.run/feed",                 // Gleam
// 	"https://odysee.com/$/rss/@Coldfusion:f", // Cold Fusion
// }

// Some example feeds to get started with
var feeds = []db.Feed{
	{
		Id:    0,
		Title: "Jack O'Shea",
		Url:   "https://www.youtube.com/feeds/videos.xml?channel_id=UCY-BxclFSs6BFNzb7InPQDw",
	},
	{
		Id:    1,
		Url:   "https://gleam.run/feed",
		Title: "Gleam",
	},
	{
		Id:    2,
		Url:   "https://odysee.com/$/rss/@Coldfusion:f",
		Title: "Cold Fusion",
	},
}

func RegisterRoutes() http.Handler {
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

		feed := feeds[value]

		parsedRss := rss.Parse(feed.Url)

		for _, item := range parsedRss.Items {
			feed.Items = append(feed.Items, db.FeedItem{
				Title:       item.Title,
				Description: item.Description,
				Url:         item.Link,
			})
		}

		return views.Render(c, views.FeedPage(feed))
	})

	return e
}
