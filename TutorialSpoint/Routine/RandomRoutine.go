package main

import (
	"fmt"
	"time"
)

func sayHello(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(str)
	}
}

func main() {
	go sayHello("Minh")

	sayHello("Vy")
	time.Sleep(time.Second)
}
