package main

import (
	"net/http"

	"github.com/nolwn/hearts/database"
)

func gameHandlers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		games(w, req)
	case http.MethodPost:
		createGame(w, req)
	}
}

func games(w http.ResponseWriter, req *http.Request) {
	var store = database.Connection()
	var games = store.Games()

	encodeResponse(w, req, games, 200)
}

func createGame(w http.ResponseWriter, req *http.Request) {
	newGame()
}
