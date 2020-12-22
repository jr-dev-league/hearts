package main

import (
	"github.com/jr-dev-league/hearts/database"
	"github.com/jr-dev-league/hearts/engine"
)

func rtsPoints(g database.GameRecord) (points uint8, cards uint8) {
	points = 26
	cards = 14
	players := g.Players

	for _, palyer := range players {
		for _, card := range palyer.Hand {
			if card.Suit == engine.Hearts {
				cards--
				points--
			} else if card.Suit == engine.Spades && card.Value == 11 {
				cards--
				points -= 13
			}
		}
	}

	return
}

func rtsBroken(points uint8, cards uint8) bool {
	var max uint8

	// Is the queen taken?
	if points/(cards-1) != 2 { // yes
		max = 13

	} else { //no
		max = 26
	}

	return points < max
}

func rtsShootable(players [4]database.Player) bool {
	withPoints := 0

	for _, player := range players {
		if player.Round > 0 {
			withPoints++
		}
	}

	return withPoints > 1
}

func rtsPlayers(g database.GameRecord) (p [4]engine.Player) {
	playing := g.Phase == database.PhasePlay
	pushCard := func(p engine.Player, card database.Card, active bool) {
		sCard := engine.Card{
			Suit:    card.Suit,
			Value:   card.Value,
			Exposed: playing && active,
			Played:  active,
		}

		p.Hand = append(p.Hand, sCard)
	}

	for i, player := range g.Players {
		for _, card := range player.Hand {
			pushCard(p[i], card, false)
		}

		for _, card := range player.Active {
			pushCard(p[i], card, true)
		}
	}

	return
}

func strPlayers(g engine.State) (p [4]database.Player) {
	for i, player := range g.Players {
		var hand []database.Card
		var active []database.Card
		var round uint8 = player.Points

		for _, card := range player.Hand {
			rCard := database.Card{
				Suit:  card.Suit,
				Value: card.Value,
			}

			if card.Played {
				active = append(active, rCard)
			} else {
				hand = append(hand, rCard)
			}
		}

		p[i] = database.Player{
			Hand:   hand,
			Active: active,
			Total:  player.Score,
			Round:  round,
		}
	}

	return
}

func toState(g database.GameRecord) (s engine.State) {
	points, cards := rtsPoints(g)
	s.Broken = rtsBroken(points, cards)
	s.Turn = g.Turn
	s.Players = rtsPlayers(g)
	s.Shootable = rtsShootable(g.Players)
	s.Readonly = false

	return
}

func toRecord(g engine.State, ID int) (r database.GameRecord) {
	r.Players = strPlayers(g)
	r.Turn = g.Turn
	r.PassDirection = g.PassDirection
	r.Phase = g.Phase

	return
}
