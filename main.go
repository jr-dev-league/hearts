package main

import (
	"log"
	"net/http"

	"github.com/jr-dev-league/go-router"
)

func main() {
	r := router.Router{}
	r.AddRoute(http.MethodPost, "/api/games", authorizeEndpoint(newGameHandler))
	r.AddRoute(http.MethodGet, "/api/games", authorizeEndpoint(gamesHandler))
	r.AddRoute(http.MethodGet, "/api/games/:id", authorizeEndpoint(gameHandler))
	r.AddRoute(http.MethodPatch, "/api/games/:id", authorizeEndpoint(playCardHandler))

	log.Fatal(http.ListenAndServe(":3000", r))
}
