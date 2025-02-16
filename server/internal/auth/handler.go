package auth

import (
	"backend/blog/configs"
	"net/http"
)

type AuthHandler struct {
	*AuthService
	config *configs.Configs
}

type AuthHandlerDeps struct {
	*AuthService
	*configs.Configs
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		config:      deps.Configs,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
