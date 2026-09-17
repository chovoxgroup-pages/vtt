package main

import (
	"math/rand"
	"strings"
)

type DungeonSize string
type DungeonEnvironment string
type DungeonFunction string
type RoomType string

type DungeonTheme struct {
	Size        DungeonSize
	Environment DungeonEnvironment
	Function    DungeonFunction
}

func determineSize(cfgSize string) DungeonSize {
	s := strings.ToLower(cfgSize)
	if s == "few" {
		return "Few"
	}
	if s == "normal" {
		return "Normal"
	}
	if s == "many" {
		return "Many"
	}

	roll := rand.Intn(100) + 1
	switch {
	case roll <= 33:
		return "Few"
	case roll <= 66:
		return "Normal"
	default:
		return "Many"
	}
}

func determineFunction(cfgFunc string) DungeonFunction {
	f := strings.ToLower(cfgFunc)
	if f == "concealment" {
		return "Concealment"
	}
	if f == "fortification" {
		return "Fortification"
	}
	if f == "restraint" {
		return "Restraint"
	}
	if f == "shelter" {
		return "Shelter"
	}
	if f == "storage" {
		return "Storage"
	}
	if f == "worship" {
		return "Worship"
	}

	roll := rand.Intn(100) + 1
	switch {
	case roll <= 15:
		return "Concealment"
	case roll <= 30:
		return "Fortification"
	case roll <= 45:
		return "Restraint"
	case roll <= 70:
		return "Shelter"
	case roll <= 88:
		return "Storage"
	default:
		return "Worship"
	}
}

func rollEnvironment() DungeonEnvironment {
	roll := rand.Intn(100) + 1
	switch {
	case roll <= 15:
		return "Alien"
	case roll <= 33:
		return "Artificial"
	case roll <= 60:
		return "Elemental"
	case roll <= 75:
		return "Light/Dark"
	case roll <= 85:
		return "Mundane"
	default:
		return "Planar"
	}
}

func getSemanticRooms(function DungeonFunction) (RoomType, []RoomType) {
	switch function {
	case "Fortification":
		return "Guard Room", []RoomType{"Armory", "Barracks", "Kitchen"}
	case "Worship":
		return "Temple", []RoomType{"Crypt", "Reliquary", "Vestry", "Library"}
	case "Restraint":
		return "Prison Block", []RoomType{"Guard Room", "Torture Chamber", "Barracks"}
	case "Storage":
		return "Main Vault", []RoomType{"Guard Room", "Sorting Room", "Barracks"}
	case "Shelter":
		return "Common Area", []RoomType{"Living Quarters", "Kitchen", "Storage"}
	default:
		return "Hidden Chamber", []RoomType{"Maze", "Guard Room", "Crypt"}
	}
}
