package main

import (
	"net/http"

	"github.com/Asker231/link-shortener-api.git/configs"
	"github.com/Asker231/link-shortener-api.git/internal/auth"
)

func main() {
	app := http.NewServeMux()
	auth.NewAuth(app)
    conf := configs.LoadConfig()
	server := http.Server{
		Addr:   conf.Port,
		Handler: app,
	}

	server.ListenAndServe()
}
