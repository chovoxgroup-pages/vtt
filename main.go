package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/harbdog/raycaster-go"
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

func testFirebaseConnection() {
	// Reach out of WebAssembly and grab the browser's global window
	window := js.Global()

	// Try to grab the database we set up in index.html
	db := window.Get("firebaseDB")

	if db.IsUndefined() {
		fmt.Println("Uh oh: Go cannot see the Firebase database.")
	} else {
		fmt.Println("Success: Go has made contact with Firebase!")
	}
}

func main() {
	testFirebaseConnection()
	sizeFlag := flag.String("size", "random", "Dungeon size: few, normal, many, random")
	funcFlag := flag.String("function", "random", "Dungeon function: fortification, worship, restraint, storage, shelter, concealment, random")
	trapFlag := flag.String("traps", "some", "Traps frequency: none, some, many")
	doorFlag := flag.String("doors", "some", "Doors frequency: none, some, many")
	lockFlag := flag.String("locked", "some", "Locked doors frequency: none, some, many")
	secretFlag := flag.Bool("secret", true, "Include secret doors: true, false")
	aplFlag := flag.Int("apl", 1, "Average Party Level: 1-20")

	flag.Parse()

	cfg := Config{
		Size:     *sizeFlag,
		Function: *funcFlag,
		Traps:    *trapFlag,
		Doors:    *doorFlag,
		Locked:   *lockFlag,
		Secret:   *secretFlag,
		APL:      *aplFlag,
	}

	rand.Seed(time.Now().UnixNano())
	gameMap := NewMap(cfg)
	gameMap.GenerateRooms()
	gameMap.PlaceDoors()
	gameMap.PlaceRoomTraps()
	gameMap.PlaceMonsters()
	gameMap.PlaceTreasure() // <-- Added Treasure Hook
	gameMap.Print()

	fmt.Println("\n--- Room Semantic Data ---")
	for _, r := range gameMap.Rooms {
		focalStr := ""
		if r.IsFocal {
			focalStr = " [FOCAL HUB]"
		}

		trapStr := ""
		if len(r.Traps) > 0 {
			trapStr = fmt.Sprintf(" | TRAP: %s (Type: %s, CR: %d)", r.Traps[0].Name, r.Traps[0].Category, r.Traps[0].CR)
		}

		monsterStr := ""
		if len(r.Monsters) > 0 {
			monsterStr = fmt.Sprintf(" | MONSTERS: %d %s (Type: %s, CR: %d)", r.Monsters[0].Count, r.Monsters[0].Name, r.Monsters[0].Type, r.Monsters[0].CR)
		}

		treasureStr := ""
		if r.Treasure.Coins != "" && r.Treasure.Coins != "None" {
			g := "None"
			i := "None"
			if len(r.Treasure.Goods) > 0 {
				g = strings.Join(r.Treasure.Goods, ", ")
			}
			if len(r.Treasure.Items) > 0 {
				i = strings.Join(r.Treasure.Items, ", ")
			}
			treasureStr = fmt.Sprintf(" | LOOT: [%s] Goods: %s | Items: %s", r.Treasure.Coins, g, i)
		}

		fmt.Printf("Room ID %d: %s%s at (%d, %d)%s%s%s\n", r.ID, r.RType, focalStr, r.X, r.Y, trapStr, monsterStr, treasureStr)
	}

	fmt.Println("\n--- Sample Door Generation Data ---")
	count := 0
	for pos, door := range gameMap.Doors {
		if count >= 5 {
			break
		}

		trapStr := "IsTrapped:false"
		if door.IsTrapped && door.Trap != nil {
			trapStr = fmt.Sprintf("IsTrapped:true | TRAP: %s (Type: %s, CR: %d)", door.Trap.Name, door.Trap.Category, door.Trap.CR)
		}

		fmt.Printf("Door at (%d, %d): &{IsOpen:%v Stuck:%v Lock:%v %s Secret:%v}\n",
			pos.X, pos.Y, door.IsOpen, door.Stuck, door.Lock, trapStr, door.Secret)
		count++
	}
	// --- Launch Ebitengine ---
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Brandon's VTT - Raycaster")

	// Find the focal room to spawn the player
	var startX, startY float64
	for _, r := range gameMap.Rooms {
		if r.IsFocal {
			cx, cy := r.Center()
			startX = float64(cx) + 0.5 // Add 0.5 to perfectly center them in the tile
			startY = float64(cy) + 0.5
			break
		}
	}

	// Wrap our dungeon grid for the Raycaster Camera
	rayMap := &RayMap{Grid: &gameMap.Grid}
	camera := raycaster.NewCamera(640, 480, 64, rayMap, rayMap)

	// Pass the generated map and camera into the engine
	game := &Game{
		DungeonMap: gameMap,
		Camera:     camera,
		PlayerX:    startX,
		PlayerY:    startY,
		PlayerA:    0,
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
