package game

import (
	"testing"

	"github.com/nzorkic/deck"
)

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
		if player.points != calculatePoints(&player.cards) {
			t.Errorf("Final result after dealing is not correct. Expected %d, got %d",
				calculatePoints(&player.cards), results[player.name])
		}
	}
	if len(dealer.cards) != nOfCardsDealt {
		t.Errorf("Player with the name %s was dealt %d cards instead of %d",
			dealer.name, len(dealer.cards), nOfCardsDealt)
	}
	if dealer.points != calculatePoints(&dealer.cards) {
		t.Errorf("Final result after dealing is not correct. Expected %d, got %d",
			calculatePoints(&dealer.cards), dealer.points)
	}
}
