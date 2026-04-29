package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"link-shortener/configs"
	"link-shortener/internal/payload"
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

		var payloadRequest payload.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payloadRequest)
		if err != nil {
			res.JSON(w, "Error encoding response body", 402)
			return
		}

		validate := validator.New()

		fmt.Println(payloadRequest)

		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
		data := payload.LoginResponse{Token: "67"}
		res.JSON(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
