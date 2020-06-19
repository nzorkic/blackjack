package main

import (
	"fmt"
	"os"

	"github.com/nzorkic/blackjack/chips"
	util "github.com/nzorkic/blackjack/internal"
	"github.com/nzorkic/deck"
)

// MaxScore is highest score player can get in Blackjack
const MaxScore = 21

// BlackjackBonus is bonus percentage for scoring blackjack
const BlackjackBonus float64 = 2

type player struct {
	name   string
	cards  []deck.Card
	points int
}

var dealer = player{name: "Mr. Dealer"}

// score represents final score of player
type score uint8

const (
	push score = iota
	won
	lost
	blackjack
)

var results = make(map[string]score)

type action uint8

const (
	stand action = iota
	hit
)

func (p player) String() string {
	return p.name
}

func main() {
	newGame()
}

func newGame() {
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
	initPoints(&newDeck)
	return newDeck
}

func initPoints(deck *deck.Deck) {
	(*deck).FacePoints(10)
}

func createPlayers(n *int) []player {
	players := make([]player, *n)
	for i := 0; i < *n; i++ {
		name := fmt.Sprintf("Player #%d", i+1)
		player := player{name: name}
		players[i] = player
		chips.AddChips(&player.name, 900)
	}
	return players
}

func offerChips(players *[]player) {
	for _, player := range *players {
		var offer int
		playerChips := chips.Chips(&player.name)
		fmt.Printf("%s place a bet (max. %d$) > ", player.name, playerChips)
		fmt.Scan(&offer)
		for offer > playerChips {
			fmt.Printf("You don't have that much chips, %s, try again (max. %d$) > ", player.name, playerChips)
			fmt.Scan(&offer)
		}
		fmt.Println()
		chips.MakeBet(&player.name, offer)
	}
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
		(*p)[idx].points = util.CalculatePoints(&player.cards)
	}
	dealer.points = util.CalculatePoints(&dealer.cards)
	printScores(p)
}

func start(p *[]player, d *deck.Deck) {
	for idx := range *p {
		invokeAction(&(*p)[idx], d)
	}
	dealer.cards[1].Visible = true
	dealer.points = util.CalculatePoints(&dealer.cards)
	invokeAction(&dealer, d)
	setResults(p)
	finalizeBets()
	printFinalScore()
}

func invokeAction(player *player, d *deck.Deck) {
	var playerAction action
	printActionChoices(player)
	if blackjackHand(&player.cards) {
		results[player.name] = blackjack
		fmt.Println("Blackjack!")
		fmt.Println()
	} else {
		fmt.Scan(&playerAction)
		fmt.Println()
		switch playerAction {
		case hit:
			hitAction(player, d)
		case stand:
			break
		}
	}
}

func blackjackHand(cards *[]deck.Card) bool {
	if len(*cards) != 2 {
		return false
	}
	if util.CalculatePoints(cards) == MaxScore {
		return true
	}
	return false
}

func hitAction(player *player, d *deck.Deck) {
	player.cards = append(player.cards, (*d).Draw(1)...)
	var playerAction action = hit
	for playerAction == hit {
		printActionChoices(player)
		points := util.CalculatePoints(&player.cards)
		player.points = points
		if points > MaxScore {
			results[player.name] = lost
			fmt.Println("It's a bust!")
			fmt.Println()
			break
		}
		fmt.Scan(&playerAction)
		fmt.Println()
		if playerAction == hit {
			player.cards = append(player.cards, (*d).Draw(1)...)
		}
	}
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
		chips.ResetBets()
		chips.ResetChips()
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
	for idx, player := range *players {
		if chips.Chips(&player.name) <= 0 {
			*players = append((*players)[:idx], (*players)[idx+1:]...)
			chips.Remove(&player.name)
		}
		(*players)[idx].cards = nil
	}
	dealer.cards = nil
	if ((len(*players) * 4) + 4) > len(*d) {
		fmt.Println("Shuffling new deck...")
		fmt.Println()
		deckSize := ((len(*players)*4)+4)/52 + 1
		*d = deck.New(deck.Size(deckSize), deck.Shuffle())
	}
	chips.ResetBets()
	offerChips(players)
	deal(players, d)
	start(players, d)
	replay(players, d)
}

func setResults(players *[]player) {
	for _, player := range *players {
		res, ok := results[player.name]
		if ok {
			if (res == blackjack || res == lost) && results[dealer.name] == res {
				results[player.name] = push
			}
		} else {
			switch results[dealer.name] {
			case blackjack:
				results[player.name] = lost
			case lost:
				results[player.name] = won
			default:
				playerPoints := player.points
				dealerPoints := dealer.points
				fmt.Println("playerPoints: ", playerPoints, " dealerPoints: ", dealerPoints)
				switch {
				case playerPoints > dealerPoints:
					results[player.name] = won
				case playerPoints < dealerPoints:
					results[player.name] = lost
				case playerPoints == dealerPoints:
					results[player.name] = push
				}
			}
		}
	}
}

func finalizeBets() {
	for name, score := range results {
		switch score {
		case push:
			chips.AddChips(&name, chips.Bet(&name))
		case blackjack:
			chips.AddChips(&name, chips.Bet(&name)+int(float64(chips.Bet(&name))*BlackjackBonus))
		case won:
			chips.AddChips(&name, chips.Bet(&name)*2)
		}
	}
}

// PRINTS

func printGameIntro() {
	fmt.Println("#############")
	fmt.Println("# BLACKJACK #")
	fmt.Println("#############")
	fmt.Println()
}

func printScores(players *[]player) {
	for _, player := range *players {
		printPlayerInfoLine(&player)
	}
	printDealerInfoLine(&dealer)
	fmt.Println()
}

func printActionChoices(player *player) {
	printPlayerStats(player)
	fmt.Println("Enter 0 to Stand or 1 to Hit")
	fmt.Print("> ")
}

func printPlayerStats(player *player) {
	fmt.Println(player)
	fmt.Println("Hand: ", util.CardsAsString(&player.cards))
	fmt.Printf("Bank: $%d\n", chips.Chips(&player.name))
	fmt.Printf("Bet: $%d\n", chips.Bet(&player.name))
	fmt.Printf("Points: %d\n", util.CalculatePoints(&player.cards))
}

func printPlayerInfoLine(player *player) {
	fmt.Printf("%s [Bet: $%d (Bank: $%d)] (%d pts): %v\n",
		player.name,
		chips.Bet(&player.name),
		chips.Chips(&player.name),
		util.CalculatePoints(&player.cards),
		util.CardsAsString(&player.cards))
}

func printDealerInfoLine(dealer *player) {
	fmt.Printf("%s (%d): %v\n",
		dealer.name,
		util.CalculatePoints(&dealer.cards),
		util.CardsAsString(&dealer.cards))
}

func printFinalScore() {
	delete(results, dealer.name)
	for name, score := range results {
		fmt.Printf("%s: ", name)
		switch score {
		case blackjack:
			fmt.Printf("Won with Blackjack! (+$%d)", int(float64(chips.Bet(&name))*BlackjackBonus))
		case won:
			fmt.Printf("Won! (+$%d)", chips.Bet(&name))
		case lost:
			fmt.Printf("Lost (-$%d)", chips.Bet(&name))
		case push:
			fmt.Print("Push")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printEndgameOptions() {
	fmt.Println("Want to play another game?")
	fmt.Println("Enter 0 for another round")
	fmt.Println("Enter 1 for new game")
	fmt.Println("Enter 2 to quit")
	fmt.Print("> ")
}
