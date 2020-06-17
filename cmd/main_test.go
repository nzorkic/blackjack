package main

import (
	"testing"

	"github.com/nzorkic/deck"
)

func TestCardsAsString(t *testing.T) {
	cards := []deck.Card{}
	cards = append(cards, deck.Card{Rank: deck.Ten, Suit: deck.Spade, Visible: true},
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true},
		deck.Card{Rank: deck.Joker, Suit: deck.Spade, Visible: true},
		deck.Card{Rank: deck.Ace, Suit: deck.Diamond, Visible: false})
	expectedString := "Ten of Spades, Queen of Hearts, Joker, FACEDOWN"
	if expectedString != cardsAsString(&cards) {
		t.Errorf("Error transforming cards to strings. Expected (%s), got (%s)", expectedString, cardsAsString(&cards))
	}
}

func TestCreatePlayers(t *testing.T) {
	n := 3
	players := createPlayers(n)
	if len(players) != n {
		t.Errorf("Error while creating players. Expected %d, got %d", n, len(players))
	}
}

func TestDeal(t *testing.T) {
	nOfPlayers := 5
	nOfCardsDealt := 2
	players := createPlayers(nOfPlayers)
	deck := deck.New(deck.Shuffle())
	deal(&players, &deck)
	dealer := players[nOfPlayers-1]
	if dealer.cards[len(dealer.cards)-1].Visible {
		t.Errorf("Dealers last cart is not face down")
	}
	for _, player := range players {
		if len(player.cards) != nOfCardsDealt {
			t.Errorf("Player with the name %s was dealt %d cards instead of %d",
				player.name, len(player.cards), nOfCardsDealt)
		}
		if player.name == dealer.name {
			continue
		}
		for _, card := range player.cards {
			if !card.Visible {
				t.Errorf("Non-dealer player with the name %s has face down card. Card should be visible",
					player.name)
			}
		}
		if results[player.name] != calculatePoints(&player.cards) {
			t.Errorf("Final result after dealing is not correct. Expected %d, got %d",
				calculatePoints(&player.cards), results[player.name])
		}
	}
}

func TestCalculatePoints(t *testing.T) {
	hand16 := []deck.Card{}
	hand16 = append(hand16,
		deck.Card{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	if calculatePoints(&hand16) != 16 {
		t.Errorf("Expected card points to be %d, got %d", 16, calculatePoints(&hand16))
	}
	hand17 := []deck.Card{}
	hand17 = append(hand17,
		deck.Card{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		deck.Card{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		deck.Card{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	)
	if calculatePoints(&hand17) != 17 {
		t.Errorf("Expected card points to be %d, got %d", 17, calculatePoints(&hand17))
	}
}
