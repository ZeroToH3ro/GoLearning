package main

import "fmt"

func receiveAndSend(c chan int) {
	fmt.Println("Receive data: ", <-c)
	fmt.Println("Send 2")
	c <- 2
}

func main() {
	myChan := make(chan int)

	go receiveAndSend(myChan)
	myChan <- 1

	fmt.Println("Display data: ", <-myChan)

}
