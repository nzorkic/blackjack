package game

import "testing"

func TestSetResults(t *testing.T) {
	players := []player{
		{name: "p1", points: 16}, {name: "p2", points: 20}, {name: "p3", points: 18}}
	dealer.points = 18
	setResults(&players)
	for idx, p := range players {
		if idx == 0 {
			if results[p.name] != lost {
				t.Errorf("Player %s should have result lost, but he didn't", p.name)
			}
		}
		if idx == 1 {
			if results[p.name] != won {
				t.Errorf("Player %s should have result won, but he didn't", p.name)
			}
		}
		if idx == 2 {
			if results[p.name] != push {
				t.Errorf("Player %s should have result push, but he didn't", p.name)
			}
		}
	}
	resetResults()
	results[dealer.name] = blackjack
	results["p2"] = blackjack
	dealer.points = 18
	setResults(&players)
	for idx, p := range players {
		if idx == 0 {
			if results[p.name] != lost {
				t.Errorf("Player %s should have result lost, but he didn't", p.name)
			}
		}
		if idx == 1 {
			if results[p.name] != push {
				t.Errorf("Player %s should have result push, but he didn't", p.name)
			}
		}
		if idx == 2 {
			if results[p.name] != lost {
				t.Errorf("Player %s should have result lost, but he didn't", p.name)
			}
		}
	}
	resetResults()
	results[dealer.name] = lost
	results["p2"] = lost
	dealer.points = 18
	setResults(&players)
	for idx, p := range players {
		if idx == 0 {
			if results[p.name] != won {
				t.Errorf("Player %s should have result won, but he didn't", p.name)
			}
		}
		if idx == 1 {
			if results[p.name] != push {
				t.Errorf("Player %s should have result push, but he didn't", p.name)
			}
		}
		if idx == 2 {
			if results[p.name] != won {
				t.Errorf("Player %s should have result won, but he didn't", p.name)
			}
		}
	}
}

func TestFinalizeResults(t *testing.T) {
	totalChips["p1"] = 900
	totalChips["p2"] = 900
	totalChips["p3"] = 900
	offeredChips["p1"] = 100
	offeredChips["p2"] = 100
	offeredChips["p3"] = 100
	results["p1"] = won
	results["p2"] = blackjack
	results["p3"] = push
	finalizeResults()
	if totalChips["p1"] != 1100 {
		t.Errorf("Total chips for p1 should be 1000, but is %d", totalChips["p1"])
	}
	if totalChips["p2"] != 1200 {
		t.Errorf("Total chips for p2 should be 1100, but is %d", totalChips["p2"])
	}
	if totalChips["p3"] != 1000 {
		t.Errorf("Total chips for p3 should be 900, but is %d", totalChips["p3"])
	}
}

func TestResetResults(t *testing.T) {
	results["test1"] = won
	results["test2"] = push
	resetResults()
	if len(results) != 0 {
		t.Error("Result table has not been emptied.")
	}
}
