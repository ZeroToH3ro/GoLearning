package main

import "fmt"

func main() {
	/* create a slice */
	numbers := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
