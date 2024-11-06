package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/link-shortener-api.git/configs"
	"github.com/Asker231/link-shortener-api.git/pkg/req"
	"github.com/Asker231/link-shortener-api.git/pkg/res"
)

type AuthDep struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: a.Secret,
		}
		res.JsonRes(w, data, 200)

	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := RegisterResponse{
			Token: a.Secret,
		}
		res.JsonRes(w, data, 200)
	}
}

func NewAuth(router *http.ServeMux, deps AuthDep) {
	authHandler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}
