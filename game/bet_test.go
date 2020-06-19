package game

import "testing"

func TestMakeBet(t *testing.T) {
	name := "p1"
	totalChips[name] = 900
	makeBet(&name, 150)
	if totalChips[name] != 750 {
		t.Errorf("Total chips for %s should be 750, got %d", name, totalChips[name])
	}
	if offeredChips[name] != 150 {
		t.Errorf("Offered chips for %s should be 150, got %d", name, offeredChips[name])
	}
}

func TestResetBets(t *testing.T) {
	offeredChips["p1"] = 1
	offeredChips["p2"] = 2
	offeredChips["p3"] = 3
	resetBets()
	if len(offeredChips) != 0 {
		t.Error("Offered chips did not reset.")
	}
}

func TestRemoveBet(t *testing.T) {
	name := "p1"
	offeredChips[name] = 1
	removeBet(&name)
	if _, ok := offeredChips[name]; ok {
		t.Error("Player was not removed from the offered bets")
	}
}
