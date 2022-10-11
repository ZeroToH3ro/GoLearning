package main

import "fmt"

// Update marsYear so that it takes earthYears
// as an int parameter
func computeMarsYears(earthYears int) int {

	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func main() {
	myAge := 20

	// Call `computeMarsYears` with `myAge`
	myMartianAge := computeMarsYears(myAge)
	fmt.Println(myMartianAge)
}
