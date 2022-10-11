package main

import "fmt"

func handlePanic() {
	//detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("Recover", a)
	}
}

func division(num1, num2 int) {
	defer handlePanic()

	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)

	}
}

func main() {
	division(4, 2)
	division(4, 0)
	division(16, 8)
}
