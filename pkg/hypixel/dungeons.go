package hypixel

import (
	"context"
	"strconv"
	"strings"

	"github.com/chickeniq/hypixel-go/pkg/math"
)

func (c *Client) GetDungeonsStats(ctx context.Context, uuid, profileName string) (*DungeonsStats, error) {
	profile, err := c.GetSkyBlockProfile(ctx, uuid, profileName)
	if err != nil {
		return nil, err
	}

	calcCata := func(xp float64) float64 {
		return math.CalcXPTable(xp, catacombsXPTable[:], 200_000_000)
	}

	dung := profile.Data.Dungeons
	cl := dung.Classes
	classLevels := map[string]float64{
		"berserk": calcCata(cl["berserk"].Experience),
		"archer":  calcCata(cl["archer"].Experience),
		"healer":  calcCata(cl["healer"].Experience),
		"mage":    calcCata(cl["mage"].Experience),
		"tank":    calcCata(cl["tank"].Experience),
	}

	selectedClass := dung.SelectedDungeonClass
	if selectedClass == "" {
		selectedClass = "none"
	}

	classTotal := 0.0
	for _, level := range classLevels {
		classTotal += level
	}

	floorStats := func(dungeonType string, floor int) DungeonFloorStats {
		data := dung.DungeonTypes[dungeonType]
		key := strconv.Itoa(floor)
		return DungeonFloorStats{
			PersonalBest: data.FastestTimeSPlus[key],
			Completions:  data.TierCompletions[key],
		}
	}

	return &DungeonsStats{
		SelectedClass:      strings.ToUpper(selectedClass[:1]) + selectedClass[1:],
		CatacombsLevel:     calcCata(dung.DungeonTypes["catacombs"].Experience),
		ClassAverage:       math.RoundToTwo(classTotal / 5),
		SelectedClassLevel: classLevels[selectedClass],
		SecretsFound:       dung.Secrets,
		ClassLevel: DungeonClasses{
			Berserk: classLevels["berserk"],
			Archer:  classLevels["archer"],
			Healer:  classLevels["healer"],
			Mage:    classLevels["mage"],
			Tank:    classLevels["tank"],
		},
		Entrance: floorStats("catacombs", 0),
		F1:       floorStats("catacombs", 1),
		F2:       floorStats("catacombs", 2),
		F3:       floorStats("catacombs", 3),
		F4:       floorStats("catacombs", 4),
		F5:       floorStats("catacombs", 5),
		F6:       floorStats("catacombs", 6),
		F7:       floorStats("catacombs", 7),
		M1:       floorStats("master_catacombs", 1),
		M2:       floorStats("master_catacombs", 2),
		M3:       floorStats("master_catacombs", 3),
		M4:       floorStats("master_catacombs", 4),
		M5:       floorStats("master_catacombs", 5),
		M6:       floorStats("master_catacombs", 6),
		M7:       floorStats("master_catacombs", 7),
	}, nil
}
