package main

import "fmt"

type ShapeCalculator interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() int {
	return (r.Width * 2) + (r.Height * 2)
}

// Verificando que la implementacion cumpla con todos los metodos
var _ ShapeCalculator = (*Rectangle)(nil)

func main() {
	rect := &Rectangle{
		Width:  10,
		Height: 3,
	}
	fmt.Println(rect.Area())
	fmt.Println(rect.Perimeter())
}
