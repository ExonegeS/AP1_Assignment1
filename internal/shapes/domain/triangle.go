package domain

import "math"

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

var _ Shape = (*Triangle)(nil)

func (s *Triangle) Area() float64 {
	var p = s.Perimeter() / 2
	return math.Sqrt(p * (p - s.SideA) * (p - s.SideB) * (p - s.SideC))

}

func (s *Triangle) Perimeter() float64 {
	return s.SideA + s.SideB + s.SideC
}
