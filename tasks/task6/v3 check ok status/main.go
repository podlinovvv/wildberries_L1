package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("end of goroutine")
				return
			}
			fmt.Println(v)
		}
	}()

	ch<-1
	ch<-2
	ch<-3
	close(ch)

	time.Sleep(1 * time.Second)
}
