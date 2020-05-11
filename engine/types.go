package engine

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

const maxHandSize = 13

// A Card represents a playing card.
type Card struct {
	suit    string
	value   uint8
	exposed bool
	played  bool
}

// A Player represents the hand and round score of a player
type Player struct {
	points    int8
	cardCount uint8
	hand      []Card
}

// A State represents the complete game state
type State struct {
	broken    bool
	shootable bool
	takenLast uint8
	players   [4]Player
}
