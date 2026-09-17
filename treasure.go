package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type TreasureData struct {
	Coins string
	Goods []string
	Items []string
}

// Roll helper: NdS
func roll(n, s int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += rand.Intn(s) + 1
	}
	return total
}

func GenerateTreasure(apl int, isLootRoom bool) TreasureData {
	if apl < 1 {
		apl = 1
	} else if apl > 20 {
		apl = 20
	}

	rollCoins := rand.Intn(100) + 1
	rollGoods := rand.Intn(100) + 1
	rollItems := rand.Intn(100) + 1

	coins := getCoins(apl, rollCoins)
	goods := getGoods(apl, rollGoods)
	items := getItems(apl, rollItems)

	// Apply 1/4 fraction if this is an undefended loot room
	if isLootRoom {
		coins = fmt.Sprintf("%s (Valued at 25%%)", coins)
		if rand.Intn(100) >= 25 {
			goods = []string{}
		}
		if rand.Intn(100) >= 25 {
			items = []string{}
		}
	}

	return TreasureData{
		Coins: coins,
		Goods: goods,
		Items: items,
	}
}

func getCoins(level, d int) string {
	switch level {
	case 1:
		if d <= 14 {
			return "None"
		}
		if d <= 29 {
			return fmt.Sprintf("%d cp", roll(1, 6)*1000)
		}
		if d <= 52 {
			return fmt.Sprintf("%d sp", roll(1, 8)*100)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(2, 8)*10)
		}
		return fmt.Sprintf("%d pp", roll(1, 4)*10)
	case 2:
		if d <= 13 {
			return "None"
		}
		if d <= 23 {
			return fmt.Sprintf("%d cp", roll(1, 10)*1000)
		}
		if d <= 43 {
			return fmt.Sprintf("%d sp", roll(2, 10)*100)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(4, 10)*10)
		}
		return fmt.Sprintf("%d pp", roll(2, 8)*10)
	case 3:
		if d <= 11 {
			return "None"
		}
		if d <= 21 {
			return fmt.Sprintf("%d cp", roll(2, 10)*1000)
		}
		if d <= 41 {
			return fmt.Sprintf("%d sp", roll(4, 8)*100)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(1, 4)*100)
		}
		return fmt.Sprintf("%d pp", roll(1, 10)*10)
	case 4:
		if d <= 11 {
			return "None"
		}
		if d <= 21 {
			return fmt.Sprintf("%d cp", roll(3, 10)*1000)
		}
		if d <= 41 {
			return fmt.Sprintf("%d sp", roll(4, 12)*1000)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(1, 6)*100)
		}
		return fmt.Sprintf("%d pp", roll(1, 8)*10)
	case 5:
		if d <= 10 {
			return "None"
		}
		if d <= 19 {
			return fmt.Sprintf("%d cp", roll(1, 4)*10000)
		}
		if d <= 38 {
			return fmt.Sprintf("%d sp", roll(1, 6)*1000)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(1, 8)*100)
		}
		return fmt.Sprintf("%d pp", roll(1, 10)*10)
	case 6:
		if d <= 10 {
			return "None"
		}
		if d <= 18 {
			return fmt.Sprintf("%d cp", roll(1, 6)*10000)
		}
		if d <= 37 {
			return fmt.Sprintf("%d sp", roll(1, 8)*1000)
		}
		if d <= 95 {
			return fmt.Sprintf("%d gp", roll(1, 10)*100)
		}
		return fmt.Sprintf("%d pp", roll(1, 12)*10)
	case 7:
		if d <= 11 {
			return "None"
		}
		if d <= 18 {
			return fmt.Sprintf("%d cp", roll(1, 10)*10000)
		}
		if d <= 35 {
			return fmt.Sprintf("%d sp", roll(1, 12)*1000)
		}
		if d <= 93 {
			return fmt.Sprintf("%d gp", roll(2, 6)*100)
		}
		return fmt.Sprintf("%d pp", roll(3, 4)*10)
	case 8:
		if d <= 10 {
			return "None"
		}
		if d <= 15 {
			return fmt.Sprintf("%d cp", roll(1, 12)*10000)
		}
		if d <= 29 {
			return fmt.Sprintf("%d sp", roll(2, 6)*1000)
		}
		if d <= 87 {
			return fmt.Sprintf("%d gp", roll(2, 8)*100)
		}
		return fmt.Sprintf("%d pp", roll(3, 6)*10)
	case 9:
		if d <= 10 {
			return "None"
		}
		if d <= 15 {
			return fmt.Sprintf("%d cp", roll(2, 6)*10000)
		}
		if d <= 29 {
			return fmt.Sprintf("%d sp", roll(2, 8)*1000)
		}
		if d <= 85 {
			return fmt.Sprintf("%d gp", roll(5, 4)*100)
		}
		return fmt.Sprintf("%d pp", roll(2, 12)*10)
	case 10:
		if d <= 10 {
			return "None"
		}
		if d <= 24 {
			return fmt.Sprintf("%d sp", roll(2, 10)*1000)
		}
		if d <= 79 {
			return fmt.Sprintf("%d gp", roll(6, 4)*100)
		}
		return fmt.Sprintf("%d pp", roll(5, 6)*10)
	case 11:
		if d <= 8 {
			return "None"
		}
		if d <= 14 {
			return fmt.Sprintf("%d sp", roll(3, 10)*1000)
		}
		if d <= 75 {
			return fmt.Sprintf("%d gp", roll(4, 8)*100)
		}
		return fmt.Sprintf("%d pp", roll(4, 10)*10)
	case 12:
		if d <= 8 {
			return "None"
		}
		if d <= 14 {
			return fmt.Sprintf("%d sp", roll(3, 12)*1000)
		}
		if d <= 75 {
			return fmt.Sprintf("%d gp", roll(1, 4)*1000)
		}
		return fmt.Sprintf("%d pp", roll(1, 4)*100)
	case 13:
		if d <= 8 {
			return "None"
		}
		if d <= 75 {
			return fmt.Sprintf("%d gp", roll(1, 4)*1000)
		}
		return fmt.Sprintf("%d pp", roll(1, 10)*100)
	case 14:
		if d <= 8 {
			return "None"
		}
		if d <= 75 {
			return fmt.Sprintf("%d gp", roll(1, 6)*1000)
		}
		return fmt.Sprintf("%d pp", roll(1, 12)*100)
	case 15:
		if d <= 3 {
			return "None"
		}
		if d <= 74 {
			return fmt.Sprintf("%d gp", roll(1, 8)*1000)
		}
		return fmt.Sprintf("%d pp", roll(3, 4)*100)
	case 16:
		if d <= 3 {
			return "None"
		}
		if d <= 74 {
			return fmt.Sprintf("%d gp", roll(1, 12)*1000)
		}
		return fmt.Sprintf("%d pp", roll(3, 4)*100)
	case 17:
		if d <= 3 {
			return "None"
		}
		if d <= 68 {
			return fmt.Sprintf("%d gp", roll(3, 4)*1000)
		}
		return fmt.Sprintf("%d pp", roll(2, 10)*100)
	case 18:
		if d <= 2 {
			return "None"
		}
		if d <= 65 {
			return fmt.Sprintf("%d gp", roll(3, 6)*1000)
		}
		return fmt.Sprintf("%d pp", roll(5, 4)*100)
	case 19:
		if d <= 2 {
			return "None"
		}
		if d <= 65 {
			return fmt.Sprintf("%d gp", roll(3, 8)*1000)
		}
		return fmt.Sprintf("%d pp", roll(3, 10)*100)
	case 20:
		if d <= 2 {
			return "None"
		}
		if d <= 65 {
			return fmt.Sprintf("%d gp", roll(4, 8)*1000)
		}
		return fmt.Sprintf("%d pp", roll(4, 10)*100)
	}
	return "None"
}

