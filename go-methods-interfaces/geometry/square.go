package geometry

type Square struct {
	Size float64
}

func (s Square) Area() float64 {
	return s.Size * s.Size
}

func (s Square) Perimeter() float64 {
	return s.Size * 4
}
