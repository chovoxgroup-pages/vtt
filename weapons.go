package main

type ProficiencyClass string

const (
	ProfSimple  ProficiencyClass = "Simple"
	ProfMartial ProficiencyClass = "Martial"
	ProfExotic  ProficiencyClass = "Exotic"
)

type WeaponCategory string

const (
	CatLight     WeaponCategory = "Light Melee"
	CatOneHanded WeaponCategory = "One-Handed Melee"
	CatTwoHanded WeaponCategory = "Two-Handed Melee"
	CatRanged    WeaponCategory = "Ranged"
)

type WeaponBlueprint struct {
	ID         string
	Name       string
	ProfClass  ProficiencyClass
	Category   WeaponCategory
	Damage     string // Standard Medium Damage (e.g., "1d8")
	Critical   string // e.g., "19-20/x2"
	Range      int    // In feet (0 if melee)
	Weight     int    // In pounds (Medium size)
	DamageType string // "Bludgeoning", "Piercing", "Slashing", or combination
}

// GlobalWeapons acts as the registry for all available 3.5e weapons
var GlobalWeapons = map[string]WeaponBlueprint{
	// --- SIMPLE LIGHT MELEE WEAPONS ---
	"wpn_unarmed":        {ID: "wpn_unarmed", Name: "Unarmed Strike", ProfClass: ProfSimple, Category: CatLight, Damage: "1d3", Critical: "x2", Range: 0, Weight: 0, DamageType: "Bludgeoning"},
	"wpn_gauntlet":       {ID: "wpn_gauntlet", Name: "Gauntlet", ProfClass: ProfSimple, Category: CatLight, Damage: "1d3", Critical: "x2", Range: 0, Weight: 1, DamageType: "Bludgeoning"},
	"wpn_dagger":         {ID: "wpn_dagger", Name: "Dagger", ProfClass: ProfSimple, Category: CatLight, Damage: "1d4", Critical: "19-20/x2", Range: 10, Weight: 1, DamageType: "Piercing or Slashing"},
	"wpn_dagger_punch":   {ID: "wpn_dagger_punch", Name: "Dagger, Punching", ProfClass: ProfSimple, Category: CatLight, Damage: "1d4", Critical: "x3", Range: 0, Weight: 1, DamageType: "Piercing"},
	"wpn_gauntlet_spike": {ID: "wpn_gauntlet_spike", Name: "Gauntlet, Spiked", ProfClass: ProfSimple, Category: CatLight, Damage: "1d4", Critical: "x2", Range: 0, Weight: 1, DamageType: "Piercing"},
	"wpn_mace_light":     {ID: "wpn_mace_light", Name: "Mace, Light", ProfClass: ProfSimple, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 0, Weight: 4, DamageType: "Bludgeoning"},
	"wpn_sickle":         {ID: "wpn_sickle", Name: "Sickle", ProfClass: ProfSimple, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 0, Weight: 2, DamageType: "Slashing"},

	// --- SIMPLE ONE-HANDED MELEE WEAPONS ---
	"wpn_club":        {ID: "wpn_club", Name: "Club", ProfClass: ProfSimple, Category: CatOneHanded, Damage: "1d6", Critical: "x2", Range: 10, Weight: 3, DamageType: "Bludgeoning"},
	"wpn_mace_heavy":  {ID: "wpn_mace_heavy", Name: "Mace, Heavy", ProfClass: ProfSimple, Category: CatOneHanded, Damage: "1d8", Critical: "x2", Range: 0, Weight: 8, DamageType: "Bludgeoning"},
	"wpn_morningstar": {ID: "wpn_morningstar", Name: "Morningstar", ProfClass: ProfSimple, Category: CatOneHanded, Damage: "1d8", Critical: "x2", Range: 0, Weight: 6, DamageType: "Bludgeoning and Piercing"},
	"wpn_shortspear":  {ID: "wpn_shortspear", Name: "Shortspear", ProfClass: ProfSimple, Category: CatOneHanded, Damage: "1d6", Critical: "x2", Range: 20, Weight: 3, DamageType: "Piercing"},

	// --- SIMPLE TWO-HANDED MELEE WEAPONS ---
	"wpn_longspear":    {ID: "wpn_longspear", Name: "Longspear", ProfClass: ProfSimple, Category: CatTwoHanded, Damage: "1d8", Critical: "x3", Range: 0, Weight: 9, DamageType: "Piercing"},
	"wpn_quarterstaff": {ID: "wpn_quarterstaff", Name: "Quarterstaff", ProfClass: ProfSimple, Category: CatTwoHanded, Damage: "1d6/1d6", Critical: "x2", Range: 0, Weight: 4, DamageType: "Bludgeoning"},
	"wpn_spear":        {ID: "wpn_spear", Name: "Spear", ProfClass: ProfSimple, Category: CatTwoHanded, Damage: "1d8", Critical: "x3", Range: 20, Weight: 6, DamageType: "Piercing"},

	// --- SIMPLE RANGED WEAPONS ---
	"wpn_crossbow_heavy": {ID: "wpn_crossbow_heavy", Name: "Crossbow, Heavy", ProfClass: ProfSimple, Category: CatRanged, Damage: "1d10", Critical: "19-20/x2", Range: 120, Weight: 8, DamageType: "Piercing"},
	"wpn_crossbow_light": {ID: "wpn_crossbow_light", Name: "Crossbow, Light", ProfClass: ProfSimple, Category: CatRanged, Damage: "1d8", Critical: "19-20/x2", Range: 80, Weight: 4, DamageType: "Piercing"},
	"wpn_dart":           {ID: "wpn_dart", Name: "Dart", ProfClass: ProfSimple, Category: CatRanged, Damage: "1d4", Critical: "x2", Range: 20, Weight: 0, DamageType: "Piercing"},
	"wpn_javelin":        {ID: "wpn_javelin", Name: "Javelin", ProfClass: ProfSimple, Category: CatRanged, Damage: "1d6", Critical: "x2", Range: 30, Weight: 2, DamageType: "Piercing"},
	"wpn_sling":          {ID: "wpn_sling", Name: "Sling", ProfClass: ProfSimple, Category: CatRanged, Damage: "1d4", Critical: "x2", Range: 50, Weight: 0, DamageType: "Bludgeoning"},

	// --- MARTIAL LIGHT MELEE WEAPONS ---
	"wpn_axe_throw":    {ID: "wpn_axe_throw", Name: "Axe, Throwing", ProfClass: ProfMartial, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 10, Weight: 2, DamageType: "Slashing"},
	"wpn_hammer_light": {ID: "wpn_hammer_light", Name: "Hammer, Light", ProfClass: ProfMartial, Category: CatLight, Damage: "1d4", Critical: "x2", Range: 20, Weight: 2, DamageType: "Bludgeoning"},
	"wpn_handaxe":      {ID: "wpn_handaxe", Name: "Handaxe", ProfClass: ProfMartial, Category: CatLight, Damage: "1d6", Critical: "x3", Range: 0, Weight: 3, DamageType: "Slashing"},
	"wpn_kukri":        {ID: "wpn_kukri", Name: "Kukri", ProfClass: ProfMartial, Category: CatLight, Damage: "1d4", Critical: "18-20/x2", Range: 0, Weight: 2, DamageType: "Slashing"},
	"wpn_pick_light":   {ID: "wpn_pick_light", Name: "Pick, Light", ProfClass: ProfMartial, Category: CatLight, Damage: "1d4", Critical: "x4", Range: 0, Weight: 3, DamageType: "Piercing"},
	"wpn_sap":          {ID: "wpn_sap", Name: "Sap", ProfClass: ProfMartial, Category: CatLight, Damage: "1d6 (Nonlethal)", Critical: "x2", Range: 0, Weight: 2, DamageType: "Bludgeoning"},
	"wpn_sword_short":  {ID: "wpn_sword_short", Name: "Sword, Short", ProfClass: ProfMartial, Category: CatLight, Damage: "1d6", Critical: "19-20/x2", Range: 0, Weight: 2, DamageType: "Piercing"},

	// --- MARTIAL ONE-HANDED MELEE WEAPONS ---
	"wpn_battleaxe":  {ID: "wpn_battleaxe", Name: "Battleaxe", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d8", Critical: "x3", Range: 0, Weight: 6, DamageType: "Slashing"},
	"wpn_flail":      {ID: "wpn_flail", Name: "Flail", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d8", Critical: "x2", Range: 0, Weight: 5, DamageType: "Bludgeoning"},
	"wpn_longsword":  {ID: "wpn_longsword", Name: "Longsword", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d8", Critical: "19-20/x2", Range: 0, Weight: 4, DamageType: "Slashing"},
	"wpn_pick_heavy": {ID: "wpn_pick_heavy", Name: "Pick, Heavy", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d6", Critical: "x4", Range: 0, Weight: 6, DamageType: "Piercing"},
	"wpn_rapier":     {ID: "wpn_rapier", Name: "Rapier", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d6", Critical: "18-20/x2", Range: 0, Weight: 2, DamageType: "Piercing"},
	"wpn_scimitar":   {ID: "wpn_scimitar", Name: "Scimitar", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d6", Critical: "18-20/x2", Range: 0, Weight: 4, DamageType: "Slashing"},
	"wpn_trident":    {ID: "wpn_trident", Name: "Trident", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d8", Critical: "x2", Range: 10, Weight: 4, DamageType: "Piercing"},
	"wpn_warhammer":  {ID: "wpn_warhammer", Name: "Warhammer", ProfClass: ProfMartial, Category: CatOneHanded, Damage: "1d8", Critical: "x3", Range: 0, Weight: 5, DamageType: "Bludgeoning"},

	// --- MARTIAL TWO-HANDED MELEE WEAPONS ---
	"wpn_falchion":    {ID: "wpn_falchion", Name: "Falchion", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "2d4", Critical: "18-20/x2", Range: 0, Weight: 8, DamageType: "Slashing"},
	"wpn_glaive":      {ID: "wpn_glaive", Name: "Glaive", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d10", Critical: "x3", Range: 0, Weight: 10, DamageType: "Slashing"},
	"wpn_greataxe":    {ID: "wpn_greataxe", Name: "Greataxe", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d12", Critical: "x3", Range: 0, Weight: 12, DamageType: "Slashing"},
	"wpn_greatclub":   {ID: "wpn_greatclub", Name: "Greatclub", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d10", Critical: "x2", Range: 0, Weight: 8, DamageType: "Bludgeoning"},
	"wpn_flail_heavy": {ID: "wpn_flail_heavy", Name: "Flail, Heavy", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d10", Critical: "19-20/x2", Range: 0, Weight: 10, DamageType: "Bludgeoning"},
	"wpn_greatsword":  {ID: "wpn_greatsword", Name: "Greatsword", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "2d6", Critical: "19-20/x2", Range: 0, Weight: 8, DamageType: "Slashing"},
	"wpn_guisarme":    {ID: "wpn_guisarme", Name: "Guisarme", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "2d4", Critical: "x3", Range: 0, Weight: 12, DamageType: "Slashing"},
	"wpn_halberd":     {ID: "wpn_halberd", Name: "Halberd", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d10", Critical: "x3", Range: 0, Weight: 12, DamageType: "Piercing or Slashing"},
	"wpn_lance":       {ID: "wpn_lance", Name: "Lance", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "1d8", Critical: "x3", Range: 0, Weight: 10, DamageType: "Piercing"},
	"wpn_ranseur":     {ID: "wpn_ranseur", Name: "Ranseur", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "2d4", Critical: "x3", Range: 0, Weight: 12, DamageType: "Piercing"},
	"wpn_scythe":      {ID: "wpn_scythe", Name: "Scythe", ProfClass: ProfMartial, Category: CatTwoHanded, Damage: "2d4", Critical: "x4", Range: 0, Weight: 10, DamageType: "Piercing or Slashing"},

	// --- MARTIAL RANGED WEAPONS ---
	"wpn_longbow":       {ID: "wpn_longbow", Name: "Longbow", ProfClass: ProfMartial, Category: CatRanged, Damage: "1d8", Critical: "x3", Range: 100, Weight: 3, DamageType: "Piercing"},
	"wpn_longbow_comp":  {ID: "wpn_longbow_comp", Name: "Longbow, Composite", ProfClass: ProfMartial, Category: CatRanged, Damage: "1d8", Critical: "x3", Range: 110, Weight: 3, DamageType: "Piercing"},
	"wpn_shortbow":      {ID: "wpn_shortbow", Name: "Shortbow", ProfClass: ProfMartial, Category: CatRanged, Damage: "1d6", Critical: "x3", Range: 60, Weight: 2, DamageType: "Piercing"},
	"wpn_shortbow_comp": {ID: "wpn_shortbow_comp", Name: "Shortbow, Composite", ProfClass: ProfMartial, Category: CatRanged, Damage: "1d6", Critical: "x3", Range: 70, Weight: 2, DamageType: "Piercing"},

	// --- EXOTIC LIGHT MELEE WEAPONS ---
	"wpn_kama":     {ID: "wpn_kama", Name: "Kama", ProfClass: ProfExotic, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 0, Weight: 2, DamageType: "Slashing"},
	"wpn_nunchaku": {ID: "wpn_nunchaku", Name: "Nunchaku", ProfClass: ProfExotic, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 0, Weight: 2, DamageType: "Bludgeoning"},
	"wpn_sai":      {ID: "wpn_sai", Name: "Sai", ProfClass: ProfExotic, Category: CatLight, Damage: "1d4", Critical: "x2", Range: 10, Weight: 1, DamageType: "Bludgeoning"},
	"wpn_siangham": {ID: "wpn_siangham", Name: "Siangham", ProfClass: ProfExotic, Category: CatLight, Damage: "1d6", Critical: "x2", Range: 0, Weight: 1, DamageType: "Piercing"},

	// --- EXOTIC ONE-HANDED MELEE WEAPONS ---
	"wpn_sword_bastard": {ID: "wpn_sword_bastard", Name: "Sword, Bastard", ProfClass: ProfExotic, Category: CatOneHanded, Damage: "1d10", Critical: "19-20/x2", Range: 0, Weight: 6, DamageType: "Slashing"},
	"wpn_waraxe_dwarf":  {ID: "wpn_waraxe_dwarf", Name: "Waraxe, Dwarven", ProfClass: ProfExotic, Category: CatOneHanded, Damage: "1d10", Critical: "x3", Range: 0, Weight: 8, DamageType: "Slashing"},
	"wpn_whip":          {ID: "wpn_whip", Name: "Whip", ProfClass: ProfExotic, Category: CatOneHanded, Damage: "1d3 (Nonlethal)", Critical: "x2", Range: 0, Weight: 2, DamageType: "Slashing"},

	// --- EXOTIC TWO-HANDED MELEE WEAPONS ---
	"wpn_axe_orc_double":  {ID: "wpn_axe_orc_double", Name: "Axe, Orc Double", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "1d8/1d8", Critical: "x3", Range: 0, Weight: 15, DamageType: "Slashing"},
	"wpn_chain_spiked":    {ID: "wpn_chain_spiked", Name: "Chain, Spiked", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "2d4", Critical: "x2", Range: 0, Weight: 10, DamageType: "Piercing"},
	"wpn_flail_dire":      {ID: "wpn_flail_dire", Name: "Flail, Dire", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "1d8/1d8", Critical: "x2", Range: 0, Weight: 10, DamageType: "Bludgeoning"},
	"wpn_hammer_gnome":    {ID: "wpn_hammer_gnome", Name: "Hammer, Gnome Hooked", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "1d8/1d6", Critical: "x3/x4", Range: 0, Weight: 6, DamageType: "Bludgeoning and Piercing"},
	"wpn_sword_two_blade": {ID: "wpn_sword_two_blade", Name: "Sword, Two-Bladed", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "1d8/1d8", Critical: "19-20/x2", Range: 0, Weight: 10, DamageType: "Slashing"},
	"wpn_urgrosh_dwarf":   {ID: "wpn_urgrosh_dwarf", Name: "Urgrosh, Dwarven", ProfClass: ProfExotic, Category: CatTwoHanded, Damage: "1d8/1d6", Critical: "x3", Range: 0, Weight: 12, DamageType: "Slashing or Piercing"},

	// --- EXOTIC RANGED WEAPONS ---
	"wpn_bolas":                 {ID: "wpn_bolas", Name: "Bolas", ProfClass: ProfExotic, Category: CatRanged, Damage: "1d4 (Nonlethal)", Critical: "x2", Range: 10, Weight: 2, DamageType: "Bludgeoning"},
	"wpn_crossbow_hand":         {ID: "wpn_crossbow_hand", Name: "Crossbow, Hand", ProfClass: ProfExotic, Category: CatRanged, Damage: "1d4", Critical: "19-20/x2", Range: 30, Weight: 2, DamageType: "Piercing"},
	"wpn_crossbow_repeat_heavy": {ID: "wpn_crossbow_repeat_heavy", Name: "Crossbow, Repeating Heavy", ProfClass: ProfExotic, Category: CatRanged, Damage: "1d10", Critical: "19-20/x2", Range: 120, Weight: 12, DamageType: "Piercing"},
	"wpn_crossbow_repeat_light": {ID: "wpn_crossbow_repeat_light", Name: "Crossbow, Repeating Light", ProfClass: ProfExotic, Category: CatRanged, Damage: "1d8", Critical: "19-20/x2", Range: 80, Weight: 6, DamageType: "Piercing"},
	"wpn_net":                   {ID: "wpn_net", Name: "Net", ProfClass: ProfExotic, Category: CatRanged, Damage: "0", Critical: "x2", Range: 10, Weight: 6, DamageType: "None"},
	"wpn_shuriken":              {ID: "wpn_shuriken", Name: "Shuriken", ProfClass: ProfExotic, Category: CatRanged, Damage: "1d2", Critical: "x2", Range: 10, Weight: 0, DamageType: "Piercing"},
}
