package rss

import (
	"github.com/mmcdole/gofeed"
)

func Parse(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	return feed
}
