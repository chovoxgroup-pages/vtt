package main

import (
	"math/rand"
)

type StuckLevel int

const (
	StuckNone StuckLevel = iota
	StuckEasy
	StuckNormal
	StuckHard
)

type LockLevel int

const (
	LockNone LockLevel = iota
	LockPoor
	LockGood
	LockHard
	LockMaster
)

type SecretLevel int

const (
	SecretNone SecretLevel = iota
	SecretPoor
	SecretWell
	SecretMaster
)

type Door struct {
	IsOpen    bool
	Stuck     StuckLevel
	Lock      LockLevel
	IsTrapped bool
	Trap      *Trap
	Secret    SecretLevel
}

func (m *Map) PlaceDoors() {
	if m.Config.Doors == "none" {
		return
	}
	baseChance, adjChance := 5, 75
	if m.Config.Doors == "some" {
		baseChance, adjChance = 5, 75
	}
	if m.Config.Doors == "many" {
		baseChance, adjChance = 15, 95
	}

	for y := 1; y < Height-1; y++ {
		for x := 1; x < Width-1; x++ {
			if m.Grid[y][x] == Floor {
				n, s, e, w := m.Grid[y-1][x], m.Grid[y+1][x], m.Grid[y][x+1], m.Grid[y][x-1]
				isHzChoke := (n == Wall && s == Wall && e == Floor && w == Floor)
				isVtChoke := (e == Wall && w == Wall && n == Floor && s == Floor)
				if isHzChoke || isVtChoke {
					adjRoom := false
					if isHzChoke {
						if m.isInRoom(x-1, y) || m.isInRoom(x+1, y) {
							adjRoom = true
						}
					} else {
						if m.isInRoom(x, y-1) || m.isInRoom(x, y+1) {
							adjRoom = true
						}
					}
					spawnChance := baseChance
					if adjRoom {
						spawnChance = adjChance
					}
					if rand.Intn(100) < spawnChance {
						m.spawnDoor(x, y)
					}
				}
			}
		}
	}
}

func (m *Map) spawnDoor(x, y int) {
	m.Grid[y][x] = DoorC
	d := &Door{}
	if m.Config.Secret {
		if rand.Intn(100) < 10 {
			d.Secret = SecretLevel(rand.Intn(3) + 1)
		}
	}

	trapChance := 0
	if m.Config.Traps == "some" {
		trapChance = 15
	}
	if m.Config.Traps == "many" {
		trapChance = 40
	}

	if rand.Intn(100) < trapChance {
		d.IsTrapped = true
		cat := []TrapCategory{TrapKill, TrapSoften}[rand.Intn(2)]
		variance := rand.Intn(4) - 1
		cr := m.Config.APL + variance
		if cr < 1 {
			cr = 1
		}

		d.Trap = &Trap{
			Name:     generateDoorTrapName(cat, cr),
			Category: cat,
			CR:       cr,
		}
	}

	lockChance := 0
	if m.Config.Locked == "some" {
		lockChance = 30
	}
	if m.Config.Locked == "many" {
		lockChance = 70
	}

	if d.Secret == SecretNone && rand.Intn(100) < 20 {
		d.IsOpen = true
	} else {
		d.IsOpen = false
		if rand.Intn(100) < 25 {
			d.Stuck = StuckLevel(rand.Intn(3) + 1)
		}
		if rand.Intn(100) < lockChance {
			d.Lock = LockLevel(rand.Intn(4) + 1)
		}
	}
	m.Doors[Node{x, y}] = d
}
