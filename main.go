package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

func main() {
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
	gameMap.PlaceTreasure()
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

	// Serialize the map for the 2D HTML Tabletop
	// Note: We create an exportable payload since maps with struct keys (Node) cannot be marshaled directly to JSON
	exportPayload := struct {
		Width  int
		Height int
		Grid   [60][120]rune // Using standard dimensions
		Rooms  []Room
		Doors  map[string]*Door
	}{
		Width:  120, // From map.go
		Height: 60,  // From map.go
		Rooms:  gameMap.Rooms,
		Doors:  make(map[string]*Door),
	}

	// Convert [Height][Width]CellType to [Height][Width]rune
	for y := 0; y < 60; y++ {
		for x := 0; x < 120; x++ {
			exportPayload.Grid[y][x] = rune(gameMap.Grid[y][x])
		}
	}

	// Convert map[Node]*Door to map[string]*Door for JSON
	for node, door := range gameMap.Doors {
		key := fmt.Sprintf("%d,%d", node.X, node.Y)
		exportPayload.Doors[key] = door
	}

	mapData, err := json.MarshalIndent(exportPayload, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("dungeon_state.json", mapData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nDungeon state exported to dungeon_state.json")
}
