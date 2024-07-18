package main

import (
	"github.com/metruzanca/rss/db"
)

func main() {
	db.Init()
	db.GetUsers()
}
