package database

const (

	// Suites

	// Spades look like: ♠
	Spades = "spades"
	// Hearts look like: ♥
	Hearts = "hearts"
	// Clubs look like: ♣
	Clubs = "clubs"
	// Diamonds look like: ♦
	Diamonds = "diamonds"

	// Phases

	// PhasePass means players are passing cards
	PhasePass = "pass"

	// PhasePlay means players are playing into the trick
	PhasePlay = "play"

	// Directions

	// PassLeft means we pass left
	PassLeft = "left"

	// PassRight means we pass right
	PassRight = "right"
)

// Card represents a card
type Card struct {
	Suit  string `json:"suit"`
	Value uint8  `json:"value"`
}

// Player represents a player as it is saved in a GameRecord.
type Player struct {
	Hand   []Card `json:"hand"`
	Active []Card `json:"active"`
	Score  int    `json:"score"`
}

// GameRecord represents a game as it exists in the "database."
type GameRecord struct {
	ID            int       `json:"id"`
	Players       [4]Player `json:"players"`
	PassDirection string
	Phase         string
	Turn          uint8
}

// GameData contains the same data as GameRecord, but does not have an ID set.
type GameData struct {
	Players       [4]Player
	PassDirection string
	Phase         string
	Turn          uint8
}

type gameTable struct {
	counter int
	data    map[int]GameRecord
}

// Store represents the saved state of a game.
type Store struct {
	games gameTable
}
