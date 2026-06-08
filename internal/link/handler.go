package link

import (
	"fmt"
	"link-shortener/configs"
	"net/http"
)

type LinkHandlerDeps struct {
	*configs.Config
}
type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := LinkHandler{deps.Config}
	router.HandleFunc("GET /{hash}", linkHandler.GoTo())
	router.HandleFunc("DELETE /link/{id}", linkHandler.Delete())
	router.HandleFunc("PATCH /link/{id}", linkHandler.Update())
	router.HandleFunc("POST  /link", linkHandler.Create())

}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
