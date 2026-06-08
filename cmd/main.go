package main

import (
	"link-shortener/configs"
	"link-shortener/internal/auth"
	"link-shortener/internal/link"
	"link-shortener/pkg/db"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	_ = db.NewDB(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
