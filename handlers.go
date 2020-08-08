package main

import (
	"net/http"

	"github.com/jr-dev-league/hearts/database"
)

func games(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getGames(w, req)
		// case http.MethodPost:
		// 	createGame(w, req)
	}
}

func getGames(w http.ResponseWriter, req *http.Request) {
	store := database.Connection()
	games := store.Games()

	writeResponse(w, req, games, 200)
}

// func createGame(w http.ResponseWriter, req *http.Request) {
// 	store := database.Connection()
// 	newGame :=
// }
