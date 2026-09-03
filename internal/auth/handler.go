package auth

import (
	"link-shortener/configs"
	"link-shortener/internal/user"
	"link-shortener/pkg/req"
	"link-shortener/pkg/res"

	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*user.UserService
}
type AuthHandler struct {
	*configs.Config
	*user.UserService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := AuthHandler{Config: deps.Config, UserService: deps.UserService}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			res.JSON(w, "login error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		handler.UserService.Login(body.Email)
		resp := LoginResponse{}
		res.JSON(w, resp, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.JSON(w, "request error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		_, err = handler.UserService.Register(body.Email, body.Password, body.Username)

		if err != nil {
			res.JSON(w, "register error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		resp := RegisterResponse{Token: "67"}
		res.JSON(w, resp, http.StatusOK)

	}
}
