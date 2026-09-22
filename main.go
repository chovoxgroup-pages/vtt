package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

type Config struct {
	Size     string
	Function string
	Traps    string
	Doors    string
	Locked   string
	Secret   bool
	APL      int
}

type Game struct {
	DungeonMap          *Map
	ActionLog           []string
	StandardRaces       []string
	SelectedRaceIndex   int
	PointBuyStats       map[string]int
	SkillRanksAllocated map[string]int
	PlayerX             int
	PlayerY             int
	PlayerFacing        int // 0=North, 1=East, 2=South, 3=West
}

func (g *Game) BroadcastNoise(x, y int, volume float64) {}

type RadialAction struct {
	Name    string
	Execute func()
}

var activeGame *Game

type MapPayload struct {
	Width        int
	Height       int
	Grid         [60][120]rune
	Rooms        []Room
	Doors        map[string]*Door
	PlayerX      int
	PlayerY      int
	PlayerFacing int
	Log          []string
}

func main() {
	fmt.Println("System: Go WebAssembly Engine Initialized.")
	rand.Seed(time.Now().UnixNano())

	js.Global().Set("generateDungeon", js.FuncOf(generateDungeonWASM))
	js.Global().Set("interactWithTile", js.FuncOf(interactWithTileWASM))
	js.Global().Set("handleKeyboard", js.FuncOf(handleKeyboardWASM)) // New Listener

	<-make(chan bool)
}

func generateDungeonWASM(this js.Value, args []js.Value) interface{} {
	cfg := Config{
		Size:     "normal",
		Function: "random",
		Traps:    "some",
		Doors:    "some",
		Locked:   "some",
		Secret:   true,
		APL:      1,
	}

	gameMap := NewMap(cfg)
	gameMap.GenerateRooms()
	gameMap.PlaceDoors()
	gameMap.PlaceRoomTraps()
	gameMap.PlaceMonsters()
	gameMap.PlaceTreasure()

	var spawnX, spawnY int
	for _, r := range gameMap.Rooms {
		if r.IsFocal {
			cx, cy := r.Center()
			spawnX = cx * 2
			spawnY = cy * 2
			break
		}
	}

	activeGame = &Game{
		DungeonMap:          gameMap,
		ActionLog:           []string{"You materialize in the dungeon."},
		StandardRaces:       []string{"race_human"},
		PointBuyStats:       map[string]int{"STR": 14, "DEX": 15, "CON": 13, "INT": 12, "WIS": 10, "CHA": 8},
		SkillRanksAllocated: map[string]int{"skill_search": 4, "skill_open_lock": 4, "skill_disable_device": 4},
		PlayerX:             spawnX,
		PlayerY:             spawnY,
		PlayerFacing:        0, // Start facing North
	}

	return generateMapStateWASM()
}

func interactWithTileWASM(this js.Value, args []js.Value) interface{} {
	if len(args) < 5 || activeGame == nil {
		return `{"error": "Engine not ready"}`
	}

	goX := args[0].Int()
	goY := args[1].Int()
	actionName := args[2].String()
	vttX := args[3].Int()
	vttY := args[4].Int()

	activeGame.ActionLog = []string{}

	doorKey := Node{goX, goY}
	if door, exists := activeGame.DungeonMap.Doors[doorKey]; exists {
		actions := door.GetAvailableActions(activeGame, goX, goY)
		for _, a := range actions {
			if a.Name == actionName {
				a.Execute()
				break
			}
		}
	} else {
		if actionName == "Move Here" {
			activeGame.PlayerX = vttX
			activeGame.PlayerY = vttY
			activeGame.ActionLog = append(activeGame.ActionLog, fmt.Sprintf("You move to [%d, %d].", vttX, vttY))
		} else if actionName == "Search Tile" || actionName == "Search Wall" {
			activeGame.ActionLog = append(activeGame.ActionLog, fmt.Sprintf("You thoroughly search the area at [%d, %d].", vttX, vttY))
		}
	}

	return generateMapStateWASM()
}

func handleKeyboardWASM(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 || activeGame == nil {
		return `{"error": "Engine not ready"}`
	}
	key := args[0].String()

	activeGame.ActionLog = []string{}
	newX := activeGame.PlayerX
	newY := activeGame.PlayerY

	// 4-Way Direction Vectors: N (0), E (1), S (2), W (3)
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	switch key {
	case "a": // Rotate Left (Counter-Clockwise)
		activeGame.PlayerFacing = (activeGame.PlayerFacing + 3) % 4
		return generateMapStateWASM()
	case "d": // Rotate Right (Clockwise)
		activeGame.PlayerFacing = (activeGame.PlayerFacing + 1) % 4
		return generateMapStateWASM()
	case "w": // Forward
		newX += dx[activeGame.PlayerFacing]
		newY += dy[activeGame.PlayerFacing]
	case "s": // Back
		backDir := (activeGame.PlayerFacing + 2) % 4
		newX += dx[backDir]
		newY += dy[backDir]
	case "q": // Strafe Left
		leftDir := (activeGame.PlayerFacing + 3) % 4
		newX += dx[leftDir]
		newY += dy[leftDir]
	case "e": // Strafe Right
		rightDir := (activeGame.PlayerFacing + 1) % 4
		newX += dx[rightDir]
		newY += dy[rightDir]
	}

	// Collision Detection mapping micro to macro grid
	goX := newX / 2
	goY := newY / 2

	if goY >= 0 && goY < 60 && goX >= 0 && goX < 120 {
		tile := activeGame.DungeonMap.Grid[goY][goX]
		doorKey := Node{goX, goY}
		door, hasDoor := activeGame.DungeonMap.Doors[doorKey]

		canMove := true
		if tile == 35 { // Wall
			canMove = false
		} else if hasDoor && !door.IsOpen { // Closed Door
			canMove = false
		}

		if canMove {
			activeGame.PlayerX = newX
			activeGame.PlayerY = newY
		} else {
			activeGame.ActionLog = append(activeGame.ActionLog, "*Thud* The way is blocked.")
		}
	}

	return generateMapStateWASM()
}

func generateMapStateWASM() string {
	payload := MapPayload{
		Width:        120,
		Height:       60,
		Rooms:        activeGame.DungeonMap.Rooms,
		Doors:        make(map[string]*Door),
		PlayerX:      activeGame.PlayerX,
		PlayerY:      activeGame.PlayerY,
		PlayerFacing: activeGame.PlayerFacing,
		Log:          activeGame.ActionLog,
	}

	for y := 0; y < 60; y++ {
		for x := 0; x < 120; x++ {
			payload.Grid[y][x] = rune(activeGame.DungeonMap.Grid[y][x])
		}
	}
	for node, door := range activeGame.DungeonMap.Doors {
		key := fmt.Sprintf("%d,%d", node.X, node.Y)
		payload.Doors[key] = door
	}

	mapData, _ := json.Marshal(payload)
	return string(mapData)
}
