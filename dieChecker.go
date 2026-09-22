package main

import (
	"math/rand"
	"time"
)

// ContestResult defines a 3-state outcome for opposed rolls.
type ContestResult int

const (
	ContestLoss ContestResult = -1
	ContestTie  ContestResult = 0
	ContestWin  ContestResult = 1
)

// DieChecker provides utilities for rolling dice and resolving 3.5e mechanics.[cite: 33]
type DieChecker struct {
	rng *rand.Rand //[cite: 33]
}

// NewDieChecker initializes a new DieChecker with a freshly seeded random number generator.[cite: 33]
func NewDieChecker() *DieChecker {
	return &DieChecker{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())), //[cite: 33]
	}
}

// Roll simulates rolling standard D&D dice (e.g., 3d6 + 2).[cite: 33]
// numDice: number of dice to roll[cite: 33]
// sides: number of sides on the die[cite: 33]
// modifier: flat bonus or penalty added to the total[cite: 33]
func (dc *DieChecker) Roll(numDice int, sides int, modifier int) int {
	if numDice <= 0 || sides <= 0 { //[cite: 33]
		return modifier //[cite: 33]
	}

	total := 0                     //[cite: 33]
	for i := 0; i < numDice; i++ { //[cite: 33]
		// Intn(sides) returns 0 to sides-1, so we add 1 to get 1 to sides[cite: 33]
		total += dc.rng.Intn(sides) + 1 //[cite: 33]
	}
	return total + modifier //[cite: 33]
}

// RollD20 is a convenience method for standard skill, attack, and saving throw checks.[cite: 33]
func (dc *DieChecker) RollD20(modifier int) int {
	return dc.Roll(1, 20, modifier) //[cite: 33]
}

// CheckDC evaluates a roll against a static Difficulty Class.[cite: 33]
// Returns true if the roll equals or exceeds the target DC.[cite: 33]
func (dc *DieChecker) CheckDC(rollResult int, targetDC int) bool {
	return rollResult >= targetDC //[cite: 33]
}

// ContestedRoll evaluates two opposed rolls (e.g., Hide vs Spot, or Grapple vs Grapple).[cite: 33]
// Returns ContestWin if the attacker/initiator wins, ContestLoss if they lose, or ContestTie if results are equal.
func (dc *DieChecker) ContestedRoll(attackerResult int, defenderResult int) ContestResult {
	if attackerResult > defenderResult {
		return ContestWin
	} else if attackerResult < defenderResult {
		return ContestLoss
	}
	return ContestTie
}
