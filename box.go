package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("Exceed capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("Index out of range")
	}
	for index := 0; index < len(b.shapes); index++ {
		if index == i && b.shapes[index] != nil {
			return b.shapes[index], nil
		}
	}
	return nil, errors.New("Not found")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("Out of range")
	}
	element := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	if element == nil {
		return nil, errors.New("No element at index")
	}
	return element, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Out of range")
	}
	if b.shapes[i] == nil {
		return nil, errors.New("No element at index")
	}
	element := b.shapes[i]
	b.shapes[i] = shape
	return element, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var area float64 = 0
	for i := 0; i < len(b.shapes); i++ {
		area += b.shapes[i].CalcPerimeter()
	}
	return area

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64 = 0
	for i := 0; i < len(b.shapes); i++ {
		area += b.shapes[i].CalcArea()
	}
	return area

}

func isCircle(t interface{}) bool {
	switch t.(type) {
	case Circle:
		return true
	default:
		return false
	}
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	result := []Shape{}
	isCircleExists := false
	for _, el := range b.shapes {
		if !isCircle(el) {
			result = append(result, el)
		} else {
			isCircleExists = true
		}
	}
	if !isCircleExists {
		return errors.New("No Circles in the list")
	}
	b.shapes = result
	return nil
}
