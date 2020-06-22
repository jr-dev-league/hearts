package database

import (
	"sync"

	"github.com/nolwn/hearts/pkg"
)

var once sync.Once
var (
	database Store
)

// Connection gives you a "connection" to the "database."
func Connection() *Store {

	once.Do(func() {
		database = Store{games: make(map[int]pkg.GameRecord)}

		database.games[0] = pkg.GameRecord{
			ID: 0,
			Players: [4]pkg.Player{
				{
					Hand: []pkg.Card{
						{Suit: pkg.Spades, Value: 10},
						{Suit: pkg.Spades, Value: 12},
						{Suit: pkg.Hearts, Value: 0},
					},
					Play:   false,
					Played: pkg.Card{},
					Pass:   false,
					Passed: []pkg.Card{},
					Score:  100,
				},
				{
					Hand: []pkg.Card{
						{Suit: pkg.Clubs, Value: 2},
						{Suit: pkg.Spades, Value: 1},
						{Suit: pkg.Diamonds, Value: 0},
					},
					Play:   true,
					Played: pkg.Card{},
					Pass:   false,
					Passed: []pkg.Card{},
					Score:  100,
				},
				{
					Hand: []pkg.Card{
						{Suit: pkg.Hearts, Value: 11},
						{Suit: pkg.Hearts, Value: 6},
						{Suit: pkg.Diamonds, Value: 8},
					},
					Play:   false,
					Played: pkg.Card{},
					Pass:   false,
					Passed: []pkg.Card{},
					Score:  100,
				},
				{
					Hand: []pkg.Card{
						{Suit: pkg.Diamonds, Value: 1},
						{Suit: pkg.Diamonds, Value: 9},
						{Suit: pkg.Hearts, Value: 4},
					},
					Play:   false,
					Played: pkg.Card{},
					Pass:   false,
					Passed: []pkg.Card{},
					Score:  100,
				},
			},
		}
	})

	return &database
}

// AddGame adds a game record to a given pkg.
func (s *Store) AddGame(gd pkg.GameData) (ID int) {
	ID = 3
	return ID
}

// Games returns all games. You aren't supposed to call it "getGames" because Go is opinionated about weird shit.
func (s *Store) Games() []pkg.GameRecord {
	games := []pkg.GameRecord{}

	for _, game := range s.games {
		games = append(games, game)
	}

	return games
}

func newStore() Store {
	return Store{games: make(map[int]pkg.GameRecord)}
}
