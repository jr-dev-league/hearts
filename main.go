package main

import (
	"log"
	"net/http"

	"github.com/jr-dev-league/go-router"
)

func main() {
	r := router.Router{}
	err := r.AddRoute(http.MethodGet, "/api/games", getGamesHandler)
	err = r.AddRoute(http.MethodGet, "/api/games/:id", getGameHandler)
	err = r.AddRoute(http.MethodPost, "/api/games", createGameHandler)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":3000", r))
}
