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

	// Exposed is true if a card is visible to other players. A card that is played is not
	// necessarily exposed. A card can be played face down, for instance when it is being
	// passed to another player.
	Exposed bool

	// Played is a card that has been selected by a player to be used in whatever way the phase
	// of the turn requires. If players are playing into the trick, then one card is played
	// exposed. If the players are passing, then three cards are simultaneously played unexposed.
	Played bool

	// Suit looks like ♠, ♥, ♣ or ♦.
	Suit string

	// Value is a number from 0 to 12 that represents the cards value. Ace is 0, King is 12.
	Value uint8
}

// A Player represents the hand and round score of a player
type Player struct {
	CardCount uint8

	// Hand is the actual cards in a players hand. Obviously, if this information should only be
	// given to the client who controls those cards. For infomation about other players hands, use
	// CardCount.
	Hand []Card

	// Points is the number of points the player has taken in the round.
	Points int8
}

// A State represents the complete game state
type State struct {

	// Broken is a flag that indicates that a heart has been taken in a trick.
	Broken bool

	// Players is an array of four Players, the four people playing the game.
	Players [4]Player

	// Readonly is...
	Readonly bool

	// Shootable is a flag that indicates that no more than one player holds any point.
	Shootable bool

	// TakenLast is an array index from 0–3 that indicates which player was the last to take a trick.
	TakenLast uint8
}
