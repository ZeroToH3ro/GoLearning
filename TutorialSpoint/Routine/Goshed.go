package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func () {
		for i:=1; i<=50; i++ {
			fmt.Println("I am goroutine 1")
			runtime.Gosched()
		}
	} ()
	
	go func () {
		for i:=1; i<=50; i++ {
			fmt.Println("I am goroutine 2")
			runtime.Gosched()
		}
	} ()
	
	time.Sleep(time.Second)
}