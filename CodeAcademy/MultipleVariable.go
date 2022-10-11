package main

import "fmt"

func main() {
	var magicNum, powerLevel int32

	magicNum = 2048
	powerLevel = 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
}