func getGoods(level, d int) []string {
	var g []string
	gems := 0
	art := 0

	switch level {
	case 1:
		if d > 90 && d <= 95 {
			gems = 1
		} else if d > 95 {
			art = 1
		}
	case 2:
		if d > 81 && d <= 95 {
			gems = roll(1, 3)
		} else if d > 95 {
			art = roll(1, 3)
		}
	case 3:
		if d > 77 && d <= 95 {
			gems = roll(1, 3)
		} else if d > 95 {
			art = roll(1, 3)
		}
	case 4:
		if d > 70 && d <= 95 {
			gems = roll(1, 4)
		} else if d > 95 {
			art = roll(1, 3)
		}
	case 5:
		if d > 60 && d <= 95 {
			gems = roll(1, 4)
		} else if d > 95 {
			art = roll(1, 4)
		}
	case 6:
		if d > 56 && d <= 92 {
			gems = roll(1, 4)
		} else if d > 92 {
			art = roll(1, 4)
		}
	case 7:
		if d > 48 && d <= 88 {
			gems = roll(1, 4)
		} else if d > 88 {
			art = roll(1, 4)
		}
	case 8:
		if d > 45 && d <= 85 {
			gems = roll(1, 6)
		} else if d > 85 {
			art = roll(1, 4)
		}
	case 9:
		if d > 40 && d <= 80 {
			gems = roll(1, 8)
		} else if d > 80 {
			art = roll(1, 4)
		}
	case 10:
		if d > 35 && d <= 79 {
			gems = roll(1, 8)
		} else if d > 79 {
			art = roll(1, 6)
		}
	case 11:
		if d > 24 && d <= 74 {
			gems = roll(1, 10)
		} else if d > 74 {
			art = roll(1, 6)
		}
	case 12:
		if d > 17 && d <= 70 {
			gems = roll(1, 10)
		} else if d > 70 {
			art = roll(1, 8)
		}
	case 13:
		if d > 11 && d <= 66 {
			gems = roll(1, 12)
		} else if d > 66 {
			art = roll(1, 10)
		}
	case 14:
		if d > 11 && d <= 66 {
			gems = roll(2, 8)
		} else if d > 66 {
			art = roll(2, 6)
		}
	case 15:
		if d > 9 && d <= 65 {
			gems = roll(2, 10)
		} else if d > 65 {
			art = roll(2, 8)
		}
	case 16:
		if d > 7 && d <= 64 {
			gems = roll(4, 6)
		} else if d > 64 {
			art = roll(2, 10)
		}
	case 17:
		if d > 4 && d <= 63 {
			gems = roll(4, 8)
		} else if d > 63 {
			art = roll(3, 8)
		}
	case 18:
		if d > 4 && d <= 54 {
			gems = roll(3, 12)
		} else if d > 54 {
			art = roll(3, 10)
		}
	case 19:
		if d > 3 && d <= 50 {
			gems = roll(6, 6)
		} else if d > 50 {
			art = roll(6, 6)
		}
	case 20:
		if d > 2 && d <= 38 {
			gems = roll(4, 10)
		} else if d > 38 {
			art = roll(7, 6)
		}
	}

	for i := 0; i < gems; i++ {
		g = append(g, getGem())
	}
	for i := 0; i < art; i++ {
		g = append(g, getArt())
	}
	return g
}

