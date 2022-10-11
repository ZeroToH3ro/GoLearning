package main

import "fmt"

func main() {
	a := "Welcome to my world"
	b := 20
	c := false

	//pass string to the function
	displayValue(a)
	//pass number to the function
	displayValue(b)
	//pass boolean to the function
	displayValue(c)
	//function call with 3 parameters
	displayAnyNumbers(a, b, c)
}

func displayValue(i interface{}) {
	fmt.Println(i)
}

func displayAnyNumbers(i ...interface{}) {
	fmt.Println(i)
}
