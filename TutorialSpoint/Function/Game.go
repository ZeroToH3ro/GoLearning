package main

import (
	"fmt"
	"math/rand"
)

var first_user, second_user string

func input_user() {
	fmt.Print("Enter name of first user: ")
	fmt.Scan(&first_user)
	fmt.Print("Enter name of second user: ")
	fmt.Scan(&second_user)
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
	fmt.Printf("\nProgress user %s: ", first_user)
	user_1 := rand_step()
	fmt.Printf("Progress user %s: ", second_user)
	user_2 := rand_step()

	if user_1 > user_2 {
		fmt.Println("User 1 win with result: ", user_1)
	} else {
		fmt.Println("User 2 win with result: ", user_2)
	}
}
