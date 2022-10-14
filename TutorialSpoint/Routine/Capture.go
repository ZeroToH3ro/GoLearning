package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 100; i++ {
		go func(value int) {
			fmt.Println(value) // value ở đây độc lập với i ở ngoài
		}(i) // value i được copy ở đây
	}

	time.Sleep(time.Second)
}