package models

import (
	"errors"
	"math"
)

const AvgTagHeight = 100.0

type Position struct {
	X      float64
	Y      float64
	Height float64
	R      float64
}

func NewPosition(x float64, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func NewPositionWithHeight(x float64, y float64, height float64) *Position {
	return &Position{
		X:      x,
		Y:      y,
		Height: height,
	}
}

func (p *Position) SetRadius(distance float64) error {
	r := math.Sqrt(math.Abs(math.Pow(distance, 2) - math.Pow(p.Height-AvgTagHeight, 2)))
	if math.IsNaN(r) {
		return errors.New("invalid radius")
	}
	p.R = r
	return nil
}

type ByRadius []*Position

func (s ByRadius) Len() int {
	return len(s)
}

func (s ByRadius) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByRadius) Less(i, j int) bool {
	return s[i].R < s[j].R
}
