package database

import (
	"errors"
	"sync"
)

var once sync.Once
var database Store
var (
	gamesTable = "games"
)

// Connection gives you a "connection" to the "database."
func Connection() *Store {

	once.Do(func() {
		database = Store{games: gameTable{counter: 1, data: make(map[int]GameRecord)}}

		database.games.data[0] = GameRecord{
			ID: 0,
			Players: [4]Player{
				{
					Hand: []Card{
						{Suit: Spades, Value: 10},
						{Suit: Spades, Value: 12},
						{Suit: Hearts, Value: 0},
					},
					Active: []Card{},
					Total:  100,
					Round:  16,
				},
				{
					Hand: []Card{
						{Suit: Clubs, Value: 2},
						{Suit: Spades, Value: 1},
						{Suit: Diamonds, Value: 0},
					},
					Active: []Card{},
					Total:  100,
					Round:  4,
				},
				{
					Hand: []Card{
						{Suit: Hearts, Value: 11},
						{Suit: Hearts, Value: 6},
						{Suit: Diamonds, Value: 8},
					},
					Active: []Card{},
					Total:  100,
					Round:  0,
				},
				{
					Hand: []Card{
						{Suit: Diamonds, Value: 1},
						{Suit: Diamonds, Value: 9},
						{Suit: Hearts, Value: 4},
					},
					Active: []Card{},
					Total:  100,
					Round:  2,
				},
			},
			Phase:         "play",
			Turn:          1,
			PassDirection: "left",
		}
	})

	return &database
}

// AddGame adds a game record to the store
func (s *Store) AddGame(data GameData) (record GameRecord) {
	record = GameRecord{
		ID:            generateID(gamesTable),
		Players:       data.Players,
		PassDirection: data.PassDirection,
		Phase:         data.Phase,
	}

	s.games.data[record.ID] = record

	return
}

// Games returns all games. You aren't supposed to call it "getGames" because Go is opinionated about weird shit.
func (s *Store) Games() (games []GameRecord) {
	games = []GameRecord{}

	for _, game := range s.games.data {
		games = append(games, game)
	}

	return games
}

// Game takes a game id and returns the corresponding game record.
func (s *Store) Game(ID int) (game GameRecord, err error) {
	err = errors.New("not found")

	for _, record := range s.games.data {
		if record.ID == ID {
			game = record
			err = nil
		}
	}

	return
}

// UpdateGame takes a GameRecord, finds the record with the same ID and replaces it.
// If the game is not found, an error is returned.
func (s *Store) UpdateGame(game GameRecord) (err error) {
	err = errors.New("not found")

	for i, record := range s.games.data {
		if record.ID == game.ID {
			s.games.data[i] = game
			err = nil
		}
	}

	return
}

func generateID(table string) (ID int) {
	switch table {
	case "games":
		return generateGameID()
	}
	return 0
}

func generateGameID() (ID int) {
	ID = database.games.counter
	database.games.counter++

	return
}
