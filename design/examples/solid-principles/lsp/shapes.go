package lsp

// BAD: Violates LSP - Square cannot truly substitute Rectangle
type Rectangle struct {
	width, height float64
}

func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r *Rectangle) SetHeight(h float64) {
	r.height = h
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type SquareBad struct {
	Rectangle
}

func (s *SquareBad) SetWidth(w float64) {
	s.width = w
	s.height = w // Breaks LSP!
}

func (s *SquareBad) SetHeight(h float64) {
	s.width = h
	s.height = h // Breaks LSP!
}

// GOOD: Follows LSP - proper abstraction
type Shape interface {
	Area() float64
}

type RectangleGood struct {
	width, height float64
}

func NewRectangle(width, height float64) *RectangleGood {
	return &RectangleGood{width: width, height: height}
}

func (r *RectangleGood) Area() float64 {
	return r.width * r.height
}

type SquareGood struct {
	side float64
}

func NewSquare(side float64) *SquareGood {
	return &SquareGood{side: side}
}

func (s *SquareGood) Area() float64 {
	return s.side * s.side
}

// Both can be used interchangeably through Shape interface
func CalculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
