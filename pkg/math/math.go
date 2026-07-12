package math

import "math"

func CalcXPTable(xp float64, table []float64, finalLevelXP ...float64) float64 {
	level := 0
	for _, levelXP := range table {
		if xp < levelXP {
			return RoundToTwo(float64(level) + xp/levelXP)
		}

		level++
		xp -= levelXP
	}

	overflow := 0.0
	if len(finalLevelXP) > 0 {
		overflow = finalLevelXP[0]
	}

	if overflow == 0 {
		slope := 600_000.0
		overflow = slope
		if len(table) > 0 {
			overflow += table[len(table)-1]
		}

		for xp >= overflow {
			level++
			xp -= overflow
			overflow += slope
			if level%10 == 0 {
				slope *= 2
			}
		}
	}

	return RoundToTwo(float64(level) + xp/overflow)
}

func RoundToTwo(value float64) float64 {
	return math.Round(value*100) / 100
}