func getItems(level, d int) []string {
	var it []string
	mnd := 0
	min := 0
	med := 0
	maj := 0

	switch level {
	case 1:
		if d > 71 && d <= 95 {
			mnd = 1
		} else if d > 95 {
			min = 1
		}
	case 2:
		if d > 49 && d <= 85 {
			mnd = 1
		} else if d > 85 {
			min = 1
		}
	case 3:
		if d > 49 && d <= 79 {
			mnd = roll(1, 3)
		} else if d > 79 {
			min = 1
		}
	case 4:
		if d > 42 && d <= 62 {
			mnd = roll(1, 4)
		} else if d > 62 {
			min = 1
		}
	case 5:
		if d > 57 && d <= 67 {
			mnd = roll(1, 4)
		} else if d > 67 {
			min = roll(1, 3)
		}
	case 6:
		if d > 54 && d <= 59 {
			mnd = roll(1, 4)
		} else if d > 59 && d <= 99 {
			min = roll(1, 3)
		} else if d == 100 {
			med = 1
		}
	case 7:
		if d > 51 && d <= 97 {
			min = roll(1, 3)
		} else if d > 97 {
			med = 1
		}
	case 8:
		if d > 48 && d <= 96 {
			min = roll(1, 4)
		} else if d > 96 {
			med = 1
		}
	case 9:
		if d > 43 && d <= 91 {
			min = roll(1, 4)
		} else if d > 91 {
			med = 1
		}
	case 10:
		if d > 40 && d <= 88 {
			min = roll(1, 4)
		} else if d > 88 && d <= 99 {
			med = 1
		} else if d == 100 {
			maj = 1
		}
	case 11:
		if d > 31 && d <= 84 {
			min = roll(1, 4)
		} else if d > 84 && d <= 98 {
			med = 1
		} else if d > 98 {
			maj = 1
		}
	case 12:
		if d > 27 && d <= 82 {
			min = roll(1, 6)
		} else if d > 82 && d <= 97 {
			med = 1
		} else if d > 97 {
			maj = 1
		}
	case 13:
		if d > 19 && d <= 73 {
			min = roll(1, 6)
		} else if d > 73 && d <= 95 {
			med = 1
		} else if d > 95 {
			maj = 1
		}
	case 14:
		if d > 19 && d <= 58 {
			min = roll(1, 6)
		} else if d > 58 && d <= 92 {
			med = 1
		} else if d > 92 {
			maj = 1
		}
	case 15:
		if d > 11 && d <= 46 {
			min = roll(1, 10)
		} else if d > 46 && d <= 90 {
			med = 1
		} else if d > 90 {
			maj = 1
		}
	case 16:
		if d > 40 && d <= 46 {
			min = roll(1, 10)
		} else if d > 46 && d <= 90 {
			med = roll(1, 3)
		} else if d > 90 {
			maj = 1
		}
	case 17:
		if d > 33 && d <= 83 {
			med = roll(1, 3)
		} else if d > 83 {
			maj = 1
		}
	case 18:
		if d > 24 && d <= 80 {
			med = roll(1, 4)
		} else if d > 80 {
			maj = 1
		}
	case 19:
		if d > 4 && d <= 70 {
			med = roll(1, 4)
		} else if d > 70 {
			maj = 1
		}
	case 20:
		if d > 25 && d <= 65 {
			med = roll(1, 4)
		} else if d > 65 {
			maj = roll(1, 3)
		}
	}

	for i := 0; i < mnd; i++ {
		it = append(it, getMundane())
	}
	for i := 0; i < min; i++ {
		it = append(it, getMagicItem("Minor"))
	}
	for i := 0; i < med; i++ {
		it = append(it, getMagicItem("Medium"))
		it = append(it, getMagicItem("Medium")) // 2 rolls per medium
	}
	for i := 0; i < maj; i++ {
		it = append(it, getMagicItem("Major"))
		it = append(it, getMagicItem("Major"))
		it = append(it, getMagicItem("Major")) // 3 rolls per major
	}
	return it
}

