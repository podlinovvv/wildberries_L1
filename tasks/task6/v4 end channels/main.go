package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	end := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		for {
			select {
			case q := <-ch:
				fmt.Println(q)
			case <-end:
				fmt.Println("goroutine done")
				return
			}
		}
	}()

	ch<-1
	ch<-2
	ch<-3

	end<-true
}
