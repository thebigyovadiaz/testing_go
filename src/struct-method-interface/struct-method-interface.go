package struct_method_interface

import (
	"math"
)

const (
	PI = math.Pi
)

type RectangleAttr struct {
	Width  float64
	Height float64
}

func (ra RectangleAttr) Perimeter() float64 {
	return 2 * (ra.Width + ra.Height)
}

func (ra RectangleAttr) Area() float64 {
	return ra.Width * ra.Height
}

type CircleAttr struct {
	Radius float64
}

func (c CircleAttr) Perimeter() float64 {
	return 2 * (c.Radius * PI)
}

func (c CircleAttr) Area() float64 {
	return PI * (c.Radius * c.Radius)
}

type TriangleAttr struct {
	Base   float64
	Height float64
}

func (t TriangleAttr) Perimeter() float64 {
	powSum := math.Pow(t.Base, 2.0) + math.Pow(t.Height, 2)
	h := math.Sqrt(powSum)
	return h + t.Base + t.Height
}

func (t TriangleAttr) Area() float64 {
	return t.Base * t.Height
}

type ShapeArea interface {
	Area() float64
}

type ShapePerimeter interface {
	Perimeter() float64
}
