package internal

import (
	"testing"

	"github.com/nzorkic/deck"
)

func TestSortedCardsCopy(t *testing.T) {
	hand1 := []deck.Card{}
	hand1 = append(hand1,
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
	)
	arr1 := SortedCardsCopy(&hand1)
	if arr1[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
	hand2 := []deck.Card{}
	hand2 = append(hand2,
		deck.Card{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	arr2 := SortedCardsCopy(&hand2)
	if arr2[1].Rank != deck.Ace || arr2[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
}
