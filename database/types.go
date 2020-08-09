package database

const (
	// Spades look like: ♠
	Spades = "spades"
	// Hearts look like: ♥
	Hearts = "hearts"
	// Clubs look like: ♣
	Clubs = "clubs"
	// Diamonds look like: ♦
	Diamonds = "diamonds"
)

// Card represents a card
type Card struct {
	Suit  string `json:"suit"`
	Value int8   `json:"value"`
}

// Player represents a player as it is saved in a GameRecord.
type Player struct {
	Hand   []Card `json:"hand"`
	Played Card   `json:"played"`
	Passed []Card `json:"passed"`
	Score  int    `json:"score"`
}

// GameRecord represents a game as it exists in the "database."
type GameRecord struct {
	ID      int       `json:"id"`
	Players [4]Player `json:"players"`
	Type    string    `json:"type"`
	Turn    int       `json:"turn"`
}

// GameData contains the same data as GameRecord, but does not have an ID set.
type GameData struct {
	Players [4]Player
	Type    string
	Turn    int
}

type gameTable struct {
	counter int
	data    map[int]GameRecord
}

// Store represents the saved state of a game.
type Store struct {
	games gameTable
}
