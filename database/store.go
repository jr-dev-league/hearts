package database

import "sync"

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
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Clubs, Value: 2},
						{Suit: Spades, Value: 1},
						{Suit: Diamonds, Value: 0},
					},
					Active: []Card{},
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Hearts, Value: 11},
						{Suit: Hearts, Value: 6},
						{Suit: Diamonds, Value: 8},
					},
					Active: []Card{},
					Score:  100,
				},
				{
					Hand: []Card{
						{Suit: Diamonds, Value: 1},
						{Suit: Diamonds, Value: 9},
						{Suit: Hearts, Value: 4},
					},
					Active: []Card{},
					Score:  100,
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

	return record
}

// Games returns all games. You aren't supposed to call it "getGames" because Go is opinionated about weird shit.
func (s *Store) Games() []GameRecord {
	games := []GameRecord{}

	for _, game := range s.games.data {
		games = append(games, game)
	}

	return games
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
