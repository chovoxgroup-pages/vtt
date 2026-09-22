package main

import (
	"fmt"
	"math/rand"
)

const (
	Width  = 120
	Height = 60
)

type CellType rune

const (
	Wall  CellType = '#'
	Floor CellType = '.'
	DoorC CellType = '+'
	DoorO CellType = 'O'
)

type Node struct {
	X, Y int
}

type Room struct {
	ID         int
	X, Y, W, H int
	RType      RoomType
	IsFocal    bool
	Traps      []Trap
	Monsters   []Monster
	Treasure   TreasureData
}

func (r Room) Center() (int, int) {
	return r.X + r.W/2, r.Y + r.H/2
}

type Map struct {
	Grid   [Height][Width]CellType
	Rooms  []Room
	Doors  map[Node]*Door
	Theme  DungeonTheme
	Config Config
}

func NewMap(cfg Config) *Map {
	m := &Map{
		Doors:  make(map[Node]*Door),
		Config: cfg,
		Theme: DungeonTheme{
			Size:        determineSize(cfg.Size),
			Environment: rollEnvironment(),
			Function:    determineFunction(cfg.Function),
		},
	}
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			m.Grid[y][x] = Wall
		}
	}
	return m
}

func (m *Map) GenerateRooms() {
	focalType, supportTypes := getSemanticRooms(m.Theme.Function)

	var targetTotalRooms int
	var baseRoomSize, varRoomSize int

	switch m.Theme.Size {
	case "Few":
		targetTotalRooms = rand.Intn(4) + 3
		if rand.Intn(2) == 0 {
			baseRoomSize, varRoomSize = 12, 10
		} else {
			baseRoomSize, varRoomSize = 8, 6
		}
	case "Normal":
		targetTotalRooms = rand.Intn(6) + 7
		baseRoomSize, varRoomSize = 8, 6
	case "Many":
		targetTotalRooms = rand.Intn(8) + 13
		baseRoomSize, varRoomSize = 5, 5
	}

	fw := rand.Intn(varRoomSize+2) + baseRoomSize + 2
	fh := rand.Intn(varRoomSize+2) + baseRoomSize + 2
	fx := (Width / 2) - (fw / 2)
	fy := (Height / 2) - (fh / 2)

	focalRoom := Room{ID: 1, X: fx, Y: fy, W: fw, H: fh, RType: focalType, IsFocal: true}
	m.carveRoom(focalRoom)
	m.Rooms = append(m.Rooms, focalRoom)

	numSupportRooms := targetTotalRooms - 1

	for i := 0; i < numSupportRooms; i++ {
		rw := rand.Intn(varRoomSize) + baseRoomSize
		rh := rand.Intn(varRoomSize) + baseRoomSize

		parent := m.Rooms[rand.Intn(len(m.Rooms))]
		px, py := parent.Center()

		distance := rand.Intn(15) + 10
		angle := rand.Float64() * 2 * 3.14159

		dirX, dirY := 0, 0
		if angle < 1.57 {
			dirX, dirY = 1, 1
		} else if angle < 3.14 {
			dirX, dirY = -1, 1
		} else if angle < 4.71 {
			dirX, dirY = -1, -1
		} else {
			dirX, dirY = 1, -1
		}

		rx := px + (dirX * distance) - (rw / 2)
		ry := py + (dirY * distance) - (rh / 2)

		if rx < 2 || ry < 2 || rx+rw > Width-2 || ry+rh > Height-2 {
			i--
			continue
		}

		sType := supportTypes[rand.Intn(len(supportTypes))]
		newRoom := Room{ID: len(m.Rooms) + 1, X: rx, Y: ry, W: rw, H: rh, RType: sType, IsFocal: false}

		if !m.overlaps(newRoom) {
			m.carveRoom(newRoom)
			m.Rooms = append(m.Rooms, newRoom)

			cx, cy := newRoom.Center()
			m.carvePathAStar(px, py, cx, cy)
		} else {
			i--
		}
	}
}

func (m *Map) carveRoom(r Room) {
	for y := r.Y; y < r.Y+r.H; y++ {
		for x := r.X; x < r.X+r.W; x++ {
			m.Grid[y][x] = Floor
		}
	}
}

func (m *Map) overlaps(newRoom Room) bool {
	for _, r := range m.Rooms {
		if newRoom.X-2 < r.X+r.W && newRoom.X+newRoom.W+2 > r.X &&
			newRoom.Y-2 < r.Y+r.H && newRoom.Y+newRoom.H+2 > r.Y {
			return true
		}
	}
	return false
}

func (m *Map) isInRoom(x, y int) bool {
	for _, r := range m.Rooms {
		if x >= r.X && x < r.X+r.W && y >= r.Y && y < r.Y+r.H {
			return true
		}
	}
	return false
}

func (m *Map) Print() {
	fmt.Printf("DUNGEON THEME -> Size: %s | Environment: %s | Function: %s | APL: %d\n", m.Theme.Size, m.Theme.Environment, m.Theme.Function, m.Config.APL)
	fmt.Println("--------------------------------------------------")
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			fmt.Printf("%c", m.Grid[y][x])
		}
		fmt.Println()
	}
}