func getGem() string {
	d := rand.Intn(100) + 1
	if d <= 25 {
		return "Gem (" + strconv.Itoa(roll(4, 4)) + " gp)"
	}
	if d <= 50 {
		return "Gem (" + strconv.Itoa(roll(2, 4)*10) + " gp)"
	}
	if d <= 70 {
		return "Gem (" + strconv.Itoa(roll(4, 4)*10) + " gp)"
	}
	if d <= 90 {
		return "Gem (" + strconv.Itoa(roll(2, 4)*100) + " gp)"
	}
	if d <= 99 {
		return "Gem (" + strconv.Itoa(roll(4, 4)*100) + " gp)"
	}
	return "Gem (" + strconv.Itoa(roll(2, 4)*1000) + " gp)"
}

func getArt() string {
	d := rand.Intn(100) + 1
	if d <= 10 {
		return "Art Object (" + strconv.Itoa(roll(1, 10)*10) + " gp)"
	}
	if d <= 25 {
		return "Art Object (" + strconv.Itoa(roll(3, 6)*10) + " gp)"
	}
	if d <= 40 {
		return "Art Object (" + strconv.Itoa(roll(1, 6)*100) + " gp)"
	}
	if d <= 50 {
		return "Art Object (" + strconv.Itoa(roll(1, 10)*100) + " gp)"
	}
	if d <= 60 {
		return "Art Object (" + strconv.Itoa(roll(2, 6)*100) + " gp)"
	}
	if d <= 70 {
		return "Art Object (" + strconv.Itoa(roll(3, 6)*100) + " gp)"
	}
	if d <= 80 {
		return "Art Object (" + strconv.Itoa(roll(4, 6)*100) + " gp)"
	}
	if d <= 85 {
		return "Art Object (" + strconv.Itoa(roll(5, 6)*100) + " gp)"
	}
	if d <= 90 {
		return "Art Object (" + strconv.Itoa(roll(1, 4)*1000) + " gp)"
	}
	if d <= 95 {
		return "Art Object (" + strconv.Itoa(roll(1, 6)*1000) + " gp)"
	}
	if d <= 99 {
		return "Art Object (" + strconv.Itoa(roll(2, 4)*1000) + " gp)"
	}
	return "Art Object (" + strconv.Itoa(roll(2, 6)*1000) + " gp)"
}

