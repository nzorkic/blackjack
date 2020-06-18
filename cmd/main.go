package main

import (
	"fmt"

	util "github.com/nzorkic/blackjack/internal"
	"github.com/nzorkic/deck/deck"
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
	numOfPlayers := 1
	players := createPlayers(numOfPlayers)
	playingDeck := deck.New(deck.Shuffle())
	initPoints(&playingDeck)
	deal(&players, &playingDeck)
	start(&players, &playingDeck)
}

func initPoints(deck *deck.Deck) {
	(*deck).FacePoints(10)
}

func createPlayers(n int) []player {
	players := make([]player, n)
	for i := 0; i < n; i++ {
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
		hitAction(player, playerAction, d)
	case stand:
		break
	}
}

func hitAction(player *player, playerAction action, d *deck.Deck) {
	player.cards = append(player.cards, (*d).Draw(1)...)
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
