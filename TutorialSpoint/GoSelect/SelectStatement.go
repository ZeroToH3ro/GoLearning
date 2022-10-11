package main

import "fmt"

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
	number <- 10
}
func chanelMessage(message chan string) {
	message <- "Learning Go Select"
}
