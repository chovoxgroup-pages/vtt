package main

// SkillBlueprint defines the core rules for a specific skill[cite: 18]
type SkillBlueprint struct {
	ID                string
	Name              string
	KeyAbility        string // "STR", "DEX", "CON", "INT", "WIS", "CHA", or "NONE"[cite: 18]
	TrainedOnly       bool
	ArmorCheckPenalty bool
}

// GlobalSkills acts as the registry for all available 3.5e skills[cite: 18]
var GlobalSkills = map[string]SkillBlueprint{
	"skill_appraise":           {ID: "skill_appraise", Name: "Appraise", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_autohypnosis":       {ID: "skill_autohypnosis", Name: "Autohypnosis", KeyAbility: "WIS", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_balance":            {ID: "skill_balance", Name: "Balance", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_bluff":              {ID: "skill_bluff", Name: "Bluff", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_climb":              {ID: "skill_climb", Name: "Climb", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_concentration":      {ID: "skill_concentration", Name: "Concentration", KeyAbility: "CON", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_control_shape":      {ID: "skill_control_shape", Name: "Control Shape", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_craft":              {ID: "skill_craft", Name: "Craft", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},                 //[cite: 23]
	"skill_decipher_script":    {ID: "skill_decipher_script", Name: "Decipher Script", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_diplomacy":          {ID: "skill_diplomacy", Name: "Diplomacy", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_disable_device":     {ID: "skill_disable_device", Name: "Disable Device", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_disguise":           {ID: "skill_disguise", Name: "Disguise", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_escape_artist":      {ID: "skill_escape_artist", Name: "Escape Artist", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_forgery":            {ID: "skill_forgery", Name: "Forgery", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_gather_info":        {ID: "skill_gather_info", Name: "Gather Information", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_handle_animal":      {ID: "skill_handle_animal", Name: "Handle Animal", KeyAbility: "CHA", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_heal":               {ID: "skill_heal", Name: "Heal", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_hide":               {ID: "skill_hide", Name: "Hide", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_intimidate":         {ID: "skill_intimidate", Name: "Intimidate", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_jump":               {ID: "skill_jump", Name: "Jump", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_knowledge":          {ID: "skill_knowledge", Name: "Knowledge", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_listen":             {ID: "skill_listen", Name: "Listen", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_martial_lore":       {ID: "skill_martial_lore", Name: "Martial Lore", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_move_silently":      {ID: "skill_move_silently", Name: "Move Silently", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_open_lock":          {ID: "skill_open_lock", Name: "Open Lock", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_perform":            {ID: "skill_perform", Name: "Perform", KeyAbility: "CHA", TrainedOnly: false, ArmorCheckPenalty: false},      //[cite: 23]
	"skill_profession":         {ID: "skill_profession", Name: "Profession", KeyAbility: "WIS", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_psicraft":           {ID: "skill_psicraft", Name: "Psicraft", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},     //[cite: 23]
	"skill_ride":               {ID: "skill_ride", Name: "Ride", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_search":             {ID: "skill_search", Name: "Search", KeyAbility: "INT", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_sense_motive":       {ID: "skill_sense_motive", Name: "Sense Motive", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_sleight_of_hand":    {ID: "skill_sleight_of_hand", Name: "Sleight of Hand", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: true},
	"skill_speak_language":     {ID: "skill_speak_language", Name: "Speak Language", KeyAbility: "NONE", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_spellcraft":         {ID: "skill_spellcraft", Name: "Spellcraft", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_spot":               {ID: "skill_spot", Name: "Spot", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_survival":           {ID: "skill_survival", Name: "Survival", KeyAbility: "WIS", TrainedOnly: false, ArmorCheckPenalty: false},
	"skill_swim":               {ID: "skill_swim", Name: "Swim", KeyAbility: "STR", TrainedOnly: false, ArmorCheckPenalty: true},
	"skill_truespeak":          {ID: "skill_truespeak", Name: "Truespeak", KeyAbility: "INT", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_tumble":             {ID: "skill_tumble", Name: "Tumble", KeyAbility: "DEX", TrainedOnly: true, ArmorCheckPenalty: true},
	"skill_use_magic_device":   {ID: "skill_use_magic_device", Name: "Use Magic Device", KeyAbility: "CHA", TrainedOnly: true, ArmorCheckPenalty: false},
	"skill_use_psionic_device": {ID: "skill_use_psionic_device", Name: "Use Psionic Device", KeyAbility: "CHA", TrainedOnly: true, ArmorCheckPenalty: false}, //[cite: 23]
	"skill_use_rope":           {ID: "skill_use_rope", Name: "Use Rope", KeyAbility: "DEX", TrainedOnly: false, ArmorCheckPenalty: false},
}
