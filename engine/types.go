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
	exposed bool
	played  bool
	suit    string
	value   uint8
}

// A Player represents the hand and round score of a player
type Player struct {
	cardCount   uint8
	hand        []Card
	points      int8
	playedCard  Card
	passedCards [3]Card
}

// A State represents the complete game state
type State struct {
	broken     bool
	players    [4]Player
	readonly   bool
	shootable  bool
	takenLast  uint8
	handNumber uint8
}

// TODO put played card, passed cards, and hand number into tests
