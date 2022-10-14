package main

import "fmt"

func GFC(i func(p, q string) string) {
	fmt.Println(i("Hello", " My name is Minh, I am a "))
}

func main() {
	//call without parameter
	value_1 := func() {
		fmt.Println("Hello world")
	}
	value_1()
	//call with parameter
	func(str string) {
		fmt.Println(str)
	}("My nam is Minh")
	//function as an argument into
	//other function
	value_2 := func(p, q string) string {
		return p + q + "Professor"
	}
	GFC(value_2)
}
