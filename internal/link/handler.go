package link

import (
	"fmt"
	"link-shortener/pkg/req"
	"link-shortener/pkg/res"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	LinkService    *LinkService
}
type LinkHandler struct {
	LinkRepository *LinkRepository
	LinkService    *LinkService
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := LinkHandler{deps.LinkRepository, deps.LinkService}
	router.HandleFunc("GET /{hash}", linkHandler.GoTo())
	router.HandleFunc("DELETE /link/{id}", linkHandler.Delete())
	router.HandleFunc("PATCH /link/{id}", linkHandler.Update())
	router.HandleFunc("POST  /link", linkHandler.Create())

}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}

		link, err := handler.LinkService.AddLink(body.URL)

		//TODO think about better err
		if err != nil {
			http.Error(w, "invalid request or link already exists", http.StatusBadRequest)
			return
		}

		res.JSON(w, link, http.StatusCreated)

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
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)

		if err != nil {
			http.Error(w, "invalid request or no such link", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}
