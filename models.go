package main

import (
	"github.com/jr-dev-league/hearts/database"
	"github.com/jr-dev-league/hearts/engine"
)

const initialScore = 100
const initialDirection = database.PassLeft

func createGame() (gameRecord database.GameRecord, err error) {
	var playerRecords [4]database.Player
	var turn uint8

	db := database.Connection()
	gameState := engine.New()
	gameState.Deal()

	for i := 0; i < len(gameState.Players); i++ {
		player := gameState.Players[i]
		playerRecord, has2Cbs := recordPlayer(&player, initialScore)
		playerRecords[i] = playerRecord

		if has2Cbs {
			turn = uint8(i)
		}
	}

	gameData := database.GameData{
		Players:       playerRecords,
		PassDirection: initialDirection,
		Phase:         database.PhasePass,
		Turn:          turn,
	}

	gameRecord = db.AddGame(gameData)

	return
}

// recordPlayer creates a player record from a player state. Because a players overall score
// cannot be determined form the state alone, a score needs to be passed in.
func recordPlayer(player *engine.Player, score int) (newPlayer database.Player, has2Cbs bool) {
	var hand []database.Card
	var activeCards []database.Card

	for i := 0; i < len(player.Hand); i++ {
		cardState := player.Hand[i]
		card := database.Card{Suit: cardState.Suit, Value: cardState.Value}

		if cardState.Suit == database.Clubs && cardState.Value == 1 {
			has2Cbs = true
		}

		if cardState.Played {
			activeCards = append(activeCards, card)
		} else {
			hand = append(hand, card)
		}
	}

	newPlayer = database.Player{
		Hand:   hand,
		Score:  score,
		Active: activeCards,
	}

	return
}
