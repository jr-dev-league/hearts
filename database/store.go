package database

import "sync"

var once sync.Once
var (
	database Store
)

// Connection gives you a "connection" to the "database."
func Connection() *Store {

	once.Do(func() {
		database = Store{games: make(map[int]GameRecord)}

		database.games[0] = GameRecord{
			ID: 0,
			Players: [4]Player{
				{
					Hand: []Card{
						{Suit: Spades, Value: 10},
						{Suit: Spades, Value: 12},
						{Suit: Hearts, Value: 0},
					},
					Play:   false,
					Played: Card{},
					Pass:   false,
					Passed: []Card{},
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Clubs, Value: 2},
						{Suit: Spades, Value: 1},
						{Suit: Diamonds, Value: 0},
					},
					Play:   true,
					Played: Card{},
					Pass:   false,
					Passed: []Card{},
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Hearts, Value: 11},
						{Suit: Hearts, Value: 6},
						{Suit: Diamonds, Value: 8},
					},
					Play:   false,
					Played: Card{},
					Pass:   false,
					Passed: []Card{},
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Diamonds, Value: 1},
						{Suit: Diamonds, Value: 9},
						{Suit: Hearts, Value: 4},
					},
					Play:   false,
					Played: Card{},
					Pass:   false,
					Passed: []Card{},
					Score:  100,
				},
			},
		}
	})

	return &database
}

// AddGame adds a game record to a given store
func (s *Store) AddGame(gr GameRecord) {
	s.games[gr.ID] = gr
}

// Games returns all games. You aren't supposed to call it "getGames" because Go is opinionated about weird shit.
func (s *Store) Games() []GameRecord {
	games := []GameRecord{}

	for _, game := range s.games {
		games = append(games, game)
	}

	return games
}

func newStore() Store {
	return Store{games: make(map[int]GameRecord)}
}
