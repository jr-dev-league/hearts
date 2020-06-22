package pkg

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
	Play   bool   `json:"play"`
	Played Card   `json:"played"`
	Pass   bool   `json:"pass"`
	Passed []Card `json:"passed"`
	Score  int    `json:"score"`
}

// GameRecord reprsents a game as it exists in the "database."
type GameRecord struct {
	ID        int       `json:"id"`
	Players   [4]Player `json:"players"`
	Broken    bool      `json:"broken"`
	Readonly  bool      `json:"readonly"`
	Shootable bool      `json:"shootable"`
	TakenLast uint8     `json:"taken"`
}

// GameData represents an initial game structure
type GameData struct {
	Players   [4]Player
	Broken    bool
	Readonly  bool
	Shootable bool
	TakenLast uint8
}
