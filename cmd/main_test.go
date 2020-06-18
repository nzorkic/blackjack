package main

import (
	"testing"

	util "github.com/nzorkic/blackjack/internal"
	"github.com/nzorkic/deck"
)

func TestInitPoints(t *testing.T) {
	newDeck := deck.New()
	initPoints(&newDeck)
	for _, card := range newDeck {
		if card.Rank == deck.Jack || card.Rank == deck.Queen || card.Rank == deck.King {
			if card.Point != 10 {
				t.Errorf("Points for  %s are %d, should be %d", card.Rank, card.Point, 10)
			}
		}
	}
}

func TestCreatePlayers(t *testing.T) {
	n := 3
	players := createPlayers(&n)
	if len(players) != n {
		t.Errorf("Error while creating players. Expected %d, got %d", n, len(players))
	}
}

func TestDeal(t *testing.T) {
	nOfPlayers := 5
	nOfCardsDealt := 2
	players := createPlayers(&nOfPlayers)
	deck := deck.New(deck.Shuffle())
	deal(&players, &deck)
	if dealer.cards[1].Visible {
		t.Errorf("Dealers last cart is not face down")
	}
	for _, player := range players {
		if len(player.cards) != nOfCardsDealt {
			t.Errorf("Player with the name %s was dealt %d cards instead of %d",
				player.name, len(player.cards), nOfCardsDealt)
		}
		for _, card := range player.cards {
			if !card.Visible {
				t.Errorf("Player with the name %s has face down card. Card should be visible",
					player.name)
			}
		}
		if results[player.name] != util.CalculatePoints(&player.cards) {
			t.Errorf("Final result after dealing is not correct. Expected %d, got %d",
				util.CalculatePoints(&player.cards), results[player.name])
		}
	}
	if len(dealer.cards) != nOfCardsDealt {
		t.Errorf("Player with the name %s was dealt %d cards instead of %d",
			dealer.name, len(dealer.cards), nOfCardsDealt)
	}
	if results[dealer.name] != util.CalculatePoints(&dealer.cards) {
		t.Errorf("Final result after dealing is not correct. Expected %d, got %d",
			util.CalculatePoints(&dealer.cards), results[dealer.name])
	}
}
