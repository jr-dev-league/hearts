package database

import "github.com/nolwn/hearts/pkg"

// Store represents the saved state of a game.
type Store struct {
	games map[int]pkg.GameRecord
}
