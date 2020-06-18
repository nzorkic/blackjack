package main

import (
	"fmt"
	"os"

	util "github.com/nzorkic/blackjack/internal"
	"github.com/nzorkic/deck"
)

// MaxScore is highest score player can get in Blackjack
const MaxScore = 21

type player struct {
	name   string
	cards  []deck.Card
	points int
}

var dealer = player{name: "Dealer"}

var results = make(map[string]int)

type action uint8

const (
	hit action = iota
	stand
)

func (p player) String() string {
	return fmt.Sprintf("%s (%d): %v", p.name, util.CalculatePoints(&p.cards), util.CardsAsString(&p.cards))
}

func main() {
	newGame()
}

func newGame() {
	fmt.Println("#############")
	fmt.Println("# BLACKJACK #")
	fmt.Println("#############")
	fmt.Println()
	fmt.Print("Enter number of players > ")
	var numerOfPlayers int
	fmt.Scanln(&numerOfPlayers)
	fmt.Println()
	players := createPlayers(&numerOfPlayers)
	playingDeck := createDeck(&numerOfPlayers)
	deal(&players, &playingDeck)
	start(&players, &playingDeck)
	replay(&players, &playingDeck)
}

func createDeck(n *int) deck.Deck {
	deckSize := ((*n*4)+4)/52 + 1
	newDeck := deck.New(deck.Size(deckSize), deck.Shuffle())
	initPoints(&newDeck)
	return newDeck
}

func initPoints(deck *deck.Deck) {
	(*deck).FacePoints(10)
}

func createPlayers(n *int) []player {
	players := make([]player, *n)
	for i := 0; i < *n; i++ {
		players[i] = player{name: fmt.Sprintf("Player #%d", i+1)}
	}
	return players
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
	for _, player := range *p {
		results[player.name] += util.CalculatePoints(&player.cards)
	}
	results[dealer.name] += util.CalculatePoints(&dealer.cards)
	printScores(p)
}

func start(p *[]player, d *deck.Deck) {
	for _, player := range *p {
		invokeAction(&player, d)
	}
	dealer.cards[1].Visible = true
	results[dealer.name] = util.CalculatePoints(&dealer.cards)
	invokeAction(&dealer, d)
	printFinalScore()
}

func invokeAction(player *player, d *deck.Deck) {
	var playerAction action
	printActionChoices(player)
	fmt.Scanln(&playerAction)
	fmt.Println()
	switch playerAction {
	case hit:
		hitAction(player, d)
	case stand:
		break
	}
}

func hitAction(player *player, d *deck.Deck) {
	player.cards = append(player.cards, (*d).Draw(1)...)
	var playerAction action = hit
	for playerAction == hit {
		printActionChoices(player)
		points := util.CalculatePoints(&player.cards)
		results[player.name] = points
		if points > MaxScore {
			fmt.Println("It's a bust!")
			fmt.Println()
			break
		}
		fmt.Scanln(&playerAction)
		fmt.Println()
		if playerAction == hit {
			player.cards = append(player.cards, (*d).Draw(1)...)
		}
	}
}

func replay(players *[]player, d *deck.Deck) {
	var choice uint8
	fmt.Println()
	fmt.Println("Want to play another game?")
	fmt.Println("Enter 0 for another round")
	fmt.Println("Enter 1 for new game")
	fmt.Println("Enter 2 to quit")
	fmt.Print("> ")
	fmt.Scanln(&choice)
	switch choice {
	case 0:
		startNewRound(players, d)
	case 1:
		for k := range results {
			delete(results, k)
		}
		dealer.cards = nil
		newGame()
	case 2:
		os.Exit(1)
	}
}

func startNewRound(players *[]player, d *deck.Deck) {
	fmt.Println()
	for k := range results {
		delete(results, k)
	}
	for idx := range *players {
		(*players)[idx].cards = nil
	}
	dealer.cards = nil
	if ((len(*players) * 4) + 4) > len(*d) {
		fmt.Println("Shuffling new deck...")
		deckSize := ((len(*players)*4)+4)/52 + 1
		*d = deck.New(deck.Size(deckSize), deck.Shuffle())
	}
	deal(players, d)
	start(players, d)
	replay(players, d)
}

func printScores(players *[]player) {
	for _, player := range *players {
		fmt.Println(player)
	}
	fmt.Println(dealer)
	fmt.Println()
}

func printActionChoices(player *player) {
	fmt.Println(player)
	fmt.Println("Enter 0 to Hit or 1 to Stand")
	fmt.Print("> ")
}

func printFinalScore() {
	var winner string
	winnerPts := 0
	runnerUpPts := 0
	for name, score := range results {
		fmt.Printf("%s: %d\n", name, score)
		if score > MaxScore {
			continue
		}
		if score >= winnerPts {
			winner = name
			runnerUpPts = winnerPts
			winnerPts = score
		}
	}
	fmt.Println()
	if runnerUpPts == winnerPts {
		fmt.Println("It's a push!")
	} else {
		fmt.Println("The winner is ", winner)
	}
}
