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

// SkillBlueprint defines the core rules for a specific skill
type SkillBlueprint struct {
	ID                string
	Name              string
	KeyAbility        string // "STR", "DEX", "CON", "INT", "WIS", "CHA", or "NONE"
	TrainedOnly       bool
	ArmorCheckPenalty bool
}

// GlobalSkills acts as the registry for all available 3.5e skills
var GlobalSkills = map[string]SkillBlueprint{
	"skill_appraise":         {ID: "skill_appraise", Name: "Appraise", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_balance":          {ID: "skill_balance", Name: "Balance", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_bluff":            {ID: "skill_bluff", Name: "Bluff", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_climb":            {ID: "skill_climb", Name: "Climb", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_concentration":    {ID: "skill_concentration", Name: "Concentration", KeyAbility: "CON", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_decipher_script":  {ID: "skill_decipher_script", Name: "Decipher Script", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_diplomacy":        {ID: "skill_diplomacy", Name: "Diplomacy", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_disable_device":   {ID: "skill_disable_device", Name: "Disable Device", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_disguise":         {ID: "skill_disguise", Name: "Disguise", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_escape_artist":    {ID: "skill_escape_artist", Name: "Escape Artist", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_forgery":          {ID: "skill_forgery", Name: "Forgery", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_gather_info":      {ID: "skill_gather_info", Name: "Gather Information", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_handle_animal":    {ID: "skill_handle_animal", Name: "Handle Animal", KeyAbility: "CHA", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_heal":             {ID: "skill_heal", Name: "Heal", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_hide":             {ID: "skill_hide", Name: "Hide", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_intimidate":       {ID: "skill_intimidate", Name: "Intimidate", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_jump":             {ID: "skill_jump", Name: "Jump", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_listen":           {ID: "skill_listen", Name: "Listen", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_move_silently":    {ID: "skill_move_silently", Name: "Move Silently", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_open_lock":        {ID: "skill_open_lock", Name: "Open Lock", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_ride":             {ID: "skill_ride", Name: "Ride", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_search":           {ID: "skill_search", Name: "Search", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_sense_motive":     {ID: "skill_sense_motive", Name: "Sense Motive", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_sleight_of_hand":  {ID: "skill_sleight_of_hand", Name: "Sleight of Hand", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: true},
	"skill_spellcraft":       {ID: "skill_spellcraft", Name: "Spellcraft", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_spot":             {ID: "skill_spot", Name: "Spot", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_survival":         {ID: "skill_survival", Name: "Survival", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_swim":             {ID: "skill_swim", Name: "Swim", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_tumble":           {ID: "skill_tumble", Name: "Tumble", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: true},
	"skill_use_magic_device": {ID: "skill_use_magic_device", Name: "Use Magic Device", KeyAbility: "CHA", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_use_rope":         {ID: "skill_use_rope", Name: "Use Rope", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: false},
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

	// Note: In a fully fleshed out engine, you would also query armor check penalties,
	// racial bonuses, and feat synergies here before returning the final total.

	return int(math.Floor(float64(ranks))) + abilityMod
}

// GetPassiveSpot implements the passive Perception rule (10 + Spot Modifier)
func (e *Entity) GetPassiveSpot() int {
	return 10 + e.GetSkillTotal("skill_spot")
}
