package auth

import (
	"net/http"
	"order-api-start/internal/user"
	"order-api-start/pkg/req"
	"order-api-start/pkg/res"
)

type AuthHandler struct {
	repo    *user.UserRepository
	service *AuthService
}

type Deps struct {
	Repo    *user.UserRepository
	Service *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps Deps) {
	handler := AuthHandler{
		repo:    deps.Repo,
		service: deps.Service,
	}

	router.HandleFunc("POST auth/login", handler.login())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := req.HandleBody[user.User](&w, r)

		if err != nil {
			return
		}

		currentSession, err := handler.service.Login(user)

		if err != nil {
			res.Json(w, err.Error(), 500)
			return
		}

		res.Json(w, LoginResponse{
			SID: currentSession.SID,
		}, 200)
	}
}
