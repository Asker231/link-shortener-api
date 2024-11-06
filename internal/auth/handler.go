package auth

import (
	"fmt"
	"net/http"
	"github.com/Asker231/link-shortener-api.git/configs"
)

type AuthDep struct{
	*configs.Config
}
type AuthHandler struct{
	*configs.Config
}

func (authlogin *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login User")
		fmt.Println(authlogin.Secret)
	}
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register")
	}
}

func NewAuth(router *http.ServeMux,deps AuthDep){
	authHandler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register",authHandler.Register())
}
