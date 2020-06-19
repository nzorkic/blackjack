package game

import (
	"testing"

	"github.com/nzorkic/deck"
)

func TestCalculatePoints(t *testing.T) {
	hand16 := []deck.Card{
		{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	}
	if calculatePoints(&hand16) != 16 {
		t.Errorf("Expected card points to be %d, got %d",
			16, calculatePoints(&hand16))
	}
	hand17 := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	}
	if calculatePoints(&hand17) != 17 {
		t.Errorf("Expected card points to be %d, got %d",
			17, calculatePoints(&hand17))
	}
}

func TestSortedCardsCopy(t *testing.T) {
	hand1 := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
		{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
	}
	arr1 := sortedCardsCopy(&hand1)
	if arr1[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
	hand2 := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	}
	arr2 := sortedCardsCopy(&hand2)
	if arr2[1].Rank != deck.Ace || arr2[2].Rank != deck.Ace {
		t.Errorf("Ace is not on the last place.")
	}
}

func TestCardsAsString(t *testing.T) {
	cards := []deck.Card{
		{Rank: deck.Ten, Suit: deck.Spade, Visible: true},
		{Rank: deck.Queen, Suit: deck.Heart, Visible: true},
		{Rank: deck.Joker, Suit: deck.Spade, Visible: true},
		{Rank: deck.Ace, Suit: deck.Diamond, Visible: false}}
	expectedString := "Ten of Spades, Queen of Hearts, Joker, FACEDOWN"
	if expectedString != cardsAsString(&cards) {
		t.Errorf("Error transforming cards to strings. Expected (%s), got (%s)",
			expectedString, cardsAsString(&cards))
	}
}

func TestBlackjackHand(t *testing.T) {
	bjackHand := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		{Rank: deck.Jack, Suit: deck.Spade, Visible: true, Point: 10},
	}
	if blackjackHand(&bjackHand) == false {
		t.Errorf("Hand should be a blackjack but it's not.")
	}
	notBlackjackHand := []deck.Card{
		{Rank: deck.Eight, Suit: deck.Heart, Visible: true, Point: 8},
		{Rank: deck.Queen, Suit: deck.Spade, Visible: true, Point: 10},
	}
	if blackjackHand(&notBlackjackHand) == true {
		t.Errorf("Hand should not be a blackjack but it is.")
	}
	notBlackjackHand21 := []deck.Card{
		{Rank: deck.Five, Suit: deck.Heart, Visible: true, Point: 5},
		{Rank: deck.King, Suit: deck.Spade, Visible: true, Point: 10},
		{Rank: deck.Six, Suit: deck.Spade, Visible: true, Point: 6},
	}
	if blackjackHand(&notBlackjackHand21) == true {
		t.Errorf("Hand should not be a blackjack but it is.")
	}
}

func TestResetHand(t *testing.T) {
	hand1 := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
		{Rank: deck.Queen, Suit: deck.Heart, Visible: true, Point: 10},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
	}
	player1 := player{cards: hand1}
	hand2 := []deck.Card{
		{Rank: deck.Ace, Suit: deck.Heart, Visible: true, Point: 1},
		{Rank: deck.Five, Suit: deck.Spade, Visible: true, Point: 5},
		{Rank: deck.Ace, Suit: deck.Spade, Visible: true, Point: 1},
	}
	player2 := player{cards: hand2}
	players := []player{player1, player2}
	resetHand(&players)
	for _, pl := range players {
		if pl.cards != nil {
			t.Errorf("Players hand has not been emptied.")
		}
	}
}

func TestCreateNewDeckIfNeeded(t *testing.T) {
	d := deck.New()
	originalLength := len(d)
	createNewDeckIfNeeded(2, &d)
	if originalLength != len(d) {
		t.Errorf("Wrong size of the deck, expected %d, got %d",
			originalLength, len(d))
	}
	createNewDeckIfNeeded(16, &d)
	if originalLength*2 != len(d) {
		t.Errorf("Wrong size of the deck, expected %d, got %d",
			originalLength*2, len(d))
	}
	createNewDeckIfNeeded(31, &d)
	if originalLength*3 != len(d) {
		t.Errorf("Wrong size of the deck, expected %d, got %d",
			originalLength*3, len(d))
	}
}

func TestSetupPoints(t *testing.T) {
	newDeck := deck.New()
	setupPoints(&newDeck)
	for _, card := range newDeck {
		if card.Rank == deck.Jack || card.Rank == deck.Queen || card.Rank == deck.King {
			if card.Point != 10 {
				t.Errorf("Points for  %s are %d, should be %d", card.Rank, card.Point, 10)
			}
		}
	}
}
