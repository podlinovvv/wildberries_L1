package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		fmt.Println("goroutine ended")
	}()

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	time.Sleep(1 * time.Second)
}
