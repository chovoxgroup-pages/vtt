package main

type ArmorClass string

const (
	ArmorLight  ArmorClass = "Light"
	ArmorMedium ArmorClass = "Medium"
	ArmorHeavy  ArmorClass = "Heavy"
	ArmorShield ArmorClass = "Shield"
)

type ArmorBlueprint struct {
	ID                string
	Name              string
	Class             ArmorClass
	ACBonus           int
	MaxDex            int
	ArmorCheckPenalty int
	ArcaneSpellFail   int // Percentage (e.g., 20 for 20%)
	Weight            int // In pounds (Medium size)
}

// GlobalArmors acts as the registry for all available 3.5e armor and shields
var GlobalArmors = map[string]ArmorBlueprint{
	// --- LIGHT ARMOR ---
	"arm_padded":          {ID: "arm_padded", Name: "Padded", Class: ArmorLight, ACBonus: 1, MaxDex: 8, ArmorCheckPenalty: 0, ArcaneSpellFail: 5, Weight: 10},
	"arm_leather":         {ID: "arm_leather", Name: "Leather", Class: ArmorLight, ACBonus: 2, MaxDex: 6, ArmorCheckPenalty: 0, ArcaneSpellFail: 10, Weight: 15},
	"arm_studded_leather": {ID: "arm_studded_leather", Name: "Studded leather", Class: ArmorLight, ACBonus: 3, MaxDex: 5, ArmorCheckPenalty: -1, ArcaneSpellFail: 15, Weight: 20},
	"arm_chain_shirt":     {ID: "arm_chain_shirt", Name: "Chain shirt", Class: ArmorLight, ACBonus: 4, MaxDex: 4, ArmorCheckPenalty: -2, ArcaneSpellFail: 20, Weight: 25},

	// --- MEDIUM ARMOR ---
	"arm_hide":        {ID: "arm_hide", Name: "Hide", Class: ArmorMedium, ACBonus: 3, MaxDex: 4, ArmorCheckPenalty: -3, ArcaneSpellFail: 20, Weight: 25},
	"arm_scale_mail":  {ID: "arm_scale_mail", Name: "Scale mail", Class: ArmorMedium, ACBonus: 4, MaxDex: 3, ArmorCheckPenalty: -4, ArcaneSpellFail: 25, Weight: 30},
	"arm_chainmail":   {ID: "arm_chainmail", Name: "Chainmail", Class: ArmorMedium, ACBonus: 5, MaxDex: 2, ArmorCheckPenalty: -5, ArcaneSpellFail: 30, Weight: 40},
	"arm_breastplate": {ID: "arm_breastplate", Name: "Breastplate", Class: ArmorMedium, ACBonus: 5, MaxDex: 3, ArmorCheckPenalty: -4, ArcaneSpellFail: 25, Weight: 30},

	// --- HEAVY ARMOR ---
	"arm_splint_mail": {ID: "arm_splint_mail", Name: "Splint mail", Class: ArmorHeavy, ACBonus: 6, MaxDex: 0, ArmorCheckPenalty: -7, ArcaneSpellFail: 40, Weight: 45},
	"arm_banded_mail": {ID: "arm_banded_mail", Name: "Banded mail", Class: ArmorHeavy, ACBonus: 6, MaxDex: 1, ArmorCheckPenalty: -6, ArcaneSpellFail: 35, Weight: 35},
	"arm_half_plate":  {ID: "arm_half_plate", Name: "Half-plate", Class: ArmorHeavy, ACBonus: 7, MaxDex: 0, ArmorCheckPenalty: -7, ArcaneSpellFail: 40, Weight: 50},
	"arm_full_plate":  {ID: "arm_full_plate", Name: "Full plate", Class: ArmorHeavy, ACBonus: 8, MaxDex: 1, ArmorCheckPenalty: -6, ArcaneSpellFail: 35, Weight: 50},

	// --- SHIELDS ---
	"shd_buckler":     {ID: "shd_buckler", Name: "Buckler", Class: ArmorShield, ACBonus: 1, MaxDex: 99, ArmorCheckPenalty: -1, ArcaneSpellFail: 5, Weight: 5},
	"shd_light_wood":  {ID: "shd_light_wood", Name: "Shield, light wooden", Class: ArmorShield, ACBonus: 1, MaxDex: 99, ArmorCheckPenalty: -1, ArcaneSpellFail: 5, Weight: 5},
	"shd_light_steel": {ID: "shd_light_steel", Name: "Shield, light steel", Class: ArmorShield, ACBonus: 1, MaxDex: 99, ArmorCheckPenalty: -1, ArcaneSpellFail: 5, Weight: 6},
	"shd_heavy_wood":  {ID: "shd_heavy_wood", Name: "Shield, heavy wooden", Class: ArmorShield, ACBonus: 2, MaxDex: 99, ArmorCheckPenalty: -2, ArcaneSpellFail: 15, Weight: 10},
	"shd_heavy_steel": {ID: "shd_heavy_steel", Name: "Shield, heavy steel", Class: ArmorShield, ACBonus: 2, MaxDex: 99, ArmorCheckPenalty: -2, ArcaneSpellFail: 15, Weight: 15},
	"shd_tower":       {ID: "shd_tower", Name: "Shield, tower", Class: ArmorShield, ACBonus: 4, MaxDex: 2, ArmorCheckPenalty: -10, ArcaneSpellFail: 50, Weight: 45},
}
