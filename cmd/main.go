package main

import (
	"net/http"

	"github.com/Asker231/link-shortener-api.git/configs"
	"github.com/Asker231/link-shortener-api.git/internal/auth"
)

func main() {
	app := http.NewServeMux()
    conf := configs.LoadConfig()
	auth.NewAuth(app,auth.AuthDep{
		Config: conf,
	})
	server := http.Server{
		Addr:   conf.Port,
		Handler: app,
	}

	server.ListenAndServe()
}
