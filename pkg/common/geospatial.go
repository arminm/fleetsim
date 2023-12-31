package common

import "math"

type Position struct {
	Latitude  float64
	Longitude float64
	Heading   float64
}

func PositionsWithinDistance(pos1 Position, pos2 Position, maxDistance float64) bool {
	return Distance(pos1, pos2) <= maxDistance
}

func Distance(pos1 Position, pos2 Position) float64 {
	x1, y1 := pos1.Latitude, pos1.Longitude
	x2, y2 := pos2.Latitude, pos2.Longitude
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
