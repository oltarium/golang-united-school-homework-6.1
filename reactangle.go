package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (t Rectangle) CalcPerimeter() float64 {
	return t.Height*2 + t.Weight*2
}

// CalcArea returns calculation result of area
func (t Rectangle) CalcArea() float64 {
	return t.Height * t.Weight
}
