package main

import (
	"net/http"

	"github.com/Asker231/link-shortener-api.git/internal/auth"
)

func main() {
	app := http.NewServeMux()
	auth.NewAuth(app)
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	server.ListenAndServe()
}
