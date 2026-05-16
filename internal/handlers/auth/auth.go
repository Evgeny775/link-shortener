package auth

import (
	"fmt"
	"link-shortener/configs"
	"link-shortener/internal/payload"
	"link-shortener/pkg/req"
	"link-shortener/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := AuthHandler{Config: deps.Config}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[payload.LoginRequest](w, r)
		if err != nil {
			res.JSON(w, "login error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(body)
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
		data := payload.LoginResponse{Token: "67"}
		res.JSON(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[payload.RegisterRequest](w, r)
		if err != nil {
			res.JSON(w, "register error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(body)
		fmt.Println("Register")
		data := payload.RegisterResponse{Token: "69"}
		res.JSON(w, data, http.StatusOK)
	}
}
