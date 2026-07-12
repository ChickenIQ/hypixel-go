package hypixel

import (
	"context"

	"github.com/chickeniq/hypixel-go/pkg/hypixel"
	pb "github.com/chickeniq/hypixel-go/proto/hypixel"
)

func (s *Server) GetDungeonsStats(ctx context.Context, req *pb.SkyBlockRequest) (*pb.DungeonsStatsResponse, error) {
	d, err := s.client.GetDungeonsStats(ctx, req.GetUuid(), req.GetProfileName())
	if err != nil {
		return nil, err
	}

	classLvl := d.ClassLevel
	floor := func(stats hypixel.DungeonFloorStats) *pb.DungeonFloorStats {
		return &pb.DungeonFloorStats{
			Completions:  stats.Completions,
			PersonalBest: stats.PersonalBest,
		}
	}

	return &pb.DungeonsStatsResponse{
		ClassAverage:       d.ClassAverage,
		CatacombsLevel:     d.CatacombsLevel,
		SelectedClass:      d.SelectedClass,
		SelectedClassLevel: d.SelectedClassLevel,
		SecretsFound:       d.SecretsFound,
		ClassLevel: &pb.DungeonClasses{
			Healer:  classLvl.Healer,
			Mage:    classLvl.Mage,
			Tank:    classLvl.Tank,
			Berserk: classLvl.Berserk,
			Archer:  classLvl.Archer,
		},
		Entrance: floor(d.Entrance),
		F1:       floor(d.F1),
		F2:       floor(d.F2),
		F3:       floor(d.F3),
		F4:       floor(d.F4),
		F5:       floor(d.F5),
		F6:       floor(d.F6),
		F7:       floor(d.F7),
		M1:       floor(d.M1),
		M2:       floor(d.M2),
		M3:       floor(d.M3),
		M4:       floor(d.M4),
		M5:       floor(d.M5),
		M6:       floor(d.M6),
		M7:       floor(d.M7),
	}, nil
}
