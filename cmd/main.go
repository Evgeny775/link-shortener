package main

import (
	"link-shortener/configs"
	"link-shortener/internal/auth"
	"link-shortener/internal/link"
	"link-shortener/internal/middleware"
	"link-shortener/pkg/db"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	dataBase := db.NewDB(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(dataBase)
	linkService := link.NewLinkService(linkRepository)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkService: linkService})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	server.ListenAndServe()
}
