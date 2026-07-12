package hypixel

var catacombsXPTable = [...]float64{
	50, 75, 110, 160, 230, 330, 470, 670, 950, 1340,
	1890, 2665, 3760, 5260, 7380, 10300, 14400, 20000, 27600, 38000,
	52500, 71500, 97000, 132000, 180000, 243000, 328000, 445000, 600000, 800000,
	1065000, 1410000, 1900000, 2500000, 3300000, 4300000, 5600000, 7200000, 9200000, 12000000,
	15000000, 19000000, 24000000, 30000000, 38000000, 48000000, 60000000, 75000000, 93000000, 116250000,
}

type DungeonClasses struct {
	Healer  float64
	Mage    float64
	Tank    float64
	Berserk float64
	Archer  float64
}

type DungeonFloorStats struct {
	Completions  float64
	PersonalBest float64
}

type DungeonsStats struct {
	ClassAverage       float64
	CatacombsLevel     float64
	SelectedClass      string
	SelectedClassLevel float64
	SecretsFound       float64
	ClassLevel         DungeonClasses
	Entrance           DungeonFloorStats
	F1                 DungeonFloorStats
	F2                 DungeonFloorStats
	F3                 DungeonFloorStats
	F4                 DungeonFloorStats
	F5                 DungeonFloorStats
	F6                 DungeonFloorStats
	F7                 DungeonFloorStats
	M1                 DungeonFloorStats
	M2                 DungeonFloorStats
	M3                 DungeonFloorStats
	M4                 DungeonFloorStats
	M5                 DungeonFloorStats
	M6                 DungeonFloorStats
	M7                 DungeonFloorStats
}
