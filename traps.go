package main

import (
	"math/rand"
)

type TrapCategory string

const (
	TrapKill   TrapCategory = "Killing"
	TrapSoften TrapCategory = "Softening"
	TrapEject  TrapCategory = "Ejecting"
	TrapAdv    TrapCategory = "Advantage"
)

type Trap struct {
	Name     string
	Category TrapCategory
	CR       int
}

type TrapData struct {
	Name string
	CR   int
}

func getTrapCategoriesForRoom(rtype RoomType) []TrapCategory {
	switch rtype {
	case "Guard Room", "Barracks":
		return []TrapCategory{TrapSoften, TrapAdv}
	case "Temple", "Main Vault":
		return []TrapCategory{TrapKill}
	case "Crypt", "Reliquary":
		return []TrapCategory{TrapKill, TrapEject}
	case "Prison Block", "Torture Chamber":
		return []TrapCategory{TrapEject, TrapSoften}
	case "Maze", "Hidden Chamber":
		return []TrapCategory{TrapAdv, TrapEject}
	case "Armory", "Sorting Room":
		return []TrapCategory{TrapSoften}
	default:
		return []TrapCategory{TrapSoften, TrapAdv}
	}
}

var trapDatabase = map[TrapCategory][]TrapData{
	TrapKill: {
		{"Scything Blade Trap", 1}, {"Wall Blade Trap", 1}, {"Ceiling Pendulum", 3}, {"Wall Scythe Trap", 4},
		{"Falling Block Trap", 5}, {"Moving Executioner Statue", 5}, {"Built-to-Collapse Wall", 6}, {"Compacting Room", 6},
		{"Flame Strike Trap", 6}, {"Deathblade Wall Scythe", 8}, {"Destruction Trap", 8}, {"Dropping Ceiling", 9},
		{"Crushing Room", 10}, {"Crushing Wall Trap", 10}, {"Wail of the Banshee Trap", 10},
	},
	TrapSoften: {
		{"Fusillade of Darts", 1}, {"Poison Dart Trap", 1}, {"Poison Needle Trap", 1}, {"Box of Brown Mold", 2},
		{"Burning Hands Trap", 2}, {"Inflict Light Wounds Trap", 2}, {"Javelin Trap", 2}, {"Poison Needle Trap", 2},
		{"Burning Hands Trap", 3}, {"Fire Trap", 3}, {"Ghoul Touch Trap", 3}, {"Hail of Needles", 3},
		{"Acid Arrow Trap", 3}, {"Poisoned Arrow Trap", 3}, {"Bestow Curse Trap", 4}, {"Glyph of Warding (Blast)", 4},
		{"Lightning Bolt Trap", 4}, {"Poisoned Dart Trap", 4}, {"Sepia Snake Sigil Trap", 4}, {"Doorknob Smeared with Contact Poison", 5},
		{"Fireball Trap", 5}, {"Fusillade of Darts", 5}, {"Phantasmal Killer Trap", 5}, {"Poison Wall Spikes", 5},
		{"Ungol Dust Vapor Trap", 5}, {"Glyph of Warding (Blast)", 6}, {"Lightning Bolt Trap", 6}, {"Whirling Poison Blades", 6},
		{"Wyvern Arrow Trap", 6}, {"Acid Fog Trap", 7}, {"Blade Barrier Trap", 7}, {"Burnt Othur Vapor Trap", 7},
		{"Chain Lightning Trap", 7}, {"Fusillade of Greenblood Oil Darts", 7}, {"Lock Covered in Dragon Bile", 7}, {"Earthquake Trap", 8},
		{"Insanity Mist Vapor Trap", 8}, {"Acid Arrow Trap", 8}, {"Prismatic Spray Trap", 8}, {"Word of Chaos Trap", 8},
		{"Drawer Handle Smeared with Contact Poison", 9}, {"Incendiary Cloud Trap", 9}, {"Energy Drain Trap", 10},
	},
	TrapEject: {
		{"Camouflaged Pit Trap", 1}, {"Deeper Pit Trap", 1}, {"Camouflaged Pit Trap", 2}, {"Pit Trap", 2},
		{"Spiked Pit Trap", 2}, {"Well-Camouflaged Pit Trap", 2}, {"Camouflaged Pit Trap", 3}, {"Pit Trap", 3},
		{"Spiked Pit Trap", 3}, {"Camouflaged Pit Trap", 4}, {"Pit Trap", 4}, {"Spiked Pit Trap", 4},
		{"Water-Filled Room Trap", 4}, {"Wide-Mouth Spiked Pit Trap", 4}, {"Camouflaged Pit Trap", 5}, {"Flooding Room Trap", 5},
		{"Pit Trap", 5}, {"Spiked Pit Trap", 5}, {"Spiked Pit Trap (80 Ft. Deep)", 5}, {"Spiked Pit Trap (100 Ft. Deep)", 6},
		{"Wide-Mouth Pit Trap", 6}, {"Water-Filled Room", 7}, {"Well-Camouflaged Pit Trap", 7}, {"Reverse Gravity Trap", 8},
		{"Well-Camouflaged Pit Trap", 8}, {"Wide-Mouth Pit Trap", 9}, {"Wide-Mouth Spiked Pit with Poisoned Spikes", 9}, {"Poisoned Spiked Pit Trap", 10},
	},
	TrapAdv: {
		{"Basic Arrow Trap", 1}, {"Portcullis Trap", 1}, {"Razor-Wire across Hallway", 1}, {"Rolling Rock Trap", 1},
		{"Spear Trap", 1}, {"Swinging Block Trap", 1}, {"Bricks from Ceiling", 2}, {"Large Net Trap", 2},
		{"Tripping Chain", 2}, {"Extended Bane Trap", 3}, {"Stone Blocks from Ceiling", 3}, {"Collapsing Column", 4},
		{"Fusillade of Spears", 6}, {"Spiked Blocks from Ceiling", 6}, {"Black Tentacles Trap", 7}, {"Summon Monster VI Trap", 7},
		{"Power Word Stun Trap", 8}, {"Forcecage and Summon Monster VII trap", 10},
	},
}

