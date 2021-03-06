package game

import (
	"fmt"
	"os"

	"github.com/nzorkic/deck"
)

// MaxScore is highest score player can get in Blackjack
const MaxScore = 21

// BlackjackBonus is bonus percentage for scoring blackjack
const BlackjackBonus float64 = 2

// NewGame initiates the game
func NewGame() {
	printGameIntro()
	fmt.Print("Enter number of players > ")
	var numerOfPlayers int
	fmt.Scan(&numerOfPlayers)
	fmt.Println()
	players := createPlayers(&numerOfPlayers)
	playingDeck := createDeck(&numerOfPlayers)
	offerChips(&players)
	deal(&players, &playingDeck)
	start(&players, &playingDeck)
	replay(&players, &playingDeck)
}

func createDeck(n *int) deck.Deck {
	deckSize := ((*n*4)+4)/52 + 1
	newDeck := deck.New(deck.Size(deckSize), deck.Shuffle())
	setupPoints(&newDeck)
	return newDeck
}

func deal(p *[]player, deck *deck.Deck) {
	playerSize := len(*p)
	turns := 2
	for i := 0; i < turns; i++ {
		for j := 0; j < playerSize; j++ {
			card := (*deck).Draw(1)[0]
			(*p)[j].cards = append((*p)[j].cards, card)
		}
		card := (*deck).Draw(1)[0]
		if i+1 == turns {
			card.Visible = false
		}
		dealer.cards = append(dealer.cards, card)
	}
	for idx, player := range *p {
		(*p)[idx].points = calculatePoints(&player.cards)
	}
	dealer.points = calculatePoints(&dealer.cards)
	printScores(p)
}

func start(p *[]player, d *deck.Deck) {
	for idx := range *p {
		invokeAction(&(*p)[idx], d)
	}
	dealer.cards[1].Visible = true
	dealer.points = calculatePoints(&dealer.cards)
	invokeAction(&dealer, d)
	setResults(p)
	finalizeResults()
	delete(results, dealer.name)
	printFinalScore()
}

func replay(players *[]player, d *deck.Deck) {
	var choice uint8
	fmt.Println()
	printEndgameOptions()
	fmt.Scan(&choice)
	switch choice {
	case 0:
		startNewRound(players, d)
	case 1:
		for k := range results {
			delete(results, k)
		}
		dealer.cards = nil
		resetBets()
		resetChips()
		NewGame()
	case 2:
		os.Exit(1)
	}
}

func startNewRound(players *[]player, d *deck.Deck) {
	fmt.Println()
	resetResults()
	removeBrokePlayers(players)
	resetHand(players)
	createNewDeckIfNeeded(len(*players), d)
	resetBets()
	offerChips(players)
	deal(players, d)
	start(players, d)
	replay(players, d)
}
