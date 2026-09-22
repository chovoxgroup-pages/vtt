package main

import (
	"math"
	"strings"
)

// AbilityScores holds the core 6 stats for a 3.5e character
type AbilityScores struct {
	STR int
	DEX int
	CON int
	INT int
	WIS int
	CHA int
}

// Speed tracks all movement types in feet per round
type Speed struct {
	Land   int
	Fly    int
	Swim   int
	Burrow int
	Climb  int
}

// DamageReduction represents a single DR instance (e.g., 5/Magic)
type DamageReduction struct {
	Amount int
	Bypass string // e.g., "magic", "silver", "adamantine", "-"
}

// Senses tracks vision and sensory ranges in feet
type Senses struct {
	NormalVision   bool
	LowLightVision bool
	Darkvision     int // Range in feet
	Tremorsense    int // Range in feet
	Blindsense     int // Range in feet
	Blindsight     int // Range in feet
}

// ClassProgression tracks a character's level and choices in a specific class
type ClassProgression struct {
	ClassID string // Maps to ClassBlueprint.ID (e.g., "cls_cleric")
	Level   int

	// --- Class-Specific Choices ---
	Domains           []string // Populated with DomainBlueprint IDs
	CombatStyle       string   // Populated with a CombatStyleBlueprint ID
	RogueAbilities    []string // Populated with RogueAbilityBlueprint IDs
	SpecialistSchool  string
	ProhibitedSchools []string
}

// Entity represents a player character, NPC, or monster
type Entity struct {
	ID        string
	Name      string
	Race      string
	Alignment string // e.g., "Lawful Good"

	// Core Attributes
	BaseStats AbilityScores
	Speed     Speed
	Senses    Senses
	DR        []DamageReduction

	// Progression
	Classes    []ClassProgression
	Experience int

	// Vitals
	MaxHP     int
	CurrentHP int
	Nonlethal int

	// Feats & Skills
	Feats      []string           // List of selected Feat IDs
	SkillRanks map[string]float32 // Skill ID to ranks (float32 allows for 0.5 cross-class increments)

	// Magic
	PreparedSpells map[string]int // Maps a Spell ID to the number of times it is prepared today
	SpellsKnown    []string       // List of Spell IDs the entity knows (for Sorcerers/Bards)
}

// NewEntity creates a new character with baseline defaults
func NewEntity(id string, name string, race string) *Entity {
	return &Entity{
		ID:   id,
		Name: name,
		Race: race,
		BaseStats: AbilityScores{
			STR: 10,
			DEX: 10,
			CON: 10,
			INT: 10,
			WIS: 10,
			CHA: 10,
		},
		Speed: Speed{
			Land: 30, // Default humanoid assumption; can be overwritten by race logic
		},
		Senses: Senses{
			NormalVision: true,
		},
		SkillRanks:     make(map[string]float32),
		PreparedSpells: make(map[string]int),
		SpellsKnown:    make([]string, 0),
		DR:             make([]DamageReduction, 0),
		Classes:        make([]ClassProgression, 0),
		Feats:          make([]string, 0),
	}
}

// GetCharacterLevel returns the sum of all class levels
func (e *Entity) GetCharacterLevel() int {
	total := 0
	for _, c := range e.Classes {
		total += c.Level
	}
	return total
}

// GetAbilityModifier pulls the current score for an ability and calculates the 3.5e modifier
func (e *Entity) GetAbilityModifier(ability string) int {
	score := 10
	switch strings.ToUpper(ability) {
	case "STR":
		score = e.BaseStats.STR
	case "DEX":
		score = e.BaseStats.DEX
	case "CON":
		score = e.BaseStats.CON
	case "INT":
		score = e.BaseStats.INT
	case "WIS":
		score = e.BaseStats.WIS
	case "CHA":
		score = e.BaseStats.CHA
	}
	return int(math.Floor(float64(score-10) / 2.0))
}

// GetSkillTotal calculates the total modifier for a specific skill check
func (e *Entity) GetSkillTotal(skillID string) int {
	blueprint, exists := GlobalSkills[skillID]
	if !exists {
		return 0 // Fail-safe if an unknown skill is requested
	}

	ranks := float32(0.0)
	if r, hasRanks := e.SkillRanks[skillID]; hasRanks {
		ranks = r
	}

	abilityMod := e.GetAbilityModifier(blueprint.KeyAbility)

	return int(math.Floor(float64(ranks))) + abilityMod
}

// GetPassiveSpot implements the passive Perception rule (10 + Spot Modifier)
func (e *Entity) GetPassiveSpot() int {
	return 10 + e.GetSkillTotal("skill_spot")
}

// EntityJSON represents the universal data structure stored in Firebase.
// This is used for Player Characters, NPCs, and Monsters.
type EntityJSON struct {
	ID      string `json:"id"`
	Type    string `json:"type"`     // "Player", "NPC", "Monster"
	OwnerID string `json:"owner_id"` // Matches the Firebase user UID (for players)

	// Identity
	Name      string `json:"name"`
	Race      string `json:"race"`  // e.g., "race_dwarf_hill"
	Class     string `json:"class"` // e.g., "cls_rogue"
	Level     int    `json:"level"`
	Alignment string `json:"alignment"`

	// Core Stats
	AbilityScores map[string]int `json:"ability_scores"` // STR, DEX, CON, INT, WIS, CHA
	MaxHP         int            `json:"max_hp"`
	CurrentHP     int            `json:"current_hp"`
	BaseAttack    int            `json:"base_attack"`
	Speed         int            `json:"speed"`

	// Armor & Defenses
	ACNormal int `json:"ac_normal"`
	ACTouch  int `json:"ac_touch"`
	ACFlat   int `json:"ac_flat"`
	SaveFort int `json:"save_fort"`
	SaveRef  int `json:"save_ref"`
	SaveWill int `json:"save_will"`

	// Progression
	SkillRanks map[string]int `json:"skill_ranks"` // e.g., {"skill_spot": 4, "skill_open_lock": 4}
	Feats      []string       `json:"feats"`       // e.g., ["feat_improved_unarmed_strike", "feat_dodge"]

	// Inventory (To be expanded)
	EquippedWeapon string `json:"equipped_weapon"`
	EquippedArmor  string `json:"equipped_armor"`
}
