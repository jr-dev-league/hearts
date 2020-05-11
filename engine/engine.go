package engine

import "errors"

// New creates a new game status
func New() State {
	game := State{
		broken:    false,
		players:   [4]Player{},
		takenLast: 4,
		shootable: true,
	}

	return game
}

// SetPlayer sets a given (by index) player's hand to a given array of cards.
func (game *State) SetPlayer(i uint8, points int8, hand []Card) error {
	if game.players[i].hand != nil {
		return errors.New("this player has already been set")
	}

	game.players[i] = Player{points, uint8(len(hand)), hand}

	return nil
}

// Player returns a card array representing the given (by index) player's hand
func (game *State) Player(i uint8) (*Player, error) {
	hand := game.players[i].hand

	if hand == nil {
		return nil, errors.New("this player had not been set")
	}

	return &game.players[i], nil
}

// ViewAs returns a game state as known by a given (by index) player
func (game *State) ViewAs(p uint8) State {
	view := New()

	for i := uint8(0); i < uint8(len(game.players)); i++ {
		player := &game.players[i]
		if i == p {
			view.SetPlayer(i, player.points, player.hand)
		} else {
			hand := make([]Card, 0, maxHandSize)

			for c := range game.players[i].hand {
				card := game.players[i].hand[c]

				if card.played {
					hand = append(hand, card)
				}
			}

			view.SetPlayer(i, player.points, hand)
		}
	}

	return view
}

// PlayUp plays one card face up by player index, p, and card index, c.
// Because played cards are still owned by the player that played them,
// there is no need to remove it from their hand. It should be up to the
// view to display the card in the middle of the table.
func (game *State) PlayUp(p uint8, c Card) error {
	player := game.players[p]
	i, err := findCard(&player.hand, c)

	if err != nil {
		return err
	}

	if player.hand[i].played {
		return errors.New("selected card was already played")
	}

	player.hand[i].played = true
	player.hand[i].exposed = true
	player.cardCount--

	return nil
}

// PRIVATE HELPER FUNCTIONS

// findCard searches a given hand for a card. Returns a reference to the card
// if found, and an error if not.
func findCard(stack *[]Card, target Card) (int, error) {
	for i, card := range *stack {
		if target.suit == card.suit &&
			target.value == card.value {
			return i, nil
		}
	}

	return -1, errors.New("could not find target card")
}
