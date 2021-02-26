package haversine

import (
	"math"
)

const RADIUS = 6371e3

func CalculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	f1 := lat1 * math.Pi / 180
	f2 := lat2 * math.Pi / 180
	d1 := (lat2 - lat1) * math.Pi / 180
	d2 := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(d1 / 2) * math.Sin(d1 / 2) + math.Cos(f1) * math.Cos(f2) * math.Sin(d2 / 2) * math.Sin(d2 / 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
	return RADIUS * c
}
