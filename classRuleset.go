package main

type ProgressionRate string

const (
	ProgGood    ProgressionRate = "Good"
	ProgAverage ProgressionRate = "Average"
	ProgPoor    ProgressionRate = "Poor"
)

type MagicType string
type CastingStyle string

const (
	Arcane MagicType = "Arcane"
	Divine MagicType = "Divine"

	Prepared    CastingStyle = "Prepared"
	Spontaneous CastingStyle = "Spontaneous"
)

// SpellcastingBlueprint defines how a specific class handles magic
type SpellcastingBlueprint struct {
	Type       MagicType
	Style      CastingStyle
	KeyAbility string // "INT", "WIS", or "CHA"

	// Maps the Class Level to a slice representing spell slots for levels 0 through 9
	SpellsPerDay map[int][]int

	// Only populated for spontaneous casters (Bards, Sorcerers)
	SpellsKnown map[int][]int

	// If true, the class receives +1 domain spell slot for each spell level it can cast (1st-9th)
	HasDomainSlots bool
}

// DomainBlueprint holds the 3.5 rules for a Cleric Domain
type DomainBlueprint struct {
	ID           string
	Name         string
	GrantedPower string         // Maps to a feature ID (e.g., "feat_turn_undead_bonus")
	Spells       map[int]string // Maps spell level to a Spell ID (e.g., 1: "spell_cure_light_wounds")
}

// ClassBlueprint holds the 3.5 rules for a specific class or monster Hit Die
type ClassBlueprint struct {
	ID           string
	Name         string
	HitDie       int
	BAB          ProgressionRate
	FortSave     ProgressionRate
	RefSave      ProgressionRate
	WillSave     ProgressionRate
	SkillPoints  int              // Base skill points per level
	Features     map[int][]string // Maps a class level to a list of feature IDs
	Spellcasting *SpellcastingBlueprint
}

// CombatStyleBlueprint defines the progression path for Ranger combat styles
type CombatStyleBlueprint struct {
	ID    string
	Name  string
	Feats map[int]string // Level -> Feat ID granted
}

// RogueAbilityBlueprint defines a selectable special ability for high-level rogues
type RogueAbilityBlueprint struct {
	ID          string
	Name        string
	Description string
}

