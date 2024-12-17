package domain

type Rectangle struct {
	Length float64
	Width  float64
}

var _ Shape = (*Rectangle)(nil)

func (s *Rectangle) Area() float64 {
	return s.Length * s.Width
}

func (s *Rectangle) Perimeter() float64 {
	return 2 * (s.Length + s.Width)
}
