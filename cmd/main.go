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
	dataBase := db.NewDB(conf)
	router := http.NewServeMux()

	//repos
	linkRepository := link.NewLinkRepository(dataBase)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
