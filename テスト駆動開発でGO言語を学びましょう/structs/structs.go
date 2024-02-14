package main

import (
	"fmt"
	"math"
)

/*
import (

	"fmt"
	"math"

)

	type Rectangle struct {
		Width  float64
		Height float64
	}

	type Circle struct {
		Radius float64
	}

	func Perimeter(rectangle Rectangle) float64 {
		return 2 * (rectangle.Width + rectangle.Height)
	}

	func Area(rectangle Rectangle) float64 {
		return rectangle.Width * rectangle.Height
	}

	func CircleCalculator(circle Circle) float64 {
		return math.Pi * circle.Radius * circle.Radius
	}

	func main() {
		fmt.Println("--result--")
		fmt.Println("Perimeter:", Perimeter(Rectangle{10.0, 10.0}))
		fmt.Println("Area:", Area(Rectangle{12.0, 6.0}))
	}
*/
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	rectangle := Rectangle{12, 6}
	fmt.Println("Area:", rectangle.Area())

	circle := Circle{10}
	fmt.Println("Area:", circle.Area())
}
