package main

import (
	"fmt"
	"math/rand"
)

func input_user() {
	fmt.Println("Enter name of first user: ")
	var first_user, second_user string
	fmt.Scanln(&first_user)
	fmt.Println("Enter name of second user: ")
	fmt.Scanln(&second_user)
}

func rand_step() int {
	result := 0
	for i := 1; i <= 5; i++ {
		step_rand := rand.Intn(10)
		fmt.Printf("Step %d: %d cm\n", i, step_rand)
		result += step_rand
	}
	return result
}
func main() {
	input_user()
	fmt.Println("\nProgress user 1: ")
	user_1 := rand_step()
	fmt.Println("Progress user 2: ")
	user_2 := rand_step()

	if user_1 > user_2 {
		fmt.Println("User 1 win with result: ", user_1)
	} else {
		fmt.Println("User 2 win with result: ", user_2)
	}
}
