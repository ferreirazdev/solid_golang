package solid

// LSP - Liskov Substitution Principle
// is a principle in object-oriented programming that states that objects of a
// superclass should be able to be replaced with objects of a subclass without
// affecting the correctness of the program.

type ShapeLSP interface {
	CalculateArea() float64
}

type RectangleLSP struct {
	width  float64
	height float64
}

func (r RectangleLSP) CalculateArea() float64 {
	return r.width * r.height
}

type SquareLSP struct {
	side float64
}

func (s SquareLSP) CalculateArea() float64 {
	return s.side * s.side
}
