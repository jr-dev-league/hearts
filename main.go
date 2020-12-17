package main

import (
	"log"
	"net/http"

	"github.com/jr-dev-league/go-router"
)

func main() {
	r := router.Router{}
	r.AddRoute(http.MethodPost, "/api/games", authorizeEndpoint(createGameHandler))
	r.AddRoute(http.MethodGet, "/api/games", authorizeEndpoint(getGamesHandler))
	r.AddRoute(http.MethodGet, "/api/games/:id", authorizeEndpoint(getGameHandler))
	// r.AddRoute(http.MethodPatch, "/api/games/:id", authorizeEndpoint(playCard))

	log.Fatal(http.ListenAndServe(":3000", r))
}
