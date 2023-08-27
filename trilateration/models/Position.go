package models

type Position struct {
	X float64
	Y float64
	R float64
}

func NewPosition(x float64, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) SetRadius(r float64) {
	p.R = r
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
