package domain

import "math"

type Circle struct {
	Radius float64
}

var _ Shape = (*Circle)(nil)

func (s *Circle) Area() float64 {
	return math.Pi * s.Radius * s.Radius
}

func (s *Circle) Perimeter() float64 {
	return 2 * s.Radius * math.Pi
}
