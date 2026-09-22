package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/harbdog/raycaster-go"
	"github.com/harbdog/raycaster-go/geom"
)

// --- TEXT WRAPPER UTILITY ---
func wrapText(text string, maxChars int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}
	currentLine := words[0]
	for _, word := range words[1:] {
		if len(currentLine)+len(word)+1 > maxChars {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

// --- FLAVOR TEXT DATABASES ---
var RaceFlavorText = map[string]string{
	"race_human":              "Most humans are the descendants of pioneers, conquerors, traders, travelers, refugees, and other people on the move. As a result, human lands are home to a mix of people physically, culturally, religiously, and politically different. Hardy or fine, light-skinned or dark, showy or austere, primitive or civilized, devout or impious, humans run the gamut. They are the most adaptable, flexible, and ambitious people among the common races.",
	"race_dwarf_hill":         "Dwarves are known for their skill in warfare, their ability to withstand physical and magical punishment, their knowledge of the earth's secrets, their hard work, and their capacity for drinking ale. Their mysterious kingdoms, carved out from the insides of mountains, are renowned for the marvelous treasures that they produce as gifts or for trade. They are slow to laugh or jest and suspicious of strangers, but they are generous to those few who earn their trust.",
	"race_elf_high":           "Elves mingle freely in human lands, always welcome yet never at home there. They are well known for their poetry, dance, song, lore, and magical arts. Elves favor things of natural and simple beauty. When danger threatens their woodland homes, however, elves reveal a more martial side, demonstrating skill with sword, bow, and battle strategy. With such a long life span, they tend to keep a broad perspective on events, remaining aloof and unfazed by petty happenstance.",
	"race_gnome_rock":         "Gnomes are welcome everywhere as technicians, alchemists, and inventors. Despite the demand for their skills, most gnomes prefer to remain among their own kind, living in comfortable burrows beneath rolling, wooded hills where animals abound. Gnomes adore animals, beautiful gems, and jokes of all kinds. They have a great sense of humor, and while they love puns, jokes, and games, they relish tricks - the more intricate the better.",
	"race_half_elf":           "Humans and elves sometimes wed, the elf attracted to the human's energy and the human to the elf's grace. These marriages end quickly as elves count years because a human's life is so brief, but they leave an enduring legacy - half-elf children. Most half-elves have the curiosity, inventiveness, and ambition of the human parent, along with the refined senses, love of nature, and artistic tastes of the elf parent. They make excellent ambassadors and go-betweens.",
	"race_half_orc":           "In the wild frontiers, tribes of human and orc barbarians live in uneasy balance, fighting in times of war and trading in times of peace. Half-orcs who are born in the frontier may live with either human or orc parents, but they are nevertheless exposed to both cultures. Half-orcs tend to be short-tempered and sullen. They would rather act than ponder and would rather fight than argue. They love simple pleasures such as feasting, drinking, boasting, and wild dancing.",
	"race_halfling_lightfoot": "Halflings are clever, capable opportunists. Halfling individuals and clans find room for themselves wherever they can. Often they are strangers and wanderers, and others react to them with suspicion or curiosity. Depending on the clan, halflings might be reliable, hard-working citizens, or they might be thieves just waiting for the opportunity to make a big score and disappear in the dead of night. Regardless, halflings are cunning, resourceful survivors.",
}

var ClassFlavorText = map[string]string{
	"cls_barbarian": "The barbarian is an excellent melee combatant. While not as disciplined or skilled as a fighter, the barbarian can fly into a berserker rage, becoming a ferociously tough and overwhelming opponent.",
	"cls_bard":      "Bards often serve as negotiators, messengers, scouts, and spies. Their magic comes from the heart. If their heart is good, a bard brings hope and courage to the downtrodden. If evil, the bard's magic inspires fear and despair.",
	"cls_cleric":    "Clerics act as intermediaries between the mortal world and the realms of the gods. As disparate as the gods they serve, clerics strive to embody their deities' motives and intentions.",
	"cls_druid":     "Druids channel the raw, untamed forces of nature. They gain their magic not from the divine, but from nature itself. They are capable of adopting the forms of animals and eventually larger, more powerful creatures.",
	"cls_fighter":   "Fighters can be knights, conquerors, royal champions, elite foot soldiers, or hardened mercenaries. Fighters are the best melee combatants in the game, mastering a wide variety of weapons and armor.",
	"cls_monk":      "Monks are martial artists whose unarmed strikes hit as hard as a weapon. Through discipline and training, they tap into the spiritual energy of their own bodies, achieving extraordinary feats of agility and strength.",
	"cls_paladin":   "Paladins are the champions of justice and destroyers of evil. Protected by their purity, they can heal wounds, cure diseases, and strike down wicked foes with divine power.",
	"cls_ranger":    "Rangers are skilled trackers and woodsmen. They often protect the fringes of civilization from the terrors of the wilderness, wielding specialized combat styles and nature magic.",
	"cls_rogue":     "Rogues share little in common with each other. Some are stealthy thieves; others are silver-tongued tricksters. They excel at striking where it hurts most and overcoming traps and locks.",
	"cls_sorcerer":  "Sorcerers cast spells through innate power rather than through careful training and study. Their magic is intuitive, derived from a draconic heritage or a strange magical anomaly in their bloodline.",
	"cls_wizard":    "Wizards are supreme magic-users, defined and united as a class by the spells they cast. Drawing on years of intense study and ancient tomes, they command powers capable of altering reality itself.",
}

// --- SPRITE IMPLEMENTATION ---
type GameSprite struct {
	X, Y        float64
	Z           float64
	ScaleFactor float64
	Img         *ebiten.Image
}

func (s *GameSprite) Pos() *geom.Vector2                     { return &geom.Vector2{X: s.X, Y: s.Y} }
func (s *GameSprite) PosZ() float64                          { return s.Z }
func (s *GameSprite) Scale() float64                         { return s.ScaleFactor }
func (s *GameSprite) VerticalAnchor() raycaster.SpriteAnchor { return raycaster.AnchorBottom }
func (s *GameSprite) Texture() *ebiten.Image                 { return s.Img }
func (s *GameSprite) TextureRect() image.Rectangle           { return s.Img.Bounds() }
func (s *GameSprite) Illumination() float64                  { return 0 }
func (s *GameSprite) SetScreenRect(rect *image.Rectangle)    {}
func (s *GameSprite) IsFocusable() bool                      { return true }

var doorOpenTex *ebiten.Image

func getDoorOpenTex() *ebiten.Image {
	if doorOpenTex == nil {
		doorOpenTex = ebiten.NewImage(64, 64)
		doorOpenTex.Fill(color.RGBA{0, 0, 0, 0}) // Transparent core
		ebitenutil.DrawRect(doorOpenTex, 0, 0, 64, 4, color.RGBA{255, 165, 0, 255})
		ebitenutil.DrawRect(doorOpenTex, 0, 60, 64, 4, color.RGBA{255, 165, 0, 255})
		ebitenutil.DrawRect(doorOpenTex, 0, 0, 4, 64, color.RGBA{255, 165, 0, 255})
		ebitenutil.DrawRect(doorOpenTex, 60, 0, 4, 64, color.RGBA{255, 165, 0, 255})
	}
	return doorOpenTex
}

// --- RAYMAP ---
type RayMap struct {
	DungeonMap     *Map
	WallTex        *ebiten.Image
	DoorTex        *ebiten.Image
	DoorTrappedTex *ebiten.Image
}

func (m *RayMap) HitCheck(x, y float64) bool {
	gridX, gridY := int(math.Floor(x)), int(math.Floor(y))
	if gridX < 0 || gridX >= Width || gridY < 0 || gridY >= Height {
		return true
	}
	tile := m.DungeonMap.Grid[gridY][gridX]
	return tile == Wall || tile == DoorC
}

func (m *RayMap) GetHitColor(x, y float64) color.RGBA {
	return color.RGBA{100, 100, 100, 255}
}

func (m *RayMap) Level(levelNum int) [][]int {
	if levelNum != 0 {
		return nil
	}
	levelArr := make([][]int, Width)
	for x := 0; x < Width; x++ {
		levelArr[x] = make([]int, Height)
		for y := 0; y < Height; y++ {
			tile := m.DungeonMap.Grid[y][x]
			if tile == Wall || tile == DoorC {
				levelArr[x][y] = 1
			} else {
				levelArr[x][y] = 0
			}
		}
	}
	return levelArr
}

func (m *RayMap) NumLevels() int {
	return 1
}

func (m *RayMap) TextureAt(x, y, levelNum, side int) *ebiten.Image {
	if m.WallTex == nil {
		m.WallTex = ebiten.NewImage(64, 64)
		m.WallTex.Fill(color.RGBA{100, 100, 100, 255})
		m.DoorTex = ebiten.NewImage(64, 64)
		m.DoorTex.Fill(color.RGBA{139, 69, 19, 255})
		m.DoorTrappedTex = ebiten.NewImage(64, 64)
		m.DoorTrappedTex.Fill(color.RGBA{200, 50, 50, 255})
	}
	if x >= 0 && x < Width && y >= 0 && y < Height {
		if m.DungeonMap.Grid[y][x] == DoorC {
			if door, ok := m.DungeonMap.Doors[Node{x, y}]; ok && door.IsTrapDetected {
				return m.DoorTrappedTex
			}
			return m.DoorTex
		}
	}
	return m.WallTex
}

func (m *RayMap) FloorTextureAt(x, y int) *image.RGBA {
	return nil
}

type Game struct {
	DungeonMap       *Map
	Camera           *raycaster.Camera
	PlayerX          float64
	PlayerY          float64
	PlayerA          float64
	ActionLog        []string
	CurrentRoomID    int
	RadialMenuOpen   bool
	RadialMenuX      int
	RadialMenuY      int
	AvailableActions []RadialAction

	// Character Creation State
	CharCreationOpen  bool
	CharCreationStep  int
	PointBuyStats     map[string]int
	PointsRemaining   int
	SelectedStatIndex int
	StatOrder         []string

	// Race Selection State
	StandardRaces     []string
	SelectedRaceIndex int
	RaceTextOffset    int

	// Class Selection State
	StandardClasses    []string
	SelectedClassIndex int
	ClassTextOffset    int

	// Skill Allocation State
	SkillPointsRemaining int
	SkillRanksAllocated  map[string]int
	SelectedSkillIndex   int
	SkillTextOffset      int
	SkillOrder           []string
}

func (g *Game) BroadcastNoise(originX, originY int, volumeRadius float64) {
	// TODO (Phase 4): Emits noise for passive perception
}

func calculatePointCost(score int) int {
	switch score {
	case 8:
		return 0
	case 9:
		return 1
	case 10:
		return 2
	case 11:
		return 3
	case 12:
		return 4
	case 13:
		return 5
	case 14:
		return 6
	case 15:
		return 8
	case 16:
		return 10
	case 17:
		return 13
	case 18:
		return 16
	default:
		return 0
	}
}

func (g *Game) Update() error {
	// --- CHARACTER CREATION INTERCEPT ---
	if ebiten.IsKeyPressed(ebiten.KeyShift) && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.CharCreationOpen = !g.CharCreationOpen
		if g.CharCreationOpen && g.PointBuyStats == nil {
			g.CharCreationStep = 1
			g.StatOrder = []string{"STR", "DEX", "CON", "INT", "WIS", "CHA"}
			g.PointBuyStats = map[string]int{"STR": 8, "DEX": 8, "CON": 8, "INT": 8, "WIS": 8, "CHA": 8}
			g.PointsRemaining = 25
			g.SelectedStatIndex = 0

			g.StandardRaces = []string{"race_human", "race_dwarf_hill", "race_elf_high", "race_gnome_rock", "race_half_elf", "race_half_orc", "race_halfling_lightfoot"}
			g.SelectedRaceIndex = 0
			g.RaceTextOffset = 0

			g.StandardClasses = []string{"cls_barbarian", "cls_bard", "cls_cleric", "cls_druid", "cls_fighter", "cls_monk", "cls_paladin", "cls_ranger", "cls_rogue", "cls_sorcerer", "cls_wizard"}
			g.SelectedClassIndex = 0
			g.ClassTextOffset = 0
		}
	}

	if g.CharCreationOpen {
		if g.CharCreationStep == 1 {
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				g.SelectedStatIndex--
				if g.SelectedStatIndex < 0 {
					g.SelectedStatIndex = 5
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				g.SelectedStatIndex++
				if g.SelectedStatIndex > 5 {
					g.SelectedStatIndex = 0
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				stat := g.StatOrder[g.SelectedStatIndex]
				current := g.PointBuyStats[stat]
				if current < 18 {
					costDiff := calculatePointCost(current+1) - calculatePointCost(current)
					if g.PointsRemaining >= costDiff {
						g.PointBuyStats[stat]++
						g.PointsRemaining -= costDiff
					}
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				stat := g.StatOrder[g.SelectedStatIndex]
				current := g.PointBuyStats[stat]
				if current > 8 {
					costDiff := calculatePointCost(current) - calculatePointCost(current-1)
					g.PointBuyStats[stat]--
					g.PointsRemaining += costDiff
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
				g.CharCreationStep = 2 // Move to Race Selection
			}
		} else if g.CharCreationStep == 2 {
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				g.SelectedRaceIndex++
				if g.SelectedRaceIndex >= len(g.StandardRaces) {
					g.SelectedRaceIndex = 0
				}
				g.RaceTextOffset = 0
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				g.SelectedRaceIndex--
				if g.SelectedRaceIndex < 0 {
					g.SelectedRaceIndex = len(g.StandardRaces) - 1
				}
				g.RaceTextOffset = 0
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				g.RaceTextOffset++
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				g.RaceTextOffset--
				if g.RaceTextOffset < 0 {
					g.RaceTextOffset = 0
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
				g.CharCreationStep = 1 // Go back to Step 1
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
				g.CharCreationStep = 3 // Move to Class Selection
			}
		} else if g.CharCreationStep == 3 {
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				g.SelectedClassIndex++
				if g.SelectedClassIndex >= len(g.StandardClasses) {
					g.SelectedClassIndex = 0
				}
				g.ClassTextOffset = 0
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				g.SelectedClassIndex--
				if g.SelectedClassIndex < 0 {
					g.SelectedClassIndex = len(g.StandardClasses) - 1
				}
				g.ClassTextOffset = 0
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				g.ClassTextOffset++
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				g.ClassTextOffset--
				if g.ClassTextOffset < 0 {
					g.ClassTextOffset = 0
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
				g.CharCreationStep = 2 // Go back to Step 2
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
				g.CharCreationStep = 6 // Move directly to Skill Allocation

				// --- Step 4/5 Background Math Execution ---
				raceID := g.StandardRaces[g.SelectedRaceIndex]
				race := GlobalRaces[raceID]

				// Calculate effective INT modifier
				intScore := g.PointBuyStats["INT"] + race.AbilityMods["INT"]
				intMod := int(math.Floor(float64(intScore-10) / 2.0))

				classID := g.StandardClasses[g.SelectedClassIndex]
				cls := GlobalClasses[classID]

				// 3.5e Rule: (Class Skill Points + INT Mod) * 4 for 1st level
				pointsPerLevel := cls.SkillPoints + intMod
				if pointsPerLevel < 1 {
					pointsPerLevel = 1
				}
				g.SkillPointsRemaining = pointsPerLevel * 4

				// Apply Human racial bonus (+4 skill points at 1st level)
				if raceID == "race_human" {
					g.SkillPointsRemaining += 4
				}

				g.SkillRanksAllocated = make(map[string]int)
				g.SelectedSkillIndex = 0
				g.SkillTextOffset = 0
				g.SkillOrder = []string{
					"skill_appraise", "skill_balance", "skill_bluff", "skill_climb", "skill_concentration",
					"skill_decipher_script", "skill_diplomacy", "skill_disable_device", "skill_disguise",
					"skill_escape_artist", "skill_forgery", "skill_gather_info", "skill_handle_animal",
					"skill_heal", "skill_hide", "skill_intimidate", "skill_jump", "skill_listen",
					"skill_move_silently", "skill_open_lock", "skill_ride", "skill_search",
					"skill_sense_motive", "skill_sleight_of_hand", "skill_spellcraft", "skill_spot",
					"skill_survival", "skill_swim", "skill_tumble", "skill_use_magic_device", "skill_use_rope",
				}
			}
		} else if g.CharCreationStep == 6 {
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				g.SelectedSkillIndex++
				if g.SelectedSkillIndex >= len(g.SkillOrder) {
					g.SelectedSkillIndex = 0
				}
				// Auto-scroll bounds logic
				if g.SelectedSkillIndex >= g.SkillTextOffset+15 {
					g.SkillTextOffset = g.SelectedSkillIndex - 14
				} else if g.SelectedSkillIndex < g.SkillTextOffset {
					g.SkillTextOffset = g.SelectedSkillIndex
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				g.SelectedSkillIndex--
				if g.SelectedSkillIndex < 0 {
					g.SelectedSkillIndex = len(g.SkillOrder) - 1
				}
				// Auto-scroll bounds logic
				if g.SelectedSkillIndex < g.SkillTextOffset {
					g.SkillTextOffset = g.SelectedSkillIndex
				} else if g.SelectedSkillIndex >= g.SkillTextOffset+15 {
					g.SkillTextOffset = g.SelectedSkillIndex - 14
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				skill := g.SkillOrder[g.SelectedSkillIndex]
				// Hard cap of 4 ranks for a 1st level character
				if g.SkillPointsRemaining > 0 && g.SkillRanksAllocated[skill] < 4 {
					g.SkillRanksAllocated[skill]++
					g.SkillPointsRemaining--
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				skill := g.SkillOrder[g.SelectedSkillIndex]
				if g.SkillRanksAllocated[skill] > 0 {
					g.SkillRanksAllocated[skill]--
					g.SkillPointsRemaining++
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
				g.CharCreationStep = 3 // Go back to Step 3
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
				// Finish Wizard and return to game
				g.CharCreationOpen = false
			}
		}
		return nil // Freeze normal 3D gameplay while the menu is open
	}

	moveSpeed := 0.08
	rotSpeed := 0.05

	newX, newY := g.PlayerX, g.PlayerY

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		newX += math.Cos(g.PlayerA) * moveSpeed
		newY += math.Sin(g.PlayerA) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		newX -= math.Cos(g.PlayerA) * moveSpeed
		newY -= math.Sin(g.PlayerA) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		newX += math.Cos(g.PlayerA+math.Pi/2) * moveSpeed
		newY += math.Sin(g.PlayerA+math.Pi/2) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		newX += math.Cos(g.PlayerA-math.Pi/2) * moveSpeed
		newY += math.Sin(g.PlayerA-math.Pi/2) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.PlayerA += rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.PlayerA -= rotSpeed
	}

	pGridX, pGridY := int(math.Floor(g.PlayerX)), int(math.Floor(g.PlayerY))
	detectedRoomID := 0
	var detectedRoom *Room

	for i := range g.DungeonMap.Rooms {
		r := &g.DungeonMap.Rooms[i]
		if pGridX >= r.X && pGridX < r.X+r.W && pGridY >= r.Y && pGridY < r.Y+r.H {
			detectedRoomID = r.ID
			detectedRoom = r
			break
		}
	}

	if detectedRoomID != g.CurrentRoomID {
		var transitionMsg string
		if g.CurrentRoomID != 0 && detectedRoomID == 0 {
			transitionMsg = fmt.Sprintf("Exited Room ID %d", g.CurrentRoomID)
		} else {
			transitionMsg = fmt.Sprintf("Entered Room ID %d (%s)", detectedRoomID, detectedRoom.RType)
		}
		g.ActionLog = append([]string{transitionMsg}, g.ActionLog...)
		if len(g.ActionLog) > 15 {
			g.ActionLog = g.ActionLog[:15]
		}
		g.CurrentRoomID = detectedRoomID
	}

	// --- RADIAL MENU INTERACTION ---
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.RadialMenuOpen {
			mx, my := ebiten.CursorPosition()
			angleStep := 2 * math.Pi / float64(len(g.AvailableActions))
			radius := 60.0

			for i, action := range g.AvailableActions {
				ax := float64(g.RadialMenuX) + math.Cos(float64(i)*angleStep)*radius
				ay := float64(g.RadialMenuY) + math.Sin(float64(i)*angleStep)*radius

				dist := math.Hypot(float64(mx)-ax, float64(my)-ay)
				if dist <= 25.0 {
					if action.Execute != nil {
						action.Execute()
					}
					// TODO: Add transition logic into action.Children if present
					break
				}
			}
			g.RadialMenuOpen = false
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && !g.CharCreationOpen {
		mx, my := ebiten.CursorPosition()
		if mx >= 0 && mx < 700 && my >= 0 && my < 768 {
			fov := 65.0 * math.Pi / 180.0
			rayAngle := g.PlayerA + math.Atan2(float64(mx-350), 350/math.Tan(fov/2))

			hit := false
			step := 0.1
			for dist := 0.0; dist < 20.0; dist += step {
				checkX := g.PlayerX + math.Cos(rayAngle)*dist
				checkY := g.PlayerY + math.Sin(rayAngle)*dist
				gridX, gridY := int(math.Floor(checkX)), int(math.Floor(checkY))

				if gridX >= 0 && gridX < Width && gridY >= 0 && gridY < Height {
					targetNode := Node{gridX, gridY}

					if door, ok := g.DungeonMap.Doors[targetNode]; ok {
						g.AvailableActions = door.GetAvailableActions(g, gridX, gridY)
						g.RadialMenuX = mx
						g.RadialMenuY = my
						g.RadialMenuOpen = true
						hit = true
						break
					} else if g.DungeonMap.Grid[gridY][gridX] == Wall {
						break
					}
				}
			}
			if !hit {
				g.RadialMenuOpen = false
			}
		}
	}

	// Basic Collision
	gridX, gridY := int(math.Floor(newX)), int(math.Floor(newY))
	if gridX >= 0 && gridX < Width && gridY >= 0 && gridY < Height {
		if g.DungeonMap.Grid[gridY][gridX] == Floor {
			g.PlayerX = newX
			g.PlayerY = newY
		}
	}

	if g.PlayerA < 0 {
		g.PlayerA += 2 * math.Pi
	} else if g.PlayerA > 2*math.Pi {
		g.PlayerA -= 2 * math.Pi
	}

	g.Camera.SetPosition(&geom.Vector2{X: g.PlayerX, Y: g.PlayerY})
	g.Camera.SetHeadingAngle(g.PlayerA)

	// --- SPRITE GENERATOR ---
	var sprites []raycaster.Sprite
	for pos, door := range g.DungeonMap.Doors {
		if door.IsOpen {
			sprites = append(sprites, &GameSprite{
				X:           float64(pos.X) + 0.5,
				Y:           float64(pos.Y) + 0.5,
				Z:           0,
				ScaleFactor: 1.0,
				Img:         getDoorOpenTex(),
			})
		}
	}

	g.Camera.Update(sprites)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})

	viewport := screen.SubImage(image.Rect(0, 0, 700, 768)).(*ebiten.Image)
	viewport.Fill(color.RGBA{30, 30, 50, 255})
	g.Camera.Draw(viewport)

	topBox := screen.SubImage(image.Rect(710, 10, 1014, 370)).(*ebiten.Image)
	topBox.Fill(color.RGBA{40, 40, 60, 255})

	bottomBox := screen.SubImage(image.Rect(710, 390, 1014, 758)).(*ebiten.Image)
	bottomBox.Fill(color.RGBA{40, 40, 60, 255})

	currentRoomDesc := "Hallway / Uncharted Corridor"
	for _, r := range g.DungeonMap.Rooms {
		pGridX, pGridY := int(math.Floor(g.PlayerX)), int(math.Floor(g.PlayerY))
		if pGridX >= r.X && pGridX < r.X+r.W && pGridY >= r.Y && pGridY < r.Y+r.H {
			currentRoomDesc = fmt.Sprintf("Room ID %d: %s", r.ID, r.RType)
			if r.IsFocal {
				currentRoomDesc += " [FOCAL HUB]"
			}
			break
		}
	}

	roomText := fmt.Sprintf("--- LOCATION INFO ---\n\nPosition:\n  X: %.2f\n  Y: %.2f\n\nCurrent Area:\n  %s",
		g.PlayerX, g.PlayerY, currentRoomDesc)
	ebitenutil.DebugPrintAt(screen, roomText, 725, 25)

	// --- UPDATED ACTION LOG WITH WORD WRAP ---
	ebitenutil.DebugPrintAt(screen, "--- ACTION LOG ---", 725, 405)

	yOffset := 435
	for _, logEntry := range g.ActionLog {
		// Wrap the text to ~45 characters to fit the right-hand panel
		wrappedLines := wrapText(logEntry, 45)

		for _, line := range wrappedLines {
			if yOffset < 740 {
				ebitenutil.DebugPrintAt(screen, line, 725, yOffset)
				yOffset += 16 // Height of each line of text
			}
		}
		yOffset += 6 // Extra visual padding between separate log entries

		if yOffset >= 740 {
			break // Stop rendering if we hit the bottom of the box
		}
	}

	if g.RadialMenuOpen {
		ebitenutil.DrawRect(screen, float64(g.RadialMenuX)-3, float64(g.RadialMenuY)-3, 6, 6, color.RGBA{255, 0, 0, 255})
		angleStep := 2 * math.Pi / float64(len(g.AvailableActions))
		radius := 60.0

		for i, action := range g.AvailableActions {
			ax := float64(g.RadialMenuX) + math.Cos(float64(i)*angleStep)*radius
			ay := float64(g.RadialMenuY) + math.Sin(float64(i)*angleStep)*radius

			ebitenutil.DrawRect(screen, ax-20, ay-20, 40, 40, color.RGBA{50, 70, 150, 220})
			textX := int(ax) - (len(action.Name) * 3)
			textY := int(ay) - 6
			ebitenutil.DebugPrintAt(screen, action.Name, textX, textY)
		}
	}

	// --- CHARACTER CREATION OVERLAY ---
	if g.CharCreationOpen {
		ebitenutil.DrawRect(screen, 112, 84, 800, 600, color.RGBA{15, 15, 25, 245})
		ebitenutil.DrawRect(screen, 112, 84, 800, 4, color.RGBA{200, 150, 50, 255})
		ebitenutil.DrawRect(screen, 112, 680, 800, 4, color.RGBA{200, 150, 50, 255})

		ebitenutil.DebugPrintAt(screen, "=== CHARACTER CREATION WIZARD ===", 400, 120)

		if g.CharCreationStep == 1 {
			ebitenutil.DebugPrintAt(screen, "STEP 1: Ability Scores (25-Point Buy)", 400, 160)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Points Remaining: %d", g.PointsRemaining), 400, 190)

			for i, stat := range g.StatOrder {
				val := g.PointBuyStats[stat]
				mod := int(math.Floor(float64(val-10) / 2.0))
				modStr := fmt.Sprintf("+%d", mod)
				if mod < 0 {
					modStr = fmt.Sprintf("%d", mod)
				}

				prefix := "  "
				if i == g.SelectedStatIndex {
					prefix = ">>"
				}

				text := fmt.Sprintf("%s %s : %2d  (Mod: %s)", prefix, stat, val, modStr)
				ebitenutil.DebugPrintAt(screen, text, 400, 240+(i*30))
			}

			ebitenutil.DebugPrintAt(screen, "Controls:", 400, 450)
			ebitenutil.DebugPrintAt(screen, "[UP/DOWN] Select Ability", 400, 470)
			ebitenutil.DebugPrintAt(screen, "[LEFT/RIGHT] Decrease/Increase Score", 400, 490)
			ebitenutil.DebugPrintAt(screen, "[ENTER] Confirm and Proceed to Step 2", 400, 520)

		} else if g.CharCreationStep == 2 {
			ebitenutil.DebugPrintAt(screen, "STEP 2: Race Selection", 400, 160)

			raceID := g.StandardRaces[g.SelectedRaceIndex]
			race := GlobalRaces[raceID]

			header := fmt.Sprintf("<<  %s  >>", race.Name)
			ebitenutil.DebugPrintAt(screen, header, 400, 200)

			details := fmt.Sprintf("Size: %s | Speed: %d ft. | Vision: %s", race.Size, race.BaseSpeed, race.Vision)
			ebitenutil.DebugPrintAt(screen, details, 400, 230)

			modText := "Ability Mods: "
			if len(race.AbilityMods) == 0 {
				modText += "None"
			} else {
				for k, v := range race.AbilityMods {
					sign := "+"
					if v < 0 {
						sign = ""
					}
					modText += fmt.Sprintf("%s%d %s  ", sign, v, k)
				}
			}
			ebitenutil.DebugPrintAt(screen, modText, 400, 260)

			rawText := RaceFlavorText[raceID]
			wrappedLines := wrapText(rawText, 70)

			maxVisibleLines := 10
			if g.RaceTextOffset > len(wrappedLines)-maxVisibleLines {
				g.RaceTextOffset = len(wrappedLines) - maxVisibleLines
			}
			if g.RaceTextOffset < 0 {
				g.RaceTextOffset = 0
			}

			for i := 0; i < maxVisibleLines; i++ {
				lineIndex := g.RaceTextOffset + i
				if lineIndex < len(wrappedLines) {
					ebitenutil.DebugPrintAt(screen, wrappedLines[lineIndex], 200, 310+(i*20))
				}
			}

			ebitenutil.DebugPrintAt(screen, "Controls:", 400, 540)
			ebitenutil.DebugPrintAt(screen, "[LEFT/RIGHT] Change Race   [UP/DOWN] Scroll Text", 400, 560)
			ebitenutil.DebugPrintAt(screen, "[BACKSPACE] Return to Step 1   [ENTER] Proceed to Step 3", 400, 580)

		} else if g.CharCreationStep == 3 {
			ebitenutil.DebugPrintAt(screen, "STEP 3: Class Selection", 400, 160)

			classID := g.StandardClasses[g.SelectedClassIndex]
			cls := GlobalClasses[classID]

			header := fmt.Sprintf("<<  %s  >>", cls.Name)
			ebitenutil.DebugPrintAt(screen, header, 400, 200)

			details := fmt.Sprintf("Hit Die: d%d | BAB: %s | Skill Pts/Lvl: %d", cls.HitDie, cls.BAB, cls.SkillPoints)
			ebitenutil.DebugPrintAt(screen, details, 400, 230)

			savesText := fmt.Sprintf("Saves -> Fort: %s | Ref: %s | Will: %s", cls.FortSave, cls.RefSave, cls.WillSave)
			ebitenutil.DebugPrintAt(screen, savesText, 400, 260)

			rawText := ClassFlavorText[classID]
			wrappedLines := wrapText(rawText, 70)

			maxVisibleLines := 10
			if g.ClassTextOffset > len(wrappedLines)-maxVisibleLines {
				g.ClassTextOffset = len(wrappedLines) - maxVisibleLines
			}
			if g.ClassTextOffset < 0 {
				g.ClassTextOffset = 0
			}

			for i := 0; i < maxVisibleLines; i++ {
				lineIndex := g.ClassTextOffset + i
				if lineIndex < len(wrappedLines) {
					ebitenutil.DebugPrintAt(screen, wrappedLines[lineIndex], 200, 310+(i*20))
				}
			}

			ebitenutil.DebugPrintAt(screen, "Controls:", 400, 540)
			ebitenutil.DebugPrintAt(screen, "[LEFT/RIGHT] Change Class   [UP/DOWN] Scroll Text", 400, 560)
			ebitenutil.DebugPrintAt(screen, "[BACKSPACE] Return to Step 2   [ENTER] Proceed to Skills", 400, 580)

		} else if g.CharCreationStep == 6 {
			ebitenutil.DebugPrintAt(screen, "STEP 6: Skill Allocation", 400, 150)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Skill Points Remaining: %d", g.SkillPointsRemaining), 400, 180)

			maxVisible := 14
			for i := 0; i < maxVisible; i++ {
				listIdx := g.SkillTextOffset + i
				if listIdx >= len(g.SkillOrder) {
					break
				}

				skillID := g.SkillOrder[listIdx]
				skillBlueprint := GlobalSkills[skillID]
				ranks := g.SkillRanksAllocated[skillID]

				prefix := "  "
				if listIdx == g.SelectedSkillIndex {
					prefix = ">>"
				}

				text := fmt.Sprintf("%s %-20s (%s): %d", prefix, skillBlueprint.Name, skillBlueprint.KeyAbility, ranks)
				ebitenutil.DebugPrintAt(screen, text, 350, 220+(i*22))
			}

			ebitenutil.DebugPrintAt(screen, "Controls:", 400, 540)
			ebitenutil.DebugPrintAt(screen, "[UP/DOWN] Select Skill   [LEFT/RIGHT] Decrease/Increase Ranks (Max 4)", 400, 560)
			ebitenutil.DebugPrintAt(screen, "[BACKSPACE] Return to Step 3   [ENTER] Finish and Return to Map", 400, 580)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 768
}
