package auth

import (
	"errors"
	"link-shortener/configs"
	"link-shortener/internal/user"
	"link-shortener/pkg/jwt"
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
			res.JSON(w, "request error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		logUser, err := handler.UserService.Login(body.Email, body.Password)

		if errors.Is(err, user.WrongCredentials) {
			res.JSON(w, "wrong email or password", http.StatusUnauthorized)
			return
		}

		if err != nil {
			res.JSON(w, "login error"+err.Error(), http.StatusUnauthorized)
			return
		}

		logJWT := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := logJWT.Create(logUser.Email)

		if err != nil {
			res.JSON(w, "jwt generation error"+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{Token: token}
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

		regUser, err := handler.UserService.Register(body.Email, body.Password, body.Username)

		if errors.Is(err, user.AlreadyExists) {
			res.JSON(w, "user already exists", http.StatusUnauthorized)
			return
		}

		if err != nil {
			res.JSON(w, "register error: "+err.Error(), http.StatusUnauthorized)
			return
		}

		regJWT := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := regJWT.Create(regUser.Email)

		if err != nil {
			res.JSON(w, "jwt generation error"+err.Error(), http.StatusUnauthorized)
			return
		}

		resp := RegisterResponse{Token: token}
		res.JSON(w, resp, http.StatusOK)

	}
}