// Global variables acting as our "Database" of 3.5 Classes and Monster Types
var (
	// --- CORE CLASSES ---
	ClassBarbarian = ClassBlueprint{
		ID:          "cls_barbarian",
		Name:        "Barbarian",
		HitDie:      12,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 4,
		Features: map[int][]string{
			1:  {"feat_fast_movement", "feat_illiteracy", "feat_rage_1"},
			2:  {"feat_uncanny_dodge"},
			3:  {"feat_trap_sense_1"},
			4:  {"feat_rage_2"},
			5:  {"feat_improved_uncanny_dodge"},
			6:  {"feat_trap_sense_2"},
			7:  {"feat_damage_reduction_1"},
			8:  {"feat_rage_3"},
			9:  {"feat_trap_sense_3"},
			10: {"feat_damage_reduction_2"},
			11: {"feat_greater_rage"},
			12: {"feat_rage_4", "feat_trap_sense_4"},
			13: {"feat_damage_reduction_3"},
			14: {"feat_indomitable_will"},
			15: {"feat_trap_sense_5"},
			16: {"feat_damage_reduction_4", "feat_rage_5"},
			17: {"feat_tireless_rage"},
			18: {"feat_trap_sense_6"},
			19: {"feat_damage_reduction_5"},
			20: {"feat_mighty_rage", "feat_rage_6"},
		},
	}

	ClassBard = ClassBlueprint{
		ID:          "cls_bard",
		Name:        "Bard",
		HitDie:      6,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 6,
		Features: map[int][]string{
			1:  {"feat_bardic_music", "feat_bardic_knowledge", "feat_countersong", "feat_fascinate", "feat_inspire_courage_1"},
			3:  {"feat_inspire_competence"},
			6:  {"feat_suggestion"},
			8:  {"feat_inspire_courage_2"},
			9:  {"feat_inspire_greatness"},
			12: {"feat_song_of_freedom"},
			14: {"feat_inspire_courage_3"},
			15: {"feat_inspire_heroics"},
			18: {"feat_mass_suggestion"},
			20: {"feat_inspire_courage_4"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Arcane,
			Style:      Spontaneous,
			KeyAbility: "CHA",
			SpellsPerDay: map[int][]int{
				1:  {2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {3, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {3, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				4:  {3, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				5:  {3, 3, 1, 0, 0, 0, 0, 0, 0, 0},
				6:  {3, 3, 2, 0, 0, 0, 0, 0, 0, 0},
				7:  {3, 3, 2, 0, 0, 0, 0, 0, 0, 0},
				8:  {3, 3, 3, 1, 0, 0, 0, 0, 0, 0},
				9:  {3, 3, 3, 2, 0, 0, 0, 0, 0, 0},
				10: {3, 3, 3, 2, 0, 0, 0, 0, 0, 0},
				11: {3, 3, 3, 3, 1, 0, 0, 0, 0, 0},
				12: {3, 3, 3, 3, 2, 0, 0, 0, 0, 0},
				13: {3, 3, 3, 3, 2, 0, 0, 0, 0, 0},
				14: {4, 3, 3, 3, 3, 1, 0, 0, 0, 0},
				15: {4, 4, 3, 3, 3, 2, 0, 0, 0, 0},
				16: {4, 4, 4, 3, 3, 2, 0, 0, 0, 0},
				17: {4, 4, 4, 4, 3, 3, 1, 0, 0, 0},
				18: {4, 4, 4, 4, 4, 3, 2, 0, 0, 0},
				19: {4, 4, 4, 4, 4, 4, 3, 0, 0, 0},
				20: {4, 4, 4, 4, 4, 4, 4, 0, 0, 0},
			},
			SpellsKnown: map[int][]int{
				1:  {4, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {5, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {6, 3, 0, 0, 0, 0, 0, 0, 0, 0},
				4:  {6, 3, 2, 0, 0, 0, 0, 0, 0, 0},
				5:  {6, 4, 3, 0, 0, 0, 0, 0, 0, 0},
				6:  {6, 4, 3, 0, 0, 0, 0, 0, 0, 0},
				7:  {6, 4, 4, 2, 0, 0, 0, 0, 0, 0},
				8:  {6, 4, 4, 3, 0, 0, 0, 0, 0, 0},
				9:  {6, 4, 4, 3, 0, 0, 0, 0, 0, 0},
				10: {6, 4, 4, 4, 2, 0, 0, 0, 0, 0},
				11: {6, 4, 4, 4, 3, 0, 0, 0, 0, 0},
				12: {6, 4, 4, 4, 3, 0, 0, 0, 0, 0},
				13: {6, 4, 4, 4, 4, 2, 0, 0, 0, 0},
				14: {6, 4, 4, 4, 4, 3, 0, 0, 0, 0},
				15: {6, 4, 4, 4, 4, 3, 0, 0, 0, 0},
				16: {6, 5, 4, 4, 4, 4, 2, 0, 0, 0},
				17: {6, 5, 5, 4, 4, 4, 3, 0, 0, 0},
				18: {6, 5, 5, 5, 4, 4, 3, 0, 0, 0},
				19: {6, 5, 5, 5, 5, 4, 4, 0, 0, 0},
				20: {6, 5, 5, 5, 5, 5, 4, 0, 0, 0},
			},
		},
	}

	ClassCleric = ClassBlueprint{
		ID:          "cls_cleric",
		Name:        "Cleric",
		HitDie:      8,           //
		BAB:         ProgAverage, //[cite: 3]
		FortSave:    ProgGood,    //[cite: 3]
		RefSave:     ProgPoor,    //[cite: 3]
		WillSave:    ProgGood,    //[cite: 3]
		SkillPoints: 2,           //
		Features: map[int][]string{
			1: {"feat_turn_or_rebuke_undead"}, //
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:           Divine,   //
			Style:          Prepared, //
			KeyAbility:     "WIS",    //
			HasDomainSlots: true,     //[cite: 25]
			SpellsPerDay: map[int][]int{
				1:  {3, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //
				2:  {4, 2, 0, 0, 0, 0, 0, 0, 0, 0}, //
				3:  {4, 2, 1, 0, 0, 0, 0, 0, 0, 0}, //
				4:  {5, 3, 2, 0, 0, 0, 0, 0, 0, 0}, //
				5:  {5, 3, 2, 1, 0, 0, 0, 0, 0, 0}, //
				6:  {5, 3, 3, 2, 0, 0, 0, 0, 0, 0}, //
				7:  {6, 4, 3, 2, 1, 0, 0, 0, 0, 0}, //
				8:  {6, 4, 3, 3, 2, 0, 0, 0, 0, 0}, //
				9:  {6, 4, 4, 3, 2, 1, 0, 0, 0, 0}, //
				10: {6, 4, 4, 3, 3, 2, 0, 0, 0, 0}, //
				11: {6, 5, 4, 4, 3, 2, 1, 0, 0, 0}, //
				12: {6, 5, 4, 4, 3, 3, 2, 0, 0, 0}, //
				13: {6, 5, 5, 4, 4, 3, 2, 1, 0, 0}, //
				14: {6, 5, 5, 4, 4, 3, 3, 2, 0, 0}, //
				15: {6, 5, 5, 5, 4, 4, 3, 2, 1, 0}, //
				16: {6, 5, 5, 5, 4, 4, 3, 3, 2, 0}, //
				17: {6, 5, 5, 5, 5, 4, 4, 3, 2, 1}, //
				18: {6, 5, 5, 5, 5, 4, 4, 3, 3, 2}, //
				19: {6, 5, 5, 5, 5, 5, 4, 4, 3, 3}, //
				20: {6, 5, 5, 5, 5, 5, 4, 4, 4, 4}, //
			},
		},
	}

	ClassDruid = ClassBlueprint{
		ID:          "cls_druid",
		Name:        "Druid",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 4,
		Features: map[int][]string{
			1:  {"feat_animal_companion", "feat_nature_sense", "feat_wild_empathy"},
			2:  {"feat_woodland_stride"},
			3:  {"feat_trackless_step"},
			4:  {"feat_resist_natures_lure"},
			5:  {"feat_wild_shape_1"},
			6:  {"feat_wild_shape_2"},
			7:  {"feat_wild_shape_3"},
			8:  {"feat_wild_shape_large"},
			9:  {"feat_venom_immunity"},
			10: {"feat_wild_shape_4"},
			11: {"feat_wild_shape_tiny"},
			12: {"feat_wild_shape_plant"},
			13: {"feat_a_thousand_faces"},
			14: {"feat_wild_shape_5"},
			15: {"feat_timeless_body", "feat_wild_shape_huge"},
			16: {"feat_wild_shape_elemental_1"},
			18: {"feat_wild_shape_6", "feat_wild_shape_elemental_2"},
			20: {"feat_wild_shape_elemental_3", "feat_wild_shape_elemental_huge"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Divine,
			Style:      Prepared,
			KeyAbility: "WIS",
			SpellsPerDay: map[int][]int{
				1:  {3, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {4, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {4, 2, 1, 0, 0, 0, 0, 0, 0, 0},
				4:  {5, 3, 2, 0, 0, 0, 0, 0, 0, 0},
				5:  {5, 3, 2, 1, 0, 0, 0, 0, 0, 0},
				6:  {5, 3, 3, 2, 0, 0, 0, 0, 0, 0},
				7:  {6, 4, 3, 2, 1, 0, 0, 0, 0, 0},
				8:  {6, 4, 3, 3, 2, 0, 0, 0, 0, 0},
				9:  {6, 4, 4, 3, 2, 1, 0, 0, 0, 0},
				10: {6, 4, 4, 3, 3, 2, 0, 0, 0, 0},
				11: {6, 5, 4, 4, 3, 2, 1, 0, 0, 0},
				12: {6, 5, 4, 4, 3, 3, 2, 0, 0, 0},
				13: {6, 5, 5, 4, 4, 3, 2, 1, 0, 0},
				14: {6, 5, 5, 4, 4, 3, 3, 2, 0, 0},
				15: {6, 5, 5, 5, 4, 4, 3, 2, 1, 0},
				16: {6, 5, 5, 5, 4, 4, 3, 3, 2, 0},
				17: {6, 5, 5, 5, 5, 4, 4, 3, 2, 1},
				18: {6, 5, 5, 5, 5, 4, 4, 3, 3, 2},
				19: {6, 5, 5, 5, 5, 5, 4, 4, 3, 3},
				20: {6, 5, 5, 5, 5, 5, 4, 4, 4, 4},
			},
		},
	}

	ClassFighter = ClassBlueprint{
		ID:          "cls_fighter",
		Name:        "Fighter",
		HitDie:      10,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
		Features: map[int][]string{
			1:  {"feat_fighter_bonus"},
			2:  {"feat_fighter_bonus"},
			4:  {"feat_fighter_bonus"},
			6:  {"feat_fighter_bonus"},
			8:  {"feat_fighter_bonus"},
			10: {"feat_fighter_bonus"},
			12: {"feat_fighter_bonus"},
			14: {"feat_fighter_bonus"},
			16: {"feat_fighter_bonus"},
			18: {"feat_fighter_bonus"},
			20: {"feat_fighter_bonus"},
		},
	}

	ClassMonk = ClassBlueprint{
		ID:          "cls_monk",
		Name:        "Monk",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 4,
		Features: map[int][]string{
			1:  {"feat_monk_bonus_choice_1", "feat_flurry_of_blows", "feat_unarmed_strike"},
			2:  {"feat_monk_bonus_choice_2", "feat_evasion"},
			3:  {"feat_still_mind", "feat_fast_movement_10"},
			4:  {"feat_ki_strike_magic", "feat_slow_fall_20"},
			5:  {"feat_purity_of_body"},
			6:  {"feat_monk_bonus_choice_3", "feat_slow_fall_30", "feat_fast_movement_20"},
			7:  {"feat_wholeness_of_body"},
			8:  {"feat_slow_fall_40"},
			9:  {"feat_improved_evasion", "feat_fast_movement_30"},
			10: {"feat_ki_strike_lawful", "feat_slow_fall_50"},
			11: {"feat_diamond_body", "feat_greater_flurry"},
			12: {"feat_abundant_step", "feat_slow_fall_60", "feat_fast_movement_40"},
			13: {"feat_diamond_soul"},
			14: {"feat_slow_fall_70"},
			15: {"feat_quivering_palm", "feat_fast_movement_50"},
			16: {"feat_ki_strike_adamantine", "feat_slow_fall_80"},
			17: {"feat_timeless_body", "feat_tongue_of_the_sun_and_moon"},
			18: {"feat_slow_fall_90", "feat_fast_movement_60"},
			19: {"feat_empty_body"},
			20: {"feat_perfect_self", "feat_slow_fall_any"},
		},
	}

	ClassPaladin = ClassBlueprint{
		ID:          "cls_paladin",
		Name:        "Paladin",
		HitDie:      10,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
		Features: map[int][]string{
			1:  {"feat_aura_of_good", "feat_detect_evil", "feat_smite_evil_1"},
			2:  {"feat_divine_grace", "feat_lay_on_hands"},
			3:  {"feat_aura_of_courage", "feat_divine_health"},
			4:  {"feat_turn_undead"},
			5:  {"feat_smite_evil_2", "feat_special_mount"},
			6:  {"feat_remove_disease_1"},
			9:  {"feat_remove_disease_2"},
			10: {"feat_smite_evil_3"},
			12: {"feat_remove_disease_3"},
			15: {"feat_remove_disease_4", "feat_smite_evil_4"},
			18: {"feat_remove_disease_5"},
			20: {"feat_smite_evil_5"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Divine,
			Style:      Prepared,
			KeyAbility: "WIS",
			SpellsPerDay: map[int][]int{
				4:  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				5:  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				6:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				7:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				8:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				9:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				10: {0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				11: {0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				12: {0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				13: {0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				14: {0, 2, 1, 1, 0, 0, 0, 0, 0, 0},
				15: {0, 2, 1, 1, 1, 0, 0, 0, 0, 0},
				16: {0, 2, 2, 1, 1, 0, 0, 0, 0, 0},
				17: {0, 2, 2, 2, 1, 0, 0, 0, 0, 0},
				18: {0, 3, 2, 2, 1, 0, 0, 0, 0, 0},
				19: {0, 3, 3, 3, 2, 0, 0, 0, 0, 0},
				20: {0, 3, 3, 3, 3, 0, 0, 0, 0, 0},
			},
		},
	}

	ClassRanger = ClassBlueprint{
		ID:          "cls_ranger",
		Name:        "Ranger",
		HitDie:      8,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 6,
		Features: map[int][]string{
			1:  {"feat_favored_enemy_1", "feat_track", "feat_wild_empathy"},
			2:  {"feat_ranger_combat_style_choice_1"},
			3:  {"feat_endurance"},
			4:  {"feat_animal_companion"},
			5:  {"feat_favored_enemy_2"},
			6:  {"feat_ranger_combat_style_choice_2"},
			7:  {"feat_woodland_stride"},
			8:  {"feat_swift_tracker"},
			9:  {"feat_evasion"},
			10: {"feat_favored_enemy_3"},
			11: {"feat_ranger_combat_style_choice_3"},
			13: {"feat_camouflage"},
			15: {"feat_favored_enemy_4"},
			17: {"feat_hide_in_plain_sight"},
			20: {"feat_favored_enemy_5"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Divine,
			Style:      Prepared,
			KeyAbility: "WIS",
			SpellsPerDay: map[int][]int{
				4:  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				5:  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				6:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				7:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				8:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				9:  {0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				10: {0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				11: {0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				12: {0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				13: {0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				14: {0, 2, 1, 1, 0, 0, 0, 0, 0, 0},
				15: {0, 2, 1, 1, 1, 0, 0, 0, 0, 0},
				16: {0, 2, 2, 1, 1, 0, 0, 0, 0, 0},
				17: {0, 2, 2, 2, 1, 0, 0, 0, 0, 0},
				18: {0, 3, 2, 2, 1, 0, 0, 0, 0, 0},
				19: {0, 3, 3, 3, 2, 0, 0, 0, 0, 0},
				20: {0, 3, 3, 3, 3, 0, 0, 0, 0, 0},
			},
		},
	}

	ClassRogue = ClassBlueprint{
		ID:          "cls_rogue",
		Name:        "Rogue",
		HitDie:      6,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 8,
		Features: map[int][]string{
			1:  {"feat_sneak_attack_1", "feat_trapfinding"},
			2:  {"feat_evasion"},
			3:  {"feat_sneak_attack_2", "feat_trap_sense_1"},
			4:  {"feat_uncanny_dodge"},
			5:  {"feat_sneak_attack_3"},
			6:  {"feat_trap_sense_2"},
			7:  {"feat_sneak_attack_4"},
			8:  {"feat_improved_uncanny_dodge"},
			9:  {"feat_sneak_attack_5", "feat_trap_sense_3"},
			10: {"feat_rogue_special_ability_choice"},
			11: {"feat_sneak_attack_6"},
			12: {"feat_trap_sense_4"},
			13: {"feat_sneak_attack_7", "feat_rogue_special_ability_choice"},
			15: {"feat_sneak_attack_8", "feat_trap_sense_5"},
			16: {"feat_rogue_special_ability_choice"},
			17: {"feat_sneak_attack_9"},
			18: {"feat_trap_sense_6"},
			19: {"feat_sneak_attack_10", "feat_rogue_special_ability_choice"},
		},
	}

	ClassSorcerer = ClassBlueprint{
		ID:          "cls_sorcerer",
		Name:        "Sorcerer",
		HitDie:      4,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 2,
		Features: map[int][]string{
			1: {"feat_summon_familiar"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Arcane,
			Style:      Spontaneous,
			KeyAbility: "CHA",
			SpellsPerDay: map[int][]int{
				1:  {5, 3, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {6, 4, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {6, 5, 0, 0, 0, 0, 0, 0, 0, 0},
				4:  {6, 6, 3, 0, 0, 0, 0, 0, 0, 0},
				5:  {6, 6, 4, 0, 0, 0, 0, 0, 0, 0},
				6:  {6, 6, 5, 3, 0, 0, 0, 0, 0, 0},
				7:  {6, 6, 6, 4, 0, 0, 0, 0, 0, 0},
				8:  {6, 6, 6, 5, 3, 0, 0, 0, 0, 0},
				9:  {6, 6, 6, 6, 4, 0, 0, 0, 0, 0},
				10: {6, 6, 6, 6, 5, 3, 0, 0, 0, 0},
				11: {6, 6, 6, 6, 6, 4, 0, 0, 0, 0},
				12: {6, 6, 6, 6, 6, 5, 3, 0, 0, 0},
				13: {6, 6, 6, 6, 6, 6, 4, 0, 0, 0},
				14: {6, 6, 6, 6, 6, 6, 5, 3, 0, 0},
				15: {6, 6, 6, 6, 6, 6, 6, 4, 0, 0},
				16: {6, 6, 6, 6, 6, 6, 6, 5, 3, 0},
				17: {6, 6, 6, 6, 6, 6, 6, 6, 4, 0},
				18: {6, 6, 6, 6, 6, 6, 6, 6, 5, 3},
				19: {6, 6, 6, 6, 6, 6, 6, 6, 6, 4},
				20: {6, 6, 6, 6, 6, 6, 6, 6, 6, 5},
			},
			SpellsKnown: map[int][]int{
				1:  {4, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {5, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {5, 3, 0, 0, 0, 0, 0, 0, 0, 0},
				4:  {6, 3, 1, 0, 0, 0, 0, 0, 0, 0},
				5:  {6, 4, 2, 0, 0, 0, 0, 0, 0, 0},
				6:  {7, 4, 2, 1, 0, 0, 0, 0, 0, 0},
				7:  {7, 5, 3, 2, 0, 0, 0, 0, 0, 0},
				8:  {8, 5, 3, 2, 1, 0, 0, 0, 0, 0},
				9:  {8, 5, 4, 3, 2, 0, 0, 0, 0, 0},
				10: {9, 5, 4, 3, 2, 1, 0, 0, 0, 0},
				11: {9, 5, 5, 4, 3, 2, 0, 0, 0, 0},
				12: {9, 5, 5, 4, 3, 2, 1, 0, 0, 0},
				13: {9, 5, 5, 4, 4, 3, 2, 0, 0, 0},
				14: {9, 5, 5, 4, 4, 3, 2, 1, 0, 0},
				15: {9, 5, 5, 4, 4, 4, 3, 2, 0, 0},
				16: {9, 5, 5, 4, 4, 4, 3, 2, 1, 0},
				17: {9, 5, 5, 4, 4, 4, 3, 3, 2, 0},
				18: {9, 5, 5, 4, 4, 4, 3, 3, 2, 1},
				19: {9, 5, 5, 4, 4, 4, 3, 3, 3, 2},
				20: {9, 5, 5, 4, 4, 4, 3, 3, 3, 3},
			},
		},
	}

	ClassWizard = ClassBlueprint{
		ID:          "cls_wizard",
		Name:        "Wizard",
		HitDie:      4,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 2,
		Features: map[int][]string{
			1:  {"feat_summon_familiar", "feat_scribe_scroll"},
			5:  {"feat_wizard_bonus_choice"},
			10: {"feat_wizard_bonus_choice"},
			15: {"feat_wizard_bonus_choice"},
			20: {"feat_wizard_bonus_choice"},
		},
		Spellcasting: &SpellcastingBlueprint{
			Type:       Arcane,
			Style:      Prepared,
			KeyAbility: "INT",
			SpellsPerDay: map[int][]int{
				1:  {3, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				2:  {4, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				3:  {4, 2, 1, 0, 0, 0, 0, 0, 0, 0},
				4:  {4, 3, 2, 0, 0, 0, 0, 0, 0, 0},
				5:  {4, 3, 2, 1, 0, 0, 0, 0, 0, 0},
				6:  {4, 3, 3, 2, 0, 0, 0, 0, 0, 0},
				7:  {4, 4, 3, 2, 1, 0, 0, 0, 0, 0},
				8:  {4, 4, 3, 3, 2, 0, 0, 0, 0, 0},
				9:  {4, 4, 4, 3, 2, 1, 0, 0, 0, 0},
				10: {4, 4, 4, 3, 3, 2, 0, 0, 0, 0},
				11: {4, 4, 4, 4, 3, 2, 1, 0, 0, 0},
				12: {4, 4, 4, 4, 3, 3, 2, 0, 0, 0},
				13: {4, 4, 4, 4, 4, 3, 2, 1, 0, 0},
				14: {4, 4, 4, 4, 4, 3, 3, 2, 0, 0},
				15: {4, 4, 4, 4, 4, 4, 3, 2, 1, 0},
				16: {4, 4, 4, 4, 4, 4, 3, 3, 2, 0},
				17: {4, 4, 4, 4, 4, 4, 4, 3, 2, 1},
				18: {4, 4, 4, 4, 4, 4, 4, 3, 3, 2},
				19: {4, 4, 4, 4, 4, 4, 4, 4, 3, 3},
				20: {4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
			},
		},
	}

	// GlobalClasses provides string-based lookups for the HTML Character Creator
	GlobalClasses = map[string]ClassBlueprint{
		"cls_barbarian": ClassBarbarian,
		"cls_bard":      ClassBard,
		"cls_cleric":    ClassCleric,
		"cls_druid":     ClassDruid,
		"cls_fighter":   ClassFighter,
		"cls_monk":      ClassMonk,
		"cls_paladin":   ClassPaladin,
		"cls_ranger":    ClassRanger,
		"cls_rogue":     ClassRogue,
		"cls_sorcerer":  ClassSorcerer,
		"cls_wizard":    ClassWizard,
	}

	// --- NPC CLASSES ---
	ClassAdept = ClassBlueprint{
		ID:          "npc_adept",
		Name:        "Adept",
		HitDie:      6,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 2,
	}

	ClassAristocrat = ClassBlueprint{
		ID:          "npc_aristocrat",
		Name:        "Aristocrat",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 4,
	}

	ClassCommoner = ClassBlueprint{
		ID:          "npc_commoner",
		Name:        "Commoner",
		HitDie:      4,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	ClassExpert = ClassBlueprint{
		ID:          "npc_expert",
		Name:        "Expert",
		HitDie:      6,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 6,
	}

	ClassMagewright = ClassBlueprint{
		ID:          "npc_magewright",
		Name:        "Magewright",
		HitDie:      4,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 2,
	}

	ClassWarrior = ClassBlueprint{
		ID:          "npc_warrior",
		Name:        "Warrior",
		HitDie:      8,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	// --- MONSTER TYPES ---
	TypeAberration = ClassBlueprint{
		ID:          "type_aberration",
		Name:        "Aberration",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 2,
	}

	TypeAnimal = ClassBlueprint{
		ID:          "type_animal",
		Name:        "Animal",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeConstruct = ClassBlueprint{
		ID:          "type_construct",
		Name:        "Construct",
		HitDie:      10,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeDragon = ClassBlueprint{
		ID:          "type_dragon",
		Name:        "Dragon",
		HitDie:      12,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 6,
	}

	TypeElementalAirFire = ClassBlueprint{
		ID:          "type_elemental_air_fire",
		Name:        "Elemental (Air/Fire)",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeElementalEarthWater = ClassBlueprint{
		ID:          "type_elemental_earth_water",
		Name:        "Elemental (Earth/Water)",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeFey = ClassBlueprint{
		ID:          "type_fey",
		Name:        "Fey",
		HitDie:      6,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 6,
	}

	TypeGiant = ClassBlueprint{
		ID:          "type_giant",
		Name:        "Giant",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeHumanoid = ClassBlueprint{
		ID:          "type_humanoid",
		Name:        "Humanoid",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeMagicalBeast = ClassBlueprint{
		ID:          "type_magical_beast",
		Name:        "Magical Beast",
		HitDie:      10,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeMonstrousHumanoid = ClassBlueprint{
		ID:          "type_monstrous_humanoid",
		Name:        "Monstrous Humanoid",
		HitDie:      8,
		BAB:         ProgGood,
		FortSave:    ProgPoor,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 2,
	}

	TypeOoze = ClassBlueprint{
		ID:          "type_ooze",
		Name:        "Ooze",
		HitDie:      10,
		BAB:         ProgAverage,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeOutsider = ClassBlueprint{
		ID:          "type_outsider",
		Name:        "Outsider",
		HitDie:      8,
		BAB:         ProgGood,
		FortSave:    ProgGood,
		RefSave:     ProgGood,
		WillSave:    ProgGood,
		SkillPoints: 8,
	}

	TypePlant = ClassBlueprint{
		ID:          "type_plant",
		Name:        "Plant",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	TypeUndead = ClassBlueprint{
		ID:          "type_undead",
		Name:        "Undead",
		HitDie:      12,
		BAB:         ProgPoor,
		FortSave:    ProgPoor,
		RefSave:     ProgPoor,
		WillSave:    ProgGood,
		SkillPoints: 4,
	}

	TypeVermin = ClassBlueprint{
		ID:          "type_vermin",
		Name:        "Vermin",
		HitDie:      8,
		BAB:         ProgAverage,
		FortSave:    ProgGood,
		RefSave:     ProgPoor,
		WillSave:    ProgPoor,
		SkillPoints: 2,
	}

	// --- CLERIC DOMAINS ---
	DomainAir = DomainBlueprint{
		ID:           "dom_air",
		Name:         "Air Domain",
		GrantedPower: "feat_domain_power_air", //[cite: 28]
		Spells: map[int]string{
			1: "spell_obscuring_mist",  //[cite: 28]
			2: "spell_wind_wall",       //[cite: 28]
			3: "spell_gaseous_form",    //[cite: 28]
			4: "spell_air_walk",        //[cite: 28]
			5: "spell_control_winds",   //[cite: 28]
			6: "spell_chain_lightning", //[cite: 28]
			7: "spell_control_weather", //[cite: 28]
			8: "spell_whirlwind",       //[cite: 28]
			9: "spell_elemental_swarm", //[cite: 28]
		},
	}

	DomainAnimal = DomainBlueprint{
		ID:           "dom_animal",
		Name:         "Animal Domain",
		GrantedPower: "feat_domain_power_animal", //[cite: 28]
		Spells: map[int]string{
			1: "spell_calm_animals",          //[cite: 28]
			2: "spell_hold_animal",           //[cite: 28]
			3: "spell_dominate_animal",       //[cite: 28]
			4: "spell_summon_natures_ally_4", //[cite: 28]
			5: "spell_commune_with_nature",   //[cite: 28]
			6: "spell_antilife_shell",        //[cite: 28]
			7: "spell_animal_shapes",         //[cite: 28]
			8: "spell_summon_natures_ally_8", //[cite: 28]
			9: "spell_shapechange",           //[cite: 28]
		},
	}

	DomainChaos = DomainBlueprint{
		ID:           "dom_chaos",
		Name:         "Chaos Domain",
		GrantedPower: "feat_domain_power_chaos", //[cite: 28]
		Spells: map[int]string{
			1: "spell_protection_from_law",      //[cite: 28]
			2: "spell_shatter",                  //[cite: 28]
			3: "spell_magic_circle_against_law", //[cite: 28]
			4: "spell_chaos_hammer",             //[cite: 28]
			5: "spell_dispel_law",               //[cite: 28]
			6: "spell_animate_objects",          //[cite: 28]
			7: "spell_word_of_chaos",            //[cite: 28]
			8: "spell_cloak_of_chaos",           //[cite: 28]
			9: "spell_summon_monster_9",         //[cite: 28]
		},
	}

	DomainDeath = DomainBlueprint{
		ID:           "dom_death",
		Name:         "Death Domain",
		GrantedPower: "feat_domain_power_death", //[cite: 28]
		Spells: map[int]string{
			1: "spell_cause_fear",            //[cite: 28]
			2: "spell_death_knell",           //[cite: 28]
			3: "spell_animate_dead",          //[cite: 28]
			4: "spell_death_ward",            //[cite: 28]
			5: "spell_slay_living",           //[cite: 28]
			6: "spell_create_undead",         //[cite: 28]
			7: "spell_destruction",           //[cite: 28]
			8: "spell_create_greater_undead", //[cite: 28]
			9: "spell_wail_of_the_banshee",   //[cite: 28]
		},
	}

	DomainDestruction = DomainBlueprint{
		ID:           "dom_destruction",
		Name:         "Destruction Domain",
		GrantedPower: "feat_domain_power_destruction", //[cite: 28]
		Spells: map[int]string{
			1: "spell_inflict_light_wounds",      //[cite: 28]
			2: "spell_shatter",                   //[cite: 28]
			3: "spell_contagion",                 //[cite: 28]
			4: "spell_inflict_critical_wounds",   //[cite: 28]
			5: "spell_inflict_light_wounds_mass", //[cite: 28]
			6: "spell_harm",                      //[cite: 28]
			7: "spell_disintegrate",              //[cite: 28]
			8: "spell_earthquake",                //[cite: 28]
			9: "spell_implosion",                 //[cite: 28]
		},
	}

	DomainEarth = DomainBlueprint{
		ID:           "dom_earth",
		Name:         "Earth Domain",
		GrantedPower: "feat_domain_power_earth", //[cite: 28]
		Spells: map[int]string{
			1: "spell_magic_stone",            //[cite: 28]
			2: "spell_soften_earth_and_stone", //[cite: 28]
			3: "spell_stone_shape",            //[cite: 28]
			4: "spell_spike_stones",           //[cite: 28]
			5: "spell_wall_of_stone",          //[cite: 28]
			6: "spell_stoneskin",              //[cite: 28]
			7: "spell_earthquake",             //[cite: 28]
			8: "spell_iron_body",              //[cite: 28]
			9: "spell_elemental_swarm",        //[cite: 28]
		},
	}

	DomainEvil = DomainBlueprint{
		ID:           "dom_evil",
		Name:         "Evil Domain",
		GrantedPower: "feat_domain_power_evil", //[cite: 28]
		Spells: map[int]string{
			1: "spell_protection_from_good",      //[cite: 28]
			2: "spell_desecrate",                 //[cite: 28]
			3: "spell_magic_circle_against_good", //[cite: 28]
			4: "spell_unholy_blight",             //[cite: 28]
			5: "spell_dispel_good",               //[cite: 28]
			6: "spell_create_undead",             //[cite: 28]
			7: "spell_blasphemy",                 //[cite: 28]
			8: "spell_unholy_aura",               //[cite: 28]
			9: "spell_summon_monster_9",          //[cite: 28]
		},
	}

	DomainFire = DomainBlueprint{
		ID:           "dom_fire",
		Name:         "Fire Domain",
		GrantedPower: "feat_domain_power_fire", //[cite: 28]
		Spells: map[int]string{
			1: "spell_burning_hands",    //[cite: 28]
			2: "spell_produce_flame",    //[cite: 28]
			3: "spell_resist_energy",    //[cite: 28]
			4: "spell_wall_of_fire",     //[cite: 28]
			5: "spell_fire_shield",      //[cite: 28]
			6: "spell_fire_seeds",       //[cite: 28]
			7: "spell_fire_storm",       //[cite: 28]
			8: "spell_incendiary_cloud", //[cite: 28]
			9: "spell_elemental_swarm",  //[cite: 28]
		},
	}

	DomainGood = DomainBlueprint{
		ID:           "dom_good",
		Name:         "Good Domain",
		GrantedPower: "feat_domain_power_good", //[cite: 28]
		Spells: map[int]string{
			1: "spell_protection_from_evil",      //[cite: 28]
			2: "spell_aid",                       //[cite: 28]
			3: "spell_magic_circle_against_evil", //[cite: 28]
			4: "spell_holy_smite",                //[cite: 28]
			5: "spell_dispel_evil",               //[cite: 28]
			6: "spell_blade_barrier",             //[cite: 28]
			7: "spell_holy_word",                 //[cite: 28]
			8: "spell_holy_aura",                 //[cite: 28]
			9: "spell_summon_monster_9",          //[cite: 28]
		},
	}

	DomainHealing = DomainBlueprint{
		ID:           "dom_healing",
		Name:         "Healing Domain",
		GrantedPower: "feat_domain_power_healing", //[cite: 28]
		Spells: map[int]string{
			1: "spell_cure_light_wounds",         //[cite: 28]
			2: "spell_cure_moderate_wounds",      //[cite: 28]
			3: "spell_cure_serious_wounds",       //[cite: 28]
			4: "spell_cure_critical_wounds",      //[cite: 28]
			5: "spell_cure_light_wounds_mass",    //[cite: 28]
			6: "spell_heal",                      //[cite: 28]
			7: "spell_regenerate",                //[cite: 28]
			8: "spell_cure_critical_wounds_mass", //[cite: 28]
			9: "spell_heal_mass",                 //[cite: 28]
		},
	}

	DomainKnowledge = DomainBlueprint{
		ID:           "dom_knowledge",
		Name:         "Knowledge Domain",
		GrantedPower: "feat_domain_power_knowledge", //[cite: 28]
		Spells: map[int]string{
			1: "spell_detect_secret_doors",        //[cite: 28]
			2: "spell_detect_thoughts",            //[cite: 28]
			3: "spell_clairaudience_clairvoyance", //[cite: 28]
			4: "spell_divination",                 //[cite: 28]
			5: "spell_true_seeing",                //[cite: 28]
			6: "spell_find_the_path",              //[cite: 28]
			7: "spell_legend_lore",                //[cite: 28]
			8: "spell_discern_location",           //[cite: 28]
			9: "spell_foresight",                  //[cite: 28]
		},
	}

	DomainLaw = DomainBlueprint{
		ID:           "dom_law",
		Name:         "Law Domain",
		GrantedPower: "feat_domain_power_law", //[cite: 28]
		Spells: map[int]string{
			1: "spell_protection_from_chaos",      //[cite: 28]
			2: "spell_calm_emotions",              //[cite: 28]
			3: "spell_magic_circle_against_chaos", //[cite: 28]
			4: "spell_orders_wrath",               //[cite: 28]
			5: "spell_dispel_chaos",               //[cite: 28]
			6: "spell_hold_monster",               //[cite: 28]
			7: "spell_dictum",                     //[cite: 28]
			8: "spell_shield_of_law",              //[cite: 28]
			9: "spell_summon_monster_9",           //[cite: 28]
		},
	}

	DomainLuck = DomainBlueprint{
		ID:           "dom_luck",
		Name:         "Luck Domain",
		GrantedPower: "feat_domain_power_luck", //[cite: 28]
		Spells: map[int]string{
			1: "spell_entropic_shield",        //[cite: 28]
			2: "spell_aid",                    //[cite: 28]
			3: "spell_protection_from_energy", //[cite: 28]
			4: "spell_freedom_of_movement",    //[cite: 28]
			5: "spell_break_enchantment",      //[cite: 28]
			6: "spell_mislead",                //[cite: 28]
			7: "spell_spell_turning",          //[cite: 28]
			8: "spell_moment_of_prescience",   //[cite: 28]
			9: "spell_miracle",                //[cite: 28]
		},
	}

	DomainMagic = DomainBlueprint{
		ID:           "dom_magic",
		Name:         "Magic Domain",
		GrantedPower: "feat_domain_power_magic", //[cite: 28]
		Spells: map[int]string{
			1: "spell_magic_aura",               //[cite: 28]
			2: "spell_identify",                 //[cite: 28]
			3: "spell_dispel_magic",             //[cite: 28]
			4: "spell_imbue_with_spell_ability", //[cite: 28]
			5: "spell_spell_resistance",         //[cite: 28]
			6: "spell_antimagic_field",          //[cite: 28]
			7: "spell_spell_turning",            //[cite: 28]
			8: "spell_protection_from_spells",   //[cite: 28]
			9: "spell_mages_disjunction",        //[cite: 28]
		},
	}

	DomainPlant = DomainBlueprint{
		ID:           "dom_plant",
		Name:         "Plant Domain",
		GrantedPower: "feat_domain_power_plant", //[cite: 28]
		Spells: map[int]string{
			1: "spell_entangle",       //[cite: 28]
			2: "spell_barkskin",       //[cite: 28]
			3: "spell_plant_growth",   //[cite: 28]
			4: "spell_command_plants", //[cite: 28]
			5: "spell_wall_of_thorns", //[cite: 28]
			6: "spell_repel_wood",     //[cite: 28]
			7: "spell_animate_plants", //[cite: 28]
			8: "spell_control_plants", //[cite: 28]
			9: "spell_shambler",       //[cite: 28]
		},
	}

	DomainProtection = DomainBlueprint{
		ID:           "dom_protection",
		Name:         "Protection Domain",
		GrantedPower: "feat_domain_power_protection", //[cite: 28]
		Spells: map[int]string{
			1: "spell_sanctuary",              //[cite: 28]
			2: "spell_shield_other",           //[cite: 28]
			3: "spell_protection_from_energy", //[cite: 28]
			4: "spell_spell_immunity",         //[cite: 28]
			5: "spell_spell_resistance",       //[cite: 28]
			6: "spell_antimagic_field",        //[cite: 28]
			7: "spell_repulsion",              //[cite: 28]
			8: "spell_mind_blank",             //[cite: 28]
			9: "spell_prismatic_sphere",       //[cite: 28]
		},
	}

	DomainStrength = DomainBlueprint{
		ID:           "dom_strength",
		Name:         "Strength Domain",
		GrantedPower: "feat_domain_power_strength", //[cite: 28]
		Spells: map[int]string{
			1: "spell_enlarge_person",  //[cite: 28]
			2: "spell_bulls_strength",  //[cite: 28]
			3: "spell_magic_vestment",  //[cite: 28]
			4: "spell_spell_immunity",  //[cite: 28]
			5: "spell_righteous_might", //[cite: 28]
			6: "spell_stoneskin",       //[cite: 28]
			7: "spell_grasping_hand",   //[cite: 28]
			8: "spell_clenched_fist",   //[cite: 28]
			9: "spell_crushing_hand",   //[cite: 28]
		},
	}

	DomainSun = DomainBlueprint{
		ID:           "dom_sun",
		Name:         "Sun Domain",
		GrantedPower: "feat_domain_power_sun", //[cite: 28]
		Spells: map[int]string{
			1: "spell_endure_elements",  //[cite: 28]
			2: "spell_heat_metal",       //[cite: 28]
			3: "spell_searing_light",    //[cite: 28]
			4: "spell_fire_shield",      //[cite: 28]
			5: "spell_flame_strike",     //[cite: 28]
			6: "spell_fire_seeds",       //[cite: 28]
			7: "spell_sunbeam",          //[cite: 28]
			8: "spell_sunburst",         //[cite: 28]
			9: "spell_prismatic_sphere", //[cite: 28]
		},
	}

	DomainTravel = DomainBlueprint{
		ID:           "dom_travel",
		Name:         "Travel Domain",
		GrantedPower: "feat_domain_power_travel", //[cite: 28]
		Spells: map[int]string{
			1: "spell_longstrider",       //[cite: 28]
			2: "spell_locate_object",     //[cite: 28]
			3: "spell_fly",               //[cite: 28]
			4: "spell_dimension_door",    //[cite: 28]
			5: "spell_teleport",          //[cite: 28]
			6: "spell_find_the_path",     //[cite: 28]
			7: "spell_teleport_greater",  //[cite: 28]
			8: "spell_phase_door",        //[cite: 28]
			9: "spell_astral_projection", //[cite: 28]
		},
	}

	DomainTrickery = DomainBlueprint{
		ID:           "dom_trickery",
		Name:         "Trickery Domain",
		GrantedPower: "feat_domain_power_trickery", //[cite: 28]
		Spells: map[int]string{
			1: "spell_disguise_self",        //[cite: 28]
			2: "spell_invisibility",         //[cite: 28]
			3: "spell_nondetection",         //[cite: 28]
			4: "spell_confusion",            //[cite: 28]
			5: "spell_false_vision",         //[cite: 28]
			6: "spell_mislead",              //[cite: 28]
			7: "spell_screen",               //[cite: 28]
			8: "spell_polymorph_any_object", //[cite: 28]
			9: "spell_time_stop",            //[cite: 28]
		},
	}

	DomainWar = DomainBlueprint{
		ID:           "dom_war",
		Name:         "War Domain",
		GrantedPower: "feat_domain_power_war", //[cite: 28]
		Spells: map[int]string{
			1: "spell_magic_weapon",     //[cite: 28]
			2: "spell_spiritual_weapon", //[cite: 28]
			3: "spell_magic_vestment",   //[cite: 28]
			4: "spell_divine_power",     //[cite: 28]
			5: "spell_flame_strike",     //[cite: 28]
			6: "spell_blade_barrier",    //[cite: 28]
			7: "spell_power_word_blind", //[cite: 28]
			8: "spell_power_word_stun",  //[cite: 28]
			9: "spell_power_word_kill",  //[cite: 28]
		},
	}

	DomainWater = DomainBlueprint{
		ID:           "dom_water",
		Name:         "Water Domain",
		GrantedPower: "feat_domain_power_water", //[cite: 28]
		Spells: map[int]string{
			1: "spell_obscuring_mist",  //[cite: 28]
			2: "spell_fog_cloud",       //[cite: 28]
			3: "spell_water_breathing", //[cite: 28]
			4: "spell_control_water",   //[cite: 28]
			5: "spell_ice_storm",       //[cite: 28]
			6: "spell_cone_of_cold",    //[cite: 28]
			7: "spell_acid_fog",        //[cite: 28]
			8: "spell_horrid_wilting",  //[cite: 28]
			9: "spell_elemental_swarm", //[cite: 28]
		},
	}

	// --- EXPANDED CLERIC DOMAINS ---

	DomainArtifice = DomainBlueprint{
		ID:           "dom_artifice",
		Name:         "Artifice Domain",
		GrantedPower: "feat_domain_power_artifice", //[cite: 28]
		Spells: map[int]string{
			1: "spell_animate_rope",     //[cite: 28]
			2: "spell_wood_shape",       //[cite: 28]
			3: "spell_stone_shape",      //[cite: 28]
			4: "spell_minor_creation",   //[cite: 28]
			5: "spell_fabricate",        //[cite: 28]
			6: "spell_major_creation",   //[cite: 28]
			7: "spell_hardening",        //[cite: 28]
			8: "spell_true_creation",    //[cite: 28]
			9: "spell_prismatic_sphere", //[cite: 28]
		},
	}

	DomainBlackwater = DomainBlueprint{
		ID:           "dom_blackwater",
		Name:         "Blackwater Domain",
		GrantedPower: "feat_domain_power_blackwater", //[cite: 28]
		Spells: map[int]string{
			1: "spell_cause_fear",                  //[cite: 28]
			2: "spell_pressure_sphere",             //[cite: 28]
			3: "spell_evards_black_tentacles",      //[cite: 28]
			4: "spell_transformation_of_the_deeps", //[cite: 28]
			5: "spell_blackwater_tentacle",         //[cite: 28]
			6: "spell_blackwater_taint",            //[cite: 28]
			7: "spell_dark_tide",                   //[cite: 28]
			8: "spell_maelstrom",                   //[cite: 28]
			9: "spell_doom_of_the_seas",            //[cite: 28]
		},
	}

	DomainCelerity = DomainBlueprint{
		ID:           "dom_celerity",
		Name:         "Celerity Domain",
		GrantedPower: "feat_domain_power_celerity", //[cite: 28]
		Spells: map[int]string{
			1: "spell_expeditious_retreat", //[cite: 28]
			2: "spell_cats_grace",          //[cite: 28]
			3: "spell_blur",                //[cite: 28]
			4: "spell_haste",               //[cite: 28]
			5: "spell_tree_stride",         //[cite: 28]
			6: "spell_wind_walk",           //[cite: 28]
			7: "spell_cats_grace_mass",     //[cite: 28]
			8: "spell_blink_improved",      //[cite: 28]
			9: "spell_time_stop",           //[cite: 28]
		},
	}

	DomainCelestial = DomainBlueprint{
		ID:           "dom_celestial",
		Name:         "Celestial Domain",
		GrantedPower: "feat_domain_power_celestial", //[cite: 28]
		Spells: map[int]string{
			1: "spell_vision_of_heaven",         //[cite: 28]
			2: "spell_consecrate",               //[cite: 28]
			3: "spell_blessed_sight",            //[cite: 28]
			4: "spell_lesser_planar_ally",       //[cite: 28]
			5: "spell_heavenly_lightning",       //[cite: 28]
			6: "spell_call_faithful_servants",   //[cite: 28]
			7: "spell_heavenly_lightning_storm", //[cite: 28]
			8: "spell_holy_aura",                //[cite: 28]
			9: "spell_gate",                     //[cite: 28]
		},
	}

	DomainCharm = DomainBlueprint{
		ID:           "dom_charm",
		Name:         "Charm Domain",
		GrantedPower: "feat_domain_power_charm", //[cite: 28]
		Spells: map[int]string{
			1: "spell_charm_person",     //[cite: 28]
			2: "spell_calm_emotions",    //[cite: 28]
			3: "spell_suggestion",       //[cite: 28]
			4: "spell_heroism",          //[cite: 28]
			5: "spell_charm_monster",    //[cite: 28]
			6: "spell_geas_quest",       //[cite: 28]
			7: "spell_insanity",         //[cite: 28]
			8: "spell_demand",           //[cite: 28]
			9: "spell_dominate_monster", //[cite: 28]
		},
	}

	DomainCity = DomainBlueprint{
		ID:           "dom_city",
		Name:         "City Domain",
		GrantedPower: "feat_domain_power_city", //[cite: 28]
		Spells: map[int]string{
			1: "spell_rooftop_strider",   //[cite: 28]
			2: "spell_city_lights",       //[cite: 28]
			3: "spell_winding_alleys",    //[cite: 28]
			4: "spell_commune_with_city", //[cite: 28]
			5: "spell_skyline_runner",    //[cite: 28]
			6: "spell_city_stride",       //[cite: 28]
			7: "spell_urban_shield",      //[cite: 28]
			8: "spell_citys_might",       //[cite: 28]
			9: "spell_animate_city",      //[cite: 28]
		},
	}

	DomainCold = DomainBlueprint{
		ID:           "dom_cold",
		Name:         "Cold Domain",
		GrantedPower: "feat_domain_power_cold", //[cite: 28]
		Spells: map[int]string{
			1: "spell_chill_touch",        //[cite: 28]
			2: "spell_chill_metal",        //[cite: 28]
			3: "spell_sleet_storm",        //[cite: 28]
			4: "spell_ice_storm",          //[cite: 28]
			5: "spell_wall_of_ice",        //[cite: 28]
			6: "spell_cone_of_cold",       //[cite: 28]
			7: "spell_control_weather",    //[cite: 28]
			8: "spell_polar_ray",          //[cite: 28]
			9: "spell_obedient_avalanche", //[cite: 28]
		},
	}

	DomainCommunity = DomainBlueprint{
		ID:           "dom_community",
		Name:         "Community Domain",
		GrantedPower: "feat_domain_power_community", //[cite: 28]
		Spells: map[int]string{
			1: "spell_bless",                             //[cite: 28]
			2: "spell_status",                            //[cite: 28]
			3: "spell_prayer",                            //[cite: 28]
			4: "spell_tongues",                           //[cite: 28]
			5: "spell_telepathic_bond",                   //[cite: 28]
			6: "spell_heroes_feast",                      //[cite: 28]
			7: "spell_refuge",                            //[cite: 28]
			8: "spell_mordenkainens_magnificent_mansion", //[cite: 28]
			9: "spell_mass_heal",                         //[cite: 28]
		},
	}

	DomainCompetition = DomainBlueprint{
		ID:           "dom_competition",
		Name:         "Competition Domain",
		GrantedPower: "feat_domain_power_competition", //[cite: 28]
		Spells: map[int]string{
			1: "spell_remove_fear",                 //[cite: 28]
			2: "spell_zeal",                        //[cite: 28]
			3: "spell_prayer",                      //[cite: 28]
			4: "spell_divine_power",                //[cite: 28]
			5: "spell_righteous_might",             //[cite: 28]
			6: "spell_zealot_pact",                 //[cite: 28]
			7: "spell_regenerate",                  //[cite: 28]
			8: "spell_moment_of_prescience",        //[cite: 28]
			9: "spell_visage_of_the_deity_greater", //[cite: 28]
		},
	}

	DomainCorruption = DomainBlueprint{
		ID:           "dom_corruption",
		Name:         "Corruption Domain",
		GrantedPower: "feat_domain_power_corruption", //[cite: 28]
		Spells: map[int]string{
			1: "spell_doom",               //[cite: 28]
			2: "spell_blindness_deafness", //[cite: 28]
			3: "spell_contagion",          //[cite: 28]
			4: "spell_morality_undone",    //[cite: 28]
			5: "spell_feeblemind",         //[cite: 28]
			6: "spell_pox",                //[cite: 28]
			7: "spell_insanity",           //[cite: 28]
			8: "spell_befoul",             //[cite: 28]
			9: "spell_despoil",            //[cite: 28]
		},
	}

	DomainCourage = DomainBlueprint{
		ID:           "dom_courage",
		Name:         "Courage Domain",
		GrantedPower: "feat_domain_power_courage", //[cite: 28]
		Spells: map[int]string{
			1: "spell_remove_fear",              //[cite: 28]
			2: "spell_aid",                      //[cite: 28]
			3: "spell_cloak_of_bravery",         //[cite: 28]
			4: "spell_heroism",                  //[cite: 28]
			5: "spell_valiant_fury",             //[cite: 28]
			6: "spell_heroes_feast",             //[cite: 28]
			7: "spell_heroism_greater",          //[cite: 28]
			8: "spell_lions_roar",               //[cite: 28]
			9: "spell_cloak_of_bravery_greater", //[cite: 28]
		},
	}

	DomainCreation = DomainBlueprint{
		ID:           "dom_creation",
		Name:         "Creation Domain",
		GrantedPower: "feat_domain_power_creation", //[cite: 28]
		Spells: map[int]string{
			1: "spell_create_water",          //[cite: 28]
			2: "spell_minor_image",           //[cite: 28]
			3: "spell_create_food_and_water", //[cite: 28]
			4: "spell_minor_creation",        //[cite: 28]
			5: "spell_major_creation",        //[cite: 28]
			6: "spell_heroes_feast",          //[cite: 28]
			7: "spell_permanent_image",       //[cite: 28]
			8: "spell_true_creation",         //[cite: 28]
			9: "spell_genesis",               //[cite: 28]
		},
	}

	DomainDarkness = DomainBlueprint{
		ID:           "dom_darkness",
		Name:         "Darkness Domain",
		GrantedPower: "feat_domain_power_darkness", //[cite: 28]
		Spells: map[int]string{
			1: "spell_obscuring_mist",    //[cite: 28]
			2: "spell_blindness",         //[cite: 28]
			3: "spell_blacklight",        //[cite: 28]
			4: "spell_armor_of_darkness", //[cite: 28]
			5: "spell_summon_monster_5",  //[cite: 28]
			6: "spell_prying_eyes",       //[cite: 28]
			7: "spell_nightmare",         //[cite: 28]
			8: "spell_power_word_blind",  //[cite: 28]
			9: "spell_power_word_kill",   //[cite: 28]
		},
	}

	DomainDeathbound = DomainBlueprint{
		ID:           "dom_deathbound",
		Name:         "Deathbound Domain",
		GrantedPower: "feat_domain_power_deathbound", //[cite: 28]
		Spells: map[int]string{
			1: "spell_chill_of_the_grave",        //[cite: 28]
			2: "spell_blade_of_pain_and_fear",    //[cite: 28]
			3: "spell_fangs_of_the_vampire_king", //[cite: 28]
			4: "spell_wither_limb",               //[cite: 28]
			5: "spell_revive_undead",             //[cite: 28]
			6: "spell_awaken_undead",             //[cite: 28]
			7: "spell_avasculate",                //[cite: 28]
			8: "spell_avascular_mass",            //[cite: 28]
			9: "spell_wail_of_the_banshee",       //[cite: 28]
		},
	}

	DomainDemonic = DomainBlueprint{
		ID:           "dom_demonic",
		Name:         "Demonic Domain",
		GrantedPower: "feat_domain_power_demonic", //[cite: 28]
		Spells: map[int]string{
			1: "spell_demonflesh",             //[cite: 28]
			2: "spell_demoncall",              //[cite: 28]
			3: "spell_demon_wings",            //[cite: 28]
			4: "spell_dimensional_anchor",     //[cite: 28]
			5: "spell_planar_binding_lesser",  //[cite: 28]
			6: "spell_planar_binding",         //[cite: 28]
			7: "spell_fiendish_clarity",       //[cite: 28]
			8: "spell_planar_binding_greater", //[cite: 28]
			9: "spell_gate",                   //[cite: 28]
		},
	}

	DomainDestiny = DomainBlueprint{
		ID:           "dom_destiny",
		Name:         "Destiny Domain",
		GrantedPower: "feat_domain_power_destiny", //[cite: 28]
		Spells: map[int]string{
			1: "spell_omen_of_peril",        //[cite: 28]
			2: "spell_augury",               //[cite: 28]
			3: "spell_delay_death",          //[cite: 28]
			4: "spell_bestow_curse",         //[cite: 28]
			5: "spell_stalwart_pact",        //[cite: 28]
			6: "spell_warp_destiny",         //[cite: 28]
			7: "spell_bestow_curse_greater", //[cite: 28]
			8: "spell_moment_of_prescience", //[cite: 28]
			9: "spell_choose_destiny",       //[cite: 28]
		},
	}

	DomainDiabolic = DomainBlueprint{
		ID:           "dom_diabolic",
		Name:         "Diabolic Domain",
		GrantedPower: "feat_domain_power_diabolic", //[cite: 28]
		Spells: map[int]string{
			1: "spell_protection_from_good",         //[cite: 28]
			2: "spell_devils_eye",                   //[cite: 28]
			3: "spell_devils_ego",                   //[cite: 28]
			4: "spell_hellfire",                     //[cite: 28]
			5: "spell_planar_binding_lesser",        //[cite: 28]
			6: "spell_planar_binding",               //[cite: 28]
			7: "spell_hellfire_storm",               //[cite: 28]
			8: "spell_demand",                       //[cite: 28]
			9: "spell_investiture_of_the_pit_fiend", //[cite: 28]
		},
	}

	DomainDomination = DomainBlueprint{
		ID:           "dom_domination",
		Name:         "Domination Domain",
		GrantedPower: "feat_domain_power_domination", //[cite: 28]
		Spells: map[int]string{
			1: "spell_command",          //[cite: 28]
			2: "spell_enthrall",         //[cite: 28]
			3: "spell_suggestion",       //[cite: 28]
			4: "spell_dominate_person",  //[cite: 28]
			5: "spell_greater_command",  //[cite: 28]
			6: "spell_geas_quest",       //[cite: 28]
			7: "spell_suggestion_mass",  //[cite: 28]
			8: "spell_true_domination",  //[cite: 28]
			9: "spell_monstrous_thrall", //[cite: 28]
		},
	}

	DomainDragon = DomainBlueprint{
		ID:           "dom_dragon",
		Name:         "Dragon Domain",
		GrantedPower: "feat_domain_power_dragon", //[cite: 28]
		Spells: map[int]string{
			1: "spell_magic_fang",          //[cite: 28]
			2: "spell_resist_energy",       //[cite: 28]
			3: "spell_magic_fang_greater",  //[cite: 28]
			4: "spell_voice_of_the_dragon", //[cite: 28]
			5: "spell_true_seeing",         //[cite: 28]
			6: "spell_stoneskin",           //[cite: 28]
			7: "spell_dragon_ally",         //[cite: 28]
			8: "spell_suggestion_mass",     //[cite: 28]
			9: "spell_dominate_monster",    //[cite: 28]
		},
	}

	DomainDream = DomainBlueprint{
		ID:           "dom_dream",
		Name:         "Dream Domain",
		GrantedPower: "feat_domain_power_dream", //[cite: 28]
		Spells: map[int]string{
			1: "spell_sleep",             //[cite: 28]
			2: "spell_augury",            //[cite: 28]
			3: "spell_deep_slumber",      //[cite: 28]
			4: "spell_phantasmal_killer", //[cite: 28]
			5: "spell_nightmare",         //[cite: 28]
			6: "spell_dream_sight",       //[cite: 28]
			7: "spell_scrying_greater",   //[cite: 28]
			8: "spell_power_word_stun",   //[cite: 28]
			9: "spell_weird",             //[cite: 28]
		},
	}

	DomainEndurance = DomainBlueprint{
		ID:           "dom_endurance",
		Name:         "Endurance Domain",
		GrantedPower: "feat_domain_power_endurance", //[cite: 28]
		Spells: map[int]string{
			1: "spell_endure_elements",          //[cite: 28]
			2: "spell_bears_endurance",          //[cite: 28]
			3: "spell_refreshment",              //[cite: 28]
			4: "spell_sustain",                  //[cite: 28]
			5: "spell_stoneskin",                //[cite: 28]
			6: "spell_mass_bears_endurance",     //[cite: 28]
			7: "spell_globe_of_invulnerability", //[cite: 28]
			8: "spell_spell_turning",            //[cite: 28]
			9: "spell_iron_body",                //[cite: 28]
		},
	}

	DomainEntropy = DomainBlueprint{
		ID:           "dom_entropy",
		Name:         "Entropy Domain",
		GrantedPower: "feat_domain_power_entropy", //[cite: 28]
		Spells: map[int]string{
			1: "spell_cause_fear",            //[cite: 28]
			2: "spell_vision_of_entropy",     //[cite: 28]
			3: "spell_ray_of_exhaustion",     //[cite: 28]
			4: "spell_fear",                  //[cite: 28]
			5: "spell_waves_of_fatigue",      //[cite: 28]
			6: "spell_disintegrate",          //[cite: 28]
			7: "spell_insanity",              //[cite: 28]
			8: "spell_scintillating_pattern", //[cite: 28]
			9: "spell_abyssal_rift",          //[cite: 28]
		},
	}

	DomainFate = DomainBlueprint{
		ID:           "dom_fate",
		Name:         "Fate Domain",
		GrantedPower: "feat_domain_power_fate", //[cite: 28]
		Spells: map[int]string{
			1: "spell_true_strike",     //[cite: 28]
			2: "spell_augury",          //[cite: 28]
			3: "spell_bestow_curse",    //[cite: 28]
			4: "spell_status",          //[cite: 28]
			5: "spell_mark_of_justice", //[cite: 28]
			6: "spell_geas_quest",      //[cite: 28]
			7: "spell_vision",          //[cite: 28]
			8: "spell_mind_blank",      //[cite: 28]
			9: "spell_foresight",       //[cite: 28]
		},
	}

	DomainFey = DomainBlueprint{
		ID:           "dom_fey",
		Name:         "Fey Domain",
		GrantedPower: "feat_domain_power_fey", //[cite: 28]
		Spells: map[int]string{
			1: "spell_faerie_fire",           //[cite: 28]
			2: "spell_charm_person",          //[cite: 28]
			3: "spell_inspired_aim",          //[cite: 28]
			4: "spell_blinding_beauty",       //[cite: 28]
			5: "spell_tree_stride",           //[cite: 28]
			6: "spell_heroes_feast",          //[cite: 28]
			7: "spell_liveoak",               //[cite: 28]
			8: "spell_unearthly_beauty",      //[cite: 28]
			9: "spell_summon_natures_ally_9", //[cite: 28]
		},
	}

	DomainForce = DomainBlueprint{
		ID:           "dom_force",
		Name:         "Force Domain",
		GrantedPower: "feat_domain_power_force", //[cite: 28]
		Spells: map[int]string{
			1: "spell_mage_armor",                  //[cite: 28]
			2: "spell_magic_missile",               //[cite: 28]
			3: "spell_blast_of_force",              //[cite: 28]
			4: "spell_otilukes_resilient_sphere",   //[cite: 28]
			5: "spell_wall_of_force",               //[cite: 28]
			6: "spell_repulsion",                   //[cite: 28]
			7: "spell_forcecage",                   //[cite: 28]
			8: "spell_otilukes_telekinetic_sphere", //[cite: 28]
			9: "spell_bigbys_crushing_hand",        //[cite: 28]
		},
	}

	DomainFury = DomainBlueprint{
		ID:           "dom_fury",
		Name:         "Fury Domain",
		GrantedPower: "feat_domain_power_fury", //[cite: 28]
		Spells: map[int]string{
			1: "spell_true_strike",         //[cite: 28]
			2: "spell_bulls_strength",      //[cite: 28]
			3: "spell_rage",                //[cite: 28]
			4: "spell_divine_power",        //[cite: 28]
			5: "spell_shout",               //[cite: 28]
			6: "spell_song_of_discord",     //[cite: 28]
			7: "spell_abyssal_frenzy",      //[cite: 28]
			8: "spell_shout_greater",       //[cite: 28]
			9: "spell_abyssal_frenzy_mass", //[cite: 28]
		},
	}

	DomainGlory = DomainBlueprint{
		ID:           "dom_glory",
		Name:         "Glory Domain",
		GrantedPower: "feat_domain_power_glory", //[cite: 28]
		Spells: map[int]string{
			1: "spell_disrupt_undead", //[cite: 28]
			2: "spell_bless_weapon",   //[cite: 28]
			3: "spell_searing_light",  //[cite: 28]
			4: "spell_holy_smite",     //[cite: 28]
			5: "spell_holy_sword",     //[cite: 28]
			6: "spell_bolt_of_glory",  //[cite: 28]
			7: "spell_sunbeam",        //[cite: 28]
			8: "spell_crown_of_glory", //[cite: 28]
			9: "spell_gate",           //[cite: 28]
		},
	}

	DomainGloryBOED = DomainBlueprint{
		ID:           "dom_glory_boed",
		Name:         "Glory Domain (BoED)",
		GrantedPower: "feat_domain_power_glory_boed", //[cite: 28]
		Spells: map[int]string{
			1: "spell_disrupt_undead",       //[cite: 28]
			2: "spell_glorious_raiment",     //[cite: 28]
			3: "spell_searing_light",        //[cite: 28]
			4: "spell_celestial_brilliance", //[cite: 28]
			5: "spell_crown_of_flame",       //[cite: 28]
			6: "spell_bolt_of_glory",        //[cite: 28]
			7: "spell_sunbeam",              //[cite: 28]
			8: "spell_crown_of_glory",       //[cite: 28]
			9: "spell_blinding_glory",       //[cite: 28]
		},
	}

	DomainGreed = DomainBlueprint{
		ID:           "dom_greed",
		Name:         "Greed Domain",
		GrantedPower: "feat_domain_power_greed", //[cite: 28]
		Spells: map[int]string{
			1: "spell_cheat",            //[cite: 28]
			2: "spell_entice_gift",      //[cite: 28]
			3: "spell_knock",            //[cite: 28]
			4: "spell_fire_trap",        //[cite: 28]
			5: "spell_fabricate",        //[cite: 28]
			6: "spell_guards_and_wards", //[cite: 28]
			7: "spell_teleport_object",  //[cite: 28]
			8: "spell_phantasmal_thief", //[cite: 28]
			9: "spell_sympathy",         //[cite: 28]
		},
	}

	DomainHatred = DomainBlueprint{
		ID:           "dom_hatred",
		Name:         "Hatred Domain",
		GrantedPower: "feat_domain_power_hatred", //[cite: 28]
		Spells: map[int]string{
			1: "spell_doom",                //[cite: 28]
			2: "spell_scare",               //[cite: 28]
			3: "spell_bestow_curse",        //[cite: 28]
			4: "spell_song_of_discord",     //[cite: 28]
			5: "spell_righteous_might",     //[cite: 28]
			6: "spell_forbiddance",         //[cite: 28]
			7: "spell_blasphemy",           //[cite: 28]
			8: "spell_antipathy",           //[cite: 28]
			9: "spell_wail_of_the_banshee", //[cite: 28]
		},
	}

	DomainHerald = DomainBlueprint{
		ID:           "dom_herald",
		Name:         "Herald Domain",
		GrantedPower: "feat_domain_power_herald", //[cite: 28]
		Spells: map[int]string{
			1: "spell_comprehend_languages",        //[cite: 28]
			2: "spell_enthrall",                    //[cite: 28]
			3: "spell_tongues",                     //[cite: 28]
			4: "spell_sending",                     //[cite: 28]
			5: "spell_greater_command",             //[cite: 28]
			6: "spell_dream",                       //[cite: 28]
			7: "spell_aspect_of_the_deity",         //[cite: 28]
			8: "spell_crown_of_glory",              //[cite: 28]
			9: "spell_greater_aspect_of_the_deity", //[cite: 28]
		},
	}

	DomainHunger = DomainBlueprint{
		ID:           "dom_hunger",
		Name:         "Hunger Domain",
		GrantedPower: "feat_domain_power_hunger", //[cite: 28]
		Spells: map[int]string{
			1: "spell_ghoul_light",      //[cite: 28]
			2: "spell_ghoul_glyph",      //[cite: 28]
			3: "spell_ghoul_gesture",    //[cite: 28]
			4: "spell_enervation",       //[cite: 28]
			5: "spell_ghoul_gauntlet",   //[cite: 28]
			6: "spell_eyes_of_the_king", //[cite: 28]
			7: "spell_field_of_ghouls",  //[cite: 28]
			8: "spell_bite_of_the_king", //[cite: 28]
			9: "spell_energy_drain",     //[cite: 28]
		},
	}

	DomainIncarnum = DomainBlueprint{
		ID:           "dom_incarnum",
		Name:         "Incarnum Domain",
		GrantedPower: "feat_domain_power_incarnum", //[cite: 28]
		Spells: map[int]string{
			1: "spell_detect_incarnum",      //[cite: 28]
			2: "spell_soul_boon",            //[cite: 28]
			3: "spell_wall_of_incarnum",     //[cite: 28]
			4: "spell_essentia_lock",        //[cite: 28]
			5: "spell_incarnum_weapon",      //[cite: 28]
			6: "spell_incarnum_vigor",       //[cite: 28]
			7: "spell_incarnum_bladestorm",  //[cite: 28]
			8: "spell_incarnum_apotheosis",  //[cite: 28]
			9: "spell_soulmeld_disjunction", //[cite: 28]
		},
	}

	DomainInquisition = DomainBlueprint{
		ID:           "dom_inquisition",
		Name:         "Inquisition Domain",
		GrantedPower: "feat_domain_power_inquisition", //[cite: 28]
		Spells: map[int]string{
			1: "spell_detect_chaos",    //[cite: 28]
			2: "spell_zone_of_truth",   //[cite: 28]
			3: "spell_detect_thoughts", //[cite: 28]
			4: "spell_discern_lies",    //[cite: 28]
			5: "spell_true_seeing",     //[cite: 28]
			6: "spell_geas_quest",      //[cite: 28]
			7: "spell_dictum",          //[cite: 28]
			8: "spell_shield_of_law",   //[cite: 28]
			9: "spell_imprisonment",    //[cite: 28]
		},
	}

	DomainJoy = DomainBlueprint{
		ID:           "dom_joy",
		Name:         "Joy Domain",
		GrantedPower: "feat_domain_power_joy", //[cite: 28]
		Spells: map[int]string{
			1: "spell_vision_of_heaven",         //[cite: 28]
			2: "spell_elation",                  //[cite: 28]
			3: "spell_distilled_joy",            //[cite: 28]
			4: "spell_good_hope",                //[cite: 28]
			5: "spell_chaavs_laugh",             //[cite: 28]
			6: "spell_greater_heroism",          //[cite: 28]
			7: "spell_starmantle",               //[cite: 28]
			8: "spell_sympathy",                 //[cite: 28]
			9: "spell_ottos_irresistible_dance", //[cite: 28]
		},
	}

	DomainLiberation = DomainBlueprint{
		ID:           "dom_liberation",
		Name:         "Liberation Domain",
		GrantedPower: "feat_domain_power_liberation", //[cite: 28]
		Spells: map[int]string{
			1: "spell_remove_fear",          //[cite: 28]
			2: "spell_remove_paralysis",     //[cite: 28]
			3: "spell_remove_curse",         //[cite: 28]
			4: "spell_freedom_of_movement",  //[cite: 28]
			5: "spell_break_enchantment",    //[cite: 28]
			6: "spell_greater_dispel_magic", //[cite: 28]
			7: "spell_refuge",               //[cite: 28]
			8: "spell_mind_blank",           //[cite: 28]
			9: "spell_freedom",              //[cite: 28]
		},
	}

	DomainMadness = DomainBlueprint{
		ID:           "dom_madness",
		Name:         "Madness Domain",
		GrantedPower: "feat_domain_power_madness", //[cite: 28]
		Spells: map[int]string{
			1: "spell_confusion_lesser",     //[cite: 28]
			2: "spell_touch_of_madness",     //[cite: 28]
			3: "spell_rage",                 //[cite: 28]
			4: "spell_confusion",            //[cite: 28]
			5: "spell_bolts_of_bedevilment", //[cite: 28]
			6: "spell_phantasmal_killer",    //[cite: 28]
			7: "spell_insanity",             //[cite: 28]
			8: "spell_maddening_scream",     //[cite: 28]
			9: "spell_weird",                //[cite: 28]
		},
	}

	DomainMind = DomainBlueprint{
		ID:           "dom_mind",
		Name:         "Mind Domain",
		GrantedPower: "feat_domain_power_mind", //[cite: 28]
		Spells: map[int]string{
			1: "spell_comprehend_languages",   //[cite: 28]
			2: "spell_detect_thoughts",        //[cite: 28]
			3: "spell_lesser_telepathic_bond", //[cite: 28]
			4: "spell_discern_lies",           //[cite: 28]
			5: "spell_rarys_telepathic_bond",  //[cite: 28]
			6: "spell_probe_thoughts",         //[cite: 28]
			7: "spell_brain_spider",           //[cite: 28]
			8: "spell_mind_blank",             //[cite: 28]
			9: "spell_weird",                  //[cite: 28]
		},
	}

	DomainMysticism = DomainBlueprint{
		ID:           "dom_mysticism",
		Name:         "Mysticism Domain",
		GrantedPower: "feat_domain_power_mysticism", //[cite: 28]
		Spells: map[int]string{
			1: "spell_divine_favor",                //[cite: 28]
			2: "spell_spiritual_weapon",            //[cite: 28]
			3: "spell_visage_of_the_deity_lesser",  //[cite: 28]
			4: "spell_weapon_of_the_deity",         //[cite: 28]
			5: "spell_righteous_might",             //[cite: 28]
			6: "spell_visage_of_the_deity",         //[cite: 28]
			7: "spell_blasphemy_holy_word",         //[cite: 28]
			8: "spell_holy_aura_unholy_aura",       //[cite: 28]
			9: "spell_visage_of_the_deity_greater", //[cite: 28]
		},
	}

	DomainNobility = DomainBlueprint{
		ID:           "dom_nobility",
		Name:         "Nobility Domain",
		GrantedPower: "feat_domain_power_nobility", //[cite: 28]
		Spells: map[int]string{
			1: "spell_divine_favor",       //[cite: 28]
			2: "spell_enthrall",           //[cite: 28]
			3: "spell_magic_vestment",     //[cite: 28]
			4: "spell_discern_lies",       //[cite: 28]
			5: "spell_greater_command",    //[cite: 28]
			6: "spell_geas_quest",         //[cite: 28]
			7: "spell_repulsion",          //[cite: 28]
			8: "spell_demand",             //[cite: 28]
			9: "spell_storm_of_vengeance", //[cite: 28]
		},
	}

	DomainOcean = DomainBlueprint{
		ID:           "dom_ocean",
		Name:         "Ocean Domain",
		GrantedPower: "feat_domain_power_ocean", //[cite: 28]
		Spells: map[int]string{
			1: "spell_endure_elements",          //[cite: 28]
			2: "spell_sound_burst",              //[cite: 28]
			3: "spell_water_breathing",          //[cite: 28]
			4: "spell_freedom_of_movement",      //[cite: 28]
			5: "spell_wall_of_ice",              //[cite: 28]
			6: "spell_otilukes_freezing_sphere", //[cite: 28]
			7: "spell_waterspout",               //[cite: 28]
			8: "spell_maelstrom",                //[cite: 28]
			9: "spell_elemental_swarm",          //[cite: 28]
		},
	}

	DomainOoze = DomainBlueprint{
		ID:           "dom_ooze",
		Name:         "Ooze Domain",
		GrantedPower: "feat_domain_power_ooze", //[cite: 28]
		Spells: map[int]string{
			1: "spell_grease",                //[cite: 28]
			2: "spell_web",                   //[cite: 28]
			3: "spell_poison",                //[cite: 28]
			4: "spell_rusting_grasp",         //[cite: 28]
			5: "spell_oozepuppet",            //[cite: 28]
			6: "spell_transmute_rock_to_mud", //[cite: 28]
			7: "spell_slime_wave",            //[cite: 28]
			8: "spell_befoul",                //[cite: 28]
			9: "spell_implosion",             //[cite: 28]
		},
	}

	DomainOracle = DomainBlueprint{
		ID:           "dom_oracle",
		Name:         "Oracle Domain",
		GrantedPower: "feat_domain_power_oracle", //[cite: 28]
		Spells: map[int]string{
			1: "spell_identify",         //[cite: 28]
			2: "spell_augury",           //[cite: 28]
			3: "spell_divination",       //[cite: 28]
			4: "spell_scrying",          //[cite: 28]
			5: "spell_commune",          //[cite: 28]
			6: "spell_legend_lore",      //[cite: 28]
			7: "spell_scrying_greater",  //[cite: 28]
			8: "spell_discern_location", //[cite: 28]
			9: "spell_foresight",        //[cite: 28]
		},
	}

	DomainPact = DomainBlueprint{
		ID:           "dom_pact",
		Name:         "Pact Domain",
		GrantedPower: "feat_domain_power_pact", //[cite: 28]
		Spells: map[int]string{
			1: "spell_command",         //[cite: 28]
			2: "spell_shield_other",    //[cite: 28]
			3: "spell_speak_with_dead", //[cite: 28]
			4: "spell_divination",      //[cite: 28]
			5: "spell_stalwart_pact",   //[cite: 28]
			6: "spell_zealot_pact",     //[cite: 28]
			7: "spell_renewal_pact",    //[cite: 28]
			8: "spell_death_pact",      //[cite: 28]
			9: "spell_gate",            //[cite: 28]
		},
	}

	DomainPestilence = DomainBlueprint{
		ID:           "dom_pestilence",
		Name:         "Pestilence Domain",
		GrantedPower: "feat_domain_power_pestilence", //[cite: 28]
		Spells: map[int]string{
			1: "spell_doom",                  //[cite: 28]
			2: "spell_summon_swarm",          //[cite: 28]
			3: "spell_contagion",             //[cite: 28]
			4: "spell_poison",                //[cite: 28]
			5: "spell_plague_of_rats",        //[cite: 28]
			6: "spell_curse_of_lycanthropy",  //[cite: 28]
			7: "spell_scourge",               //[cite: 28]
			8: "spell_create_greater_undead", //[cite: 28]
			9: "spell_otyugh_swarm",          //[cite: 28]
		},
	}

	DomainPlanning = DomainBlueprint{
		ID:           "dom_planning",
		Name:         "Planning Domain",
		GrantedPower: "feat_domain_power_planning", //[cite: 28]
		Spells: map[int]string{
			1: "spell_deathwatch",                 //[cite: 28]
			2: "spell_augury",                     //[cite: 28]
			3: "spell_clairaudience_clairvoyance", //[cite: 28]
			4: "spell_status",                     //[cite: 28]
			5: "spell_detect_scrying",             //[cite: 28]
			6: "spell_heroes_feast",               //[cite: 28]
			7: "spell_scrying_greater",            //[cite: 28]
			8: "spell_discern_location",           //[cite: 28]
			9: "spell_time_stop",                  //[cite: 28]
		},
	}

	DomainPleasure = DomainBlueprint{
		ID:           "dom_pleasure",
		Name:         "Pleasure Domain",
		GrantedPower: "feat_domain_power_pleasure", //[cite: 28]
		Spells: map[int]string{
			1: "spell_remove_fear",           //[cite: 28]
			2: "spell_lastais_caress",        //[cite: 28]
			3: "spell_hearts_ease",           //[cite: 28]
			4: "spell_remove_fatigue",        //[cite: 28]
			5: "spell_mass_eagles_splendor",  //[cite: 28]
			6: "spell_celestial_blood",       //[cite: 28]
			7: "spell_empyreal_ecstasy",      //[cite: 28]
			8: "spell_spread_of_contentment", //[cite: 28]
			9: "spell_sublime_revelry",       //[cite: 28]
		},
	}

	DomainPurification = DomainBlueprint{
		ID:           "dom_purification",
		Name:         "Purification Domain",
		GrantedPower: "feat_domain_power_purification", //[cite: 28]
		Spells: map[int]string{
			1: "spell_nimbus_of_light",                 //[cite: 28]
			2: "spell_deific_vengeance",                //[cite: 28]
			3: "spell_recitation",                      //[cite: 28]
			4: "spell_castigate",                       //[cite: 28]
			5: "spell_dance_of_the_unicorn",            //[cite: 28]
			6: "spell_fires_of_purity",                 //[cite: 28]
			7: "spell_righteous_wrath_of_the_faithful", //[cite: 28]
			8: "spell_sunburst",                        //[cite: 28]
			9: "spell_visage_of_the_deity_greater",     //[cite: 28]
		},
	}

	DomainRepose = DomainBlueprint{
		ID:           "dom_repose",
		Name:         "Repose Domain",
		GrantedPower: "feat_domain_power_repose", //[cite: 28]
		Spells: map[int]string{
			1: "spell_deathwatch",          //[cite: 28]
			2: "spell_gentle_repose",       //[cite: 28]
			3: "spell_speak_with_dead",     //[cite: 28]
			4: "spell_death_ward",          //[cite: 28]
			5: "spell_slay_living",         //[cite: 28]
			6: "spell_undeath_to_death",    //[cite: 28]
			7: "spell_destruction",         //[cite: 28]
			8: "spell_surelife",            //[cite: 28]
			9: "spell_wail_of_the_banshee", //[cite: 28]
		},
	}

	DomainRune = DomainBlueprint{
		ID:           "dom_rune",
		Name:         "Rune Domain",
		GrantedPower: "feat_domain_power_rune", //[cite: 28]
		Spells: map[int]string{
			1: "spell_erase",                    //[cite: 28]
			2: "spell_secret_page",              //[cite: 28]
			3: "spell_glyph_of_warding",         //[cite: 28]
			4: "spell_explosive_runes",          //[cite: 28]
			5: "spell_lesser_planar_binding",    //[cite: 28]
			6: "spell_greater_glyph_of_warding", //[cite: 28]
			7: "spell_instant_summons",          //[cite: 28]
			8: "spell_symbol_of_death",          //[cite: 28]
			9: "spell_teleportation_circle",     //[cite: 28]
		},
	}

	DomainSand = DomainBlueprint{
		ID:           "dom_sand",
		Name:         "Sand Domain",
		GrantedPower: "feat_domain_power_sand", //[cite: 28]
		Spells: map[int]string{
			1: "spell_waste_strider",        //[cite: 28]
			2: "spell_black_sand",           //[cite: 28]
			3: "spell_haboob",               //[cite: 28]
			4: "spell_blast_of_sand",        //[cite: 28]
			5: "spell_flaywind_burst",       //[cite: 28]
			6: "spell_awaken_sand",          //[cite: 28]
			7: "spell_vitrify",              //[cite: 28]
			8: "spell_desert_binding",       //[cite: 28]
			9: "spell_summon_desert_ally_9", //[cite: 28]
		},
	}

	DomainSeafolk = DomainBlueprint{
		ID:           "dom_seafolk",
		Name:         "Seafolk Domain",
		GrantedPower: "feat_domain_power_seafolk", //[cite: 28]
		Spells: map[int]string{
			1: "spell_quickswim",             //[cite: 28]
			2: "spell_fins_to_feet",          //[cite: 28]
			3: "spell_scales_of_the_sealord", //[cite: 28]
			4: "spell_sirens_call",           //[cite: 28]
			5: "spell_commune_with_nature",   //[cite: 28]
			6: "spell_airy_water",            //[cite: 28]
			7: "spell_megalodon_empowerment", //[cite: 28]
			8: "spell_depthsurge",            //[cite: 28]
			9: "spell_foresight",             //[cite: 28]
		},
	}

	DomainSkalykind = DomainBlueprint{
		ID:           "dom_skalykind",
		Name:         "Skalykind Domain",
		GrantedPower: "feat_domain_power_skalykind", //[cite: 28]
		Spells: map[int]string{
			1: "spell_magic_fang",         //[cite: 28]
			2: "spell_animal_trance",      //[cite: 28]
			3: "spell_greater_magic_fang", //[cite: 28]
			4: "spell_poison",             //[cite: 28]
			5: "spell_animal_growth",      //[cite: 28]
			6: "spell_eyebite",            //[cite: 28]
			7: "spell_creeping_doom",      //[cite: 28]
			8: "spell_animal_shapes",      //[cite: 28]
			9: "spell_shapechange",        //[cite: 28]
		},
	}

	DomainSky = DomainBlueprint{
		ID:           "dom_sky",
		Name:         "Sky Domain",
		GrantedPower: "feat_domain_power_sky", //[cite: 28]
		Spells: map[int]string{
			1: "spell_raptors_sight",      //[cite: 28]
			2: "spell_summon_dire_hawk",   //[cite: 28]
			3: "spell_enduring_flight",    //[cite: 28]
			4: "spell_aerial_alacrity",    //[cite: 28]
			5: "spell_control_winds",      //[cite: 28]
			6: "spell_wind_walk",          //[cite: 28]
			7: "spell_reverse_gravity",    //[cite: 28]
			8: "spell_mastery_of_the_sky", //[cite: 28]
			9: "spell_summon_devoted_roc", //[cite: 28]
		},
	}

	DomainSpite = DomainBlueprint{
		ID:           "dom_spite",
		Name:         "Spite Domain",
		GrantedPower: "feat_domain_power_spite", //[cite: 28]
		Spells: map[int]string{
			1: "spell_bestow_wound",          //[cite: 28]
			2: "spell_rage",                  //[cite: 28]
			3: "spell_vampiric_touch",        //[cite: 28]
			4: "spell_pronouncement_of_fate", //[cite: 28]
			5: "spell_fire_in_the_blood",     //[cite: 28]
			6: "spell_cloak_of_hate",         //[cite: 28]
			7: "spell_pact_of_return",        //[cite: 28]
			8: "spell_mantle_of_pure_spite",  //[cite: 28]
			9: "spell_imprison_soul",         //[cite: 28]
		},
	}

	DomainStorm = DomainBlueprint{
		ID:           "dom_storm",
		Name:         "Storm Domain",
		GrantedPower: "feat_domain_power_storm", //[cite: 28]
		Spells: map[int]string{
			1: "spell_entropic_shield",      //[cite: 28]
			2: "spell_gust_of_wind",         //[cite: 28]
			3: "spell_call_lightning",       //[cite: 28]
			4: "spell_sleet_storm",          //[cite: 28]
			5: "spell_ice_storm",            //[cite: 28]
			6: "spell_call_lightning_storm", //[cite: 28]
			7: "spell_control_weather",      //[cite: 28]
			8: "spell_whirlwind",            //[cite: 28]
			9: "spell_storm_of_vengeance",   //[cite: 28]
		},
	}

	DomainSummer = DomainBlueprint{
		ID:           "dom_summer",
		Name:         "Summer Domain",
		GrantedPower: "feat_domain_power_summer", //[cite: 28]
		Spells: map[int]string{
			1: "spell_impede_suns_brilliance",      //[cite: 28]
			2: "spell_sunstroke",                   //[cite: 28]
			3: "spell_protection_from_dessication", //[cite: 28]
			4: "spell_skin_of_the_cactus",          //[cite: 28]
			5: "spell_unearthly_heat",              //[cite: 28]
			6: "spell_sunbeam",                     //[cite: 28]
			7: "spell_control_weather",             //[cite: 28]
			8: "spell_sunburst",                    //[cite: 28]
			9: "spell_storm_of_vengeance",          //[cite: 28]
		},
	}

	DomainSummoner = DomainBlueprint{
		ID:           "dom_summoner",
		Name:         "Summoner Domain",
		GrantedPower: "feat_domain_power_summoner", //[cite: 28]
		Spells: map[int]string{
			1: "spell_summon_monster_1",    //[cite: 28]
			2: "spell_summon_monster_2",    //[cite: 28]
			3: "spell_summon_monster_3",    //[cite: 28]
			4: "spell_lesser_planar_ally",  //[cite: 28]
			5: "spell_summon_monster_5",    //[cite: 28]
			6: "spell_planar_ally",         //[cite: 28]
			7: "spell_summon_monster_7",    //[cite: 28]
			8: "spell_greater_planar_ally", //[cite: 28]
			9: "spell_gate",                //[cite: 28]
		},
	}

	DomainTemptation = DomainBlueprint{
		ID:           "dom_temptation",
		Name:         "Temptation Domain",
		GrantedPower: "feat_domain_power_temptation", //[cite: 28]
		Spells: map[int]string{
			1: "spell_charm_person",     //[cite: 28]
			2: "spell_beckoning_call",   //[cite: 28]
			3: "spell_suggestion",       //[cite: 28]
			4: "spell_charm_monster",    //[cite: 28]
			5: "spell_dominate_person",  //[cite: 28]
			6: "spell_mass_suggestion",  //[cite: 28]
			7: "spell_soul_link",        //[cite: 28]
			8: "spell_sympathy",         //[cite: 28]
			9: "spell_dominate_monster", //[cite: 28]
		},
	}

	DomainThirst = DomainBlueprint{
		ID:           "dom_thirst",
		Name:         "Thirst Domain",
		GrantedPower: "feat_domain_power_thirst", //[cite: 28]
		Spells: map[int]string{
			1: "spell_parching_touch",    //[cite: 28]
			2: "spell_desiccate",         //[cite: 28]
			3: "spell_tormenting_thirst", //[cite: 28]
			4: "spell_dispel_water",      //[cite: 28]
			5: "spell_desiccate_mass",    //[cite: 28]
			6: "spell_symbol_of_thirst",  //[cite: 28]
			7: "spell_mephit_mob",        //[cite: 28]
			8: "spell_horrid_wilting",    //[cite: 28]
			9: "spell_energy_drain",      //[cite: 28]
		},
	}

	DomainTyranny = DomainBlueprint{
		ID:           "dom_tyranny",
		Name:         "Tyranny Domain",
		GrantedPower: "feat_domain_power_tyranny", //[cite: 28]
		Spells: map[int]string{
			1: "spell_command",              //[cite: 28]
			2: "spell_enthrall",             //[cite: 28]
			3: "spell_discern_lies",         //[cite: 28]
			4: "spell_fear",                 //[cite: 28]
			5: "spell_command_greater",      //[cite: 28]
			6: "spell_geas_quest",           //[cite: 28]
			7: "spell_bigbys_grasping_hand", //[cite: 28]
			8: "spell_charm_monster_mass",   //[cite: 28]
			9: "spell_dominate_monster",     //[cite: 28]
		},
	}

	DomainUndeath = DomainBlueprint{
		ID:           "dom_undeath",
		Name:         "Undeath Domain",
		GrantedPower: "feat_domain_power_undeath", //[cite: 28]
		Spells: map[int]string{
			1: "spell_detect_undead",         //[cite: 28]
			2: "spell_desecrate",             //[cite: 28]
			3: "spell_animate_dead",          //[cite: 28]
			4: "spell_death_ward",            //[cite: 28]
			5: "spell_circle_of_death",       //[cite: 28]
			6: "spell_create_undead",         //[cite: 28]
			7: "spell_control_undead",        //[cite: 28]
			8: "spell_create_greater_undead", //[cite: 28]
			9: "spell_energy_drain",          //[cite: 28]
		},
	}

	DomainVileDarkness = DomainBlueprint{
		ID:           "dom_vile_darkness",
		Name:         "Vile Darkness Domain",
		GrantedPower: "feat_domain_power_vile_darkness", //[cite: 28]
		Spells: map[int]string{
			1: "spell_darkvision",             //[cite: 28]
			2: "spell_darkbolt",               //[cite: 28]
			3: "spell_deeper_darkness",        //[cite: 28]
			4: "spell_damning_darkness",       //[cite: 28]
			5: "spell_evards_black_tentacles", //[cite: 28]
			6: "spell_wall_of_force",          //[cite: 28]
			7: "spell_shadow_walk",            //[cite: 28]
			8: "spell_utterdark",              //[cite: 28]
			9: "spell_screen",                 //[cite: 28]
		},
	}

	DomainWealth = DomainBlueprint{
		ID:           "dom_wealth",
		Name:         "Wealth Domain",
		GrantedPower: "feat_domain_power_wealth", //[cite: 28]
		Spells: map[int]string{
			1: "spell_alarm",                 //[cite: 28]
			2: "spell_obscure_object",        //[cite: 28]
			3: "spell_glyph_of_warding",      //[cite: 28]
			4: "spell_detect_scrying",        //[cite: 28]
			5: "spell_leomunds_secret_chest", //[cite: 28]
			6: "spell_forbiddance",           //[cite: 28]
			7: "spell_sequester",             //[cite: 28]
			8: "spell_discern_location",      //[cite: 28]
			9: "spell_antipathy",             //[cite: 28]
		},
	}

	DomainWeather = DomainBlueprint{
		ID:           "dom_weather",
		Name:         "Weather Domain",
		GrantedPower: "feat_domain_power_weather", //[cite: 28]
		Spells: map[int]string{
			1: "spell_obscuring_mist",       //[cite: 28]
			2: "spell_fog_cloud",            //[cite: 28]
			3: "spell_call_lightning",       //[cite: 28]
			4: "spell_sleet_storm",          //[cite: 28]
			5: "spell_call_lightning_storm", //[cite: 28]
			6: "spell_control_winds",        //[cite: 28]
			7: "spell_control_weather",      //[cite: 28]
			8: "spell_whirlwind",            //[cite: 28]
			9: "spell_storm_of_vengeance",   //[cite: 28]
		},
	}

	DomainWinter = DomainBlueprint{
		ID:           "dom_winter",
		Name:         "Winter Domain",
		GrantedPower: "feat_domain_power_winter", //[cite: 28]
		Spells: map[int]string{
			1: "spell_snowsight",       //[cite: 28]
			2: "spell_snow_walk",       //[cite: 28]
			3: "spell_winters_embrace", //[cite: 28]
			4: "spell_ice_storm",       //[cite: 28]
			5: "spell_blizzard",        //[cite: 28]
			6: "spell_death_hail",      //[cite: 28]
			7: "spell_control_weather", //[cite: 28]
			8: "spell_summon_giants",   //[cite: 28]
			9: "spell_fimbulwinter",    //[cite: 28]
		},
	}

	DomainWrath = DomainBlueprint{
		ID:           "dom_wrath",
		Name:         "Wrath Domain",
		GrantedPower: "feat_domain_power_wrath", //[cite: 28]
		Spells: map[int]string{
			1: "spell_doom",               //[cite: 28]
			2: "spell_energize_potion",    //[cite: 28]
			3: "spell_affliction",         //[cite: 28]
			4: "spell_radiant_shield",     //[cite: 28]
			5: "spell_righteous_might",    //[cite: 28]
			6: "spell_vengeance_halo",     //[cite: 28]
			7: "spell_righteous_smite",    //[cite: 28]
			8: "spell_last_judgment",      //[cite: 28]
			9: "spell_storm_of_vengeance", //[cite: 28]
		},
	}

	// --- PLANAR DOMAINS ---

	DomainTheAbyss = DomainBlueprint{
		ID:           "dom_the_abyss",
		Name:         "The Abyss Domain",
		GrantedPower: "feat_domain_power_the_abyss", //[cite: 28]
		Spells: map[int]string{
			1: "spell_align_weapon",        //[cite: 28]
			2: "spell_bulls_strength",      //[cite: 28]
			3: "spell_babau_slime",         //[cite: 28]
			4: "spell_balor_nimbus",        //[cite: 28]
			5: "spell_slay_living",         //[cite: 28]
			6: "spell_bulls_strength_mass", //[cite: 28]
			7: "spell_destruction",         //[cite: 28]
			8: "spell_finger_of_death",     //[cite: 28]
			9: "spell_implosion",           //[cite: 28]
		},
	}

	DomainArborea = DomainBlueprint{
		ID:           "dom_arborea",
		Name:         "Arborea Domain",
		GrantedPower: "feat_domain_power_arborea", //[cite: 28]
		Spells: map[int]string{
			1: "spell_endure_elements",   //[cite: 28]
			2: "spell_aid",               //[cite: 28]
			3: "spell_heroism",           //[cite: 28]
			4: "spell_neutralize_poison", //[cite: 28]
			5: "spell_break_enchantment", //[cite: 28]
			6: "spell_heroes_feast",      //[cite: 28]
			7: "spell_spell_turning",     //[cite: 28]
			8: "spell_heroism_greater",   //[cite: 28]
			9: "spell_freedom",           //[cite: 28]
		},
	}

	DomainBaator = DomainBlueprint{
		ID:           "dom_baator",
		Name:         "Baator Domain",
		GrantedPower: "feat_domain_power_baator", //[cite: 28]
		Spells: map[int]string{
			1: "spell_bane",             //[cite: 28]
			2: "spell_darkness",         //[cite: 28]
			3: "spell_detect_thoughts",  //[cite: 28]
			4: "spell_deeper_darkness",  //[cite: 28]
			5: "spell_spell_resistance", //[cite: 28]
			6: "spell_dominate_person",  //[cite: 28]
			7: "spell_repulsion",        //[cite: 28]
			8: "spell_demand",           //[cite: 28]
			9: "spell_imprisonment",     //[cite: 28]
		},
	}

	DomainCelestia = DomainBlueprint{
		ID:           "dom_celestia",
		Name:         "Celestia Domain",
		GrantedPower: "feat_domain_power_celestia", //[cite: 28]
		Spells: map[int]string{
			1: "spell_light_of_lunia",  //[cite: 28]
			2: "spell_bears_endurance", //[cite: 28]
			3: "spell_magic_vestment",  //[cite: 28]
			4: "spell_divine_power",    //[cite: 28]
			5: "spell_righteous_might", //[cite: 28]
			6: "spell_blade_barrier",   //[cite: 28]
			7: "spell_regenerate",      //[cite: 28]
			8: "spell_power_word_stun", //[cite: 28]
			9: "spell_foresight",       //[cite: 28]
		},
	}

	DomainElysium = DomainBlueprint{
		ID:           "dom_elysium",
		Name:         "Elysium Domain",
		GrantedPower: "feat_domain_power_elysium", //[cite: 28]
		Spells: map[int]string{
			1: "spell_charm_person",              //[cite: 28]
			2: "spell_enthrall",                  //[cite: 28]
			3: "spell_magic_circle_against_evil", //[cite: 28]
			4: "spell_charm_monster",             //[cite: 28]
			5: "spell_dispel_evil",               //[cite: 28]
			6: "spell_find_the_path",             //[cite: 28]
			7: "spell_control_weather",           //[cite: 28]
			8: "spell_holy_aura",                 //[cite: 28]
			9: "spell_heal_mass",                 //[cite: 28]
		},
	}

	DomainHades = DomainBlueprint{
		ID:           "dom_hades",
		Name:         "Hades Domain",
		GrantedPower: "feat_domain_power_hades", //[cite: 28]
		Spells: map[int]string{
			1: "spell_doom",                      //[cite: 28]
			2: "spell_resist_planar_alignment",   //[cite: 28]
			3: "spell_magic_circle_against_good", //[cite: 28]
			4: "spell_contagion",                 //[cite: 28]
			5: "spell_crushing_despair",          //[cite: 28]
			6: "spell_mind_fog",                  //[cite: 28]
			7: "spell_blasphemy",                 //[cite: 28]
			8: "spell_unholy_aura",               //[cite: 28]
			9: "spell_energy_drain",              //[cite: 28]
		},
	}

	DomainLimbo = DomainBlueprint{
		ID:           "dom_limbo",
		Name:         "Limbo Domain",
		GrantedPower: "feat_domain_power_limbo", //[cite: 28]
		Spells: map[int]string{
			1: "spell_lesser_confusion",         //[cite: 28]
			2: "spell_entropic_shield",          //[cite: 28]
			3: "spell_magic_circle_against_law", //[cite: 28]
			4: "spell_chaos_hammer",             //[cite: 28]
			5: "spell_baleful_polymorph",        //[cite: 28]
			6: "spell_animate_objects",          //[cite: 28]
			7: "spell_song_of_discord",          //[cite: 28]
			8: "spell_cloak_of_chaos",           //[cite: 28]
			9: "spell_perinarch_planar",         //[cite: 28]
		},
	}

	DomainMechanus = DomainBlueprint{
		ID:           "dom_mechanus",
		Name:         "Mechanus Domain",
		GrantedPower: "feat_domain_power_mechanus", //[cite: 28]
		Spells: map[int]string{
			1: "spell_command",                    //[cite: 28]
			2: "spell_calm_emotions",              //[cite: 28]
			3: "spell_magic_circle_against_chaos", //[cite: 28]
			4: "spell_discern_lies",               //[cite: 28]
			5: "spell_dispel_chaos",               //[cite: 28]
			6: "spell_hold_monster",               //[cite: 28]
			7: "spell_dictum",                     //[cite: 28]
			8: "spell_iron_body",                  //[cite: 28]
			9: "spell_call_marut",                 //[cite: 28]
		},
	}
)

// Global variables for Ranger Combat Styles
var (
	// --- STANDARD RANGER STYLES ---
	StyleArchery = CombatStyleBlueprint{
		ID:   "style_archery",
		Name: "Archery",
		Feats: map[int]string{
			2:  "feat_rapid_shot",
			6:  "feat_manyshot",
			11: "feat_improved_precise_shot",
		},
	}

	StyleTwoWeapon = CombatStyleBlueprint{
		ID:   "style_twoweapon",
		Name: "Two-Weapon Combat",
		Feats: map[int]string{
			2:  "feat_two_weapon_fighting",
			6:  "feat_improved_two_weapon_fighting",
			11: "feat_greater_two_weapon_fighting",
		},
	}

	// --- ALTERNATIVE RANGER STYLES ---
	StyleBeastWrestling = CombatStyleBlueprint{
		ID:   "style_beast_wrestling",
		Name: "Beast-Wrestling Style",
		Feats: map[int]string{
			2:  "feat_improved_unarmed_strike",
			6:  "feat_improved_grapple",
			11: "feat_stunning_fist",
		},
	}

	StyleMountedCombat = CombatStyleBlueprint{
		ID:   "style_mounted_combat",
		Name: "Mounted-Combat Style",
		Feats: map[int]string{
			2:  "feat_ride_by_attack",
			6:  "feat_spirited_charge",
			11: "feat_trample",
		},
	}

	StylePiscator = CombatStyleBlueprint{
		ID:   "style_piscator",
		Name: "Piscator Style",
		Feats: map[int]string{
			2:  "feat_exotic_weapon_proficiency_net",
			6:  "feat_improved_trip",
			11: "feat_improved_critical",
		},
	}

	StyleStrongArm = CombatStyleBlueprint{
		ID:   "style_strong_arm",
		Name: "Strong-Arm Style",
		Feats: map[int]string{
			2:  "feat_power_attack",
			6:  "feat_improved_sunder",
			11: "feat_great_cleave",
		},
	}

	StyleThrowing = CombatStyleBlueprint{
		ID:   "style_throwing",
		Name: "Throwing Style",
		Feats: map[int]string{
			2:  "feat_quick_draw",
			6:  "feat_point_blank_shot",
			11: "feat_far_shot",
		},
	}

	// --- ROGUE SPECIAL ABILITIES ---
	RogueAbilityCripplingStrike = RogueAbilityBlueprint{
		ID:          "rogue_ab_crippling_strike",
		Name:        "Crippling Strike",
		Description: "Sneak attacks deal 2 points of temporary Strength damage.", //
	}

	RogueAbilityDefensiveRoll = RogueAbilityBlueprint{
		ID:          "rogue_ab_defensive_roll",
		Name:        "Defensive Roll",
		Description: "Once per day, roll with a lethal blow to take half damage.", //
	}

	RogueAbilityImprovedEvasion = RogueAbilityBlueprint{
		ID:          "rogue_ab_improved_evasion",
		Name:        "Improved Evasion",
		Description: "Take no damage on successful Reflex save, half damage on failure.", //
	}

	RogueAbilityOpportunist = RogueAbilityBlueprint{
		ID:          "rogue_ab_opportunist",
		Name:        "Opportunist",
		Description: "Make one extra attack of opportunity per round against an ally's target.", //
	}

	RogueAbilitySkillMastery = RogueAbilityBlueprint{
		ID:          "rogue_ab_skill_mastery",
		Name:        "Skill Mastery",
		Description: "Take 10 on selected skills even under stress or distraction.", //
	}

	RogueAbilitySlipperyMind = RogueAbilityBlueprint{
		ID:          "rogue_ab_slippery_mind",
		Name:        "Slippery Mind",
		Description: "Attempt a second saving throw 1 round later against enchantment effects.", //
	}
	// --------------------------
)
