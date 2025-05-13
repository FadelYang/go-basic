package main

import "fmt"

type Shape interface {
	Area()
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	rectangle := Rectangle{Width: 4, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Printf("Luas dari rectangle adalah: %v\n\n", rectangle.Area())
	fmt.Printf("Luas dari circle adalah: %v", circle.Area())
}
