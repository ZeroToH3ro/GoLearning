package main

import "fmt"

func greetingPlanet(namePlanet *string) {
	*namePlanet = "Angular"
	fmt.Println(*namePlanet)
}
func main() {
	star := "Polaris"

	starAddress := &star
	fmt.Println("The address of star is ", starAddress)
	*starAddress = "Sirius"
	fmt.Println("The value of star is ", star)

	greetingPlanet(&star)
}
