package server

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/metruzanca/rss/lib/db"
	"github.com/metruzanca/rss/lib/utils"
)

func NewServer() *http.Server {
	sqlDb, _ := sql.Open("sqlite3", "./sqlite.db")
	PORT := utils.Env("PORT", "3000")

	appState := db.NewSqliteAppState(sqlDb)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: RegisterRoutes(appState),
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	return server
}
