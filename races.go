package main

type SizeCategory string

const (
	SizeSmall  SizeCategory = "Small"
	SizeMedium SizeCategory = "Medium"
)

type VisionType string

const (
	VisionNormal  VisionType = "Normal"
	VisionLow     VisionType = "Low-Light"
	VisionDark60  VisionType = "Darkvision 60 ft."
	VisionDark90  VisionType = "Darkvision 90 ft."
	VisionDark120 VisionType = "Darkvision 120 ft."
)

// RaceBlueprint holds the 3.5 rules for a specific race or subrace
type RaceBlueprint struct {
	ID           string
	Name         string
	Size         SizeCategory
	BaseSpeed    int
	AbilityMods  map[string]int
	FavoredClass string
	Vision       VisionType
}

// GlobalRaces acts as our database for the Character Creation wizard
var GlobalRaces = map[string]RaceBlueprint{
	// --- HUMANS ---
	"race_human": {
		ID:           "race_human",
		Name:         "Human",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{},
		FavoredClass: "Any",
		Vision:       VisionNormal,
	},

	// --- DWARVES ---
	"race_dwarf_hill": {
		ID:           "race_dwarf_hill",
		Name:         "Dwarf (Hill)",
		Size:         SizeMedium,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "CHA": -2},
		FavoredClass: "Fighter",
		Vision:       VisionDark60,
	},
	"race_dwarf_deep": {
		ID:           "race_dwarf_deep",
		Name:         "Dwarf (Deep)",
		Size:         SizeMedium,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "CHA": -2},
		FavoredClass: "Fighter",
		Vision:       VisionDark90,
	},
	"race_dwarf_duergar": {
		ID:           "race_dwarf_duergar",
		Name:         "Dwarf (Duergar)",
		Size:         SizeMedium,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "CHA": -4},
		FavoredClass: "Fighter",
		Vision:       VisionDark120,
	},
	"race_dwarf_gold": {
		ID:           "race_dwarf_gold",
		Name:         "Dwarf (Gold)",
		Size:         SizeMedium,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "DEX": -2},
		FavoredClass: "Fighter",
		Vision:       VisionDark60,
	},

	// --- ELVES ---
	"race_elf_high": {
		ID:           "race_elf_high",
		Name:         "Elf (High)",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"DEX": 2, "CON": -2},
		FavoredClass: "Wizard",
		Vision:       VisionLow,
	},
	"race_elf_aquatic": {
		ID:           "race_elf_aquatic",
		Name:         "Elf (Aquatic)",
		Size:         SizeMedium,
		BaseSpeed:    30, // Also has Swim 40 ft, handled in entity instantiation
		AbilityMods:  map[string]int{"DEX": 2, "INT": -2},
		FavoredClass: "Fighter",
		Vision:       VisionLow,
	},
	"race_elf_drow": {
		ID:           "race_elf_drow",
		Name:         "Elf (Drow)",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"DEX": 2, "CON": -2, "INT": 2, "CHA": 2},
		FavoredClass: "Wizard", // Or Cleric for females
		Vision:       VisionDark120,
	},
	"race_elf_gray": {
		ID:           "race_elf_gray",
		Name:         "Elf (Gray)",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"DEX": 2, "CON": -2, "INT": 2, "STR": -2},
		FavoredClass: "Wizard",
		Vision:       VisionLow,
	},
	"race_elf_wild": {
		ID:           "race_elf_wild",
		Name:         "Elf (Wild)",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"DEX": 2, "INT": -2},
		FavoredClass: "Sorcerer",
		Vision:       VisionLow,
	},
	"race_elf_wood": {
		ID:           "race_elf_wood",
		Name:         "Elf (Wood)",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"DEX": 2, "CON": -2, "STR": 2, "INT": -2},
		FavoredClass: "Ranger",
		Vision:       VisionLow,
	},

	// --- GNOMES ---
	"race_gnome_rock": {
		ID:           "race_gnome_rock",
		Name:         "Gnome (Rock)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "STR": -2},
		FavoredClass: "Bard",
		Vision:       VisionLow,
	},
	"race_gnome_svirfneblin": {
		ID:           "race_gnome_svirfneblin",
		Name:         "Gnome (Svirfneblin)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"DEX": 2, "WIS": 2, "STR": -2, "CHA": -4},
		FavoredClass: "Rogue",
		Vision:       VisionDark120,
	},
	"race_gnome_forest": {
		ID:           "race_gnome_forest",
		Name:         "Gnome (Forest)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"CON": 2, "STR": -2},
		FavoredClass: "Bard",
		Vision:       VisionLow,
	},

	// --- HALF-ELVES ---
	"race_half_elf": {
		ID:           "race_half_elf",
		Name:         "Half-Elf",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{},
		FavoredClass: "Any",
		Vision:       VisionLow,
	},

	// --- HALF-ORCS ---
	"race_half_orc": {
		ID:           "race_half_orc",
		Name:         "Half-Orc",
		Size:         SizeMedium,
		BaseSpeed:    30,
		AbilityMods:  map[string]int{"STR": 2, "INT": -2, "CHA": -2},
		FavoredClass: "Barbarian",
		Vision:       VisionDark60,
	},

	// --- HALFLINGS ---
	"race_halfling_lightfoot": {
		ID:           "race_halfling_lightfoot",
		Name:         "Halfling (Lightfoot)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"DEX": 2, "STR": -2},
		FavoredClass: "Rogue",
		Vision:       VisionNormal,
	},
	"race_halfling_tallfellow": {
		ID:           "race_halfling_tallfellow",
		Name:         "Halfling (Tallfellow)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"DEX": 2, "STR": -2},
		FavoredClass: "Rogue",
		Vision:       VisionNormal,
	},
	"race_halfling_deep": {
		ID:           "race_halfling_deep",
		Name:         "Halfling (Deep)",
		Size:         SizeSmall,
		BaseSpeed:    20,
		AbilityMods:  map[string]int{"DEX": 2, "STR": -2},
		FavoredClass: "Rogue",
		Vision:       VisionDark60,
	},
}
