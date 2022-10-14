package main

func main() {
	myChan := make(chan int)
	myChan <- 1 // deadlock here
}
