package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/harbdog/raycaster-go"
	"github.com/harbdog/raycaster-go/geom"
)

// RayMap wraps our dungeon grid to satisfy the raycaster Map and TextureHandler interfaces
type RayMap struct {
	Grid    *[Height][Width]CellType
	WallTex *ebiten.Image
	DoorTex *ebiten.Image
}

// HitCheck determines if a ray hits a wall at the given (x,y) coordinates
func (m *RayMap) HitCheck(x, y float64) bool {
	gridX, gridY := int(math.Floor(x)), int(math.Floor(y))

	if gridX < 0 || gridX >= Width || gridY < 0 || gridY >= Height {
		return true
	}

	return m.Grid[gridY][gridX] == Wall || m.Grid[gridY][gridX] == DoorC
}

// GetHitColor returns the color of the object the ray hit
func (m *RayMap) GetHitColor(x, y float64) color.RGBA {
	gridX, gridY := int(math.Floor(x)), int(math.Floor(y))

	if gridX >= 0 && gridX < Width && gridY >= 0 && gridY < Height {
		if m.Grid[gridY][gridX] == DoorC {
			return color.RGBA{139, 69, 19, 255} // Brown for doors
		}
	}

	return color.RGBA{100, 100, 100, 255} // Gray for standard walls
}

// Level returns a 2D array representation of the map grid for the raycaster
func (m *RayMap) Level(levelNum int) [][]int {
	if levelNum != 0 {
		return nil
	}

	levelArr := make([][]int, Width)
	for x := 0; x < Width; x++ {
		levelArr[x] = make([]int, Height)
		for y := 0; y < Height; y++ {
			if m.Grid[y][x] == Wall || m.Grid[y][x] == DoorC {
				levelArr[x][y] = 1
			} else {
				levelArr[x][y] = 0
			}
		}
	}
	return levelArr
}

// NumLevels returns the total number of vertical levels in the map
func (m *RayMap) NumLevels() int {
	return 1
}

// TextureAt returns the wall texture at the given map coordinate
func (m *RayMap) TextureAt(x, y, levelNum, side int) *ebiten.Image {
	if m.WallTex == nil {
		m.WallTex = ebiten.NewImage(64, 64)
		m.WallTex.Fill(color.RGBA{100, 100, 100, 255})

		m.DoorTex = ebiten.NewImage(64, 64)
		m.DoorTex.Fill(color.RGBA{139, 69, 19, 255})
	}

	if x >= 0 && x < Width && y >= 0 && y < Height {
		if m.Grid[y][x] == DoorC {
			return m.DoorTex
		}
	}
	return m.WallTex
}

// FloorTextureAt returns the floor texture (nil for default flat color)
func (m *RayMap) FloorTextureAt(x, y int) *image.RGBA {
	return nil
}

// Game implements the ebiten.Game interface
type Game struct {
	DungeonMap    *Map
	Camera        *raycaster.Camera
	PlayerX       float64
	PlayerY       float64
	PlayerA       float64
	ActionLog     []string
	CurrentRoomID int // Track current room to detect entries/exits
}

