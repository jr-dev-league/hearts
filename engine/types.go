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
	Exposed bool
	Played  bool
	Suit    string
	Value   uint8
}

// A Player represents the hand and round score of a player
type Player struct {
	CardCount uint8
	Hand      []Card
	Points    int8
}

// A State represents the complete game state
type State struct {
	Broken    bool
	Players   [4]Player
	Readonly  bool
	Shootable bool
	TakenLast uint8
}
