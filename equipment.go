package main

type EquipmentCategory string

const (
	EquipAdventuringGear   EquipmentCategory = "Adventuring Gear"
	EquipSpecialSubstance  EquipmentCategory = "Special Substances and Items"
	EquipToolsAndKits      EquipmentCategory = "Tools and Skill Kits"
	EquipClothing          EquipmentCategory = "Clothing"
	EquipFoodDrinkLodging  EquipmentCategory = "Food, Drink, and Lodging"
	EquipMountsRelatedGear EquipmentCategory = "Mounts and Related Gear"
	EquipTransport         EquipmentCategory = "Transport"
)

type ItemBlueprint struct {
	ID       string
	Name     string
	Category EquipmentCategory
	CostCP   int     // Stored in Copper Pieces for exact math (1 gp = 100 cp, 1 sp = 10 cp)
	Weight   float32 // In pounds
}

// GlobalItems acts as the registry for all standard adventuring gear
var GlobalItems = map[string]ItemBlueprint{
	// --- ADVENTURING GEAR ---
	"itm_backpack":       {ID: "itm_backpack", Name: "Backpack (empty)", Category: EquipAdventuringGear, CostCP: 200, Weight: 2},           //[cite: 31]
	"itm_barrel":         {ID: "itm_barrel", Name: "Barrel (empty)", Category: EquipAdventuringGear, CostCP: 200, Weight: 30},              //[cite: 31]
	"itm_basket":         {ID: "itm_basket", Name: "Basket (empty)", Category: EquipAdventuringGear, CostCP: 40, Weight: 1},                //[cite: 31]
	"itm_bedroll":        {ID: "itm_bedroll", Name: "Bedroll", Category: EquipAdventuringGear, CostCP: 10, Weight: 5},                      //[cite: 31]
	"itm_bell":           {ID: "itm_bell", Name: "Bell", Category: EquipAdventuringGear, CostCP: 100, Weight: 0},                           //[cite: 31]
	"itm_blanket_winter": {ID: "itm_blanket_winter", Name: "Blanket, winter", Category: EquipAdventuringGear, CostCP: 50, Weight: 3},       //[cite: 31]
	"itm_block_tackle":   {ID: "itm_block_tackle", Name: "Block and tackle", Category: EquipAdventuringGear, CostCP: 500, Weight: 5},       //[cite: 31]
	"itm_bottle_wine":    {ID: "itm_bottle_wine", Name: "Bottle, wine glass", Category: EquipAdventuringGear, CostCP: 200, Weight: 0},      //[cite: 31]
	"itm_bucket":         {ID: "itm_bucket", Name: "Bucket (empty)", Category: EquipAdventuringGear, CostCP: 50, Weight: 2},                //[cite: 31]
	"itm_caltrops":       {ID: "itm_caltrops", Name: "Caltrops", Category: EquipAdventuringGear, CostCP: 100, Weight: 2},                   //[cite: 31]
	"itm_candle":         {ID: "itm_candle", Name: "Candle", Category: EquipAdventuringGear, CostCP: 1, Weight: 0},                         //[cite: 31]
	"itm_canvas":         {ID: "itm_canvas", Name: "Canvas (sq. yd.)", Category: EquipAdventuringGear, CostCP: 10, Weight: 1},              //[cite: 31]
	"itm_case_map":       {ID: "itm_case_map", Name: "Case, map or scroll", Category: EquipAdventuringGear, CostCP: 100, Weight: 0.5},      //[cite: 31]
	"itm_chain":          {ID: "itm_chain", Name: "Chain (10 ft.)", Category: EquipAdventuringGear, CostCP: 3000, Weight: 2},               //[cite: 31]
	"itm_chalk":          {ID: "itm_chalk", Name: "Chalk, 1 piece", Category: EquipAdventuringGear, CostCP: 1, Weight: 0},                  //[cite: 31]
	"itm_chest":          {ID: "itm_chest", Name: "Chest (empty)", Category: EquipAdventuringGear, CostCP: 200, Weight: 25},                //[cite: 31]
	"itm_crowbar":        {ID: "itm_crowbar", Name: "Crowbar", Category: EquipAdventuringGear, CostCP: 200, Weight: 5},                     //[cite: 31]
	"itm_firewood":       {ID: "itm_firewood", Name: "Firewood (per day)", Category: EquipAdventuringGear, CostCP: 1, Weight: 20},          //[cite: 31]
	"itm_fishhook":       {ID: "itm_fishhook", Name: "Fishhook", Category: EquipAdventuringGear, CostCP: 10, Weight: 0},                    //[cite: 31]
	"itm_fishing_net":    {ID: "itm_fishing_net", Name: "Fishing net, 25 sq. ft.", Category: EquipAdventuringGear, CostCP: 400, Weight: 5}, //[cite: 31]
	"itm_flask":          {ID: "itm_flask", Name: "Flask (empty)", Category: EquipAdventuringGear, CostCP: 3, Weight: 1.5},                 //[cite: 31]
	"itm_flint_steel":    {ID: "itm_flint_steel", Name: "Flint and steel", Category: EquipAdventuringGear, CostCP: 100, Weight: 0},         //[cite: 31]
	"itm_grappling_hook": {ID: "itm_grappling_hook", Name: "Grappling hook", Category: EquipAdventuringGear, CostCP: 100, Weight: 4},       //[cite: 31]
	"itm_hammer":         {ID: "itm_hammer", Name: "Hammer", Category: EquipAdventuringGear, CostCP: 50, Weight: 2},                        //[cite: 31]
	"itm_ink":            {ID: "itm_ink", Name: "Ink (1 oz. vial)", Category: EquipAdventuringGear, CostCP: 800, Weight: 0},                //[cite: 31]
	"itm_inkpen":         {ID: "itm_inkpen", Name: "Inkpen", Category: EquipAdventuringGear, CostCP: 10, Weight: 0},                        //[cite: 31]
	"itm_jug_clay":       {ID: "itm_jug_clay", Name: "Jug, clay", Category: EquipAdventuringGear, CostCP: 3, Weight: 9},                    //[cite: 31]
	"itm_ladder":         {ID: "itm_ladder", Name: "Ladder, 10-foot", Category: EquipAdventuringGear, CostCP: 5, Weight: 20},               //[cite: 31]
	"itm_lamp_common":    {ID: "itm_lamp_common", Name: "Lamp, common", Category: EquipAdventuringGear, CostCP: 10, Weight: 1},             //[cite: 31]
	"itm_lantern_bulls":  {ID: "itm_lantern_bulls", Name: "Lantern, bullseye", Category: EquipAdventuringGear, CostCP: 1200, Weight: 3},    //[cite: 31]
	"itm_lantern_hood":   {ID: "itm_lantern_hood", Name: "Lantern, hooded", Category: EquipAdventuringGear, CostCP: 700, Weight: 2},        //[cite: 31]
	"itm_lock_simple":    {ID: "itm_lock_simple", Name: "Lock, Very simple", Category: EquipAdventuringGear, CostCP: 2000, Weight: 1},      //[cite: 31]
	"itm_lock_average":   {ID: "itm_lock_average", Name: "Lock, Average", Category: EquipAdventuringGear, CostCP: 4000, Weight: 1},         //[cite: 31]
	"itm_lock_good":      {ID: "itm_lock_good", Name: "Lock, Good", Category: EquipAdventuringGear, CostCP: 8000, Weight: 1},               //[cite: 31]
	"itm_lock_amazing":   {ID: "itm_lock_amazing", Name: "Lock, Amazing", Category: EquipAdventuringGear, CostCP: 15000, Weight: 1},        //[cite: 31]
	"itm_manacles":       {ID: "itm_manacles", Name: "Manacles", Category: EquipAdventuringGear, CostCP: 1500, Weight: 2},                  //[cite: 31]
	"itm_manacles_mwk":   {ID: "itm_manacles_mwk", Name: "Manacles, masterwork", Category: EquipAdventuringGear, CostCP: 5000, Weight: 2},  //[cite: 31]
	"itm_mirror":         {ID: "itm_mirror", Name: "Mirror, small steel", Category: EquipAdventuringGear, CostCP: 1000, Weight: 0.5},       //[cite: 31]
	"itm_mug":            {ID: "itm_mug", Name: "Mug/Tankard, clay", Category: EquipAdventuringGear, CostCP: 2, Weight: 1},                 //[cite: 31]
	"itm_oil":            {ID: "itm_oil", Name: "Oil (1-pint flask)", Category: EquipAdventuringGear, CostCP: 10, Weight: 1},               //[cite: 31]
	"itm_paper":          {ID: "itm_paper", Name: "Paper (sheet)", Category: EquipAdventuringGear, CostCP: 40, Weight: 0},                  //[cite: 31]
	"itm_parchment":      {ID: "itm_parchment", Name: "Parchment (sheet)", Category: EquipAdventuringGear, CostCP: 20, Weight: 0},          //[cite: 31]
	"itm_pick_miner":     {ID: "itm_pick_miner", Name: "Pick, miner's", Category: EquipAdventuringGear, CostCP: 300, Weight: 10},           //[cite: 31]
	"itm_pitcher_clay":   {ID: "itm_pitcher_clay", Name: "Pitcher, clay", Category: EquipAdventuringGear, CostCP: 2, Weight: 5},            //[cite: 31]
	"itm_piton":          {ID: "itm_piton", Name: "Piton", Category: EquipAdventuringGear, CostCP: 10, Weight: 0.5},                        //[cite: 31]
	"itm_pole":           {ID: "itm_pole", Name: "Pole, 10-foot", Category: EquipAdventuringGear, CostCP: 20, Weight: 8},                   //[cite: 31]
	"itm_pot_iron":       {ID: "itm_pot_iron", Name: "Pot, iron", Category: EquipAdventuringGear, CostCP: 50, Weight: 10},                  //[cite: 31]
	"itm_pouch_belt":     {ID: "itm_pouch_belt", Name: "Pouch, belt (empty)", Category: EquipAdventuringGear, CostCP: 100, Weight: 0.5},    //[cite: 31]
	"itm_ram_portable":   {ID: "itm_ram_portable", Name: "Ram, portable", Category: EquipAdventuringGear, CostCP: 1000, Weight: 20},        //[cite: 31]
	"itm_rations":        {ID: "itm_rations", Name: "Rations, trail (per day)", Category: EquipAdventuringGear, CostCP: 50, Weight: 1},     //[cite: 31]
	"itm_rope_hempen":    {ID: "itm_rope_hempen", Name: "Rope, hempen (50 ft.)", Category: EquipAdventuringGear, CostCP: 100, Weight: 10},  //[cite: 31]
	"itm_rope_silk":      {ID: "itm_rope_silk", Name: "Rope, silk (50 ft.)", Category: EquipAdventuringGear, CostCP: 1000, Weight: 5},      //[cite: 31]
	"itm_sack":           {ID: "itm_sack", Name: "Sack (empty)", Category: EquipAdventuringGear, CostCP: 10, Weight: 0.5},                  //[cite: 31]
	"itm_sealing_wax":    {ID: "itm_sealing_wax", Name: "Sealing wax", Category: EquipAdventuringGear, CostCP: 100, Weight: 1},             //[cite: 31]
	"itm_sewing_needle":  {ID: "itm_sewing_needle", Name: "Sewing needle", Category: EquipAdventuringGear, CostCP: 50, Weight: 0},          //[cite: 31]
	"itm_signal_whistle": {ID: "itm_signal_whistle", Name: "Signal whistle", Category: EquipAdventuringGear, CostCP: 80, Weight: 0},        //[cite: 31]
	"itm_signet_ring":    {ID: "itm_signet_ring", Name: "Signet ring", Category: EquipAdventuringGear, CostCP: 500, Weight: 0},             //[cite: 31]
	"itm_sledge":         {ID: "itm_sledge", Name: "Sledge", Category: EquipAdventuringGear, CostCP: 100, Weight: 10},                      //[cite: 31]
	"itm_soap":           {ID: "itm_soap", Name: "Soap (per lb.)", Category: EquipAdventuringGear, CostCP: 50, Weight: 1},                  //[cite: 31]
	"itm_spade":          {ID: "itm_spade", Name: "Spade or shovel", Category: EquipAdventuringGear, CostCP: 200, Weight: 8},               //[cite: 31]
	"itm_spyglass":       {ID: "itm_spyglass", Name: "Spyglass", Category: EquipAdventuringGear, CostCP: 100000, Weight: 1},                //[cite: 31]
	"itm_tent":           {ID: "itm_tent", Name: "Tent", Category: EquipAdventuringGear, CostCP: 1000, Weight: 20},                         //[cite: 31]
	"itm_torch":          {ID: "itm_torch", Name: "Torch", Category: EquipAdventuringGear, CostCP: 1, Weight: 1},                           //[cite: 31]
	"itm_vial":           {ID: "itm_vial", Name: "Vial ink or potion", Category: EquipAdventuringGear, CostCP: 100, Weight: 0.1},           //[cite: 31]
	"itm_waterskin":      {ID: "itm_waterskin", Name: "Waterskin", Category: EquipAdventuringGear, CostCP: 100, Weight: 4},                 //[cite: 31]
	"itm_whetstone":      {ID: "itm_whetstone", Name: "Whetstone", Category: EquipAdventuringGear, CostCP: 2, Weight: 1},                   //[cite: 31]

	// --- SPECIAL SUBSTANCES AND ITEMS ---
	"itm_acid":              {ID: "itm_acid", Name: "Acid (flask)", Category: EquipSpecialSubstance, CostCP: 1000, Weight: 1},                        //[cite: 31]
	"itm_alchemists_fire":   {ID: "itm_alchemists_fire", Name: "Alchemist's fire (flask)", Category: EquipSpecialSubstance, CostCP: 2000, Weight: 1}, //[cite: 31]
	"itm_antitoxin":         {ID: "itm_antitoxin", Name: "Antitoxin (vial)", Category: EquipSpecialSubstance, CostCP: 5000, Weight: 0},               //[cite: 31]
	"itm_everburning_torch": {ID: "itm_everburning_torch", Name: "Everburning torch", Category: EquipSpecialSubstance, CostCP: 11000, Weight: 1},     //[cite: 31]
	"itm_holy_water":        {ID: "itm_holy_water", Name: "Holy water (flask)", Category: EquipSpecialSubstance, CostCP: 2500, Weight: 1},            //[cite: 31]
	"itm_smokestick":        {ID: "itm_smokestick", Name: "Smokestick", Category: EquipSpecialSubstance, CostCP: 2000, Weight: 0.5},                  //[cite: 31]
	"itm_sunrod":            {ID: "itm_sunrod", Name: "Sunrod", Category: EquipSpecialSubstance, CostCP: 200, Weight: 1},                             //[cite: 31]
	"itm_tanglefoot_bag":    {ID: "itm_tanglefoot_bag", Name: "Tanglefoot bag", Category: EquipSpecialSubstance, CostCP: 5000, Weight: 4},            //[cite: 31]
	"itm_thunderstone":      {ID: "itm_thunderstone", Name: "Thunderstone", Category: EquipSpecialSubstance, CostCP: 3000, Weight: 1},                //[cite: 31]
	"itm_tindertwig":        {ID: "itm_tindertwig", Name: "Tindertwig", Category: EquipSpecialSubstance, CostCP: 100, Weight: 0},                     //[cite: 31]

	// --- TOOLS AND SKILL KITS ---
	"itm_alchemists_lab":   {ID: "itm_alchemists_lab", Name: "Alchemist's lab", Category: EquipToolsAndKits, CostCP: 50000, Weight: 40},                 //[cite: 31]
	"itm_artisans_tools":   {ID: "itm_artisans_tools", Name: "Artisan's tools", Category: EquipToolsAndKits, CostCP: 500, Weight: 5},                    //[cite: 31]
	"itm_artisans_tools_m": {ID: "itm_artisans_tools_m", Name: "Artisan's tools, masterwork", Category: EquipToolsAndKits, CostCP: 5500, Weight: 5},     //[cite: 31]
	"itm_climbers_kit":     {ID: "itm_climbers_kit", Name: "Climber's kit", Category: EquipToolsAndKits, CostCP: 8000, Weight: 5},                       //[cite: 31]
	"itm_disguise_kit":     {ID: "itm_disguise_kit", Name: "Disguise kit", Category: EquipToolsAndKits, CostCP: 5000, Weight: 8},                        //[cite: 31]
	"itm_healers_kit":      {ID: "itm_healers_kit", Name: "Healer's kit", Category: EquipToolsAndKits, CostCP: 5000, Weight: 1},                         //[cite: 31]
	"itm_holy_symbol_wood": {ID: "itm_holy_symbol_wood", Name: "Holy symbol, wooden", Category: EquipToolsAndKits, CostCP: 100, Weight: 0},              //[cite: 31]
	"itm_holy_symbol_slvr": {ID: "itm_holy_symbol_slvr", Name: "Holy symbol, silver", Category: EquipToolsAndKits, CostCP: 2500, Weight: 1},             //[cite: 31]
	"itm_musical_inst":     {ID: "itm_musical_inst", Name: "Musical instrument, common", Category: EquipToolsAndKits, CostCP: 500, Weight: 3},           //[cite: 31]
	"itm_musical_inst_mwk": {ID: "itm_musical_inst_mwk", Name: "Musical instrument, masterwork", Category: EquipToolsAndKits, CostCP: 10000, Weight: 3}, //[cite: 31]
	"itm_spell_comp_pouch": {ID: "itm_spell_comp_pouch", Name: "Spell component pouch", Category: EquipToolsAndKits, CostCP: 500, Weight: 2},            //[cite: 31]
	"itm_spellbook_blank":  {ID: "itm_spellbook_blank", Name: "Spellbook, wizard's (blank)", Category: EquipToolsAndKits, CostCP: 1500, Weight: 3},      //[cite: 31]
	"itm_thieves_tools":    {ID: "itm_thieves_tools", Name: "Thieves' tools", Category: EquipToolsAndKits, CostCP: 3000, Weight: 1},                     //[cite: 31]
	"itm_thieves_tools_m":  {ID: "itm_thieves_tools_m", Name: "Thieves' tools, masterwork", Category: EquipToolsAndKits, CostCP: 10000, Weight: 2},      //[cite: 31]
	"itm_tool_mwk":         {ID: "itm_tool_mwk", Name: "Tool, masterwork", Category: EquipToolsAndKits, CostCP: 5000, Weight: 1},                        //[cite: 31]

	// --- CLOTHING ---
	"itm_outfit_artisan":  {ID: "itm_outfit_artisan", Name: "Artisan's outfit", Category: EquipClothing, CostCP: 100, Weight: 4},    //[cite: 31]
	"itm_outfit_cleric":   {ID: "itm_outfit_cleric", Name: "Cleric's vestments", Category: EquipClothing, CostCP: 500, Weight: 6},   //[cite: 31]
	"itm_outfit_cold":     {ID: "itm_outfit_cold", Name: "Cold weather outfit", Category: EquipClothing, CostCP: 800, Weight: 7},    //[cite: 31]
	"itm_outfit_courtier": {ID: "itm_outfit_courtier", Name: "Courtier's outfit", Category: EquipClothing, CostCP: 3000, Weight: 6}, //[cite: 31]
	"itm_outfit_explorer": {ID: "itm_outfit_explorer", Name: "Explorer's outfit", Category: EquipClothing, CostCP: 1000, Weight: 8}, //[cite: 31]
	"itm_outfit_monk":     {ID: "itm_outfit_monk", Name: "Monk's outfit", Category: EquipClothing, CostCP: 500, Weight: 2},          //[cite: 31]
	"itm_outfit_noble":    {ID: "itm_outfit_noble", Name: "Noble's outfit", Category: EquipClothing, CostCP: 7500, Weight: 10},      //[cite: 31]
	"itm_outfit_peasant":  {ID: "itm_outfit_peasant", Name: "Peasant's outfit", Category: EquipClothing, CostCP: 10, Weight: 2},     //[cite: 31]
	"itm_outfit_royal":    {ID: "itm_outfit_royal", Name: "Royal outfit", Category: EquipClothing, CostCP: 20000, Weight: 15},       //[cite: 31]
	"itm_outfit_scholar":  {ID: "itm_outfit_scholar", Name: "Scholar's outfit", Category: EquipClothing, CostCP: 500, Weight: 6},    //[cite: 31]
	"itm_outfit_traveler": {ID: "itm_outfit_traveler", Name: "Traveler's outfit", Category: EquipClothing, CostCP: 100, Weight: 5},  //[cite: 31]
}
