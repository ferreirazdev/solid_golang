package solid

import "fmt"

// OCP - Open-Closed Principle
// software should be designed in such a way that new features
// can be added without modifying the existing code.

// Suppose we have an existing Shape interface with a Draw method.
// We also have a Rectangle struct that implements the Shape interface.

type Shape interface {
	Draw() string
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing rectangle with width %d and height %d", r.width, r.height)
}

// Now suppose we need to add a new shape, such as a Circle.
// Instead of modifying the existing Shape interface and the Rectangle struct,
// we can create a new Circle struct that also implements the Shape interface.

type Circle struct {
	radius int
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing circle with radius %d", c.radius)
}
