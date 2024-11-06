package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Asker231/link-shortener-api.git/configs"
	"github.com/Asker231/link-shortener-api.git/pkg/res"
)

type AuthDep struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func (authlogin *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var LoginRequest LoginRequest

		err := json.NewDecoder(r.Body).Decode(&LoginRequest)
		if err != nil{
			res.JsonRes(w,err.Error(),402)
		}

		fmt.Println(LoginRequest)

		data := LoginResponse{
			Token: authlogin.Secret,
		}
		res.JsonRes(w, data, 200)

	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register")
	}
}

func NewAuth(router *http.ServeMux, deps AuthDep) {
	authHandler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}
