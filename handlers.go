package main

import (
	"net/http"

	"github.com/nolwn/hearts/database"
)

func games(w http.ResponseWriter, req *http.Request) {
	var store = database.Connection()
	var games = store.Games()

	encodeResponse(w, req, games, 200)
}
