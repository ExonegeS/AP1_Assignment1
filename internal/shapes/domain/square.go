package domain

type Square struct {
	Length float64
}

var _ Shape = (*Square)(nil)

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func (s *Square) Perimeter() float64 {
	return 4 * s.Length
}
