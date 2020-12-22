package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jr-dev-league/go-router"
	"github.com/jr-dev-league/hearts/database"
	"github.com/jr-dev-league/hearts/engine"
)

type errorResponse struct {
	statusCode int
	message    string
}

type gameResponse struct {
	ID int `json:"ID"`
}

func newGameHandler(w http.ResponseWriter, req *http.Request) {
	newGame := createGame()
	resBody := gameResponse{ID: newGame.ID}

	writeResponse(w, req, resBody, http.StatusCreated)
}

func gamesHandler(w http.ResponseWriter, req *http.Request) {
	db := database.Connection()
	games := db.Games()

	writeResponse(w, req, games, http.StatusOK)
}

func gameHandler(w http.ResponseWriter, req *http.Request) {
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

func playCardHandler(w http.ResponseWriter, req *http.Request) {
	user := getUser(req)
	params := router.PathParams(req)
	cardInput := req.Body
	gameID := params["id"]

	var card engine.Card
	cardDecoder := json.NewDecoder(cardInput)
	err := cardDecoder.Decode(&card)

	if err != nil {
		writeResponse(w, req, resError{"Bad Request. No body found."}, 400)

		return
	}

	db := database.Connection()
	ID, err := strconv.ParseInt(gameID, 0, 0)

	if err != nil {
		writeResponse(w, req, resError{"Not Found."}, 404)

		return
	}

	game, err := db.Game(int(ID))

	state := toState(game)
	state.PlayUp(uint8(user.ID), card)

	db.UpdateGame(toRecord(state, game.ID))
}
