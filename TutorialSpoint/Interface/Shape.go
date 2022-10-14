package main

import "fmt"

// interface
type Shape interface {
	area()
}

// Rectangle struct implements the interface
type Rectangle struct {
	length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// Triangle struct implements the interface
type Triangle struct {
	base, height float32
}

// Triangle provides implementation for area()
func (t Triangle) area() int {
	return 0.5 * t.base * t.height
}

// access method of the interface
//func calculate(s Shape)  {
//	return s.area()
//}

// main function
func main() {

	// assigns value to struct members
	r := Rectangle{7, 4}
	t := Triangle{8, 12}

	// call calculate() with struct variable rect
	//rectangleArea := calculate(r)
	fmt.Println("Area of Rectangle:")

	triangleArea := calculate(t)
	fmt.Println("Area of Triangle:", triangleArea)

}
