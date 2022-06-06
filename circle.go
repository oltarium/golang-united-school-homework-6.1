package shape

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (t Circle) CalcPerimeter() float64 {
	return t.Radius * 2 * math.Pi
}

// CalcArea returns calculation result of area
func (t Circle) CalcArea() float64 {
	return math.Pi * math.Pow(t.Radius, 2)
}
