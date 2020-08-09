package main

import (
	"net/http"

	"github.com/jr-dev-league/hearts/database"
)

type gameResponse struct {
	ID int "json:id"
}

type errorResponse struct {
	statusCode int
	message    string
}

func gamesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getGamesHandler(w, req)
		// case http.MethodPost:
		// 	createGame(w, req)
	}
}

func getGamesHandler(w http.ResponseWriter, req *http.Request) {
	db := database.Connection()
	games := db.Games()

	writeResponse(w, req, games, http.StatusOK)
}

func createGameHandler(w http.ResponseWriter, req *http.Request) {
	// newGame, err := createGame()
	// resBody := gameResponse{ID: newGame.ID}

	// if err != nil {
	// 	statusCode := http.StatusInternalServerError
	// 	message := "Internal Server Error."
	// 	errBody := errorResponse{statusCode, message}
	// 	writeResponse(w, req, errBody, http.StatusInternalServerError)
	// }

	// writeResponse(w, req, resBody, http.StatusCreated)
}
