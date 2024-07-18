package main

import (
	"rss-reader/db"
)

func main() {
	db.Init()
	db.GetUsers()
}
