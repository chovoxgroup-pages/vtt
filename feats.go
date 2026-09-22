package main

type FeatType string

const (
	FeatGeneral      FeatType = "General"       //[cite: 17]
	FeatItemCreation FeatType = "Item creation" //[cite: 17]
	FeatMetamagic    FeatType = "Metamagic"     //[cite: 17]
	FeatSpecial      FeatType = "Special"       //[cite: 17]
	FeatPsionic      FeatType = "Psionic"       //[cite: 17]
	FeatMetapsionic  FeatType = "Metapsionic"   //[cite: 17]
)

type FeatBlueprint struct {
	ID            string
	Name          string
	Type          FeatType
	Prerequisites string
	Source        string
}

// GlobalFeats acts as the registry for all available 3.5e feats[cite: 17]
var GlobalFeats = map[string]FeatBlueprint{
	"feat_acrobatic":           {ID: "feat_acrobatic", Name: "Acrobatic", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                            //[cite: 17]
	"feat_agile":               {ID: "feat_agile", Name: "Agile", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                                    //[cite: 17]
	"feat_alertness":           {ID: "feat_alertness", Name: "Alertness", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                            //[cite: 17]
	"feat_animal_affinity":     {ID: "feat_animal_affinity", Name: "Animal Affinity", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                //[cite: 17]
	"feat_armor_prof_light":    {ID: "feat_armor_prof_light", Name: "Armor Proficiency (Light)", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                     //[cite: 17]
	"feat_armor_prof_medium":   {ID: "feat_armor_prof_medium", Name: "Armor Proficiency (Medium)", Type: FeatGeneral, Prerequisites: "Armor Proficiency (light)", Source: "Core"},                           //[cite: 17]
	"feat_armor_prof_heavy":    {ID: "feat_armor_prof_heavy", Name: "Armor Proficiency (Heavy)", Type: FeatGeneral, Prerequisites: "Armor Proficiency (light), Armor Proficiency (medium)", Source: "Core"}, //[cite: 17]
	"feat_athletic":            {ID: "feat_athletic", Name: "Athletic", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                              //[cite: 17]
	"feat_augment_summoning":   {ID: "feat_augment_summoning", Name: "Augment Summoning", Type: FeatGeneral, Prerequisites: "Spell Focus (conjuration)", Source: "Core"},                                    //[cite: 17]
	"feat_blind_fight":         {ID: "feat_blind_fight", Name: "Blind-Fight", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                        //[cite: 17]
	"feat_brew_potion":         {ID: "feat_brew_potion", Name: "Brew Potion", Type: FeatItemCreation, Prerequisites: "Caster level 3rd", Source: "Core"},                                                    //[cite: 17]
	"feat_cleave":              {ID: "feat_cleave", Name: "Cleave", Type: FeatGeneral, Prerequisites: "Str 13, Power Attack", Source: "Core"},                                                               //[cite: 17]
	"feat_cleave_great":        {ID: "feat_cleave_great", Name: "Cleave, Great", Type: FeatGeneral, Prerequisites: "Str 13, Cleave, Power Attack, base attack bonus +4", Source: "Core"},                    //[cite: 17]
	"feat_combat_casting":      {ID: "feat_combat_casting", Name: "Combat Casting", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                  //[cite: 17]
	"feat_combat_expertise":    {ID: "feat_combat_expertise", Name: "Combat Expertise", Type: FeatGeneral, Prerequisites: "Int 13", Source: "Core"},                                                         //[cite: 17]
	"feat_combat_reflexes":     {ID: "feat_combat_reflexes", Name: "Combat Reflexes", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                //[cite: 17]
	"feat_craft_magic_arms":    {ID: "feat_craft_magic_arms", Name: "Craft Magic Arms and Armor", Type: FeatItemCreation, Prerequisites: "Caster level 5th", Source: "Core"},                                //[cite: 17]
	"feat_deceitful":           {ID: "feat_deceitful", Name: "Deceitful", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                                            //[cite: 17]
	"feat_deflect_arrows":      {ID: "feat_deflect_arrows", Name: "Deflect Arrows", Type: FeatGeneral, Prerequisites: "Dex 13, Improved Unarmed Strike", Source: "Core"},                                    //[cite: 17]
	"feat_dodge":               {ID: "feat_dodge", Name: "Dodge", Type: FeatGeneral, Prerequisites: "Dex 13", Source: "Core"},                                                                               //[cite: 17]
	"feat_empower_spell":       {ID: "feat_empower_spell", Name: "Empower Spell", Type: FeatMetamagic, Prerequisites: "-", Source: "Core"},                                                                  //[cite: 17]
	"feat_improved_initiative": {ID: "feat_improved_initiative", Name: "Improved Initiative", Type: FeatGeneral, Prerequisites: "-", Source: "Core"},                                                        //[cite: 17]
	"feat_power_attack":        {ID: "feat_power_attack", Name: "Power Attack", Type: FeatGeneral, Prerequisites: "Str 13", Source: "Core"},                                                                 //[cite: 17]
	"feat_weapon_finesse":      {ID: "feat_weapon_finesse", Name: "Weapon Finesse", Type: FeatGeneral, Prerequisites: "Base attack bonus +1", Source: "Core"},                                               //[cite: 17]
	"feat_weapon_focus":        {ID: "feat_weapon_focus", Name: "Weapon Focus", Type: FeatGeneral, Prerequisites: "Proficiency with selected weapon, base attack bonus +1", Source: "Core"},                 //[cite: 17]
}
