package server

import (
	"fmt"
	"net/http"

	"github.com/metruzanca/rss/lib/utils"
)

func NewServer() *http.Server {
	PORT := utils.Env("PORT", "3000")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: RegisterRoutes(),
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	return server
}
