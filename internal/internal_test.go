package internal

import (
	"testing"

	"github.com/nzorkic/deck/deck"
)

func TestCalculatePoints(t *testing.T) {
	hand16 := []deck.Card{}
	hand16 = append(hand16,
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	if CalculatePoints(&hand16) != 16 {
		t.Errorf("Expected card points to be %d, got %d", 16, CalculatePoints(&hand16))
	}
	hand17 := []deck.Card{}
	hand17 = append(hand17,
		deck.Card{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	if CalculatePoints(&hand17) != 17 {
		t.Errorf("Expected card points to be %d, got %d", 17, CalculatePoints(&hand17))
	}
}

func TestSortedCardsCopy(t *testing.T) {
	hand1 := []deck.Card{}
	hand1 = append(hand1,
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
	)
	arr1 := sortedCardsCopy(&hand1)
	if arr1[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
	hand2 := []deck.Card{}
	hand2 = append(hand2,
		deck.Card{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	arr2 := sortedCardsCopy(&hand2)
	if arr2[1].Rank != deck.Ace || arr2[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
}

func TestCardsAsString(t *testing.T) {
	cards := []deck.Card{}
	cards = append(cards, deck.Card{Rank: deck.Ten, Suit: deck.Spade, Visible: true},
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true},
		deck.Card{Rank: deck.Joker, Suit: deck.Spade, Visible: true},
		deck.Card{Rank: deck.Ace, Suit: deck.Diamond, Visible: false})
	expectedString := "Ten of Spades, Queen of Hearts, Joker, FACEDOWN"
	if expectedString != CardsAsString(&cards) {
		t.Errorf("Error transforming cards to strings. Expected (%s), got (%s)", expectedString, CardsAsString(&cards))
	}
}
