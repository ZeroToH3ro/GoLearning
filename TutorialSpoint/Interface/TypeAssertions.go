package main

import "fmt"

func main() {
	//create an empty interface
	var a interface{}
	//store integer to an empty interface
	a = 10
	//type assertion
	interfaceValue, booleanValue := a.(int)

	fmt.Println(interfaceValue)
	fmt.Println(booleanValue)
	/*
		b = "go"
		interfaceValue2 := b.(int)
		fmt.Println(interfaceValue2)
	*/

}
