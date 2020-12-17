package main

import (
	"net/http"
	"strconv"

	"github.com/jr-dev-league/go-router"
	"github.com/jr-dev-league/hearts/database"
)

type errorResponse struct {
	statusCode int
	message    string
}

type gameResponse struct {
	ID int `json:"ID"`
}

func createGameHandler(w http.ResponseWriter, req *http.Request) {
	newGame, err := createGame()
	resBody := gameResponse{ID: newGame.ID}

	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Internal Server Error."
		errBody := errorResponse{statusCode, message}
		writeResponse(w, req, errBody, http.StatusInternalServerError)

		return
	}

	writeResponse(w, req, resBody, http.StatusCreated)
}

func getGamesHandler(w http.ResponseWriter, req *http.Request) {
	db := database.Connection()
	games := db.Games()

	writeResponse(w, req, games, http.StatusOK)
}

func getGameHandler(w http.ResponseWriter, req *http.Request) {
	params := router.PathParams(req)
	gameIDParam := params["id"]
	gameID, err := strconv.ParseInt(gameIDParam, 10, 0)

	if err != nil {
		writeResponse(w, req, nil, http.StatusBadRequest)

		return
	}

	db := database.Connection()
	game, err := db.Game(int(gameID))

	if err != nil {
		writeResponse(w, req, nil, 404)

		return
	}

	writeResponse(w, req, game, http.StatusOK)
}

// func playCard(w http.ResponseWriter, req *http.Request) {

// }