func (g *Game) Update() error {
	moveSpeed := 0.08
	rotSpeed := 0.05

	newX, newY := g.PlayerX, g.PlayerY

	// Forward and backward movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		newX += math.Cos(g.PlayerA) * moveSpeed
		newY += math.Sin(g.PlayerA) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		newX -= math.Cos(g.PlayerA) * moveSpeed
		newY -= math.Sin(g.PlayerA) * moveSpeed
	}

	// Strafe Left (Q) and Right (E)
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		newX += math.Cos(g.PlayerA+math.Pi/2) * moveSpeed
		newY += math.Sin(g.PlayerA+math.Pi/2) * moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		newX += math.Cos(g.PlayerA-math.Pi/2) * moveSpeed
		newY += math.Sin(g.PlayerA-math.Pi/2) * moveSpeed
	}

	// Rotation
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.PlayerA += rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.PlayerA -= rotSpeed
	}

	// Track Room Entry / Exit
	pGridX, pGridY := int(math.Floor(g.PlayerX)), int(math.Floor(g.PlayerY))
	detectedRoomID := 0 // 0 means hallway / corridor
	var detectedRoom *Room

	for i := range g.DungeonMap.Rooms {
		r := &g.DungeonMap.Rooms[i]
		if pGridX >= r.X && pGridX < r.X+r.W && pGridY >= r.Y && pGridY < r.Y+r.H {
			detectedRoomID = r.ID
			detectedRoom = r
			break
		}
	}

	// If our room ID has changed, log the transition
	if detectedRoomID != g.CurrentRoomID {
		var transitionMsg string
		if g.CurrentRoomID != 0 && detectedRoomID == 0 {
			transitionMsg = fmt.Sprintf("Exited Room ID %d", g.CurrentRoomID)
		} else if g.CurrentRoomID == 0 && detectedRoomID != 0 {
			transitionMsg = fmt.Sprintf("Entered Room ID %d (%s)", detectedRoomID, detectedRoom.RType)
		} else {
			transitionMsg = fmt.Sprintf("Entered Room ID %d (%s)", detectedRoomID, detectedRoom.RType)
		}

		g.ActionLog = append([]string{transitionMsg}, g.ActionLog...)
		if len(g.ActionLog) > 15 {
			g.ActionLog = g.ActionLog[:15]
		}
		g.CurrentRoomID = detectedRoomID
	}

	// Door Interaction (Left Mouse Click) - Opens or Closes within 1 unit
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		reach := 1.0
		targetX := int(math.Floor(g.PlayerX + math.Cos(g.PlayerA)*reach))
		targetY := int(math.Floor(g.PlayerY + math.Sin(g.PlayerA)*reach))

		if targetX >= 0 && targetX < Width && targetY >= 0 && targetY < Height {
			doorNode := Node{X: targetX, Y: targetY}

			if door, ok := g.DungeonMap.Doors[doorNode]; ok {
				var msg string
				if door.IsOpen {
					door.IsOpen = false
					g.DungeonMap.Grid[targetY][targetX] = DoorC
					msg = fmt.Sprintf("Closed door at (%d, %d)", targetX, targetY)
				} else {
					door.IsOpen = true
					g.DungeonMap.Grid[targetY][targetX] = Floor
					msg = fmt.Sprintf("Opened door at (%d, %d)", targetX, targetY)
				}

				// Prepend new message to the top of the log slice
				g.ActionLog = append([]string{msg}, g.ActionLog...)
				if len(g.ActionLog) > 15 {
					g.ActionLog = g.ActionLog[:15] // Keep last 15 lines max
				}
			}
		}
	}

	// Basic Collision Check against the floor grid before allowing movement
	gridX, gridY := int(math.Floor(newX)), int(math.Floor(newY))
	if gridX >= 0 && gridX < Width && gridY >= 0 && gridY < Height {
		if g.DungeonMap.Grid[gridY][gridX] == Floor {
			g.PlayerX = newX
			g.PlayerY = newY
		}
	}

	// Normalize angle
	if g.PlayerA < 0 {
		g.PlayerA += 2 * math.Pi
	} else if g.PlayerA > 2*math.Pi {
		g.PlayerA -= 2 * math.Pi
	}

	g.Camera.SetPosition(&geom.Vector2{X: g.PlayerX, Y: g.PlayerY})
	g.Camera.SetHeadingAngle(g.PlayerA)
	g.Camera.Update(nil)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 1. Fill background panel
	screen.Fill(color.RGBA{20, 20, 30, 255})

	// 2. Left 3D Viewport (Width: 700, Height: 768)
	viewport := screen.SubImage(image.Rect(0, 0, 700, 768)).(*ebiten.Image)
	viewport.Fill(color.RGBA{30, 30, 50, 255})
	g.Camera.Draw(viewport)

	// 3. Right UI Column Backgrounds
	// Top Right Box Background
	topBox := screen.SubImage(image.Rect(710, 10, 1014, 370)).(*ebiten.Image)
	topBox.Fill(color.RGBA{40, 40, 60, 255})

	// Bottom Right Box Background
	bottomBox := screen.SubImage(image.Rect(710, 390, 1014, 758)).(*ebiten.Image)
	bottomBox.Fill(color.RGBA{40, 40, 60, 255})

	// Find which room the player is currently inside
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

	// Draw Top Right Text directly onto the main screen
	roomText := fmt.Sprintf("--- LOCATION INFO ---\n\nPosition:\n  X: %.2f\n  Y: %.2f\n\nCurrent Area:\n  %s",
		g.PlayerX, g.PlayerY, currentRoomDesc)
	ebitenutil.DebugPrintAt(screen, roomText, 725, 25)

	// Draw Bottom Right Text (Action Log) descending from top to bottom
	ebitenutil.DebugPrintAt(screen, "--- ACTION LOG ---", 725, 405)
	for i, logEntry := range g.ActionLog {
		// Stagger each line down by 18 pixels
		yOffset := 435 + (i * 18)
		if yOffset < 740 { // Prevent overflow past bottom box
			ebitenutil.DebugPrintAt(screen, logEntry, 725, yOffset)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 768
}
