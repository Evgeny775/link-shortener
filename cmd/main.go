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

	
	linkRepository := link.NewLinkRepository(dataBase)
	linkService := link.NewLinkService(linkRepository)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository, LinkService: linkService})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
