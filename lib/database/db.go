// https://docs.turso.tech/sdk/go/quickstart#local-embedded-replicas
package database

import "database/sql"

// OLD
type FeedItem struct {
	Title       string
	Description string
	Url         string
}

type Feed struct {
	Id    uint
	Url   string
	Title string
	Items []FeedItem
}

// END OLD

type Database interface {
	CreateUser(username string) error
}

type AppState struct {
	db *sql.DB
}

func NewSqliteAppState(db *sql.DB) *AppState {
	return &AppState{db}
}

func (state *AppState) CreateUser(username string) error {
	println("Creating user: %s", username)
	return nil
}
