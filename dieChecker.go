package main

import (
	"math/rand"
	"time"
)

// DieChecker provides utilities for rolling dice and resolving 3.5e mechanics.
type DieChecker struct {
	rng *rand.Rand
}

// NewDieChecker initializes a new DieChecker with a freshly seeded random number generator.
func NewDieChecker() *DieChecker {
	return &DieChecker{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Roll simulates rolling standard D&D dice (e.g., 3d6 + 2).
// numDice: number of dice to roll
// sides: number of sides on the die
// modifier: flat bonus or penalty added to the total
func (dc *DieChecker) Roll(numDice int, sides int, modifier int) int {
	if numDice <= 0 || sides <= 0 {
		return modifier
	}

	total := 0
	for i := 0; i < numDice; i++ {
		// Intn(sides) returns 0 to sides-1, so we add 1 to get 1 to sides
		total += dc.rng.Intn(sides) + 1
	}
	return total + modifier
}

// RollD20 is a convenience method for standard skill, attack, and saving throw checks.
func (dc *DieChecker) RollD20(modifier int) int {
	return dc.Roll(1, 20, modifier)
}

// CheckDC evaluates a roll against a static Difficulty Class.
// Returns true if the roll equals or exceeds the target DC.
func (dc *DieChecker) CheckDC(rollResult int, targetDC int) bool {
	return rollResult >= targetDC
}

// ContestedRoll evaluates two opposed rolls (e.g., Hide vs Spot, or Grapple vs Grapple).
// Returns true if the attacker/initiator wins. On a tie, the defender wins.
func (dc *DieChecker) ContestedRoll(attackerResult int, defenderResult int) bool {
	return attackerResult > defenderResult
}
