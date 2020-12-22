package engine

import (
	"errors"
)

// New creates a new game status
func New() (game State) {
	game = State{
		Broken:    false,
		Players:   [4]Player{},
		Turn:      4,
		Shootable: true,
		Readonly:  false,
	}

	return
}

// SetPlayer sets a player at the given index. It takes and index, i, which is the player index;
// it takes points, which is the number of points the plater has; t takes cardCount which is
// the number of cards a player has; and it takes hand which are the cards at the player.
func (game *State) SetPlayer(i uint8, points uint8, score uint8, cardCount uint8, hand []Card) error {
	if game.Readonly {
		return errors.New("cannot edit a readonly game")
	}

	if game.Players[i].Hand != nil {
		return errors.New("this player has already been set")
	}

	game.Players[i] = Player{cardCount, hand, points, score}

	return nil
}

// Player returns a player by the given index, i.
func (game *State) Player(i uint8) (*Player, error) {
	hand := game.Players[i].Hand

	// TODO: is this a useful error? Is this a useful check?
	if hand == nil {
		return nil, errors.New("this player had not been set")
	}

	return &game.Players[i], nil
}

// ViewAs returns a game state as known by a given (by index) player
func (game *State) ViewAs(p uint8) (view State) {
	view = New()

	for i := uint8(0); i < uint8(len(game.Players)); i++ {
		player := &game.Players[i]
		if i == p { // if we are viewing from the given player...
			view.SetPlayer(i, player.Points, player.Score, maxHandSize, player.Hand)
		} else { // if we are viewing an opponent...
			hand := make([]Card, 0, maxHandSize)

			for c := range game.Players[i].Hand {
				card := game.Players[i].Hand[c]

				if card.Exposed { // only show exposed cards
					hand = append(hand, card)
				}
			}

			view.SetPlayer(i, player.Points, player.Score, maxHandSize, hand)
		}
	}

	view.Readonly = true // TODO: I still don't understand why this is useful

	return
}

// PlayUp plays one card face up by player index, p, and card index, c.
// Because played cards are still owned by the player that played them,
// there is no need to remove it from their hand. It should be up to the
// client to display the card in the middle of the table.
//
// Returns an error if the state is read only, or if the selected card has already been
// played.
func (game *State) PlayUp(p uint8, c Card) error {
	if game.Readonly { // Why would the engine try to play on a game it can't play?
		return errors.New("cannot edit a readonly game")
	}

	player := game.Players[p]
	i, err := findCard(player.Hand, c)

	if err != nil {
		return err
	}

	if player.Hand[i].Played {
		return errors.New("selected card was already played")
	}

	player.Hand[i].Played = true
	player.Hand[i].Exposed = true
	player.CardCount--

	return nil
}

// Deal shuffles a deck of cards and deals it out to each player. returns an error if cards
// have already been dealt.
func (game *State) Deal() error {
	deck := stdDeck()

	shuffle(deck)

	for i := 0; i < len(game.Players); i++ {
		player := &game.Players[i]
		if len(player.Hand) > 0 {
			return errors.New("one or more players already has cards")
		}

		hand, err := dealHand(13, &deck)

		if err != nil {
			return err
		}

		player.Hand = hand
	}

	return nil
}

// Discard deletes a card from the hand of the player p
func (game *State) Discard(p uint8, card Card) error {
	pl := &game.Players[p]
	i, err := findCard(pl.Hand, card)

	if err != nil {
		return errors.New("engine.Game.Discard: card not found")
	}

	if i != -1 {
		pl.discard(i, card)
	}

	return err
}

// DiscardPlayed played deletes all cards from the game that have been played,
// and returns a slice of the cards deleted in this way
func (game *State) DiscardPlayed() (stack []Card) {
	for p := range game.Players {
		pl := &game.Players[p]
		for i, card := range pl.Hand {
			if card.Played == true {
				pl.discard(i, card)
				stack = append(stack, card)
			}
		}
	}

	return
}

// PRIVATE HELPER FUNCTIONS

// discard (lowercase) is a helper method to call on a player to remove a
// card at index i in their hand
func (player *Player) discard(i int, card Card) {
	player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
	player.CardCount--
}

// findCard searches a given hand for a card. Returns the card index if found
// and an error if not.
func findCard(stack []Card, target Card) (int, error) {
	for i, card := range stack {
		if target.Suit == card.Suit &&
			target.Value == card.Value {
			return i, nil
		}
	}

	return -1, errors.New("could not find target card")
}
