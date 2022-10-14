package main

import (
	"fmt"
	"time"
)

func main() {
	//create two chanel
	number := make(chan int)
	message := make(chan string)

	//function call with goroutine
	go chanelNumber(number)
	go chanelMessage(message)

	//selects and executes a chanel
	select {

	case firstChannel := <-number:
		fmt.Println("Channel Data: ", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data: ", secondChannel)
	}

}

// go routine send integer data to chanel
func chanelNumber(number chan int) {
	time.Sleep(2 * time.Second)
	number <- 10
}

// sleep the process by 2 second
func chanelMessage(message chan string) {
	time.Sleep(2 * time.Second)
	message <- "Learning Go Select"
}