func getMundane() string {
	d := rand.Intn(100) + 1
	if d <= 17 {
		return "Alchemical Item"
	}
	if d <= 50 {
		return "Mundane Armor/Shield"
	}
	if d <= 83 {
		return "Mundane Weapon"
	}
	return "Tools/Gear"
}

func getMagicItem(class string) string {
	d := rand.Intn(100) + 1
	switch {
	case d <= 20:
		sd := rand.Intn(6) + 1
		if sd <= 4 {
			return fmt.Sprintf("Potion or Oil (%s)", class)
		}
		return "Elixir or Potion (Special)"
	case d <= 25:
		return fmt.Sprintf("Ring (%s)", class)
	case d == 26:
		return "Rod"
	case d <= 41:
		return fmt.Sprintf("Scroll (%s)", class)
	case d == 42:
		return "Staff"
	case d <= 45:
		return fmt.Sprintf("Wand (%s)", class)
	case d == 46:
		return "Wondrous: Book, Libram, Manual, Tome"
	case d <= 48:
		return "Wondrous: Jewel, Phylactery, Amulet"
	case d <= 50:
		return "Wondrous: Cloak or Robe"
	case d <= 52:
		return "Wondrous: Boot, Bracer, Glove"
	case d == 53:
		return "Wondrous: Girdle, Hat, Helm"
	case d <= 55:
		return "Wondrous: Bag, Bottle, Container"
	case d == 56:
		return "Wondrous: Candle, Dust, Stone"
	case d == 57:
		return "Wondrous: Household Item or Tool"
	case d == 58:
		return "Wondrous: Musical Instrument"
	case d <= 60:
		return "Wondrous: Miscellaneous Item"
	case d <= 75:
		return fmt.Sprintf("Magic Armor or Shield (%s)", class)
	default:
		return fmt.Sprintf("Magic Weapon (%s)", class)
	}
}

func (m *Map) PlaceTreasure() {
	for i := range m.Rooms {
		r := &m.Rooms[i]

		hasMonsters := len(r.Monsters) > 0
		isLootRoom := !hasMonsters && (r.RType == "Main Vault" || r.RType == "Reliquary" || r.IsFocal)

		// Place treasure if there's an encounter, or if it's an undefended vault/hub
		if hasMonsters || isLootRoom {
			r.Treasure = GenerateTreasure(m.Config.APL, isLootRoom)
		}
	}
}
