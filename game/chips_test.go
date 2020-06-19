package game

import "testing"

func TestResetChips(t *testing.T) {
	totalChips["p1"] = 1
	totalChips["p2"] = 2
	totalChips["p3"] = 3
	resetChips()
	if len(totalChips) != 0 {
		t.Error("Total chips did not reset.")
	}
}

func TestRemoveChip(t *testing.T) {
	name := "p1"
	totalChips[name] = 1
	removeChip(&name)
	if _, ok := totalChips[name]; ok {
		t.Error("Player was not removed from the total bets")
	}
}
