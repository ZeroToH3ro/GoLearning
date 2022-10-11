package main

import "fmt"

func main() {
	coolSneakers := 65.99
	niceNecklace := 45.50

	var taxCalculation float64

	taxCalculation += coolSneakers

	taxCalculation += niceNecklace

	// Compute the NYC sales tax
	// 8.875% of the purchase here:
	taxCalculation *= 0.08875
	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)
}
