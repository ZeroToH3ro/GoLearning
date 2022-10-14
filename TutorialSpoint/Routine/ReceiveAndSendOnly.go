package main

import "fmt"

func receiveOnly(c <-chan int) {
	fmt.Println("Receive data: ", <-c)
}

func sendOnly(c chan<- int) {
	c <- 3
}
func main() {
	myChan := make(chan int)

	go sendOnly(myChan)
	go receiveOnly(myChan)

	myChan <- 5
	fmt.Println("Display data: ", <-myChan)

}
