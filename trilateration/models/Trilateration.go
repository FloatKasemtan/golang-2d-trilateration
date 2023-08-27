package models

import (
	"errors"
	"math"
	"sort"
)

type Trilateration struct {
	beacons []*Position
}

func NewTrilateration(beacons []*Position) Trilateration {
	return Trilateration{
		beacons: beacons,
	}
}

func (t *Trilateration) GetBeacons() []*Position {
	return t.beacons
}

// Calculation source: https://www.101computing.net/cell-phone-trilateration-algorithm/
func (t *Trilateration) GetPosition() (*Position, error) {
	// TODO: Implement trilateration algorithm
	if len(t.GetBeacons()) < 3 {
		return nil, errors.New("not enough beacons to calculate position")
	}
	// Find most accurate beacon
	beacon := t.GetBeacons()
	sort.Sort(ByRadius(beacon))
	selectedBeacon := beacon[0:3]

	// Calculate position
	a, b, c, d, e, f := prepareVariable(selectedBeacon)
	x := (c*e - f*b) / (e*a - b*d)
	y := (c*d - a*f) / (b*d - a*e)

	return NewPosition(x, y), nil
}

func prepareVariable(beacons []*Position) (float64, float64, float64, float64, float64, float64) {
	a := 2*beacons[1].X - 2*beacons[0].X
	b := 2*beacons[1].Y - 2*beacons[0].Y
	c := math.Pow(beacons[0].R, 2) - math.Pow(beacons[1].R, 2) - math.Pow(beacons[0].X, 2) + math.Pow(beacons[1].X, 2) - math.Pow(beacons[0].Y, 2) + math.Pow(beacons[1].Y, 2)
	d := 2*beacons[2].X - 2*beacons[1].X
	e := 2*beacons[2].Y - 2*beacons[1].Y
	f := math.Pow(beacons[1].R, 2) - math.Pow(beacons[2].R, 2) - math.Pow(beacons[1].X, 2) + math.Pow(beacons[2].X, 2) - math.Pow(beacons[1].Y, 2) + math.Pow(beacons[2].Y, 2)
	return a, b, c, d, e, f
}