// Additional subset designed specifically for doors/handles/locks
var doorTrapDatabase = map[TrapCategory][]TrapData{
	TrapKill: {
		{"Scything Blade Trap", 1}, {"Wall Blade Trap", 1}, {"Wall Scythe Trap", 4}, {"Deathblade Wall Scythe", 8}, {"Destruction Trap", 8},
	},
	TrapSoften: {
		{"Poison Needle Trap", 1}, {"Box of Brown Mold", 2}, {"Inflict Light Wounds Trap", 2}, {"Poison Needle Trap", 2},
		{"Fire Trap", 3}, {"Ghoul Touch Trap", 3}, {"Acid Arrow Trap", 3}, {"Bestow Curse Trap", 4},
		{"Glyph of Warding (Blast)", 4}, {"Sepia Snake Sigil Trap", 4}, {"Doorknob Smeared with Contact Poison", 5}, {"Fireball Trap", 5},
		{"Phantasmal Killer Trap", 5}, {"Glyph of Warding (Blast)", 6}, {"Lock Covered in Dragon Bile", 7}, {"Acid Arrow Trap", 8},
		{"Word of Chaos Trap", 8}, {"Drawer Handle Smeared with Contact Poison", 9}, {"Energy Drain Trap", 10},
	},
}

func getTrapData(cat TrapCategory, cr int, isDoor bool) string {
	if cr < 1 {
		cr = 1
	}

	var list []TrapData
	if isDoor {
		// Doors strongly favor Kill or Soften
		if cat != TrapKill && cat != TrapSoften {
			cat = TrapSoften // default to soften for doors if somehow another category is passed
		}
		list = doorTrapDatabase[cat]
	} else {
		list = trapDatabase[cat]
	}

	if len(list) == 0 {
		return "Unknown Hazard"
	}

	var valid []TrapData
	for _, t := range list {
		if t.CR <= cr {
			valid = append(valid, t)
		}
	}

	if len(valid) == 0 {
		return list[0].Name
	}

	// Filter to most challenging appropriate traps
	maxValidCR := 0
	for _, t := range valid {
		if t.CR > maxValidCR {
			maxValidCR = t.CR
		}
	}

	var challenging []TrapData
	for _, t := range valid {
		if t.CR >= maxValidCR-2 { // Give a little variance for traps (within 2 CR)
			challenging = append(challenging, t)
		}
	}

	return challenging[rand.Intn(len(challenging))].Name
}

func generateTrapName(cat TrapCategory, cr int) string {
	return getTrapData(cat, cr, false)
}

func generateDoorTrapName(cat TrapCategory, cr int) string {
	return getTrapData(cat, cr, true)
}

func (m *Map) PlaceRoomTraps() {
	if m.Config.Traps == "none" {
		return
	}

	trapChance := 15
	if m.Config.Traps == "many" {
		trapChance = 40
	}
	focalTrapChance := trapChance + 25

	for i := range m.Rooms {
		chance := trapChance
		if m.Rooms[i].IsFocal {
			chance = focalTrapChance
		}

		if rand.Intn(100) < chance {
			categories := getTrapCategoriesForRoom(m.Rooms[i].RType)
			cat := categories[rand.Intn(len(categories))]
			variance := rand.Intn(4) - 1
			cr := m.Config.APL + variance
			if cr < 1 {
				cr = 1
			}

			trapName := generateTrapName(cat, cr)
			m.Rooms[i].Traps = append(m.Rooms[i].Traps, Trap{
				Name:     trapName,
				Category: cat,
				CR:       cr,
			})
		}
	}
}
