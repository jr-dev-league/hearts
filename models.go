package main

import (
	"github.com/jr-dev-league/hearts/database"
	"github.com/jr-dev-league/hearts/engine"
)

// initialScore is the number of points to count down from. When a player has less than 0 at the
// end of a round, the game ends.
const initialScore = 100

// initialDirection is the direction that players pass their cards during the first round.
const initialDirection = database.PassLeft

// createGame creates a new game with the default initial values.
func createGame() (gameRecord database.GameRecord) {
	var playerRecords [4]database.Player
	var turn uint8

	// db represents whatever the db is that we're using. For v1, the database is actually just a
	// simple in map in memory. It is standing in for a real databse in v2.
	db := database.Connection()
	gameState := engine.New()
	gameState.Deal()

	// loop over the game players...
	for i := 0; i < len(gameState.Players); i++ {
		player := gameState.Players[i]
		playerRecord, has2Cbs := recordPlayer(&player, initialScore)
		playerRecords[i] = playerRecord

		// has2Cbs indicates that the player has the two of clubs, and goes first.
		if has2Cbs {
			turn = uint8(i) // make it that player's turn
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
//
// player is a player object as understood by the engine. Score is the overall player score.
//
// newPlayer is the record to be stored in the database. has2Cbs is a flag that signals that the
// player has the two of clubs.
func recordPlayer(player *engine.Player, score uint8) (newPlayer database.Player, has2Cbs bool) {
	var hand []database.Card

	// activeCards is ithe card or cards that have been played. They remain with the player until
	// the round has finished.
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
		Total:  score,
		Active: activeCards,
	}

	return
}
