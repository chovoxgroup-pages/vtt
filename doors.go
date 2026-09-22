package main

import (
	"fmt"
	"math"
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
	IsOpen         bool
	Stuck          StuckLevel
	Lock           LockLevel
	IsTrapped      bool
	IsTrapDetected bool
	Trap           *Trap
	Secret         SecretLevel
	Material       string
	HP             int
	MaxHP          int
	Hardness       int
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

	// Assign Material, Hardness, and HP
	matRoll := rand.Intn(100)
	if matRoll < 60 {
		d.Material = "Wood"
		d.Hardness = 5
		d.MaxHP = 15
	} else if matRoll < 90 {
		d.Material = "Stone"
		d.Hardness = 8
		d.MaxHP = 60
	} else {
		d.Material = "Metal"
		d.Hardness = 10
		d.MaxHP = 60
	}
	d.HP = d.MaxHP

	// Traps and Locks based on APL
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

// Helper function to extract active ability modifier based on Point Buy + Race
func getAbilityMod(g *Game, stat string) int {
	// SAFETY FALLBACK: If the player tests a door before creating a character, default to a +0 mod.
	if len(g.StandardRaces) == 0 || g.PointBuyStats == nil {
		return 0
	}

	raceID := g.StandardRaces[g.SelectedRaceIndex]
	race := GlobalRaces[raceID]
	score := g.PointBuyStats[stat] + race.AbilityMods[stat]
	return int(math.Floor(float64(score-10) / 2.0))
}

// GetAvailableActions satisfies the Interactable interface for Doors
func (d *Door) GetAvailableActions(g *Game, targetX, targetY int) []RadialAction {
	var actions []RadialAction

	// 1. Perception/Inspect (Search check)
	actions = append(actions, RadialAction{
		Name: "Inspect",
		Execute: func() {
			searchMod := g.SkillRanksAllocated["skill_search"] + getAbilityMod(g, "INT")
			roll := rand.Intn(20) + 1
			total := roll + searchMod

			// Determine Lock text
			lockText := "Unlocked"
			if d.Lock != LockNone {
				lockText = "Locked"
			}

			msg := fmt.Sprintf("You inspect the %s door. (HP: %d/%d, Hardness: %d, %s)",
				d.Material, d.HP, d.MaxHP, d.Hardness, lockText)

			// Trap Search Check
			if d.IsTrapped && !d.IsTrapDetected {
				trapDC := 20 + d.Trap.CR // Standard 3.5e scaling
				if total >= trapDC {
					d.IsTrapDetected = true
					msg += fmt.Sprintf(" [Search %d vs DC %d] You spot a hidden hazard: %s!", total, trapDC, d.Trap.Name)
				} else {
					msg += fmt.Sprintf(" [Search %d] You don't notice any immediate hazards.", total)
				}
			}
			g.ActionLog = append([]string{msg}, g.ActionLog...)
		},
	})

	// 2. Open or Close
	if d.IsOpen {
		if d.HP > 0 { // Prevent closing a destroyed door
			actions = append(actions, RadialAction{
				Name: "Close",
				Execute: func() {
					d.IsOpen = false
					g.DungeonMap.Grid[targetY][targetX] = DoorC
					g.ActionLog = append([]string{fmt.Sprintf("Closed door at (%d, %d)", targetX, targetY)}, g.ActionLog...)
				},
			})
		}
	} else {
		actions = append(actions, RadialAction{
			Name: "Open",
			Execute: func() {
				if d.Lock != LockNone {
					g.ActionLog = append([]string{"You rattle the handle, but the door is firmly locked."}, g.ActionLog...)
				} else if d.Stuck != StuckNone {
					g.ActionLog = append([]string{"The door seems to be stuck in its frame. You'll need to force it."}, g.ActionLog...)
				} else {
					d.IsOpen = true
					g.DungeonMap.Grid[targetY][targetX] = Floor
					g.ActionLog = append([]string{fmt.Sprintf("Opened door at (%d, %d)", targetX, targetY)}, g.ActionLog...)

					// Trigger trap if opening an untrapped, unbroken door
					if d.IsTrapped {
						g.ActionLog = append([]string{fmt.Sprintf("*CLICK* The %s triggers in your face!", d.Trap.Name)}, g.ActionLog...)
						d.IsTrapped = false // Disarmed by triggering
						d.IsTrapDetected = false
					}
				}
			},
		})
	}

	// 3. Unlock (Open Lock check)
	if !d.IsOpen && d.Lock != LockNone {
		actions = append(actions, RadialAction{
			Name: "Unlock",
			Execute: func() {
				openLockMod := g.SkillRanksAllocated["skill_open_lock"] + getAbilityMod(g, "DEX")
				roll := rand.Intn(20) + 1
				total := roll + openLockMod

				dc := 15
				switch d.Lock {
				case LockGood:
					dc = 20
				case LockHard:
					dc = 25
				case LockMaster:
					dc = 30
				}

				if total >= dc {
					d.Lock = LockNone
					g.ActionLog = append([]string{fmt.Sprintf("[Open Lock %d vs DC %d] *Click.* You successfully pick the lock.", total, dc)}, g.ActionLog...)
				} else {
					g.ActionLog = append([]string{fmt.Sprintf("[Open Lock %d vs DC %d] You struggle with the mechanism but fail to open it.", total, dc)}, g.ActionLog...)
				}
			},
		})
	}

	// 4. Disable Device (Only appears if a trap is detected)
	if !d.IsOpen && d.IsTrapped && d.IsTrapDetected {
		actions = append(actions, RadialAction{
			Name: "Disable Trap",
			Execute: func() {
				disableMod := g.SkillRanksAllocated["skill_disable_device"] + getAbilityMod(g, "INT")
				roll := rand.Intn(20) + 1
				total := roll + disableMod
				dc := 20 + d.Trap.CR

				if total >= dc {
					d.IsTrapped = false
					d.IsTrapDetected = false
					d.Trap = nil
					g.ActionLog = append([]string{fmt.Sprintf("[Disable Device %d vs DC %d] You carefully jam the mechanism. The trap is disarmed.", total, dc)}, g.ActionLog...)
				} else if total <= dc-5 {
					// Failed by 5 or more triggers the trap
					d.IsTrapped = false
					d.IsTrapDetected = false
					g.ActionLog = append([]string{fmt.Sprintf("[Disable Device %d vs DC %d] You slip! *SNAP* The %s triggers in your face!", total, dc, d.Trap.Name)}, g.ActionLog...)
				} else {
					g.ActionLog = append([]string{fmt.Sprintf("[Disable Device %d vs DC %d] You fail to disarm the trap, but realize your mistake before setting it off.", total, dc)}, g.ActionLog...)
				}
			},
		})
	}

	// 5. Bash (Deals damage to HP through Hardness, emits noise)
	if !d.IsOpen && d.HP > 0 && (d.Stuck != StuckNone || d.Lock != LockNone) {
		actions = append(actions, RadialAction{
			Name: "Bash",
			Execute: func() {
				// Temporary simulated player damage (1d8 + STR mod)
				dmg := rand.Intn(8) + 1 + getAbilityMod(g, "STR")
				netDmg := dmg - d.Hardness

				g.ActionLog = append([]string{"*CRASH!* You violently bash the door. The noise echoes down the halls..."}, g.ActionLog...)
				g.BroadcastNoise(targetX, targetY, 50.0)

				if netDmg > 0 {
					d.HP -= netDmg
					if d.HP <= 0 {
						d.HP = 0
						d.IsOpen = true
						g.DungeonMap.Grid[targetY][targetX] = Floor
						g.ActionLog = append([]string{fmt.Sprintf("The %s door splinters and breaks open!", d.Material)}, g.ActionLog...)

						// Bashing triggers the trap if it wasn't disabled
						if d.IsTrapped {
							g.ActionLog = append([]string{fmt.Sprintf("The impact jars the mechanism! The %s triggers!", d.Trap.Name)}, g.ActionLog...)
							d.IsTrapped = false
							d.IsTrapDetected = false
						}
					} else {
						g.ActionLog = append([]string{fmt.Sprintf("You deal %d damage to the %s door. (HP: %d/%d)", netDmg, d.Material, d.HP, d.MaxHP)}, g.ActionLog...)
					}
				} else {
					g.ActionLog = append([]string{fmt.Sprintf("Your bash glances harmlessly off the %s door's hardness (%d).", d.Material, d.Hardness)}, g.ActionLog...)
				}
			},
		})
	}

	return actions
}
