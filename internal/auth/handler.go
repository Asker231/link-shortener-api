package auth

import (
	"fmt"
	"net/http"
)

type AuthLogin struct{}

func (authlogin *AuthLogin) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login User")
	}
}

func NewAuth(router *http.ServeMux) {
	login := &AuthLogin{}
	router.HandleFunc("/login", login.Login())
}
